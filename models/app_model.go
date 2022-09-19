package models

type AppBottom struct {
	Id int64
	Name string
	FlutterIcon string
	AndroidIcon string
	IosIcon string
	//index person_center ...
	Flag string
}

type AppPage struct {
	Id int64
}
type AppTab struct {
	Id int64
	Catagories []AppTabCatagory
}
type AppTabCatagory struct {
	Id int64
	Name string
	FlutterIcon string
	AndroidIcon string
	IosIcon string
	//catagory
	Flag string
}