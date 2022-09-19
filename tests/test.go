package tests

import (
	"os"
	"testing"
)

func _TestRegister2(t *testing.T) {
	t.Log("test")
	d,err:=os.Getwd()
	if err!=nil{
		t.Log("test get dir fail")
		return
	}
	t.Logf("test get dir suc,dir:%s",d)
}
