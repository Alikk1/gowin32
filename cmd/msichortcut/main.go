package main

import (
	"fmt"
	"github.com/Alikk1/gowin32"
)

func main() {
	target := gowin32.GetShortcutTargetPath("C:\\Users\\Public\\Desktop\\MobaXterm.lnk")
	fmt.Println(target)
}
