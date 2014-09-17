package main

import (
    "fmt"
)

type Handler interface {
    
}

type TokenParser interface {
    SetHandler(handler Handler)
    Parse()
}

type HtmlParser struct {
    handler Handler
    Data string
    buffer []rune
    length int
    current int
}

func (p* HtmlParser) SetData(str string) {
    p.Data = str
    p.buffer = []rune(p.Data)}
    p.length = len(p.buffer)
    p.current = 0
}

func (p *HtmlParser) SetHandler(handler Handler) {
    p.handler = handler
}

func (p *HtmlParser) Parse() {
    for ; p.current < p.length; p.current++ {
        
    }
}
