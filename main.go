package main

import (
	"github.com/FranklinThree/male-vocal-dev/config"
	"github.com/FranklinThree/male-vocal-dev/server"
	"github.com/FranklinThree/male-vocal-dev/utils"
	"os"
	"os/signal"
	"syscall"

	// 新增module后请在下方引入
	_ "github.com/FranklinThree/male-vocal-dev/module/newUser"         // 新关注的欢迎模块（示例）
	_ "github.com/FranklinThree/male-vocal-dev/module/pong"            // gin的ping-pong模块
	_ "github.com/FranklinThree/male-vocal-dev/module/templateMessage" // 集中管理发送模板消息的模组
	_ "github.com/FranklinThree/male-vocal-dev/module/testPong"        // three 开发的第一个模块
	_ "github.com/FranklinThree/male-vocal-dev/module/wechatPong"      // 微信ping-pong模块
)

func init() {
	utils.WriteLogToFS()
	config.Init()

}

func main() {
	server.Init()

	server.StartService()

	server.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	server.Stop()
}
