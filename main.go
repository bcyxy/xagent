package main

import (
	"os"
	"path"
	"path/filepath"
	"xagent/common"
	"xagent/glbval"
	"xagent/shelltool"
)

var (
	gitCommitID    string
	buildTime      string
	buildGoVersion string
)

func main() {
	// 处理编译参数
	glbval.GitCommitID = gitCommitID
	glbval.BuildTime = buildTime
	glbval.BuildGoVersion = buildGoVersion

	// 程序root路径
	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	glbval.RootPath = path.Join(binDir, "..")

	// xxx
	common.TestConf()
	common.TestLog()
	shelltool.TestSellTool()
}
