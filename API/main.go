package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	app := appWithHttpClient()
	res, _ := app.getWithQueryString()
	fmt.Println(string(res))
}

type App struct {
	client *http.Client
}

func appWithHttpClient() *App {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			MaxIdleConnsPerHost: 10,
		},
	}

	return &App{client: client}
}

func (a *App) getFromAPI() ([]byte, error) {
	resp, err := a.client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *App) getFromAPIWithClient() ([]byte, error) {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1/comments", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *App) getWithQueryString() ([]byte, error) {

	//id := "1"
	//params := "postId=" + url.QueryEscape(id)
	//path := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?%s", params)

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/comments", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("postId", "1")
	req.URL.RawQuery = q.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}