package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var digitsRegexp = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)

func handleUrl(url string) {
	l := analysis(fetchResponse(url))
	urlList := list.New()
	for e := l.Front(); e != nil; e = e.Next() {
		str := fmt.Sprintf("%v", e.Value)
		res := findUrlByLine(str)
		for i := 0; i < len(res); i++ {
			urlList.PushBack(res[i])
		}
	}

	for e := urlList.Front(); e != nil; e = e.Next() {
		fmt.Println(fmt.Sprintf("%v", e.Value))
	}

	removeDuplicatelist := removeDuplicate(urlList)
	fmt.Println("after remove duplicate:")
	for e := removeDuplicatelist.Front(); e != nil; e = e.Next() {
		fmt.Println(fmt.Sprintf("%v", e.Value))
	}

	imageUrls := filterImages(removeDuplicatelist)
	fmt.Println("image urls:")
	for e := imageUrls.Front(); e != nil; e = e.Next() {
		fmt.Println(fmt.Sprintf("%v", e.Value))
	}

	download(imageUrls)
}

func fetchResponse(url string) io.Reader {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res.Body
}

func analysis(reader io.Reader) *list.List {
	lst := list.New()
	bufReader := bufio.NewReader(reader)
	for {
		line, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		}
		lineBystring := string(line)
		lst.PushBack(lineBystring)
	}
	return lst
}

func findUrlByLine(line string) []string {
	res := digitsRegexp.FindAllString(line, -1)
	return res
}
