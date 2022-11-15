package main

/*
go to mmt
download content of web
num of mmt keyword in downloaded content
do it efficiently with golang concepts

Print Time taken to fidn word


url = https://www.makemytrip.com/
hit url
get content
search content
parallization possiblities
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func parallelizeCOunt(content, word string) int {

	totalWords := 0
	var lck sync.Mutex

	for _, line := range content {
		go func(line string) {
			cnt := countWord(string(line), word)
			lck.Lock()
			totalWords += cnt
			lck.Unlock()
		}(string(line))
	}
	return totalWords
}

func countWord(content, word string) int {
	return strings.Count(content, word)
}

func main() {
	//fmt.Println()

	tDownloadSt := time.Now()
	sContent := OnPage("https://www.makemytrip.com/")
	tDownloadend := time.Now().Sub(tDownloadSt)
	tProcessingSt := time.Now()
	count := parallelizeCOunt(sContent, "makemytrip")
	tProcessingEnd := time.Now().Sub(tProcessingSt)//tProcessingSt.Sub(time.Now())
	fmt.Println(tDownloadend, " TOtal processing time ", tProcessingEnd)
	fmt.Println(count)
}
