package acl

import "fmt"

//ErrorEmptyName module not found error
const ErrorEmptyName = "Empty name is not accepted."

//ErrorModuleNotFound module not found error
const ErrorModuleNotFound = "%s Module Not Found"

//ErrorModuleRegistered if module is registered error.
const ErrorModuleRegistered = "%s Module is Registered . Maybe need UpdateModule?"

//ErrorRoleNotFound module not found error
const ErrorRoleNotFound = "%s Role Not Found"

//ErrorRoleRegistered if module is registered error.
const ErrorRoleRegistered = "%s Role is Registered . Maybe need UpdateModule?"

// ModuleManager type
type ModuleManager map[string]*Module

// RoleManager type
type RoleManager map[string]*Role

// ACL acl manager
type ACL struct {
	ModuleManager ModuleManager
	RoleManager   RoleManager
}

// New new acl manager
func New() *ACL {
	return &ACL{
		ModuleManager: make(ModuleManager),
		RoleManager:   make(RoleManager),
	}
}

// AddModule Register New Module
func (acl *ACL) AddModule(module *Module) error {
	if !module.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	if acl.HasModule(module.GetName()) {
		return fmt.Errorf(ErrorModuleRegistered, module.GetName())
	}
	acl.ModuleManager[module.GetName()] = module
	return nil
}

// UpdateModule Update Module
func (acl *ACL) UpdateModule(module *Module) error {
	if !module.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	acl.ModuleManager[module.GetName()] = module
	return nil
}

// GetModuleWithName Get Module Object
func (acl *ACL) GetModuleWithName(name string) (*Module, error) {
	if name == "" {
		return nil, fmt.Errorf(ErrorEmptyName)
	}
	if acl.HasModule(name) {
		v, _ := acl.ModuleManager[name]
		return v, nil
	}
	return nil, fmt.Errorf(ErrorModuleNotFound, name)
}

// RemoveModule Remove Module
func (acl *ACL) RemoveModule(module *Module) error {
	return acl.RemoveModuleWithName(module.GetName())
}

// RemoveModuleWithName Remove With Name
func (acl *ACL) RemoveModuleWithName(name string) error {
	if name == "" {
		return fmt.Errorf(ErrorEmptyName)
	}
	if acl.HasModule(name) {
		delete(acl.ModuleManager, name)
		return nil
	}
	return fmt.Errorf(ErrorModuleNotFound, name)
}

// HasModule Check Module exists
func (acl *ACL) HasModule(name string) bool {
	_, ok := acl.ModuleManager[name]
	return ok
}

// HasPerm Check Module exists
func (acl *ACL) HasPerm(moduleName, permName string) bool {
	module, err := acl.GetModuleWithName(moduleName)
	if err != nil {
		return false
	}
	return module.HasPermWithName(permName)
}

// AddRole Register New Role
func (acl *ACL) AddRole(Role *Role) error {
	if !Role.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	if acl.HasRole(Role.GetName()) {
		return fmt.Errorf(ErrorRoleRegistered, Role.GetName())
	}
	acl.RoleManager[Role.GetName()] = Role
	return nil
}

// UpdateRole Update Role
func (acl *ACL) UpdateRole(Role *Role) error {
	if !Role.Verify() {
		return fmt.Errorf(ErrorEmptyName)
	}
	acl.RoleManager[Role.GetName()] = Role
	return nil
}

// GetRoleWithName Get Role Object
func (acl *ACL) GetRoleWithName(name string) (*Role, error) {
	if name == "" {
		return nil, fmt.Errorf(ErrorEmptyName)
	}
	if acl.HasRole(name) {
		v, _ := acl.RoleManager[name]
		return v, nil
	}
	return nil, fmt.Errorf(ErrorRoleNotFound, name)
}

// RemoveRole Remove Role
func (acl *ACL) RemoveRole(Role *Role) error {
	return acl.RemoveRoleWithName(Role.GetName())
}

// RemoveRoleWithName Remove With Name
func (acl *ACL) RemoveRoleWithName(name string) error {
	if name == "" {
		return fmt.Errorf(ErrorEmptyName)
	}
	if acl.HasRole(name) {
		delete(acl.RoleManager, name)
		return nil
	}
	return fmt.Errorf(ErrorRoleNotFound, name)
}

// HasRole Check Role exists
func (acl *ACL) HasRole(name string) bool {
	_, ok := acl.RoleManager[name]
	return ok
}

// Has Check Role exists
func (acl *ACL) Has(RoleName, module, permName string) bool {
	Role, err := acl.GetRoleWithName(RoleName)
	if err != nil {
		return false
	}
	return Role.HasPermWithName(module, permName)
}
