package dtos

type RoleOutPut struct {


	//角色Id
	Id int64 ` json:"id" form:"id"  xml:"id" `

	//角色名称
	Name string ` json:"name" form:"name"  xml:"name" `

	ParentId int64 ` json:"-" form:"-"  xml:"-" `


	//角色子集
	Children []*RoleOutPut ` gorm:"-" json:"children" form:"children"  xml:"children" `
}
