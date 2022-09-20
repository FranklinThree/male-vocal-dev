package testPong

import (
	"github.com/FranklinThree/male-vocal-dev/server"
	"github.com/getsentry/sentry-go"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"strconv"
	"sync"
)

func init() {
	instance = &testPong{}
	server.RegisterModule(instance)
}

var instance *testPong

type testPong struct {
}

func (m *testPong) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       server.NewModuleID("FranklinThree", "testPong"),
		Instance: instance,
	}
}

func (m *testPong) Init() {
}

func (m *testPong) PostInit() {

}

func (m *testPong) Serve(s *server.Server) {
	s.MsgEngine.Group("^testPing$", func(msg *server.Message) {
		msg.Reply = &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("testPong\n" +
			"replyCount :" + strconv.Itoa(s.ReplyCount) + "\n" +
			"by FranklinThree")}
	}).MsgText("", 1).EventClick("")
}

func (m *testPong) Start(s *server.Server) {
	defer sentry.Recover()
}

func (m *testPong) Stop(s *server.Server, wg *sync.WaitGroup) {
	defer wg.Done()

}
