package idp

import (
	"fmt"
	"io/ioutil"
	"log"
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
	req, err := c.newRequest("GET", "organization/users")
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
	req, err := c.newRequest("GET", "organization/departments")
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

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	url := c.Endpoint + "/" + path

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Sansan-Api-Key", c.APIKey)
	return req, nil
}
