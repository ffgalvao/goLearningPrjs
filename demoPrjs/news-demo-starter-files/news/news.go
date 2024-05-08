package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ResultsTpl struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
}

type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

func (a *Article) FormatPublishedDate() string {
	year, month, day := a.PublishedAt.Date()
	return fmt.Sprintf("%v %d, %d", month, day, year)
}

type Results struct {
	ResultsTpl
	Articles []Article `json:"articles"`
}

func NewClient(httpClient *http.Client, apikey string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}
	baseUrl := "https://newsapi.org/v2/everything"

	return &Client{httpClient, baseUrl, apikey, pageSize}
}

type Client struct {
	http     *http.Client
	baseUrl  string
	apikey   string
	PageSize int
}

func (c *Client) FetchEverything(query string, page string) (*Results, error) {

	endpoint := fmt.Sprintf(
		"?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=en",
		url.QueryEscape(query),
		c.PageSize,
		page,
		c.apikey,
	)

	resp, err := c.http.Get(c.baseUrl + endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Results{}
	return res, json.Unmarshal(body, res)
}
