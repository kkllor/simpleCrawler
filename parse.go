package main

import (
	"container/list"
	"fmt"
	"strings"
)

var PNG = "png"
var JPG = "jpg"
var JPEG = "jpeg"

func removeDuplicate(ori *list.List) *list.List {
	l := list.New()
	m := make(map[string]string)
	for e := ori.Front(); e != nil; e = e.Next() {
		value := fmt.Sprintf("%v", e.Value)
		_, ok := m[value]
		if !ok {
			m[value] = value
		}
	}
	for _, value := range m {
		l.PushBack(value)
	}
	return l
}

func filterImages(ori *list.List) *list.List {
	list := list.New()
	for e := ori.Front(); e != nil; e = e.Next() {
		url := fmt.Sprintf("%v", e.Value)
		if strings.Contains(url, PNG) || strings.Contains(url, JPG) || strings.Contains(url, JPEG) {
			list.PushBack(url)
		}
	}
	return list
}
