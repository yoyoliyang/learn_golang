package main

import (
	// "compress/zlib"
	"encoding/json"
	"fmt"
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

func errLog(e error) {
	if e != nil {
		log.Fatalln(e)
		return
	}
}
func main() {

	var musicKeyword string
	fmt.Printf("input music name: ")
	fmt.Scanln(&musicKeyword)
	rn := 20
	API := fmt.Sprintf("http://search.kuwo.cn/r.s?client=kt&all=%vpn=1&rn=%v&uid=794762570&ver=kwplayer_ar_9.2.2.1&vipver=1&show_copyright_off=1&newver=1&ft=music&cluster=0&strategy=2012&encoding=utf8&rformat=json&vermerge=1&mobi=1&issubtitle=01", musicKeyword, rn)

	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	errLog(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36")
	resp, err := client.Do(req)
	// resp.Header.Set("Content-Type", "application/json; charset=utf-8")
	errLog(err)

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
	errLog(errAbslistData)
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

	// 从切片中获取mp3Info结构体数据
	var musicrid mp3Info
	if chooseID-1 < len(result) && chooseID >= 0 {
		musicrid = result[chooseID]
		fmt.Println(musicrid)
		// 利用musicrid获取MP3 url
		resp, err := http.Get(fmt.Sprintf("http://www.kuwo.cn/url?rid=%v&type=convert_url3&br=320kmp3", musicrid.musicrid))
		errLog(err)
		body, err := ioutil.ReadAll(resp.Body)
		errLog(err)
		body = []byte(string(body))
		msg := msgData{}
		errUrl := json.Unmarshal(body, &msg)
		errLog(errUrl)
		fmt.Println(msg.URL)
		musicrid.url = msg.URL
	} else {
		fmt.Println("error chooseID", len(result), chooseID)
		return
	}
	mp3Req, err := http.Get(musicrid.url)
	errLog(err)
	// 创建MP3文件
	filePath := fmt.Sprintf("/data/python/music_downloader/%v-%v.mp3", musicrid.name, musicrid.artist)
	mp3File, err := os.Create(filePath)
	errLog(err)
	defer mp3File.Close()
	_, err = io.Copy(mp3File, mp3Req.Body)
	errLog(err)
	fmt.Println("done -> ", filePath)

}
