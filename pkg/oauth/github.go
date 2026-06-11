package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GitHubUser GitHub API返回的用户信息
type GitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}

// GitHubClient GitHub OAuth客户端
type GitHubClient struct {
	ClientID     string
	ClientSecret string
	HTTPClient   *http.Client
}

// NewGitHubClient 创建GitHub OAuth客户端
func NewGitHubClient(clientID, clientSecret string) *GitHubClient {
	return &GitHubClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   http.DefaultClient,
	}
}

// GetAccessToken 用授权码换取access_token
func (c *GitHubClient) GetAccessToken(code string) (string, error) {
	url := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		c.ClientID,
		c.ClientSecret,
		code,
	)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		AccessToken      string `json:"access_token"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.Error != "" {
		msg := result.Error
		if result.ErrorDescription != "" {
			msg = result.ErrorDescription
		}
		return "", fmt.Errorf("github oauth error: %s", msg)
	}

	return result.AccessToken, nil
}

// GetUser 获取GitHub用户信息
func (c *GitHubClient) GetUser(accessToken string) (*GitHubUser, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user GitHubUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
