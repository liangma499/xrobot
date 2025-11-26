package user

import (
	"context"
	optionChannelDao "xrobot/internal/dao/option-channel"
	userDao "xrobot/internal/dao/user"
	userBaseDao "xrobot/internal/dao/user-base"
	"xrobot/internal/event/message"
	"xrobot/internal/model"

	"xbase/log"
	"xbase/utils/xrand"
	optiontelegramcmd "xrobot/internal/option/option-telegram-cmd"
	"xrobot/internal/service/user/pb"
	tgmsg "xrobot/internal/xtelegram/tg-msg"
	tgtemplate "xrobot/internal/xtelegram/tg-template"
	tgtypes "xrobot/internal/xtelegram/tg-types"
	waitforinput "xrobot/internal/xtelegram/wait-for-input"
	"xrobot/internal/xtypes"
)

func (s *Server) doSubscribeMessageStart(uuid string, payload *message.MessageStart) {
	if payload == nil {
		return
	}
	if payload.Type != message.MessageType_Private {
		return
	}
	ctx := context.Background()
	if payload.UserID == 0 {
		return
	}

	user, err := userDao.Instance().FindOneByTelegramUserID(ctx, payload.UserID)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	channelCfg, err := optionChannelDao.Instance().GetChannel(ctx, payload.ChannelCode)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if channelCfg == nil {
		log.Errorf("channelCfg is nil")
		return
	}

	var userBase *model.UserBase
	if user == nil {

		user, userBase, err = s.doRegister(ctx, &registerArgs{
			telegramuserName:  payload.UserName,
			telegramPwd:       xrand.Str(xrand.LetterSeed+xrand.DigitSeed, 32),
			registerLoginType: xtypes.RegisterLoginType_TG,
			telegramUserID:    payload.UserID,
			channelName:       channelCfg.Name,
			inviteCode:        payload.InviteCode,
			userType:          xtypes.UserTypeTg,
			avatar:            "",
			clientIP:          payload.ClientIP,
			isVerifiedEmail:   xtypes.UserVerifiedEmailStatusNo,
			channelCode:       payload.ChannelCode,
		}, &pb.DeviceInfo{
			DeviceType: 2,
		})
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		if userBase == nil {
			log.Warnf("%v", "userBase == nil")
			return
		}

	} else {
		userBase, err = userBaseDao.Instance().GetUserBase(ctx, user.ID)
		if userBase == nil {
			log.Errorf("%v", err)
			return
		}
		if user.TelegramPwd == "" {
			user.TelegramPwd = xrand.Str(xrand.LetterSeed+xrand.DigitSeed, 32)
			userDao.Instance().Update(ctx, func(cols *userDao.Columns) any {
				return map[string]any{
					cols.ID: user.ID,
				}
			}, func(cols *userDao.Columns) any {
				return map[string]any{
					cols.TelegramPwd: user.TelegramPwd,
				}
			})
		}

	}

	_, err = s.doLogin(ctx, user, payload.ClientIP, &pb.DeviceInfo{
		DeviceType: 2,
	}, xtypes.RegisterLoginType_TG)
	if err != nil {
		log.Errorf("%v", err)
	}
	cmdMsg := optiontelegramcmd.GetChanCodeCmd(payload.ChannelCode, tgtypes.XTelegramCmd_Start)
	if cmdMsg == nil {
		return
	}
	xMsg, err := tgmsg.NewXTelegramMessage(channelCfg.TelegramCfg.MainRobotToken,
		tgmsg.WithText(cmdMsg.Text),
		tgmsg.WithDebug(true),
		tgmsg.WithCmd(cmdMsg.Cmd),
		tgmsg.WithMsgType(cmdMsg.Type),
		tgmsg.WithParseMode(cmdMsg.ParseMode),
		tgmsg.WithKeyboard(cmdMsg.Keyboard))
	/*
		ChannelCode string                   `json:"channelCode"` //渠道code
		Cmd         tgtypes.XTelegramCmd     `json:"cmd"`         //推送方式
		Keyboard    *tgbutton.TelegramButton `json:"keyboard"`
		PictureUrl  string                   `json:"picture_url"` //图片
	*/
	if err != nil {
		return
	}
	if xMsg == nil {
		return
	}
	if _, err := xMsg.SendMessage(payload.ChatID, map[string]string{
		tgtemplate.CustomerKey:      channelCfg.ChannelCfg.Customer,
		tgtemplate.EnergySavingsKey: channelCfg.ChannelCfg.EnergySavings,
	}); err != nil {
		log.Warnf("sendMessage:%v", err)
	} else {
		waitforinput.SetWaitForinputKey(&waitforinput.WaitforinputInfo{
			UserID:        payload.UserID, //用户TG
			InPutMsg:      payload.WaitforinputMsg.InPutMsg,
			UserBottonKey: payload.Button.WaitForInputKey(),
			Extended:      nil,
		})
	}

}
