package acl

import "encoding/json"

//Perm : Permission Object
type Perm struct {
	BaseACL         // BaseACL Object Name,Desc Setter Getter
	Module  *Module `json:"module"`
}

// SetModule : Set Permission module
func (perm *Perm) SetModule(module *Module) {
	perm.Module = module
}

// GetModule : Get Permission module
func (perm *Perm) GetModule() *Module {
	return perm.Module
}

// NewPerm : Create New Permission
func NewPerm(name string, desc string) *Perm {
	r := &Perm{}
	r.SetName(name)
	r.SetDesc(desc)
	return r
}

// NewPermWithModule : Create New Permission with Module
func NewPermWithModule(name string, desc string, module *Module) (*Perm, error) {
	r := &Perm{}
	r.SetName(name)
	r.SetDesc(desc)
	r.SetModule(module)
	err := module.AddPerm(r)
	return r, err
}

// MarshalJSON : Perm Object to Json
// Need custom function because *Module is problem.
func (perm *Perm) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name   string   `json:"name"`
		Desc   string   `json:"desc"`
		Module *BaseACL `json:"module"`
	}{
		Name: perm.GetName(),
		Desc: perm.GetDesc(),
		Module: &BaseACL{
			Name: perm.Module.GetName(),
			Desc: perm.Module.GetDesc(),
		},
	})
}
