package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("input the url:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	url := input.Text()
	if !validate(url) {
		return
	}
	handleUrl(url)
}

func validate(url string) bool {
	url = strings.Trim(url, " ")
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		fmt.Printf("%s is not validate \n", url)
		return false
	}
	fmt.Printf("%s is validate \n", url)
	return true
}
