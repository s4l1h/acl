package acl_test

import (
	"testing"

	"github.com/s4l1h/acl"
)

func TestPerm(t *testing.T) {
	perm := acl.NewPerm("add", "add Perm")

	if perm.GetModule() != nil {
		t.Error("Error Perm Module must be nil : ", perm.GetModule())
	}

	mname := "members"
	mdesc := "members module"
	module := acl.NewModule(mname, mdesc)
	// Set Module
	perm.SetModule(module)
	// Check Perm Module
	if perm.GetModule() == nil {
		t.Error("Error Perm SetModule : ", perm.GetModule())
	}
	// Check Perm Module Name
	if perm.GetModule().GetName() != mname {
		t.Error("Error Perm GetModule.GetName : ", perm.GetModule())
	}
	// Check Perm Module Desc
	if perm.GetModule().GetDesc() != mdesc {
		t.Error("Error Perm GetModule.GetDesc : ", perm.GetModule())
	}

	checkNewModule(perm, t)

}

func TestPermWithModule(t *testing.T) {

	mname := "members"
	mdesc := "members module"
	module := acl.NewModule(mname, mdesc)

	perm, err := acl.NewPermWithModule("add", "add Perm", module)
	if err != nil {
		t.Errorf("Error NewPermWithModule %s", err)
	}

	// Check Perm Module
	if perm.GetModule() == nil {
		t.Error("Error Perm SetModule : ", perm.GetModule())
	}
	// Check Perm Module Name
	if perm.GetModule().GetName() != mname {
		t.Error("Error Perm GetModule.GetName : ", perm.GetModule())
	}
	// Check Perm Module Desc
	if perm.GetModule().GetDesc() != mdesc {
		t.Error("Error Perm GetModule.GetDesc : ", perm.GetModule())
	}
	checkNewModule(perm, t)

}

func checkNewModule(perm *acl.Perm, t *testing.T) {

	name := "testModule"
	desc := "testModule Desc"
	module := acl.NewModule(name, desc)

	perm.SetModule(module)
	// Check Perm Module Name
	if perm.GetModule().GetName() != name {
		t.Error("Error Perm GetModule.GetName : ", perm.GetModule())
	}
	// Check Perm Module Desc
	if perm.GetModule().GetDesc() != desc {
		t.Error("Error Perm GetModule.GetDesc : ", perm.GetModule())
	}
}
