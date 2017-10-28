package acl

//BaseACL : Setter Getter or Name and Description Base Object
type BaseACL struct {
	Name string `json:"name"` // Name
	Desc string `json:"desc"` // Description
}

//String : toString to Name And Description
func (base *BaseACL) String() string {
	return "Name: " + base.GetName() + " , Desc: " + base.GetDesc()
}

//GetName : Getter Name
func (base *BaseACL) GetName() string {
	return base.Name
}

//SetName : Setter Name
func (base *BaseACL) SetName(name string) {
	base.Name = name
}

//GetDesc : Getter Description
func (base *BaseACL) GetDesc() string {
	return base.Desc
}

//SetDesc : Setter Description
func (base *BaseACL) SetDesc(desc string) {
	base.Desc = desc
}

//Verify : Check Object name is not empty etc..
func (base *BaseACL) Verify() bool {
	return base.GetName() != ""
}
