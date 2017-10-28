package acl

import "fmt"

// Role acl role object
type Role struct {
	BaseACL
	Data map[string]map[string]*Perm
}

// NewRole create new role
func NewRole(name string, desc string) *Role {
	v := &Role{}
	v.SetName(name)
	v.SetDesc(desc)
	v.Data = map[string]map[string]*Perm{}
	return v
}

// Add add perm role
func (role *Role) Add(module *Module, perm *Perm) error {
	if !module.Verify() || !perm.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	// LOLOLO :=> panic: assignment to entry in nil map
	mm, ok := role.Data[module.GetName()]
	if !ok {
		mm = make(map[string]*Perm)
		role.Data[module.GetName()] = mm
	}

	role.Data[module.GetName()][perm.GetName()] = perm
	return nil
}

// Remove check role has perms??
func (role *Role) Remove(module, perm interface{}) error {

	mName := ""
	pName := ""

	switch v := module.(type) {
	case string:
		mName = v // Create New Perm
	case *Module:
		mName = v.GetName()
	}

	switch v := perm.(type) {
	case string:
		pName = v // Create New Perm
	case *Perm:
		pName = v.GetName()
	}
	return role.RemoveWithName(mName, pName)

}

// RemoveWithObject add perm role
func (role *Role) RemoveWithObject(module *Module, perm *Perm) error {
	if !module.Verify() || !perm.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}

	delete(role.Data[module.GetName()], perm.GetName())
	//role.Data[module.GetName()][perm.GetName()] = perm
	return nil
}

// RemoveWithName add perm role
func (role *Role) RemoveWithName(module, perm string) error {
	if module == "" || perm == "" {
		return fmt.Errorf(ErrorEmptyName)
	}
	delete(role.Data[module], perm)
	return nil
}

// HasModule : role has module?
func (role *Role) HasModule(module interface{}) bool {

	switch v := module.(type) {
	case string:
		_, ok := role.Data[v]
		return ok
	case *Module:
		_, ok := role.Data[v.GetName()]
		return ok
	}
	return false
}

// RemoveModule : Remove module and perms from role
func (role *Role) RemoveModule(module interface{}) {

	switch v := module.(type) {
	case string:
		role.RemoveModuleWithName(v) // Remove Module
	case *Module:
		role.RemoveModuleWithObject(v)
	}
}

// RemoveModuleWithObject : Remove module and perms from role
func (role *Role) RemoveModuleWithObject(module *Module) {
	role.RemoveModuleWithName(module.GetName())
}

// RemoveModuleWithName : Remove module and perms from role
func (role *Role) RemoveModuleWithName(module string) {
	delete(role.Data, module)
}

// HasPerm check role has perms??
func (role *Role) HasPerm(module, perm interface{}) bool {

	mName := ""
	pName := ""

	switch v := module.(type) {
	case string:
		mName = v // get Module Name
	case *Module:
		mName = v.GetName() //get Module Name
	}

	switch v := perm.(type) {
	case string:
		pName = v // get Perm Name
	case *Perm:
		pName = v.GetName() // get Perm Name
	}
	//fmt.Println(mName, pName)
	return role.HasPermWithName(mName, pName)

}

// HasPermWithObject check perm exists
func (role *Role) HasPermWithObject(module *Module, perm *Perm) bool {
	if !module.Verify() || !perm.Verify() {
		return false
	}
	mm, ok := role.Data[module.GetName()]
	// Module Not Found
	if !ok {
		return false
	}
	// Module not have perm
	if _, ok := mm[perm.GetName()]; !ok {
		return false
	}
	return true
}

// HasPermWithName check perm exists
func (role *Role) HasPermWithName(module, perm string) bool {
	if module == "" || perm == "" {
		return false
	}
	mm, ok := role.Data[module]
	// Module Not Found
	if !ok {
		return false
	}
	// Module not have perm
	if _, ok := mm[perm]; !ok {
		return false
	}
	return true
}
