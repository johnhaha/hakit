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

func Get(url string, header map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	return bodyRes, err
}

func Delete(url string, header map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	return bodyRes, err
}

func Post(url string, body interface{}, header map[string]string) ([]byte, error) {
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", url, responseBody)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	return bodyRes, err
}

func Put(url string, body interface{}, header map[string]string) ([]byte, error) {
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("PUT", url, responseBody)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	return bodyRes, err
}

//post with body and url, get decoded res
func FastPost(body interface{}, url string, resData interface{}) error {
	res, err := Post(url, body, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, resData)
	return err
}

//get with url, get decoded res
func FastGet(url string, resData interface{}) error {
	res, err := Get(url, nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, resData)
	return err
}

//post with body, url and auth, get decoded res
func AuthFastPost(body interface{}, url string, resData interface{}, th string) error {
	res, err := Post(url, body, map[string]string{
		"Authorization": th,
		"Content-Type":  "application/json; charset=utf-8",
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, resData)
	return err

}

func AuthFastGet(url string, resData interface{}, th string) error {
	res, err := Get(url, map[string]string{
		"Authorization": th,
		"Content-Type":  "application/json; charset=utf-8",
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, resData)
	return err
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
