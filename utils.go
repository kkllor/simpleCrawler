package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os/user"
	"strings"
)

func getHomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func getDownLoadName(url string) string {
	md5 := Md5(url)
	if strings.Contains(url, PNG) {
		return md5 + "." + PNG
	} else if strings.Contains(url, JPEG) {
		return md5 + "." + JPEG
	} else if strings.Contains(url, JPG) {
		return md5 + "." + JPEG
	}
	return ""
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}
