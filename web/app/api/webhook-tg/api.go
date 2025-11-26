package webhooktg

import (
	"xbase/log"
	"xbase/transport"
	"xrobot/internal/code"
	"xrobot/internal/http/response"

	"xrobot/internal/xtelegram/telegram/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type API struct {
	validate    *validator.Validate
	transporter transport.Transporter
}

func NewAPI(transporter transport.Transporter) *API {
	return &API{
		validate:    validator.New(),
		transporter: transporter,
	}
}

// tgWebhook监听
func (a *API) TelegramWebhook(ctx *gin.Context) {

	req := &types.Update{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Warnf("test002:%v", "elegramWebhook")
		response.Fail(ctx, code.InvalidArgument)
	}
	channelCode := ctx.GetHeader("X-Telegram-Bot-Api-Secret-Token")
	log.Warnf("X-Telegram-Bot-Api-Secret-Token:%v", channelCode)
	if req.Message != nil {
		if req.Message.Text == "" || req.Message.Chat.ID == 0 {
			response.Success(ctx, 200)
			return
		}
		if err := a.botMsg(req.Message, ctx.ClientIP(), channelCode); err != nil {
			log.Warnf("errCommand:%v", err)
			response.Fail(ctx, code.InternalError)
		}
	} else if req.CallbackQuery != nil {
		//log.Warnf("callbackQuery:%#v message:%v", *req.CallbackQuery, req.CallbackQuery.Message)
		if err := a.botCallBackMsg(req.CallbackQuery, ctx.ClientIP(), channelCode); err != nil {
			log.Warnf("errCommand:%v", err)
			response.Fail(ctx, code.InternalError)
		}
	}

	response.Success(ctx, 200)
}
