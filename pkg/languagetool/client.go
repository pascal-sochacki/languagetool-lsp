package languagetool

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"strconv"

	"go.uber.org/zap"
)

type LanguagetoolApi interface {
	CheckText(ctx context.Context, text string, language string) (CheckResult, error)
}

type Client struct {
	log     *zap.Logger
	baseURL string
	client  *http.Client
}

func NewClient(logger *zap.Logger) *Client {
	return &Client{
		baseURL: "https://api.languagetoolplus.com/v2/",
		client:  &http.Client{},
		log:     logger,
	}
}

type CheckResult struct {
	Matches []Match `json:"matches"`
}

type Match struct {
	Message      string        `json:"message"`
	Offset       int           `json:"offset"`
	Length       int           `json:"length"`
	Context      MatchContext  `json:"context"`
	Sentence     string        `json:"sentence"`
	Replacements []Replacement `json:"replacements"`
}

type Replacement struct {
	Value string `json:"value"`
}

type MatchContext struct {
	Text   string `json:"text"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

func (c Client) CheckText(ctx context.Context, text string, language string) (CheckResult, error) {
	c.log.Debug(text)
	c.log.Debug(fmt.Sprintf("%d", len(text)))

	result := CheckResult{}
	fullUrl := c.baseURL + "check"

	formData := url.Values{
		"text":        {text},
		"username":    {""},
		"apiKey":      {""},
		"language":    {"auto"},
		"enabledOnly": {strconv.FormatBool(false)},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullUrl, strings.NewReader(formData.Encode()))

	if err != nil {
		c.log.Error(err.Error())
		return result, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	c.log.Debug(resp.Status)
	if err != nil {
		c.log.Error(err.Error())
		return result, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.log.Error(err.Error())
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, nil

}
