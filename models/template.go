package models

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

type LangeFalg int
//https://www.nhooo.com/tool/json2go/
//https://www.qetool.com/sql_json_go/json.html
const(
	LangeNone LangeFalg=iota
	LangeChines
	LangeEnglish
)

/**
{
"id":0,
"name":"",
"langes":[
	{
"id":0,
"name":"",
"flag":0
	}
]
}
*/

type Remark struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Names []string ` orm:"-" json:"names"`
	Langes []RemarkLange `orm:"-" json:"langes,omitempty"`
}
type RemarkLange struct {
	Id int64 `json:"id,omitempty"`
	RemarkId int64 `json:"remark_id,omitempty"`
	Name string `json:"name,omitempty"`
	Names []string ` orm:"-" json:"names"`
	Flag LangeFalg `json:"flag,omitempty"`
}
type TableBean struct {
	Id int64 `json:"id,omitempty"`
	Xml bool `json:"xml,omitempty"`
	BeeOrm bool `json:"bee_orm,omitempty"`
	Gorm bool `json:"gorm,omitempty"`
	Names []string `  json:"names"`
	Json bool `json:"json,omitempty"`
	Form bool `json:"form,omitempty"`
	Bson bool `json:"bson,omitempty"` //mong
	Table   string        `json:"table,omitempty"`
	Comment string        `json:"comment,omitempty"`
	Key     string        `json:"key,omitempty"`
	RefKey  string        `json:"ref_key,omitempty"`
	Class   string        `json:"class,omitempty"`
	TempClass   string        `json:"temp_class,omitempty"`

	Columns []*ColumnBean `json:"columns,omitempty"`
}
type ColumnBean struct {
	Id int64 `json:"id,omitempty"`
	DefaultColumn  string `json:"default_column,omitempty"`
	Column  string `json:"column,omitempty"`
	Xml string `json:"xml,omitempty"`
	Json string `json:"json,omitempty"`
	Form string `json:"form,omitempty"`
	Bson string `json:"bson,omitempty"` //mong
	ProName string `json:"pro_name,omitempty"`
	IsPro bool `json:"is_pro,omitempty"`
	Type    string `json:"type,omitempty"`
	Comment string `json:"comment,omitempty"`
	Length  int    `json:"length,omitempty"`
	Default string `json:"default,omitempty"`
	Names []string `  json:"names"`
}
type TemplateHelper struct {
	Remarks *[]Remark
	Tabs []*TableBean
	NewTabs []*TableBean
	File bool
	One bool
	Strs []string
	Str string
	ClassPrefixes []string
	ClassSuffixes []string
	Lange LangeFalg
	Dir string
	user *user
}
type user struct {
	classes []string
	//login register
	/*
	{"account":1,"pwd":1}
	{"account":1,"password":1}
	*/
	accountes []map[string]bool
	userNames []map[string]bool
	phones []map[string]bool
	emails []map[string]bool
	//exists
	existsAccountes map[string]bool
	existsUserNames map[string]bool
	existsPhones map[string]bool
	existsEmails map[string]bool
	//update
	account map[string]bool
	userName map[string]bool
	phone map[string]bool
	email map[string]bool
}
var TemplateInstance =&TemplateHelper{}

