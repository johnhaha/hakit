package hamsg

import wxworkbot "github.com/vimsucks/wxwork-bot-go"

type WeChatBot struct {
	Bot *wxworkbot.WxWorkBot
}

func NewWechatBot(token string) *WeChatBot {
	bot := wxworkbot.New(token)
	return &WeChatBot{Bot: bot}
}

func (bot *WeChatBot) Send(text string) error {
	t := wxworkbot.Text{
		Content: text,
	}
	err := bot.Bot.Send(t)
	return err
}
