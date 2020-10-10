package service

import (
	"clusters/pkg/consul"
	"errors"
	"os"
	"path/filepath"
)

func RemovePath(absname string)(rerr error)  {
	return os.RemoveAll(absname)
}

func Rename(oldpath,newpath string)(rerr error)  {
	dir := filepath.Dir(newpath)
	if _,err := os.Stat(dir);err != nil || !os.IsExist(err){
		if rerr = os.MkdirAll(dir,os.ModePerm);rerr == nil{
			rerr = os.Rename(oldpath,newpath)
		}
	}else{
		rerr = os.Rename(oldpath,newpath)
	}
	return
}

func NewFolder(dirname string)(rerr error)  {
	rerr = os.MkdirAll(dirname,os.ModePerm)
	return
}

func NewFile(filename string)(rerr error)  {
	if _,err := os.Stat(filename);err == nil || os.IsExist(err){
		rerr = errors.New("File is exist")
		return
	}
	if _,err := os.Stat(filepath.Dir(filename));err != nil || !os.IsExist(err){
		if err = os.MkdirAll(filepath.Dir(filename),os.ModePerm);err != nil{
			rerr = err
			return
		}
	}
	if f,err := os.Create(filename);err==nil{
		defer f.Close()
	}else{
		rerr = err
	}
	return
}

func Discovery()(nodes []map[string]string,rerr error)  {
	nodes,rerr = consul.ServiceDiscovery()
	return
}