func (helper *TemplateHelper)generaotrUser(tabs []*TableBean)  {
	users:=make([]*TableBean,0)
	for _, tab := range tabs {
		for _, user := range helper.user.classes {
			if strings.Index(tab.Table,user)>-1{
				users=append(users,tab)
				break
			}
		}
	}

}
func (helper *TemplateHelper) UpdateInfrastructureGo(dir string)error{
	dirs,err:=os.ReadDir(dir)
	if err!=nil{
		return err
	}
	for _, v := range dirs {
		log.Println("dir name or file name:"+v.Name())
		if v.IsDir(){

		}else{

		}
	}
	return nil
}
func (helper *TemplateHelper) Init(){
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("GetCurrentDir,err:%s", err.Error())
		return
	}
	log.Printf("GetCurrentDir,dir:%s", dir)
	helper.Dir=dir
	helper.Lange= LangeChines
	//helper.ClassPrefixes=make([]string,0)
	//helper.ClassPrefixes=append(helper.ClassPrefixes,"LiveBroadcast")
	//helper.ClassPrefixes=append(helper.ClassPrefixes,"ShortVideo")

	helper.ClassPrefixes=make([]string,2)
	helper.ClassPrefixes[0]="LiveBroadcast"
	helper.ClassPrefixes[1]="ShortVideo"

	helper.ClassSuffixes=make([]string,4)
	helper.ClassSuffixes[0]="Model"
	helper.ClassSuffixes[1]="Bean"
	helper.ClassSuffixes[2]="Entity"
	helper.ClassSuffixes[3]="Entry"

	cfg := dir + "/conf/remark.json"
	cfg = strings.ReplaceAll(cfg, "\\", "/")
	log.Printf(cfg)
	buffer, err := os.ReadFile(cfg)
	if err != nil {
		log.Printf("read remark file %s,err:%s", cfg, err.Error())
		return
	}
	err=json.Unmarshal(buffer,&helper.Remarks)
	if err != nil {
		log.Printf("bind remark file %s,err:%s", cfg, err.Error())
		return
	}
	//helper.updateRemark()
}
/**
{
	"class":"",
	"columns":[
		"pro_name":"",
		"type":""
	]
}
*/
func (helper *TemplateHelper)ParseClass2(tabs []*TableBean){
	for i := 0; i < len(tabs); i++ {
		helper.UpdateTable(tabs[i])
		helper.UpdateClass(tabs[i])
		helper.UpdateColumn(tabs[i])
	}
	//helper.NewTabs=tabs
	helper.UpdateRemark(tabs)
}
func (helper *TemplateHelper)ParseClass(tabs []TableBean){
	helper.Tabs=make([]*TableBean,len(tabs))
	for i := 0; i < len(tabs); i++ {
		helper.Tabs[i]=&tabs[i]
	}
	helper.ParseClass2(helper.Tabs)
}
//names:a_b_c_d a b c d
//name count:5 com names count:5+2+3+2=12
//com names:a_b_c_d a b c d
//a_b a_b_c
//b_c b_c_d
//c_d
func (helper *TemplateHelper)UpdateRemark(tabs []*TableBean){
	helper.updateRemark()
	l:=len(tabs)
	for i := 0; i < l; i++ {
		t:=tabs[i]
		if t.Comment==""||t.Comment==t.Class{
			names:=helper.updateRemarkNames(t.Table)
			names=helper.getComRemark(names)
			t.Names=names
			r:=helper.getRemark(names,*helper.Remarks)
			if r!=""{
				t.Comment=r
			}
		}

		n:=len(t.Columns)
		for j := 0; j < n; j++ {
			c:=t.Columns[j]
			if c.Comment!=""&&c.Comment!=c.DefaultColumn{
				continue
			}
			names:=helper.updateRemarkNames(c.Column)

			names=helper.getComRemark(names)
			c.Names=names
			r:=helper.getRemark(names,*helper.Remarks)
			if r!=""{
				c.Comment=r
			}
		}
	}
}

//user_name user name
func (helper *TemplateHelper)updateRemarkNames(name string)[]string{
	names:=strings.Split(name,"_")
	l:=len(names)
	if l==1{
		return  names
	}

	temps:=make([]string,l+1)
	temps[0]=name
	for i := 1; i <= l; i++ {
		temps[i]=names[i-1]
	}
	log.Println("name:=>%s,names:%q,update names:%q",
		name,names,temps)
	return temps
}

//names:a_b_c_d a b c d
//name count:5 com names count:5+2+3+2=12
//com names:a_b_c_d a b c d
//a_b a_b_c
//b_c b_c_d
//c_d
func (helper *TemplateHelper)getComRemark(names []string)[]string{
	if len(names)==1{
		return  names
	}
	var temps=make([]string,len(names))
	temps[0]=names[0]
	i := 1
	temps[len(names)-1]=names[len(names)-1]
	for ; i <len(names)-1 ; i++ {
		temps[i]=names[i]
		na:=names[i]
		for j := i+1; j < len(names); j++ {
			if i==1&&j==len(names)-1{
				break
			}
			na+="_"+names[j]
			temps=append(temps,na)
		}
	}
	return  temps
}

