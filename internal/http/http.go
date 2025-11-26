package http

import (
	"xbase/component"
	"xbase/etc"
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/http/middleware"

	"xbase/log/zap"

	"github.com/gin-gonic/gin"
)

type RouteHandler func(engine *gin.Engine)

type Http struct {
	component.Base
	engine       *gin.Engine
	logger       *logger
	routeHandler RouteHandler
}

func NewHttp(routeHandler RouteHandler) *Http {
	gin.DefaultWriter = newLogger()

	h := &Http{}
	h.logger = gin.DefaultWriter.(*logger)
	h.engine = gin.New()
	h.engine.Use(middleware.Logger(), middleware.Recover(), middleware.Cors())
	h.routeHandler = routeHandler

	return h
}

// Name 组件名称
func (h *Http) Name() string {
	return "http"
}

// Init 初始化组件
func (h *Http) Init() {
	if h.routeHandler != nil {
		h.routeHandler(h.engine)
	}
}

// Start 启动组件
func (h *Http) Start() {
	go func() {
		addr := etc.Get("etc.http.addr", ":8080").String()
		if err := h.engine.Run(addr); err != nil {
			log.Fatal("http server startup failed: %v", err)
		}
	}()
}

// Destroy 销毁组件
func (h *Http) Destroy() {
	_ = h.logger.Close()
}

// Engine 获取gin.Engine
func (h *Http) Engine() *gin.Engine {
	return h.engine
}

type logger struct {
	logger *zap.Logger
}

func newLogger() *logger {
	logFile := etc.Get("etc.http.log", "./log/http.log").String()

	return &logger{
		logger: zap.NewLogger(
			zap.WithFile(logFile),
			zap.WithCallerSkip(2),
			zap.WithStackLevel(log.PanicLevel+1),
		),
	}
}

func (l *logger) Write(p []byte) (n int, err error) {
	l.logger.Print(log.InfoLevel, xconv.String(p))
	return
}

func (l *logger) Close() error {
	return l.logger.Close()
}
