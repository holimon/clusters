// +build windows

package service

import (
	"clusters/pkg/sysinfo"
	"io/ioutil"
	"path/filepath"
)

func PathList(dirname string) (list []StPathList,rerr error)  {
	if dirname == "/"{
		if par,err := sysinfo.GetPartion();err==nil{
			for _,item := range par{
				temp:=StPathList{}
				temp.Isdir = true
				temp.Title = item
				temp.Path = item
				list = append(list,temp)
			}
		}else{
			rerr = err
			return
		}
	}else {
		dirname,_ = filepath.Abs(dirname)
		if files,err := ioutil.ReadDir(dirname);err == nil{
			for _,item := range files{
				temp:=StPathList{}
				if item.IsDir(){
					temp.Isdir = true
				}else{
					temp.Isdir = false
				}
				temp.Title = item.Name()
				temp.Path = filepath.Join(dirname,item.Name())
				list = append(list,temp)
			}
		}else{
			rerr = err
		}
	}
	return
}