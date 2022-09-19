package models

import (
	"encoding/xml"
	"reflect"
)

type Root struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Classes []ClassTpl `json:"classes" xml:"classes>class"`
}
type ClassTpl struct {
	XMLName xml.Name `json:"-" xml:"class"`
	Cur bool `json:"cur" xml:"cur,attr"`
	Curd bool `json:"curd" xml:"curd,attr"`
	C bool `json:"c" xml:"c,attr"`
	Cu bool `json:"cu" xml:"cu,attr"`
	U bool `json:"u" xml:"u,attr"`
	R bool `json:"r" xml:"r,attr"`
	D bool `json:"d" xml:"d,attr"`
	Fk bool `json:"fk" xml:"fk,attr"`
	Fs bool `json:"fs" xml:"fs,attr"`
	Skip string `json:"skip" xml:"skip,attr"`

	Name string `json:"name" xml:"name,attr"`
	Key string `json:"key" xml:"key,attr"`
	Refkey string `json:"refkey" xml:"refkey,attr"`
	Isinterface bool `json:"isinterface" xml:"isinterface,attr"`
	Interfaces string `json:"interfaces" xml:"interfaces,attr"`
	Isenum bool `json:"isenum" xml:"isenum,attr"`
	Values string `json:"values" xml:"values,attr"`
	Class1 string `json:"class1" xml:"class1,attr"`
	Table string `json:"table" xml:"table,attr"`
	Group string `json:"group" xml:"group,attr"`
	Comment string `json:"comment" xml:"comment,attr"`
	Fields []FieldTpl `json:"fields" xml:"fields>field"`
}
type FieldTpl struct {
	XMLName xml.Name `json:"-" xml:"field"`
	Id bool `json:"id" xml:"id,attr"`
	Name string `json:"name" xml:"name,attr"`
	Type string `json:"type" xml:"type,attr"`
	Comment string `json:"comment" xml:"comment,attr"`
	Length int `json:"length" xml:"length,attr"`
	Validate string `json:"validate" xml:"validate,attr"`
	End bool `json:"end" xml:"end,attr"`
	Remark string `json:"remark" xml:"remark,attr"`
}

func ParseClassTpl(types ...reflect.Type)*Root  {

	if len(types)>0{
		classes:=make([]ClassTpl,len(types))
		for i,v := range types {
			classes[i]=ClassTpl{
				Class1: v.Name(),
				Table: TemplateInstance.ParseString(v.Name()),
			}
			classes[i].Fields=make([]FieldTpl,v.NumField())
			for j := 0; j < v.NumField(); j++ {
				f:=v.Field(j)
				classes[i].Fields[j]=FieldTpl{
					Name: TemplateInstance.ParseString(f.Name),
					Type: f.Type.Name(),
				}
			}
		}
		return  &Root{Classes: classes}
	}
	return  nil
}