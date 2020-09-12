package main

import (
	// "bufio"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"sync"

	"github.com/schollz/progressbar/v3"
)

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

func downloader(file *os.File, w int, block int64, lastBlock int64, url string, done chan<- bool, m *sync.Mutex, bar *progressbar.ProgressBar) {
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

		// 数据写入偏移位置
		offset := block * (int64(w) - 1)
		if w != 1 {
			offset = offset + 1
		}
		//此处有问题尚未解决，seek后，由于其他协程的操作，可能指针立刻变成了其他位置，下方的copy便会错位,必须要保证指针在copy之前不可变
		buf := make([]byte, 1024)

		for {
			nr, err := resp.Body.Read(buf)
			if nr > 0 {
				file.WriteAt(buf[0:nr], offset)
				offset += int64(nr)
			}
			if err != nil {
				break
			}
			if err == io.EOF {
				break
			}
		}

		fmt.Println("worker: ", w, "end")
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
	// createZeroFile(length, f)
	b, lb := block(length)

	bar := progressbar.DefaultBytes(
		length,
		"downloading",
	)

	// 初始化互斥锁
	mutex := &sync.Mutex{}

	done := make(chan bool, 10)
	// 下发任务
	for w := 1; w <= 10; w++ {
		go downloader(f, w, b, lb, url, done, mutex, bar)
	}

	// 阻塞任务
	for j := 1; j <= 10; j++ {
		<-done
	}
	close(done)
	fmt.Println("end")

}
