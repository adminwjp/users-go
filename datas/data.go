package datas

type DataRepository interface {
	Add(obj interface{})(int,error)
	AddBatch(length int,objs interface{})(int,error)

	Update(obj interface{})(int,error)
	UpdateBatch(length int,objs interface{})(int,error)

	Delete(id interface{},model interface{})(int,error)
	DeleteBatch(length int,id interface{},model interface{})(int,error)

	Get(id interface{},obj interface{})(error)

	One(obj interface{})error
	//db
	ExecuteSql(sql string,values ... interface{})
}
