package hareq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

//post with body and url, get decoded res
func FastPost(body interface{}, url string, resData interface{}) error {
	postBody, _ := json.Marshal(body)
	bodyBuffer := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", bodyBuffer)
	if err != nil {
		log.Printf("An Error Occured %v", err)
		return err
	}

	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return err
	}
	sb := string(bodyRes)
	err = json.Unmarshal([]byte(sb), resData)
	if err != nil {
		println(err.Error())
	}
	defer recover()
	defer resp.Body.Close()
	return nil

}

//get with url, get decoded res
func FastGet(url string, resData interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return err
	}
	sb := string(bodyRes)
	err = json.Unmarshal([]byte(sb), resData)
	return err
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
		log.Printf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	sb := string(bodyRes)
	log.Print(sb)

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
		log.Printf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	sb := string(bodyRes)
	log.Print(sb)

	json.Unmarshal([]byte(sb), resData)
	return nil
}

func FastUpload(filePath string, fileName string, url string, body map[string]string, resData interface{}) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for key, val := range body {
		err := bodyWriter.WriteField(key, val)
		if err != nil {
			panic(err)
		}
	}
	fileWriter, err := bodyWriter.CreateFormFile(fileName, filePath)
	if err != nil {
		fmt.Println("error writing to buffer")
		panic(err)
	}

	fh, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		panic(err)
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resp_body, resData)
	if err != nil {
		panic(err)
	}
	return nil
}
