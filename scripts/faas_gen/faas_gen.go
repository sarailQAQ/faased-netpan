package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

var srcDir = flag.String("d", "./", "存放函数的文件夹")
var tarDir = flag.String("t", "", "生成函数的目录")

func main() {
	flag.Parse()
	if !strings.HasSuffix(*srcDir, "/") {
		*srcDir = *srcDir + "/"
	}
	if !strings.HasSuffix(*tarDir, "/") {
		*tarDir = *tarDir + "/"
	}

	// 创建 handlers 文件夹
	err := exec.Command("mkdir", *tarDir+"handlers").Run()
	if err != nil {
		log.Fatal(err)
	}

	// 根据文件名生成创建对应的 openfaas 函数
	funcPaths, _ := filepath.Glob(*srcDir + "*.go")
	var funcNames []string
	for i, funcPath := range funcPaths {
		split := strings.Split(funcPath, "/")
		funcFileName := split[len(split)-1]
		funcName := substr(funcFileName, 0, len(funcFileName)-3)
		fmt.Println(i, funcName)
		funcName = strings.ReplaceAll(funcName, "_", "-")
		cmd := exec.Command("faas", "new", "--lang", "golang-middleware", "--handler", *tarDir+"handlers/"+funcName+"-handler", funcName)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("combined out:\n%s\n", string(out))
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		funcNames = append(funcNames, funcName)
	}

	for _, funcName := range funcNames {
		path := funcName + ".yml"
		cmd := exec.Command("mv", path, *tarDir+path)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	//for _, funcPath := range funcPaths {
	//	newFuncPath := fileRename(funcPath)
	//	fmt.Println(funcPath, newFuncPath)
	//	if funcPath == newFuncPath {
	//		continue
	//	}
	//
	//	cmd := exec.Command("mv", funcPath, newFuncPath)
	//	err := cmd.Run()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}

func substr(str string, idx, length int) string {
	split := strings.Split(str, "")
	return strings.Join(split[idx:idx+length], "")
}

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
