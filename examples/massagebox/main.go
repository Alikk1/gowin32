package main

import (
	"github.com/Alikk1/gowin32/wrappers"
)

func main() {
	err := wrappers.MessageBox("hello world", "你好")
	if err != nil {
		panic(err)
	}

	//gowin32.GetWindowText()
}
