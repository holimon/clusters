// +build linux

package service

import (
	"io/ioutil"
	"path/filepath"
)

func PathList(dirname string) (list []StPathList, rerr error)  {
	if dirname == "/"{
		temp:=StPathList{}
		temp.Path = "/"
		temp.Title = "/"
		temp.Isdir = true
		list = append(list,temp)
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