package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

var srcDir = flag.String("d", "./", "存放函数的文件夹")

func main() {
	flag.Parse()

	funcPaths, _ := filepath.Glob(*srcDir + "*")

	for _, funcPath := range funcPaths {
		newFuncPath := fileRename(funcPath)
		fmt.Println(funcPath, newFuncPath)
		if funcPath == newFuncPath {
			continue
		}

		cmd := exec.Command("mv", funcPath, newFuncPath)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 将驼峰的文件名冲命名
func fileRename(s string) (new string) {
	new = ""
	for i, r := range s {
		if 'A' <= r && r <= 'Z' {
			if i != 0 {
				new += "_"
			}
			r += 'a' - 'A'
		}
		new += string(r)
	}
	return new
}
