package hamsg

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type DiscordHook struct {
	Url string
}

func NewDiscordHook(url string) *DiscordHook {
	return &DiscordHook{Url: url}
}

func (hook *DiscordHook) Send(content string) error {
	c := http.Client{}
	form := url.Values{}
	form.Set("content", content)
	req, err := http.NewRequest("POST", hook.Url, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	return nil
}
