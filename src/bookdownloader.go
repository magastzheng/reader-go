package main

import (
    "fmt"
    "net/http"
    "io/util"
)

type Charpter struct {
    charpterId string
    charpterName string
    charpterHtml string
}

type BookDownloader struct {
    id string
    name string
    baseUrl string
    charpters map[string] Charpters
}


//http://www.dhzw.com/book/0/983/9090780.html
func (b *BookDownloader) GetIndexUrl() string {
    return "http://www.dhzw.com/book/0/983/"
}

func (b *BookDownloader) GetCharpterUrl(charpterId string) string {
    charpterHtml = b.Charpters[charpterId].charpterHtml
    return b.baseUrl + "/" + charpterHtml
}




