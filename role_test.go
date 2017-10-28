package acl_test

import (
	"fmt"
	"testing"

	"github.com/s4l1h/acl"
)

func TestRole(t *testing.T) {

	// New Module : members
	membersModule := acl.NewModule("members", "members module")
	// New add Perm With Module
	addPerm, err := acl.NewPermWithModule("add", "can add new member?", membersModule)
	if err != nil {
		t.Error("Error Create Perm")
	}
	// New remove Perm With Module
	removePerm, err := acl.NewPermWithModule("remove", "can remove member?", membersModule)
	if err != nil {
		t.Error("Error Create Perm")
	}
	// New remove Perm With Module
	listPerm, err := acl.NewPermWithModule("list", "can show member list?", membersModule)
	if err != nil {
		t.Error("Error Create Perm")
	}
	// New remove Perm With Module
	updatePerm, err := acl.NewPermWithModule("update", "can update member?", membersModule)
	if err != nil {
		t.Error("Error Create Perm")
	}
	// Create New Role
	role := acl.NewRole("ROLE_ADMIN", "admin role")
	// Add Permission to Role
	if err := role.Add(membersModule, addPerm); err != nil {
		t.Error(err)
	}
	// Add Permission to Role
	if err := role.Add(membersModule, removePerm); err != nil {
		t.Error(err)
	}
	// Add Permission to Role
	if err := role.Add(membersModule, listPerm); err != nil {
		t.Error(err)
	}
	// Add Permission to Role
	if err := role.Add(membersModule, updatePerm); err != nil {
		t.Error(err)
	}
	// Check module has list perm?
	if ok := role.HasPermWithObject(membersModule, listPerm); ok != true {
		t.Error("Role hasperm Problem...s..", ok)
	}
	// this role have members list role?
	if ok := role.HasPermWithName("members", "list"); ok != true {
		t.Error("Role HasPermWithName Problem...s..", ok)
	}

	// this role have members list role?
	if ok := role.HasPerm(membersModule, "list"); ok != true {
		t.Error("Role HasPermWithName Problem...s..", ok)
	}
	// this role have members list role?
	if ok := role.HasPerm("members", listPerm); ok != true {
		t.Error("Role HasPermWithName Problem...s..", ok)
	}

	// Now remove listPerm from role
	if err := role.Remove(membersModule, listPerm); err != nil {
		t.Error("Role Remove Problem...s..", err)
	}
	// Now remove addPerm from role
	if err := role.Remove("members", addPerm); err != nil {
		t.Error("Role Remove Problem...s..", err)
	}
	// Now remove addPerm from role
	if err := role.RemoveWithObject(membersModule, updatePerm); err != nil {
		t.Error("Role Remove Problem...s..", err)
	}

	// Check again module has list perm?
	if ok := role.HasPerm(membersModule, listPerm); ok != false {
		t.Error("After Remove Role hasperm Problem...s..", ok)
	}

	// this role have members add role?
	if ok := role.HasPermWithName("members", "add"); ok != false {
		t.Error("Role HasPermWithName Problem...s..", ok)
	}
	// this role have members list role?
	if ok := role.HasPermWithName("members", "list"); ok != false {
		t.Error("Role HasPermWithName Problem...s..", ok)
	}

	/*
		GENERATE DUMMY DATA
	*/
	testList := make(map[string]string)
	testList["add"] = "%s add Role"
	testList["delete"] = "%s delete Role"
	testList["update"] = "%s update Role"
	testList["list"] = "%s list Role"

	// New Module : articles
	articlesModule := acl.NewModule("articles", "articles module")
	// New Module : faq
	faqModule := acl.NewModule("faq", "faq module")
	// New Module : blog
	blogModule := acl.NewModule("blog", "blog module")

	for k, v := range testList {
		// Create dummy data
		// articles add,delete,update,list role
		role.Add(articlesModule, acl.NewPerm(k, fmt.Sprintf(v, articlesModule.GetName())))
		// faq add,delete,update,list role
		role.Add(faqModule, acl.NewPerm(k, fmt.Sprintf(v, faqModule.GetName())))
		// blog add,delete,update,list role
		role.Add(blogModule, acl.NewPerm(k, fmt.Sprintf(v, blogModule.GetName())))
	}

	if role.HasModule("articles") != true {
		t.Error("Role HasModule Problem...s..")
	}
	if role.HasModule(faqModule) != true {
		t.Error("Role HasModule Problem...s..")
	}

	role.RemoveModule(articlesModule)
	role.RemoveModule("faq")

	if role.HasModule(articlesModule) != false {
		t.Error("Role HasModule Problem...s..")
	}
	if role.HasModule("faq") != false {
		t.Error("Role HasModule Problem...s..")
	}

	if role.HasModule("blog") != true {
		t.Error("Role HasModule Problem...s..")
	}
	if role.HasModule(blogModule) != true {
		t.Error("Role HasModule Problem...s..")
	}
}
