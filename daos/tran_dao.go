package daos

var EmptyTranInstance =&EmptyTranDao{}
type EmptyTranDao struct {


}
func (dao *EmptyTranDao)Begin(){}
func (dao *EmptyTranDao)Commit(){}
func (dao *EmptyTranDao)Rollback(){}
type TranDao interface {
	Begin()
	Commit()
	Rollback()
}