//user_name user name
func (helper *TemplateHelper)getRemark(names []string,remarks []Remark)string{
	//short_video_name short video name
	//short_video_name short_video name
	ns:=make([]string,len(names)-1)
	//vs:=make([]string,len(names)-1)
	ms:=make(map[string]string,len(names)-1)
	for i := 0; i < len(names); i++ {
		n:=names[i]


		for j := 0; j <len(remarks) ; j++ {
			r:=remarks[j]
			na:=helper.getLangeRemarkByName(r,n)
			if na!=""{
				if i==0{
					return  na
				}
				//ns[i-1]=na
				ms[n]=na
				//break
			}
		}
	}
	if len(ms)==0 {
		return ""
	}
	//hand edit comment
	//pass
	//短视频关注
	//short_video_follow short video
	//follow short_video short_video_follow video_follow

	//eror
	//"table": "short_video_follow",
	//"comment": "视频关注短视频",
	msl:=make(map[string]int,len(names)-1)
	ls:=make([]int,len(names)-1)
	ll:=0
	for k, _ := range ms {
		ns[ll]=k
		ls[ll]=len(k)
		msl[k]=len(k)
		ll++
	}
	//name split sort question
	sort.Ints(ls) //length sort
	//string length sort
	for i := 0; i <ll ; i++ {
		for k, v := range msl {
			if v==ls[i]{
				for j := 0; j < ll; j++ {
					if ns[j]==k{
						temp:=ns[i]
						ns[i]=ns[j]
						ns[j]=temp
						break
					}
				}
			}
		}

	}
	//delete short string(equal)
	for i := ll-2; i >=0 ; i-- {
		for j := i-1; j >= 0; j-- {
			if strings.Index(ns[i],ns[j])>-1{
				msl[ns[i]]=0
				break
			}
		}
	}

	str:=""
	log.Println("names:")
	log.Println(ns)
	log.Println("map name:lange=>")
	log.Println(ms)
	log.Println("map name:length=>")
	log.Println(msl)
	for k, v := range msl {
		if v!=0{
			str+=ms[k]
		}
	}

	//sort
	/*str=ns[0]
	for i := 1; i < len(ns); i++ {
		str+=ns[i]
	}*/
	return str
}



func (helper *TemplateHelper)getLangeRemarkByName(remark Remark,name string)string{
	if remark.Names!=nil{
		for k := 0; k <len(remark.Names) ; k++ {
			na:=remark.Names[k]
			if na==name{
				r :=  helper.getLangeRemark(remark)
				if r!=""{
					return  r
				}
			}
		}
	}else{
		if remark.Name==name{
			return  helper.getLangeRemark(remark)
		}
	}
	return ""
}

func (helper *TemplateHelper)getLangeRemark(remark Remark)string{
	for l1 := 0; l1 <len(remark.Langes) ; l1++ {
		l:=remark.Langes[l1]
		if l.Flag==helper.Lange{
			return l.Name
		}
	}
	return ""
}

func (helper *TemplateHelper)updateRemark(){
	for i := 0; i <len(*helper.Remarks) ; i++ {
		r:=(*helper.Remarks)[i]
		if strings.Index(r.Name,",")>-1{
			r.Names=strings.Split(r.Name,",")
		}
		for j := 0; j < len(r.Langes); j++ {
			ll:=r.Langes[j]
			if strings.Index(ll.Name,",")>-1{
				ll.Names=strings.Split(ll.Name,",")
			}
		}
	}
}

func (helper *TemplateHelper) Parse(types ...reflect.Type )  {
	l:=len(types)
	helper.NewTabs=make([]*TableBean,l)
	for i := 0; i < l; i++ {
		t:=types[i]
		n:=t.NumField()
		helper.NewTabs[i]=&TableBean{
			Class:t.Name(),Columns: make([]*ColumnBean,n),
		}
		helper.UpdateTable(helper.NewTabs[i])
		helper.UpdateClass(helper.NewTabs[i])

		for j := 0; j < n; j++ {
			f:=t.Field(j)
			helper.NewTabs[i].Columns[j]=&ColumnBean{
				ProName: f.Name,Type: f.Type.Name(),
			}
		}
		helper.UpdateColumn(helper.NewTabs[i])
	}
}

func (helper *TemplateHelper)UpdateTable(tab *TableBean)  {
	if tab.Class==""{
		if tab.Table==""{
			return
		}
		tab.Class=helper.ParseString(tab.Table)
	}
	if tab.Comment==""{
		tab.Comment=tab.Class
	}
	tab.TempClass=tab.Class
	for i := 0; i <len(helper.ClassPrefixes) ; i++ {
		p:=helper.ClassPrefixes[i]
		if strings.Index(tab.TempClass,p)==0{
			tab.TempClass=strings.Replace(tab.TempClass,p,"",-1)
		}
	}
	for i := 0; i <len(helper.ClassSuffixes) ; i++ {
		p:=helper.ClassSuffixes[i]
		if strings.Index(tab.TempClass,p)==len(helper.ClassSuffixes)-len(p){
			tab.TempClass=strings.Replace(tab.TempClass,p,"",-1)
		}
	}
}

func (helper *TemplateHelper)UpdateClass(tab *TableBean)  {
	if tab.Table==""{
		if tab.Class==""{
			return
		}
		tab.Table=helper.ParseString(tab.Class)
	}
}

