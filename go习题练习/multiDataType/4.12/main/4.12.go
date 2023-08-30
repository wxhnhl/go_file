package main

import (
	"fmt"
	"net/http"
	"os"
	//"io/ioutil"
	"encoding/json"
	"strings"
)

type Xvcd struct {
	Title      string
	Img        string
	Transcript string
}

/*
练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写>
一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
1.暂时没有存索引
2.使用协程实现,很快
*/
func main() {
	var urls []string
	for i := 0; i < 1000; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		urls = append(urls, url)
	}
	//var content []Xvcd
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	fmt.Println(<-ch)
}
func fetch(url string, ch chan<- string) {
	var result Xvcd
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	titles := strings.Split(result.Transcript, " ")

	for _, v := range titles {
		if v == string(os.Args[1]) {
			ch <- result.Img
			break
		}
	}
}
