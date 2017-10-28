package acl

import "fmt"

//ErrorPermRegistered if module is registered error.
const ErrorPermRegistered = "%s Perm is Registered in %s. Maybe need UpdatePerm?"

//ErrorInvalidPerm Invalid Perm name or object
const ErrorInvalidPerm = "%s %v Invalid Perm Name"

//Module : Modül
type Module struct {
	BaseACL
	Perms map[string]*Perm `json:"perms"` // Perms
}

// NewModule create new module
func NewModule(name string, desc string) *Module {
	v := &Module{}
	v.SetName(name)
	v.SetDesc(desc)
	v.Perms = make(map[string]*Perm)
	return v
}

// AddPermObject : Add Perm to Module
func (module *Module) AddPermObject(perm *Perm) error {
	if !perm.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	if module.HasPerm(perm) {
		return fmt.Errorf(ErrorPermRegistered, perm.GetName(), module.GetName())
	}
	module.Perms[perm.GetName()] = perm
	perm.SetModule(module)
	return nil
}

// AddPerm : Add Perm to Module
// Accept: Perm name and Perm object perm ...*Perm
func (module *Module) AddPerm(name interface{}) error {
	switch v := name.(type) {
	case string:
		perm := NewPerm(v, "") // Create New Perm
		return module.AddPermObject(perm)
	case *Perm:
		return module.AddPermObject(v)
	}
	return fmt.Errorf(ErrorInvalidPerm, name, name)
}

// UpdatePerm : Add Perm to Module
func (module *Module) UpdatePerm(perm *Perm) error {
	if !perm.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	module.Perms[perm.GetName()] = perm
	perm.SetModule(module)
	return nil
}

// GetPerms : Get Module Perms
func (module *Module) GetPerms() map[string]*Perm {
	return module.Perms
}

// RemovePerm : Remove Perms From Module.
// Accept: Perm name and Perm object perm ...*Perm
// We can combine string perm name and perm object : RemovePerms("add",&Perm{Name:"delete"})
func (module *Module) RemovePerm(perm ...interface{}) {
	for _, v := range perm {
		switch t := v.(type) {
		case string:
			module.RemovePermWithName(t)
		case *Perm:
			module.RemovePermWithName(t.GetName())
		}
	}
}

// RemovePermWithName : Remove Perm From Module use Perm name
func (module *Module) RemovePermWithName(name string) {
	delete(module.Perms, name)
}

// HasPermObject : Module Has Perm?
func (module *Module) HasPermObject(perm *Perm) bool {
	return module.HasPermWithName(perm.GetName())
}

// HasPermWithName : Module Has Perm check with name?
func (module *Module) HasPermWithName(name string) bool {
	_, ok := module.Perms[name]
	return ok
}

// HasPerm : Module Has Perm? Accept perm name and *Perm object
func (module *Module) HasPerm(name interface{}) bool {
	switch v := name.(type) {
	case string:
		return module.HasPermWithName(v)
	case *Perm:
		return module.HasPermObject(v)
	}
	return false
}
