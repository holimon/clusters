package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetAppDir() (appdir string) {
	App, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(App)
	appdir, _ = filepath.Split(AppPath)
	return
}