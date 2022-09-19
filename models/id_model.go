package models
type IdModel struct{
	Id int64 `gorm:"id;primary_key" orm:"id" json:"id" form:"id" xml:"id"  bson:"id" `
	Db string `gorm:"db;default:''" orm:"db;" json:"db" form:"db" xml:"db"  bson:"db" `
	Dbs []string `gorm:"-" orm:"-" json:"db" form:"db" xml:"db"  bson:"-" `
	Table string `gorm:"table;default:''"  json:"table" form:"table" xml:"table"  bson:"table" `
	Tables []string `gorm:"-" orm:"-"  json:"table" form:"table" xml:"table"  bson:"-"`
	Value string `gorm:"value;default:''" json:"value" form:"value" xml:"value"  bson:"value" `
	MinId int64 `gorm:"min_id" json:"min_id" form:"min_id" xml:"min_id"  bson:"min_id" `
	MaxId int64 `gorm:"max_id" json:"max_id" form:"max_id" xml:"max_id"  bson:"max_id" `

	Created string `orm:"column(created)" gorm:"created" json:"created" form:"-"  xml:"-"  bson:"created"  `

	Updated string `orm:"column(updated)" gorm:"updated" json:"updated" form:"-"  xml:"-"  bson:"updated" `
}
func (*IdModel)GetIdProName() string{
	return  "Id"
}
func (obj *IdModel)GetId() interface{}{
	return  obj.Id
}
func (obj *IdModel)GetIdColName() string{
	return  "id"
}
func (*IdModel)GetDescription() string{
	return  "ids"
}
func (*IdModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*IdModel)GetTable() string{
	return  "ids"
}

func (*IdModel)GetCollection() string{
	return  "ids"
}

func (*IdModel)GetDoc() string{
	return  "ids"
}

func (*IdModel)GetIdName() string{
	return  "id"
}