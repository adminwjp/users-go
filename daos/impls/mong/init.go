package dao_mong_impl


func intResult(err error)(int,error)  {
	if err!=nil{
		return  0,err
	}
	return 1,err
}


