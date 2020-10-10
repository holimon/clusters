package service

//go:generate go run github.com/GeertJohan/go.rice/rice embed-go

import (
	"clusters/config"
	rice "github.com/GeertJohan/go.rice"
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
)

func Service() {
	e := echo.New()
	//静态资源
	assetHandler := http.FileServer(rice.MustFindBox("../../dist").HTTPBox())
	e.GET("/*", echo.WrapHandler(assetHandler))
	//健康检查
	e.GET("/api/v1/check", HandlerCheck)
	//节点发现
	e.GET("/api/v1/explorer/nodes", HandlerNodes)
	//文件管理
	e.GET("/api/v1/explorer/dirlist", HandlerPathList)
	e.GET("/api/v1/explorer/rename", HandlerRenamePath)
	e.GET("/api/v1/explorer/remove", HandlerRemovePath)
	e.GET("/api/v1/explorer/newfile", HandlerNewFile)
	e.GET("/api/v1/explorer/newfolder", HandlerNewFolder)
	e.GET("/api/v1/explorer/download", HandlerDownload)
	e.POST("/api/v1/explorer/upload", nil)
	//进程管理
	e.GET("/api/v1/process/list", HandlerProcesslist)
	e.GET("/api/v1/process/kill", HandlerProcessKill)
	e.GET("/api/v1/process/suspend", HandlerProcessSuspend)
	e.GET("/api/v1/process/resume", HandlerProcessResume)
	//系统信息
	e.GET("/api/v1/system/brief", HandlerHostBrief)
	e.GET("/api/v1/system/cpuuse", HandlerCpuuse)
	e.GET("/api/v1/system/diskuse", HandlerDiskuse)
	e.GET("/api/v1/system/memuse", HandlerMemuse)
	e.GET("/api/v1/system/net/brief", nil)
	e.GET("/api/v1/system/net/speed", nil)
	//服务启动
	panic(e.Start(fmt.Sprintf(":%d", config.Configs.Port)))
}
