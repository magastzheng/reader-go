package main

import (
    "fmt"
    "os"
    "io"
    "bytes"
)

const (
    ElementOpen = iota
    ElementClose
    Text
    Comment
    CData
    Other
)

const (
    Lt rune = '<'
    Gt rune = '>'
    Slash rune = '/'
    And rune = '&'
    Apos rune = '\''
    Quot rune = '"'
    Blank rune = ' '
    Exclam = '!'
    Dash = '-'
    //LeftBracket = '['
    //RightBracket = ']'
)

const (
    CDataStr = "CDATA"
)

type TextReader interface {
    Read() rune 
    //ReadCharacter() rune
    ReadElement() string
    ReadText() string
    IsSpecialCh(r rune) bool
    ClearStatus()
}

type HtmlReader struct {
    htmlData string
    buffer []rune
    current int
    name string
    attrs map[string] string
    eltype int
}

func (p* HtmlReader) IsSpecialCh(r rune) bool {
    return r == Lt || r == Gt
}

func (p *HtmlReader) ClearStatus(){
    p.name = ""
    for k := range p.attrs {
        delete(p.attrs, k)
    }
}

func (p* HtmlReader) Read () rune {
  r := p.buffer[p.current]
  
  nextPos := p.current + 1
  next := p.buffer[nextPos]

  if r == Lt {
    if next == Slash {
        p.eltype = ElementClose
        p.current = nextPos + 1
    }else if next == Exclam {
        thirdCh := p.buffer[nextPos + 1]
        if thirdCh == Dash {
            p.eltype = Comment
            p.current = nextPos + 2
        }else{
            p.eltype = CData
            p.current = nextPos + 7
        }
    }else{
        p.eltype = ElementOpen
        p.current = nextPos
    }

    r = p.buffer[p.current]
  }else if r == Slash {
    if next == Gt {
        p.current = nextPos + 1
        p.ClearStatus()
    }else{
        p.current = nextPos 
    }

    r = p.buffer[p.current]
  }else if r == Gt{
    p.ClearStatus()
    p.current = nextPos + 1
    r = p.buffer[p.current]
  }else{
    p.current = nextPos
    //fmt.Println(""
  }

  return r
}

//func (p* HtmlReader) ReadCharacter() rune {
//
//}


//func (p* HtmlReader) ReadElement() string {
//    var element []rune = make([]rune, 1024)
//    ch := p.buffer[current]
//     
//}

//func (p* HtmlReader) ReadCharacter() rune {
//
//}

//func (p* HtmlReader) ReadText() string {
//
//}

func (p *HtmlReader) Parse(htmlData string) bool {
    p.htmlData = htmlData
    p.buffer = []rune(p.htmlData)
    length := len(p.buffer)
    fmt.Println(length)
    p.current = 0
    for r := p.Read(); p.current < length - 1; r = p.Read(){
        
        fmt.Println(p.current, string(r))
    }
    

        //r := buffer[i]
        //r, n := utf8.DecodeRune(buffer)
        //fmt.Printf("%c", r)
        //buffer = buffer[n:]

    return true
}


func main() {
    file, err := os.Open("test.xml")

    if err !=  nil {
        fmt.Println("Cannot read file!")
        panic(err)
    }
    defer file.Close()
    
    //reader := bufio.NewReader(file)

    //chunks := make([]byte, 1024, 1024)
    //buf := make([]byte, 1024)
    //for{
    //    n, err := reader.Read(buf)
    //    if err != nil && err != io.EOF{
    //        panic(err)
    //    }
    //    if n == 0 {
    //        break
    //    }

    //    chunks = append(chunks, buf[:n])
    //}

    //chunks := ioutil.ReadFile("test.xml")
    chunks := bytes.NewBuffer(nil)
    io.Copy(chunks, file)
    str := string(chunks.Bytes())
    fmt.Println(str)

    fmt.Println("Start ... ")
    parser := new(HtmlReader)
    parser.Parse(str)
}
