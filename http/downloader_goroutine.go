package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"sync"
)

func createZeroFile(size int64, file *os.File) {

	//创建空文件（方便下面文件指针偏移）
	fmt.Println("正在创建空文件 size: ", size)
	var zeroSize int64 = 0
	zeroBytes := make([]byte, 1024)
	if size <= 1024 {
		zeroBytes := make([]byte, size)
		zeroReader := bytes.NewReader(zeroBytes)
		n, err := io.Copy(file, zeroReader)
		if err != nil {
			panic(err)
		}
		fmt.Println("ok. zero file size ：", n)
		return
	}
	for {
		zeroReader := bytes.NewReader(zeroBytes)
		n, err := io.Copy(file, zeroReader)
		if err != nil {
			panic(err)
		}
		zeroSize += n
		// fmt.Printf("写入了: %v/%v\n", zeroSize, size)

		if zeroSize == size {
			fmt.Println("ok. zero file size ：", n)
			break
		}
	}
}

func block(length int64) (int64, int64) {
	// 划分文件为10份
	b := length
	var inc int64 = 0
	var block int64
	for {
		// 取余数
		y := math.Mod(float64(b), 10.0)
		if y != 0 {
			b++
			inc++
		} else {
			block = b / 10
			fmt.Println("block is : ", block, "last block is: ", block-inc)
			break
		}
	}
	return block, block - inc
	// fmt.Println(block*9 + block - inc)
}

func downloader(file *os.File, w int, block int64, lastBlock int64, url string, done chan<- bool, m *sync.Mutex) {
	fmt.Println("starting worker: ", w)

	// 启动http请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	jobs := 10
	if w == 1 {
		req.Header.Set(
			"Range",
			fmt.Sprintf("bytes=0-%v", block))
	} else if w > 1 && w < jobs {
		req.Header.Set(
			"Range",
			fmt.Sprintf("bytes=%v-%v", block*(int64(w)-1)+1, block*int64(w)))
	} else if w == jobs {
		req.Header.Set(
			"Range",
			fmt.Sprintf("bytes=%v-%v", block*(int64(w)-1)+1, block*(int64(w)-1)+1+lastBlock))
	}

	if resp, err := client.Do(req); err == nil {
		if resp.StatusCode != 206 {
			panic("error 206 response,不支持的分段下载操作\nview: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range")
		}
		defer resp.Body.Close()

		m.Lock()
		// 数据写入偏移位置
		if w == 1 {
			file.Seek(block*(int64(w)-1), 0)
		} else {
			file.Seek(block*(int64(w)-1)+1, 0)
		}
		n, err := io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
		m.Unlock()
		fmt.Println("worker: ", w, " end ", "writed ", n)
		done <- true
	}
}

func getFileInfo() (url string, filename string, length int64) {

	//获取下载文件信息

	// url := "https://mirrors.aliyun.com/archlinux/iso/archboot/2018.06/archlinux-2018.06-1-archboot-network.iso"
	// url := "https://mirrors.aliyun.com/archlinux/iso/archboot/2018.06/md5sum.txt"

	// 获取url和文件名称
	args := os.Args
	if len(args) != 3 {
		panic("error download link\nuse ./downloader {LINK} file.name")
	}

	// 获取文件大小
	resp, err := http.Get(args[1])
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 404 {
		panic("error 404, not found")
	}
	fmt.Println(resp.Status)
	fmt.Println("file size: ", resp.ContentLength)

	return args[1], args[2], int64(resp.ContentLength)
}

func main() {

	url, filename, length := getFileInfo()

	// 创建下载文件
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 写入空值到下载文件中
	createZeroFile(length, f)
	b, lb := block(length)

	// 初始化互斥锁
	mutex := &sync.Mutex{}

	done := make(chan bool, 10)
	// 下发任务
	for w := 1; w <= 10; w++ {
		go downloader(f, w, b, lb, url, done, mutex)
	}

	// 阻塞任务
	for j := 1; j <= 10; j++ {
		<-done
	}
	close(done)

}
