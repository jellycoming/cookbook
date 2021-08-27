package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type SafeUrlMap struct {
	v   map[string]bool
	mux sync.Mutex
}

func (m *SafeUrlMap) Put(key string) {
	m.mux.Lock()
	m.v[key] = true
	m.mux.Unlock()
}

func (m *SafeUrlMap) Contains(key string) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	_, ok := m.v[key]
	return ok
}

var urlMap = &SafeUrlMap{v: make(map[string]bool)}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		urlMap.Put(url)
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if urlMap.Contains(u) {
			continue
		}
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	fetcher := &HttpFetcher{}
	Crawl("https://github.com/jellycoming", 4, fetcher)
}

type HttpFetcher struct {
}

func (f *HttpFetcher) Fetch(url string) (string, []string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	client := http.DefaultClient
	res, e := client.Do(req)
	if e != nil {
		return "", nil, fmt.Errorf("request %s error: %s", url, e)
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		fmt.Printf("body from url %s: \n%s", url, resStr)
		return resStr, nil, nil
	} else {
		return "", nil, fmt.Errorf("StatusCode error: %d", res.StatusCode)
	}
}
