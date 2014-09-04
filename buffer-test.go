package main

import (
    //"bytes"
    "fmt"
)

const (
    Lt rune = '<'
    Gt rune = '>'
)

func main(){
    str := "Hello你好！"
    //buf := []byte(s)
    length := len(str)
    fmt.Printf("%d", length)
    for i := 0; i < length; i++{
        b := str[i] //buf.ReadByte()
        fmt.Println("%c", b)
    }

    fmt.Println("======================")
    for i,s := range str {
        fmt.Println(i, "Unicode(",s,") string=", string(s))
    }

    r := []rune(str)
    fmt.Println("rune=", r)
    fmt.Println("len=", len(r))
    for i := 0; i < len(r); i++ {
        fmt.Println("r[",i,"]=",r[i],"string=",string(r[i]))
    }

    

    fmt.Println("lt=", string(Lt))
    xmlstr := "<>/'\""
    xmlr := []rune(xmlstr)
    for i := 0; i < len(xmlr); i++ {
        fmt.Println("xmlr[",i,"]=", xmlr[i], "string=",string(xmlr[i]))
    }
}
