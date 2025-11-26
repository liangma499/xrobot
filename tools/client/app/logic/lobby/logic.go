package lobby

import (
	"context"
	"time"

	//lobbyProto "xrobot/internal/xprotocol/lobby"
	//"xrobot/internal/xprotocol/msgid"
	"xbase/cluster"
	"xbase/cluster/client"
	"xbase/eventbus"
	"xbase/log"
	"xrobot/tools/client/app/topic"
)

type UserInfo struct {
}
type Logic struct {
	ctx   context.Context
	proxy *client.Proxy
	Conn  *client.Conn
}

func NewLogic(proxy *client.Proxy) *Logic {
	return &Logic{
		ctx:   context.Background(),
		proxy: proxy,
	}
}

func (l *Logic) Init() {
	// 初始化事件
	l.initHook()
	// 初始化事件
	l.initEvent()
	// 初始化路由
	l.initRoute()
	// 初始化事件总线
	l.initEventbus()
}

// 初始化客户端事件
func (l *Logic) initHook() {
	// 初始化
	l.proxy.AddHookListener(cluster.Init, l.InitHook)
	// 打开连接
	l.proxy.AddHookListener(cluster.Start, l.StartHook)
	// 打开连接
	//l.proxy.AddHookListener(cluster.Restart, l.RestartHook)
	// 打开连接
	l.proxy.AddHookListener(cluster.Destroy, l.DestroyHook)

}

// 初始化客户端事件
func (l *Logic) initEvent() {
	// 打开连接
	l.proxy.AddEventListener(cluster.Connect, l.connectHandle)
	// 重新连接
	l.proxy.AddEventListener(cluster.Reconnect, l.reconnectHandle)
	// 断开连接
	l.proxy.AddEventListener(cluster.Disconnect, l.disconnectHandle)
	// 心跳事件
	//l.proxy.AddEventListener(cluster.Heartbeat, l.heartHandle)
}

// 初始化路由
func (l *Logic) initRoute() {
	//注册
	//l.proxy.AddRouteHandler(int32(msgid.RummyMsgID_LobbyRegisterReq), l.lobbyRegister)
	// 登录响应
	//l.proxy.AddRouteHandler(int32(msgid.RummyMsgID_LobbyLoginEmailReq), l.lobbyLoginEmail)

}

// 初始化事件总线
func (l *Logic) initEventbus() {
	// 监听登录成功事件
	_ = eventbus.Subscribe(l.ctx, topic.LoginSuccess, l.loginSuccessHandle)
}

// 链接初始化
func (l *Logic) InitHook(proxy *client.Proxy) {
	log.Warnf("InitHook")
}

// 链接初始化
func (l *Logic) StartHook(proxy *client.Proxy) {
	log.Warnf("StartHook")
	conn, err := proxy.Dial()
	if err != nil {
		log.Errorf("%v", err)
	}
	l.Conn = conn
}

// 链接初始化
func (l *Logic) RestartHook(proxy *client.Proxy) {
	log.Warnf("RestartHook")
}

// 链接初始化
func (l *Logic) DestroyHook(proxy *client.Proxy) {
	log.Warnf("DestroyHook")
}

// 处理连接打开
func (l *Logic) connectHandle(conn *client.Conn) {
	l.RegisterRes(conn)
}

// 处理重新连接
func (l *Logic) reconnectHandle(_ *client.Conn) {
	log.Infof("connection is reconnect")
}

// 处理断开连接
func (l *Logic) disconnectHandle(_ *client.Conn) {
	log.Infof("connection is closed")
	/*
		err := l.proxy.Reconnect()
		if err != nil {
			log.Errorf("reconnect failed: %v", err)
		}*/
}

// 用户心跳
func (l *Logic) heartHandle(conn *client.Conn) {
	if conn == nil {
		return
	}
	log.Warnf("heartHandle: %v", conn.UID())

}

const (
	email    = "test9999@gmail.com"
	password = "123456789"
)

// doTest
func (l *Logic) RegisterRes(conn *client.Conn) {
	/*
		err := conn.Push(&cluster.Message{
			Route: int32(msgid.RummyMsgID_LobbyRegisterReq),
			Data: &lobbyProto.LobbyRegisterReq{
				Email:       email,
				Password:    password,
				InviteCode:  "",
				ChannelCode: "",
				DeviceID:    "123456789",
				DeviceType:  1,
				DeviceModel: "123456789",
				Captcha:     "123456",
			},
		})
		if err != nil {
			log.Errorf("send login message failed: %v", err)
		}
	*/
}

// 登录
func (l *Logic) lobbyRegister(ctx *client.Context) {
	/*
		res := &lobbyProto.LobbyRegisterRes{}
		if err := ctx.Parse(res); err != nil {
			log.Warnf("%v", err)
			return
		}

		log.Warnf("res: %#v", res)

		if res.Code == 0 || res.Code == 106 {
			//用户登录
			err := ctx.Conn().Push(&cluster.Message{
				Route: int32(msgid.RummyMsgID_LobbyLoginEmailRes),
				Data: &lobbyProto.LobbyLoginEmailReq{
					Email:       email,
					Password:    password,
					DeviceID:    "123456789",
					DeviceType:  1,
					DeviceModel: "123456789",
				},
			})
			if err != nil {
				log.Errorf("send login message failed: %v", err)
			}
		} else {
			log.Errorf("%v", res.Code)
		}
	*/
}

