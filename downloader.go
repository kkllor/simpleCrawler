package main

import (
	"container/list"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var pathPrefix = getHomePath() + "/Downloads/simpleCrawler/"

func download(l *list.List) {
	if l.Len() == 0 {
		return
	}
	c := make(chan string, l.Len())
	path := pathPrefix + fmt.Sprintf("%v", time.Now().Unix())
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for e := l.Front(); e != nil; e = e.Next() {
		go downloadFromNet(fmt.Sprintf("%v", e.Value), path, c)
	}
	for i := 0; i < l.Len(); i++ {
		fmt.Printf("%s 下载完成!\n", <-c)
	}
}

func downloadFromNet(url string, path string, c chan string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	path += string(os.PathSeparator) + getDownLoadName(url)
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	fmt.Println("%s 下载完成!", url)
	c <- url
	return err
}