func (helper *TemplateHelper)UpdateColumn(tab *TableBean)  {
	for i := 0; i < len(tab.Columns); i++ {

		c:=tab.Columns[i]
		if c.DefaultColumn!=""{
			if c.Column==""{
				c.Column=c.DefaultColumn
			}
		}
		if c.Type==""{
			c.Type="string"
		}
		if c.ProName==""{
			if c.Column==""{
				continue
			}
			c.ProName=helper.ParseString(c.Column)
		}
		if c.Column==""{
			if c.ProName==""{
				continue
			}
			c.Column=helper.ParseString(c.ProName)
		}
		if c.DefaultColumn==""{
			c.DefaultColumn=c.Column
		}
		if c.Comment==""{
			c.Comment=c.DefaultColumn
		}
		if c.Xml==""{
			c.Xml=c.DefaultColumn
		}
		if c.Json==""{
			c.Json=c.DefaultColumn
		}
		if c.Form==""{
			c.Form=c.DefaultColumn
		}
		if c.Bson==""{
			c.Bson=c.DefaultColumn
		}
	}
}

func (helper *TemplateHelper) GeneratorBeanByJson(tpl ,jsonFile string,testOutput bool){
	dir:=helper.Dir
	cfg := dir + jsonFile
	cfg = strings.ReplaceAll(cfg, "\\", "/")
	log.Printf(cfg)
	buffer, err := os.ReadFile(cfg)
	if err != nil {
		log.Printf("read file %s,err:%s", cfg, err.Error())
		return
	}
	//test ``
	if testOutput {
		str := string(buffer)
		log.Printf(str)
	}
	var tabs []*TableBean
	err = json.Unmarshal(buffer, &tabs)
	if err != nil {
		log.Printf("read file %s,parse json fail,err:%s", cfg, err.Error())
		return
	}
	log.Printf("count:%d", len(tabs))
	helper.Tabs=tabs
	helper.GeneratorBean(tpl,testOutput)
}

func (helper *TemplateHelper) GeneratorBean(tpl string,testOutput bool) {
	tabs:=helper.Tabs
	newTabs := helper.UpdateGeneratorBean(tabs)
	helper.NewTabs=newTabs
	buffer, _:= json.Marshal(newTabs)
	cfg := strings.ReplaceAll(helper.Dir, "cfg", "") + "cfg/test.json"
	if testOutput{
		log.Println(string(buffer))
		cfg = strings.ReplaceAll(cfg, "\\", "/")
		os.WriteFile(cfg, buffer, fs.FileMode(fs.ModeType))

		cfg = strings.ReplaceAll(helper.Dir, "cfg", "") + "cfg"
		cfg = strings.ReplaceAll(cfg, "\\", "/")
		buffer, _ := os.ReadFile(cfg + "/mode.tpl")
		rstr := strings.ReplaceAll(string(buffer), "\"", "\\\"")
		rstr = strings.ReplaceAll(rstr, "\r", "\\r")
		rstr = strings.ReplaceAll(rstr, "\n", "\\n")
		helper.DeleteDir(cfg+"/const")
		con := "package consts \r\n \r\n const GoModel=\"" + rstr + "\""
		os.WriteFile(cfg+"/const/const.go", []byte(con), fs.FileMode(fs.ModeType))

	}

	for _, tab := range newTabs {
		tab.Class = tab.Class
	}

	cfg = strings.ReplaceAll(helper.Dir, "cfg", "") + "model"
	cfg = strings.ReplaceAll(cfg, "\\", "/")
	helper.DeleteDir(cfg)
	helper.GeneratorModel(newTabs, cfg,tpl)
}

func (helper *TemplateHelper)DeleteDir(dir string) {
	log.Println("current dir:"+dir)
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, d := range dirs {
		log.Println("dir:"+d.Name())
		if d.IsDir() {
			helper.DeleteDir(dir + "/" + d.Name())
		} else {
			os.Remove(dir + "/" + d.Name())
		}
	}
}

//`` define syanc error


func (helper *TemplateHelper)GeneratorModelByXmlJson(tabs []*TableBean, dir string) {

}

//ABC -> a_b_c
//a_b_c -> ABC
func(helper *TemplateHelper) ParseString(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i, c := range str {
		if '_' == c {
			continue
		}
		//65 90
		if 'A' <= c && 'Z' >= c {
			c += 'a' - 'A'
			if i > 0 {
				b.WriteByte('_')
				b.WriteByte(byte(c))
				continue
			}
		} else if 'a' <= c && c <= 'z' { //97 122
			if (i-1 > 0 && str[i-1] == '_') ||
				(i+1 < len(str) && str[i+1] == '_')||
				(i==0&&i+1 < len(str) &&'a' <= str[i+1] && str[i+1] <= 'z'){
				c -= 'a' - 'A'
			}
		}
		b.WriteByte(byte(c))
	}
	str1 := b.String()
	log.Println("string parse:"+str1)
	return str1
}

