package service

type CrudService interface {
	Add(obj interface{})(int,error)
	Update(obj interface{})(int,error)
	Delete(id interface{})(int,error)
	DeleteBatch(ids interface{})(int,error)
	List(list interface{})(int64,error)
	ListByPage(page int,size int,list interface{})(int64,error)
}
