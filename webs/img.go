package web

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

type ImgCtrl struct {
	
}
func (ctrl *ImgCtrl) Get(img string,request *http.Request,response http.ResponseWriter){

	fileName := "static/imgs/"+img
	file, err := os.Open(fileName)
	if err != nil {
		io.WriteString(response, "")
		return
	}
	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	fileStat, err := file.Stat()
	if err != nil {
		io.WriteString(response, "")
		return
	}
	response.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	response.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	response.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))
	file.Seek(0, 0)
	io.Copy(response, file)
}

