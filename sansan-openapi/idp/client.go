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
func (c *Client) GetUsers(delimiter string) (string, error) {
	url := c.Endpoint + "/organization/users?delimiter=" + delimiter

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Failed to prepare to get users: %s", err)
	}
	req.Header.Set("X-Sansan-Api-Key", c.APIKey)

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
