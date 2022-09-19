package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type ToolHelper struct {
	system string
}

func (t *ToolHelper) Init()  {
	arch:=runtime.GOARCH
	os:=runtime.GOOS
	t.system=os
	log.Printf("arch:%s,os:%s",arch,os)
}
func (t *ToolHelper) Exec(cmd string)(string,error) {
	os:=runtime.GOOS
	commandName:="cmd"
	commandPath:="/c"
	switch os {
	case "windows":
		break
	case "linux":
		commandName="bash"
		commandPath="-c"
		break
	}
	cmd1 := exec.Command(commandName,commandPath, cmd)
	output,err := cmd1.Output()
	if err != nil {
		log.Println("error:"+err.Error())
		return "",err
	}
	res:=string(output)
	log.Printf("exec cmd :"+cmd)
	log.Printf("exec cmd result:")
	log.Println(res)
	return  res,nil
}
func (t *ToolHelper) FindPortToPId(port int)int {
	os:=runtime.GOOS
	cmdStr:=""
	switch os {
		case "windows":
			cmdStr=fmt.Sprintf("netstat -ano -p tcp | findstr %d",port)
			break
		case "linux":
			cmdStr=fmt.Sprintf("lsof  -i:%d",port)
			break
	}
	res,err :=t.Exec(cmdStr)
	if err != nil {
		return -1
	}
	r:=regexp.MustCompile(`\s\d+\s`).FindAllString(res,-1)
	if len(r)>0{
		pId,err:=strconv.Atoi(strings.TrimSpace(r[0]))
		if err!=nil{
			return -1
		}
		return pId
	}
	return -1

}
func (t *ToolHelper) QueryAllPort() ([]int,error) {
	os:=runtime.GOOS
	cmdStr:=""
	switch os {
	case "windows":
		cmdStr="netstat -a"
		break
	case "linux":
		break
	}
	res,err :=t.Exec(cmdStr)
	if err != nil {
		return nil,err
	}
	log.Println(res)
	r:=regexp.MustCompile(`\\d+:\s\d+\s`).FindAllString(res,-1)
	if len(r)>0{
		ps:=make([]int,len(r))
		for i, v := range r {
			p,err:=strconv.Atoi(strings.Split(strings.TrimSpace(v),":")[0])
			if err!=nil{
				log.Println("parse string -> int fail,string => "+v)
				continue
			}
			ps[i]=p
		}
		return ps,err
	}
	return nil,err
}

func (t *ToolHelper) SkillPocess(pid int)(string,error)  {
	cmdStr:="taskkill /t /f /im "+strconv.Itoa(pid) //windows
	res,err :=t.Exec(cmdStr)
	return res, err
}

func (t *ToolHelper) OperatorService(name string,opr string)(string,error) {
	os:=runtime.GOOS
	cmdStr:=""
	switch os {
	case "windows":
		cmdStr=fmt.Sprintf("net %s  %s",opr,name)
		break
	case "linux":
		cmdStr=fmt.Sprintf("systemctl %s  %s",opr,name)
		break
	}
	res,err :=t.Exec(cmdStr)
	return res, err

}
func (t *ToolHelper) StartService(name string)(string,error) {
	return  t.OperatorService(name,"start")
}
func (t *ToolHelper) StopService(name string)(string,error) {
	return  t.OperatorService(name,"stop")
}
func (t *ToolHelper) ReStartService(name string)(string,error) {
	os:=runtime.GOOS
	if os=="windows"{
		str,err:=  t.OperatorService(name,"stop")
		if err!=nil{
			return str, err
		}
		str1,err:=  t.OperatorService(name,"start")
		return str+"\r\n"+str1, err
	}
	return  t.OperatorService(name,"restart")
}
func (t *ToolHelper) DeleteService(name string)(string,error) {
	os:=runtime.GOOS
	if os=="windows"{
		return  t.OperatorService(name,"sc")
	}
	return "", nil
}

func (t *ToolHelper) ExecCmd(cmdPath string)(*os.Process,error) {
	return os.StartProcess(cmdPath,[]string{},&os.ProcAttr{})
}