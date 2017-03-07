package cpu

import (
	"io/ioutil"
	"fmt"
	"strings"
	"errors"
)

type CpuWatch struct {
	filecontent string
}


func (this *CpuWatch)Refresh()  {
	filename:="/proc/cpuinfo"

	content,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
		return
	}
	this.filecontent=string(content)
	fmt.Println(this.filecontent)
}

func (this *CpuWatch)ParamState(column string) (string,error) {
	this.Refresh()
	states:=strings.Split(this.filecontent,"\n")
	for _,state:=range states{
		param:=strings.Split(state,"\t")
		if strings.EqualFold(param[0],column){
			res:=param[1:]
			return strings.Join(res,"\t"),nil
		}
	}
	return "",errors.New("没有找到参数")
}

func (this *CpuWatch)ParamStateBatch(columns map[string]string) (map[string]string,error) {
	this.Refresh()
	states:=strings.Split(this.filecontent,"\n")
	for _,state:=range states{
		param:=strings.Split(state,"\t")
		for key,_:=range columns{
			if strings.EqualFold(param[0],key){
				res:=param[1:]
				columns[key]= strings.Join(res,"\t")
			}
		}
	}
	return columns,nil
}

func (this *CpuWatch)Start()  {

}

