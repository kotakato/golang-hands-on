package idp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

// DefaultEndpoint はデフォルトのAPIエンドポイントを表す。
const DefaultEndpoint = "https://api.sansan.com/v2.1"

// Client はSansan OpenAPIの組織APIクライアントを表す。
type Client struct {
	Endpoint string
	APIKey   string
}

// NewClient はClientオブジェクトのポインタを取得する。
func NewClient(endpoint, apiKey string) *Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	if strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint[:len(endpoint)-1]
	}
	return &Client{
		Endpoint: endpoint,
		APIKey:   apiKey,
	}
}

// GetUsers はユーザー一覧のCSV文字列を取得する。
func (c *Client) GetUsers() (string, error) {
	req, err := c.newRequest("GET", "organization/users", nil)
	if err != nil {
		return "", fmt.Errorf("Failed to prepare to get users: %s", err)
	}

	log.Printf("Requesting [%s] %s\n", req.Method, req.URL)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to get users: %s", err)
	}
	defer resp.Body.Close()

	log.Printf("Status: %d\n", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to get users: %s", err)
	}

	return string(body), nil
}

// GetDepartments は部署一覧のCSV文字列を取得する。
func (c *Client) GetDepartments() (string, error) {
	req, err := c.newRequest("GET", "organization/departments", nil)
	if err != nil {
		return "", fmt.Errorf("Failed to prepare to get departments: %s", err)
	}

	log.Printf("Requesting [%s] %s\n", req.Method, req.URL)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to get departments: %s", err)
	}
	defer resp.Body.Close()

	log.Printf("Status: %d\n", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to get departments: %s", err)
	}

	return string(body), nil
}

// Check はCSVファイルが正しいことを検証する。
func (c *Client) Check(files map[string]io.Reader) (string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for fieldname, file := range files {
		err := addFile(w, fieldname, file)
		if err != nil {
			return "", fmt.Errorf("Failed to prepare to check: %s", err)
		}
	}
	w.Close()

	req, err := c.newRequest("POST", "organization/check", &b)
	if err != nil {
		return "", fmt.Errorf("Failed to prepare to check: %s", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	log.Printf("Requesting [%s] %s\n", req.Method, req.URL)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to get departments: %s", err)
	}
	defer resp.Body.Close()

	log.Printf("Status: %d\n", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to get departments: %s", err)
	}

	return string(body), nil
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := c.Endpoint + "/" + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Sansan-Api-Key", c.APIKey)
	return req, nil
}

func addFile(w *multipart.Writer, filedname string, file io.Reader) error {
	fw, err := w.CreateFormField(filedname)
	if err != nil {
		return fmt.Errorf("Failed to prepare to get departments: %s", err)
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return fmt.Errorf("Failed to prepare to get departments: %s", err)
	}
	return nil
}
