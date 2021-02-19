package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"time"
	"xagent/src/common"
	"xagent/src/glbval"
	"xagent/src/httpser"
)

// 在编译时赋值的变量
var (
	gitCommitID    string
	buildTime      string
	buildGoVersion string
)

func main() {
	var err error

	// 赋值全局参数
	glbval.GitCommitID = gitCommitID
	glbval.BuildTime = buildTime
	glbval.BuildGoVersion = buildGoVersion
	glbval.StartTime = time.Now().Format("2006-01-02 03:04:05")
	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	glbval.RootPath = path.Join(binDir, "..")

	// 加载配置
	err = common.InitConf()
	if err != nil {
		fmt.Printf("[Error]: InitConf failed. err=%v", err)
		return
	}

	// 初始化日志
	err = common.InitLogger()
	if err != nil {
		fmt.Printf("[Error]: InitLogger failed. err=%v", err)
		return
	}
	common.LogInfo("Program git commitID: %s", glbval.GitCommitID)
	common.LogInfo("Program build time:   %s", glbval.BuildTime)

	// 启动HTTP服务
	err = httpser.Start()
	if err != nil {
		common.LogError("Start http service failed. err=%v", err)
		return
	}
	common.LogInfo("Start http service success.")

	// 等待停止信号
	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt, os.Kill)
	<-chanSignal

	// 释放资源
	common.LogInfo("Program stop.")
}
