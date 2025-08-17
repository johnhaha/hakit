package hamsg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FeishuHook struct {
	Url string
}

func NewFeishuHook(url string) *FeishuHook {
	return &FeishuHook{Url: url}
}

func (hook *FeishuHook) SendText(content string) error {
	body := map[string]any{
		"msg_type": "text",
		"content": map[string]string{
			"text": content,
		},
	}
	err := hook.Send(body)
	if err != nil {
		return err
	}
	return nil
}

func (hook *FeishuHook) SendRichText(title string, content string) error {
	body := map[string]any{
		"msg_type": "post",
		"content": map[string]any{
			"post": map[string]any{
				"zh_cn": map[string]any{
					"title": title,
					"content": [][]map[string]string{
						{
							{
								"tag":  "text",
								"text": content,
							},
						},
					},
				},
			},
		},
	}
	err := hook.Send(body)
	if err != nil {
		return err
	}
	return nil
}

func (hook *FeishuHook) Send(body map[string]any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(hook.Url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("send feishu msg failed, error code is: %d", resp.StatusCode)
	}
}