func(helper *TemplateHelper) UpdateGeneratorBean(tabs []*TableBean) []*TableBean {
	for _, tab := range tabs {

		for _, column := range tab.Columns {
			if column.Type==""{
				column.Type="string"
			}
		}
	}
	newTabs := make([]*TableBean, 1)
	dics := make(map[string]*TableBean, len(tabs))
	index := 0
	log.Println(len(newTabs))
	for _, tab := range tabs {
		if tab.Key != "" {
			dics[tab.Key] = tab
		} else {
			if index == 0 {
				log.Println("new table:"+tab.Table)
				newTabs[index] = tab
			} else {
				//log.Println("update tab")
				newTabs = append(newTabs, tab)
			}
			var temp *TableBean
			for  {
				if tab.RefKey == "" {
					break
				}
				temp = dics[tab.RefKey]

				/*for _, column := range temp.Columns {
					tab.Columns=append(tab.Columns,column)
				}*/
				cols := temp.Columns
				for _, column := range tab.Columns {
					//tab.Columns=append(temp.Columns,column)
					cols = append(cols, column)
				}
				tab.Columns = cols
			}

			index++
		}
	}

	return newTabs
}

func (helper *TemplateHelper) Execute(tpl *template.Template, wr io.Writer,data interface{})error{
	return 	tpl.Execute(wr, data)
}

func (helper *TemplateHelper) getOneTplData(tabs []*TableBean)interface{}{
	var temps = make(map[string]interface{}, 3)
	temps["package"] = "model"
	temps["a1"] = '`'
	temps["tabs"] = tabs
	return  temps
}

func (helper *TemplateHelper) getTplData(tab *TableBean)interface{}{
	var temps = make(map[string]interface{},3)
	temps["package"] = "model"
	temps["a1"] = '`'
	var tabsTemp = make([]*TableBean, 1)
	tabsTemp[0] = tab
	temps["tabs"] = tabsTemp
	return  temps
}

func (helper *TemplateHelper) ExecuteStringResult(tpl *template.Template,
	wr *bytes.Buffer,tabs []*TableBean)error{
	var err error
	if helper.One{
		var temps=helper.getOneTplData(tabs)
		err=tpl.Execute(wr, temps)
		helper.Str=wr.String()
		wr.Reset()
		return err
	}
	for i, tab := range tabs {
		var temps=helper.getTplData(tab)
		err=tpl.Execute(wr, temps)
		helper.Strs[i]=wr.String()
		wr.Reset()
	}
	return err
}

func (helper *TemplateHelper) ExecuteFileResult(tpl *template.Template,
	tabs []*TableBean,modelName string) error{
	var err error
	var file string = helper.Dir + "/"+modelName+".go"
	log.Println(file)
	var f *os.File
	if fi, err := os.Stat(file); err == nil {
		fi.Size()
		f, err = os.OpenFile(file, //os.O_APPEND|
			os.O_RDWR|os.O_CREATE|os.O_WRONLY, 0666) //打开文件
	} else {
		f, err = os.Create(file) //创建文件
	}
	if err != nil {
		log.Printf("create or get file err:%s", err.Error())
		return err
	}
	var temps=helper.getOneTplData(tabs)
	err=tpl.Execute(f, temps)
	f.Close()
	return err
}

func (helper *TemplateHelper)GeneratorModel(tabs []*TableBean, dir string,tpl string)error {
	t1 := template.New("model")
	log.Println(tpl)
	tModel, err := t1.Parse(tpl)
	if err != nil {
		log.Printf("parse template err:%s", err.Error())
		return nil
	}
	var bu *bytes.Buffer
	if !helper.File {
		bu = &bytes.Buffer{}
		bu.Grow(len(tpl))
	}
	if helper.One&& helper.File {
		log.Printf("parse template to %s","file")
		return helper.ExecuteFileResult(tModel, tabs, "mode")
	}else if !helper.File{
		log.Printf("parse template to %s","string")
		if !helper.One{
			helper.Strs=make([]string,len(tabs))
		}
		return helper.ExecuteStringResult(tModel,bu,tabs)
	}
	log.Printf("parse template to %s","file")
	for _, tab := range tabs {
		var tabsTemp = make([]*TableBean, 1)
		tabsTemp[0] = tab
		err=helper.ExecuteFileResult(tModel, tabsTemp, helper.ParseString(tab.Class))
	}
	return err
}

func (helper *TemplateHelper)GeneratorDao(tabs []*TableBean) {

}