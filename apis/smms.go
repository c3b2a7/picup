package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type SMMS struct {
	Token string
}

type SMMSV2Response struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		FileId    int    `json:"file_id"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Filename  string `json:"filename"`
		Storename string `json:"storename"`
		Size      int    `json:"size"`
		Path      string `json:"path"`
		Hash      string `json:"hash"`
		Url       string `json:"url"`
		Delete    string `json:"delete"`
		Page      string `json:"page"`
	} `json:"data"`
	Images    string `json:"images"`
	RequestId string `json:"RequestId"`
}

type SMMSV2RespWriter struct {
	resp *SMMSV2Response
}

func (w *SMMSV2RespWriter) Write(bytes []byte) (n int, err error) {
	resp := &SMMSV2Response{}
	json.Unmarshal(bytes, resp)
	w.resp = resp
	return len(bytes), nil
}

func (smms *SMMS) Up(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := new(bytes.Buffer)
	writer := multipart.NewWriter(reader)
	formFile, err := writer.CreateFormFile("smfile", file.Name())
	if err != nil {
		return "", err
	}
	io.Copy(formFile, file)
	writer.Close()
	request, err := http.NewRequest("POST", "https://sm.ms/api/v2/upload", reader)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", smms.Token)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	respWriter := &SMMSV2RespWriter{}
	io.Copy(respWriter, response.Body)
	resp := respWriter.resp
	if resp.Success {
		return resp.Data.Url, nil
	} else if resp.Images != "" {
		return resp.Images, nil
	}
	return "", errors.New("upload failed")
}
