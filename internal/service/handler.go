package service

import (
	"clusters/pkg/sysinfo"
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
	"strconv"
)

func HandlerCheck(c echo.Context) error {
	return c.String(http.StatusOK, "health")
}

func HandlerNodes(c echo.Context) error  {
	if nodes,err := Discovery();err==nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"data":nodes,"message":"Success"})
	}else {
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{"message":"An error occurred while discovering"})
	}
}

func HandlerPathList(c echo.Context) error {
	dirname := c.QueryParam("abspath")
	if list,err := PathList(dirname);err==nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"data":list,"message":"Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{"message":"An error occurred while reading the folder"})
	}
}

func HandlerRenamePath(c echo.Context) error {
	oldpath := c.QueryParam("oldpath")
	newpath := c.QueryParam("newpath")
	if err := Rename(oldpath, newpath); err == nil {
		return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
	} else {
		return c.JSON(http.StatusExpectationFailed, map[string]string{"message": "An error occurred during the rename process"})
	}
}

func HandlerRemovePath(c echo.Context) error {
	abspath := c.QueryParam("abspath")
	if err := RemovePath(abspath); err == nil {
		return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
	} else {
		return c.JSON(http.StatusExpectationFailed, map[string]string{"message": "An error occurred during deletion"})
	}
}

func HandlerNewFile(c echo.Context) error  {
	abspath := c.QueryParam("abspath")
	if err := NewFile(abspath);err == nil{
		return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]string{"message": "An error occurred during create file"})
	}
}

func HandlerNewFolder(c echo.Context) error  {
	abspath := c.QueryParam("abspath")
	if err := NewFolder(abspath);err == nil{
		return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]string{"message": "An error occurred during create folder"})
	}
}

func HandlerHostBrief(c echo.Context) error {
	if info,err:=sysinfo.GetHostBrief();err==nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"data":info,"message":"Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]string{"message": "An error occurred"})
	}
}

func HandlerCpuuse(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data":sysinfo.GetCpuPercent(),"message":"Success"})
}

func HandlerMemuse(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data":sysinfo.GetMemPercent(),"message":"Success"})
}

func HandlerDiskuse(c echo.Context) error {
	if disks,err := sysinfo.GetDiskPercent();err==nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"data":disks,"message":"Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{"message":"An error occurred"})
	}
}

func HandlerProcesslist(c echo.Context) error {
	if processes,err := sysinfo.GetProcessList();err == nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"data":processes,"message":"Success"})
	}else {
		return c.JSON(http.StatusExpectationFailed,map[string]interface{}{"message":"An error occurred"})
	}
}

func HandlerProcessKill(c echo.Context) error  {
	pid,_ := strconv.ParseInt(c.QueryParam("pid"),10,32)
	if err := sysinfo.ProcessKill(int32(pid));err == nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"message":"Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{"message":"An error occurred during kill the process"})
	}
}

func HandlerProcessSuspend(c echo.Context) error  {
	pid,_ := strconv.ParseInt(c.QueryParam("pid"),10,32)
	if err := sysinfo.ProcessSuspend(int32(pid));err == nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"message":"Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{"message":"An error occurred during suspend the process"})
	}
}

func HandlerProcessResume(c echo.Context) error  {
	pid,_ := strconv.ParseInt(c.QueryParam("pid"),10,32)
	if err := sysinfo.ProcessResume(int32(pid));err == nil{
		return c.JSON(http.StatusOK, map[string]interface{}{"message":"Success"})
	}else{
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{"message":"An error occurred during resume the process"})
	}
}

func HandlerDownload(c echo.Context) error  {
	file := c.QueryParam("abspath")
	c.Response().Header().Set(echo.HeaderContentDisposition,"attachment; filename="+filepath.Base(file))
	return c.File(file)
}