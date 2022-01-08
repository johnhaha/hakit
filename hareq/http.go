package hareq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/johnhaha/hakit/hadata"
	"github.com/johnhaha/hakit/hafile"
)

//post with body and url, get decoded res
func FastPost(body interface{}, url string, resData interface{}) error {
	postBody, _ := json.Marshal(body)
	bodyBuffer := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", bodyBuffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(bodyRes)
	err = json.Unmarshal([]byte(sb), resData)
	if err != nil {
		return err
	}
	return nil

}

//get with url, get decoded res
func FastGet(url string, resData interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(bodyRes)
	err = json.Unmarshal([]byte(sb), resData)
	return err
}

//get with url, get decoded res
func DataGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	sb := string(bodyRes)
	return sb, nil
}

//post with body, url and auth, get decoded res
func AuthFastPost(body interface{}, url string, resData interface{}, th string) error {
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", url, responseBody)
	req.Header.Set("Authorization", th)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(bodyRes)

	json.Unmarshal([]byte(sb), resData)
	return nil
}

func AuthFastGet(url string, resData interface{}, th string) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", th)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(bodyRes)

	json.Unmarshal([]byte(sb), resData)
	return nil
}

func FastUpload(url string, body map[string]string, resData interface{}, files ...FileToUpload) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//create body data
	for key, val := range body {
		err := bodyWriter.WriteField(key, val)
		if err != nil {
			return err
		}
	}
	//create file data
	for _, file := range files {
		fileWriter, err := bodyWriter.CreateFormFile(file.Name, file.Path)
		if err != nil {
			return err
		}

		fh, err := os.Open(file.Path)
		if err != nil {
			return err
		}
		defer fh.Close()
		_, err = io.Copy(fileWriter, fh)
		if err != nil {
			return err
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp_body, resData)
	if err != nil {
		return err
	}
	return nil
}

type FileToUpload struct {
	Name string
	Path string
}

func DownloadFileFromUrl(url string, saveIn string, name string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("can not get from url, status code" + hadata.GetStringFromInt(response.StatusCode))
	}
	hafile.CheckFolder(saveIn)
	file, err := os.Create(saveIn + "/" + name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
