package main

import (
	"flag"
	"fmt"
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
	funcNames, _ := filepath.Glob(*srcDir + "*.go")
	for _, funcPath := range funcNames {
		split := strings.Split(funcPath, "/")
		funcFileName := split[len(split)-1]
		funcName := substr(funcFileName, 0, len(funcFileName)-3)
		fmt.Println(funcName)
		cmd := exec.Command("faas new", "--lang", "golang-middleware", "--handler", *tarDir+"handlers", funcName)
		fmt.Println(cmd.String())
	}

}

func substr(str string, idx, length int) string {
	split := strings.Split(str, "")
	return strings.Join(split[idx:idx+length], "")
}
