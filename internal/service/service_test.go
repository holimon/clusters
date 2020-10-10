package service

import (
	"fmt"
	"testing"
)

func TestPathList(t *testing.T) {
	fmt.Println(PathList("/"))
}

func TestRemovePath(t *testing.T) {
	RemovePath("D:\\tmp\\consul-test")
}

func TestRename(t *testing.T) {
	Rename("D:\\tmp\\test.ima","D:\\tmp\\consul\\test.ima")
}