// 登录成功发送心跳
func (l *Logic) lobbyLoginEmail(ctx *client.Context) {
	/*
		res := &lobbyProto.LobbyLoginEmailRes{}
		if err := ctx.Parse(res); err != nil {
			log.Warnf("%v", err)
			return
		}

		//ctx.Conn().Bind(29)
		log.Warnf("lobbyLoginEmail: %#v,:%v:uid", res, ctx.UID())

		//登录成功发送心跳
		task.AddTask(func() {
			l.doInitHeartData(ctx)

		})
	*/
}

// 用户心跳
func (l *Logic) doInitHeartData(ctx *client.Context) {
	conn := ctx.Conn()
	l.doSendHeartData(conn)
	timer := time.NewTicker(5 * time.Second)
	for {
		<-timer.C
		l.doSendHeartData(conn)
	}

}
func (l *Logic) doSendHeartData(conn *client.Conn) {
	if conn == nil {
		log.Warnf("conn is nil")
		return
	}

	err := conn.Push(
		&cluster.Message{
			Route: 1,
			Data:  nil,
		})
	if err != nil {
		log.Warnf("conn is err:%v", err)
		return
	}
	log.Warnf("doSendHeartData succes")
}

// 登录成功事件处理
func (l *Logic) loginSuccessHandle(event *eventbus.Event) {
	log.Info("login success")

	for i := 0; i < 10; i++ {
		// 拉取游戏房间
		l.doFetchRooms(3032)

		// 拉取游戏房间
		l.doFetchRooms(1100)
	}

	// 拉取邮件信息
	l.doFetchMails()

	// 拉取公告信息
	l.doFetchAnnouncements()
	// 拉取游戏
	l.doFetchGames()
	//拉取游戏记录
	l.doFetchGameRecordList()
}

// 拉取资产结果
func (l *Logic) fetchAssetHandle(ctx *client.Context) {
	/*
		res := &logic.FetchAssetRes{}

		err := ctx.Parse(res)
		if err != nil {
			log.Errorf("invalid response message, err: %v", err)
			return
		}

		if res.Code != 0 {
			log.Errorf("fetch user's asset failed, code: %d", res.Code)
			return
		}

		log.Infof("user's asset, gold = %d diamond = %d ticket = %d", res.Data.Gold, res.Data.Diamond, res.Data.Ticket)
	*/
}

// 拉取游戏房间
func (l *Logic) doFetchRooms(gameID int) {
	/*
		err := l.proxy.Push(&cluster.Message{
			Route: route.FetchRooms,
			Data: logic.FetchRoomsReq{
				GameID: gameID,
			},
		})
		if err != nil {
			log.Errorf("fetch game's rooms failed: %v", err)
		}
	*/
}

func (l *Logic) fetchRoomsHandle(ctx *client.Context) {
	/*
		res := &logic.FetchRoomsRes{}

		err := ctx.Parse(res)
		if err != nil {
			log.Errorf("invalid response message, err: %v", err)
			return
		}

		if res.Code != 0 {
			log.Errorf("fetch game's rooms failed, code: %d", res.Code)
			return
		}

		log.Info(xconv.Json(res.Data))
	*/
}

func (l *Logic) doFetchMails() {
	/*
		err := l.proxy.Push(&cluster.Message{
			Route: route.FetchMails,
		})
		if err != nil {
			log.Errorf("fetch user's mails failed: %v", err)
		}
	*/
}

// 拉取邮件
func (l *Logic) fetchMailsHandle(ctx *client.Context) {
	/*
		res := &logic.FetchMailsRes{}

		err := ctx.Parse(res)
		if err != nil {
			log.Errorf("invalid response message, err: %v", err)
			return
		}

		if res.Code != 0 {
			log.Errorf("fetch user's mails failed, code: %d", res.Code)
			return
		}

		log.Info(xconv.Json(res.Data))
	*/
}

func (l *Logic) doFetchAnnouncements() {
	/*
		err := l.proxy.Push(&cluster.Message{
			Route: route.FetchAnnouncements,
		})
		if err != nil {
			log.Errorf("fetch announcements failed: %v", err)
		}
	*/
}

// 拉取公告处理
func (l *Logic) fetchAnnouncementsHandle(ctx *client.Context) {
	/*
		res := &logic.FetchAnnouncementsRes{}

		err := ctx.Parse(res)
		if err != nil {
			log.Errorf("invalid response message, err: %v", err)
			return
		}

		if res.Code != 0 {
			log.Errorf("fetch announcements failed, code: %d", res.Code)
			return
		}

		log.Info(xconv.Json(res.Data))
	*/
}

func (l *Logic) doFetchGames() {
	/*
		err := l.proxy.Push(&cluster.Message{
			Route: route.FetchGames,
			Data: &logic.FetchGamesReq{
				AgentID: 1,
			},
		})
		if err != nil {
			log.Errorf("fetch games failed: %v", err)
		}
	*/
}

// 拉取游戏处理
func (l *Logic) fetchGamesHandle(ctx *client.Context) {
	/*
		res := &logic.FetchGamesRes{}

		err := ctx.Parse(res)
		if err != nil {
			log.Errorf("invalid response message, err: %v", err)
			return
		}

		if res.Code != 0 {
			log.Errorf("fetch games failed, code: %d", res.Code)
			return
		}

		log.Info(xconv.Json(res.Data))
	*/
}

func (l *Logic) doFetchGameRecordList() {
	/*
		err := l.proxy.Push(&cluster.Message{
			Route: route.FetchGameRecordList,
			Data:  &logic.FetchGameRecordListReq{},
		})
		if err != nil {
			log.Errorf("fetch games failed: %v", err)
		}
	*/
}
