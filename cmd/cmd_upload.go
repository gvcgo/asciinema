package cmd

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	Upload_API = "https://asciinema.org/api/asciicasts"
)

func (r *Runner) Upload() (resp string, err error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	filePart, err := writer.CreateFormFile("asciicast", "ascii.cast")
	_, err = io.Copy(filePart, file)
	writer.Close()
	req, _ := http.NewRequest("POST", Upload_API, buf)
	req.SetBasicAuth("goAsciinema", cfg.ApiToken())
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("User-Agent", "goAsciinema/1.0.0")
	client := &http.Client{
		Timeout: time.Second * 600,
	}
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(rsp.Body)
	if err != nil {
		return "", err
	}
	rsp.Body.Close()
	resp = body.String()
	return resp, err
}
