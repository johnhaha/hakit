package hamsg

import "github.com/johnhaha/hakit/hareq"

type DiscordHook struct {
	Url string
}

func NewDiscordHook(url string) *DiscordHook {
	return &DiscordHook{Url: url}
}

func (hook *DiscordHook) Send(content string) error {
	body := map[string]string{
		"content": content,
	}
	var res interface{}
	err := hareq.FastPost(body, hook.Url, res)
	return err
}
