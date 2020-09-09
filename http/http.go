// kuwo api下载脚本 仅限本人学习使用
package main

import (
	// "compress/zlib"
	"encoding/json"
	"fmt"
	// "time"

	// "internal/reflectlite"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "strings"
)

// json-abslist
type JsonData struct {
	Abslist []struct {
		ALBUM    string `json:"ALBUM"`
		ARTIST   string `json:"ARTIST"`
		MUSICRID string `json:"MUSICRID"`
		NAME     string `json:"NAME"`
	} `json:"abslist"`
}

// rid返回json数据结构体
type msgData struct {
	MSG string `json:"msg"`
	URL string `json:"url"`
}

type mp3Info struct {
	name     string
	musicrid string
	artist   string
	url      string
}

func errLog(msg string, e error) {
	if e != nil {
		fmt.Printf(msg, "\t")
		log.Fatalln(e)
		return
	}
}
func main() {

	var musicKeyword string
	fmt.Printf("input music name: ")
	fmt.Scanln(&musicKeyword)
	rn := 30
	pn := 1
	API := fmt.Sprintf("http://search.kuwo.cn/r.s?client=kt&all=%vpn=%v&rn=%v&uid=794762570&ver=kwplayer_ar_9.2.2.1&vipver=1&show_copyright_off=1&newver=1&ft=music&cluster=0&strategy=2012&encoding=utf8&rformat=json&vermerge=1&mobi=1&issubtitle=01", musicKeyword, pn, rn)

	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	errLog("API get", err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36")
	resp, err := client.Do(req)
	// resp.Header.Set("Content-Type", "application/json; charset=utf-8")
	errLog("resp", err)

	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status, resp.Header)

	// fmt.Println(r)

	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var abslistData JsonData

	// str := strings.ReplaceAll(string(body), "'", "\"")
	// fmt.Println(str)
	bytes := []byte(string(body))
	errAbslistData := json.Unmarshal(bytes, &abslistData)
	errLog("abslistData unmarshal", errAbslistData)
	// chooseMap := make(map[int]string)
	// 创建结构体切片,保存MP3信息列表
	result := make([]mp3Info, len(abslistData.Abslist))
	for index, item := range abslistData.Abslist {
		fmt.Println(index, "\t", item.NAME, "\t", item.ALBUM, "\t", item.ARTIST, "\t", item.MUSICRID)
		// chooseMap[index] = item.MUSICRID
		result[index] = mp3Info{
			name:     item.NAME,
			artist:   item.ARTIST,
			musicrid: item.MUSICRID,
		}
	}

	var chooseID int
	fmt.Printf("select music id: ")
	fmt.Scanf("%d", &chooseID)

	done := make(chan bool)
	if chooseID-1 < len(result) && chooseID >= 0 {
		go downloader(chooseID, done, result)
		<-done
	} else if chooseID == 100 {
		// 协程任务下发
		for j := 0; j < len(result); j++ {
			go downloader(j, done, result)
		}
		for w := 1; w <= len(result); w++ {
			// 阻塞等待任务完毕
			<-done
		}
	} else {
		fmt.Println("error chooseID", len(result), chooseID)
		return
	}

}

func downloader(index int, done chan<- bool, result []mp3Info) {
	// 从切片中获取mp3Info结构体数据
	musicrid := result[index]
	fmt.Println(musicrid)
	// 利用musicrid获取MP3 url
	resp, err := http.Get(fmt.Sprintf("http://www.kuwo.cn/url?rid=%v&type=convert_url3&br=320kmp3", musicrid.musicrid))
	errLog("musicrid get", err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errLog("musicrid read body", err)
	// fmt.Println(string(body))
	body = []byte(string(body))
	msg := msgData{}
	errURL := json.Unmarshal(body, &msg)
	if errURL != nil {
		done <- false
		return
	}
	// errLog("musicrid Unmarshal", errURL)
	fmt.Println(msg.URL)
	musicrid.url = msg.URL
	fmt.Printf("downloader %v started\n", index)
	mp3Req, err := http.Get(musicrid.url)
	errLog("mp3 url get", err)
	defer mp3Req.Body.Close()
	// 创建MP3文件
	filePath := fmt.Sprintf("%v-%v.mp3", musicrid.name, musicrid.artist)
	mp3File, err := os.Create(filePath)
	errLog("create mp3 file", err)
	defer mp3File.Close()
	_, err = io.Copy(mp3File, mp3Req.Body)
	errLog("copy bytes to file", err)
	fmt.Println("done -> ", filePath)
	done <- true

}
