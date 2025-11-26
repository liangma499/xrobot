package webhooktg

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"
	"tron_robot/internal/event/message"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	waitforinput "tron_robot/internal/xtelegram/wait-for-input"
	"xbase/log"

	"github.com/mr-tron/base58"
)

func (a *API) isValidTronAddress(address string) bool {
	// TRON 地址正则：以 T 开头，后跟 33 个字符（可为字母和数字）
	var tronAddressPattern = regexp.MustCompile(`^T[a-zA-Z0-9]{33}$`)
	return tronAddressPattern.MatchString(address)
}

// decodeBase58 解码 Base58 地址
func (a *API) decodeBase58(address string) ([]byte, error) {
	return base58.Decode(address)
}

// validateAddress 验证地址的合法性
func (a *API) validateAddress(address string) bool {
	if !a.isValidTronAddress(address) {
		return false
	}

	decoded, err := a.decodeBase58(address)
	if err != nil {
		return false
	}

	// 对解码后的数据进行 SHA256 哈希
	checksum := decoded[len(decoded)-4:] // 获取最后4个字节作为校验和
	body := decoded[:len(decoded)-4]     // 取出有效数据部分

	// 计算校验和
	hash := sha256.Sum256(body)
	hash = sha256.Sum256(hash[:])

	// 验证校验和
	calculatedChecksum := hash[:4]
	return hex.EncodeToString(calculatedChecksum) == hex.EncodeToString(checksum)
}

func (a *API) doGetUserMsg(text string, userID int64) (tgtypes.XTelegramButton, message.WaitforinputMsg, string) {
	inputMsg := message.WaitforinputMsg{
		InPutMsg: false,
		PutMsg:   "", //接收到的消息
	}
	inviteCode := ""
	buttonKeyStr := text
	if strings.HasPrefix(text, "start") || strings.HasPrefix(text, "/start") {
		msgInfos := strings.Split(text, "&")
		lenMsgInfo := len(msgInfos)
		if lenMsgInfo == 0 {
			log.Warnf("err new: %v", msgInfos)
			return tgtypes.XTelegramButton_None, inputMsg, inviteCode
		} else if lenMsgInfo > 1 {
			inviteCode = msgInfos[1]
		}
		buttonKeyStr = "/start"
	}

	button := tgtypes.StringToXTelegramButton(buttonKeyStr)
	inputMsg.PutMsg = text
	if button == tgtypes.XTelegramButton_None {
		if inputInfo := waitforinput.GetWaitForinputKey(userID); inputInfo != nil {
			if inputInfo.UserBottonKey != "" {
				inputMsg.InPutMsg = true
				inputMsg.Extended = inputInfo.Extended.Clone()
				buttonKeyStr = inputInfo.UserBottonKey
				button = tgtypes.StringToXTelegramButton(buttonKeyStr)
			}
		} else if a.validateAddress(text) {
			button = tgtypes.XTelegramButton_AddressDetail
		}

	}

	return button, inputMsg, inviteCode
}
