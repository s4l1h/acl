package acl_test

import (
	"testing"

	"github.com/s4l1h/acl"
)

const (
	MODULE_NAME = "Members"
	MODULE_DESC = "Members Module"
	PERM_NAME   = "add"
	PERM_DESC   = "Add Perms"
)

// TestModuleNameDesc test module setter getter
func TestModuleNameDesc(t *testing.T) {

	module := acl.NewModule(MODULE_NAME, MODULE_DESC)

	if module.GetName() != MODULE_NAME {
		t.Errorf("Module Name Error : %v != %v", module.GetName(), MODULE_NAME)
	}
	if module.GetDesc() != MODULE_DESC {
		t.Errorf("Module desc Error : %v != %v", module.GetDesc(), MODULE_DESC)
	}

	//Test Set Name,Desc
	nname := MODULE_NAME + " New"
	ndesc := MODULE_DESC + " New"
	module.SetName(nname)
	module.SetDesc(ndesc)

	if module.GetName() != nname {
		t.Errorf("Module SetName Error : %v != %v", module.GetName(), nname)
	}
	if module.GetDesc() != ndesc {
		t.Errorf("Module SetDesc Error : %v != %v", module.GetDesc(), ndesc)
	}

}

// TestPermNameDesc test perm setter getter
func TestPermNameDesc(t *testing.T) {

	perm := acl.NewPerm(PERM_NAME, PERM_DESC)
	if perm.GetName() != PERM_NAME {
		t.Errorf("Perm Name Error : %v != %v", perm.GetName(), PERM_NAME)
	}
	if perm.GetDesc() != PERM_DESC {
		t.Errorf("Perm desc Error : %v != %v", perm.GetDesc(), PERM_DESC)
	}
	//Test Set Name,Desc
	nname := PERM_NAME + " New"
	ndesc := PERM_DESC + " New"
	perm.SetName(nname)
	perm.SetDesc(ndesc)

	if perm.GetName() != nname {
		t.Errorf("Perm SetName Error : %v != %v", perm.GetName(), nname)
	}
	if perm.GetDesc() != ndesc {
		t.Errorf("Perm SetDesc Error : %v != %v", perm.GetDesc(), ndesc)
	}
}
