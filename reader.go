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
    STAT_NONE = iota
    STAT_AFTER_LT
    STAT_START_TAG
    STAT_END_TAG
    STAT_TEXT
    STAT_PRE_COMMENT1
    STAT_PRE_COMMENT2
    STAT_COMMENT
    STAT_PROCESS_INSTRUCTION
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
    CR = '\n'
    RE = '\r'
    Tab = '\t'
    Question = '?'
    Underscore = '_'
    Eq = '='
    //LeftBracket = '['
    //RightBracket = ']'
)

const(
    End rune = rune(-1)
)

const (
    CDataStr = "CDATA"
)

type Parser interface {
    Parse() void
}

type TextReader interface {
    Read() rune 
    //ReadCharacter() rune
    ReadElement() string
    ReadText() string
    IsSpecialCh(r rune) bool
    ClearStatus()
}

type TextParser struct {
    Data string
    buffer []rune
    status int
    length int
    current int
}

func (p *TextParser) IsSpace(ch rune) bool {
    return ch == Blank || ch == Tab
}

func (p *TextParser) IsAlpha(ch rune) bool {
    return ('a' < ch && ch < 'z') || ( 'A' < ch && ch < 'Z')
}

func (p *TextParser) Parse(){
    
    p.status = STAT_NONE
    for p.current = 0; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch p.status {
            case STAT_NONE:
                if ch == Lt {
                    //reset_buffer
                    p.status = STAT_AFTER_LT
                } else if !p.IsSpace(ch) {
                    p.status = STAT_TEXT
                }
            case STAT_AFTER_LE:
                if ch == Question {
                    p.status = STAT_PROCESS_INSTRUCTION
                } else if ch == Slash {
                    p.status = STAT_END_TAG
                } else if ch == Exclam {
                    p.status = STAT_PRE_COMMENT1
                } else if p.IsAlpha(ch) || ch == Underscore {
                    p.status = STAT_STAT_TAG
                } else {
                    //do nothing
                }
            case STAT_START_TAG:
                //parse start tag
                p.status = STAT_NONE
            case STAT_END_TAG:
                //parse end tag
                p.status = STAT_NONE
            case STAT_PROCESS_INSTRUCTION:
                //parse process instruction
                p.status = STAT_NONE
            case STAT_TEXT:
                //parse text
                p.status = STAT_NONE
            case STAT_PRE_COMMENT1:
                if ch == Dash {
                    p.status = STAT_PRE_COMMENT2
                } else {
                    //do nothing
                }
            case STAT_PRE_COMMENT2:
                if ch == Dash {
                    p.status = STAT_COMMEN
                } else {
                    //do nothing
                }
            case STAT_COMMENT:
                //parse comment
                p.status = STAT_NONE
        }
    }
}

const (
    STAT_PRE_KEY = iota
    STAT_KEY
    STAT_PRE_VALUE
    STAT_VALUE
    STAT_NAME
    STAT_ATTR
    STAT_END
)

const MAX_ATTR_NR = 1024

func (p *TextParser) ParseAttributes(endch rune) {
    status := STAT_PRE_KEY
    valueEnd := Quote
    start := 0
    attrNR := 0
    for ; p.current < p.length && attrNR < MAX_ATTR_NR; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_PRE_KEY:
                if ch == endch || ch == Gt {
                    //read '/' or '>' then go to end status
                    status = STAT_END
                } else if !p.IsSpace(ch) {
                    status = STAT_KEY
                    start = p.current
                }
            case STAT_KEY:
                if ch == Eq {
                    //read the name (p.current - start)
                    names = p.buffer[start: p.current - 1]
                    fmt.Println("attr name ", string(names))
                    status = STAT_PRE_VALUE
                }
            case STAT_PRE_VALUE:
                if ch == Quote || ch == Apos {
                    //read " or '
                    status = STAT_VALUE
                    valueEnd = ch
                    start = p.current + 1
                }
            case STAT_VALUE:
                if ch == valueEnd {
                    values = p.buffer[p.current - start]
                    fmt.Println("attr value", string(values))
                    status = STAT_PRE_KEY
                }
        }

        if status == STAT_END {
            break
        }
    }
}

func (p *TextParser) ParseStartTag() {
    status := STAT_NAME
    start := p.current - 1
    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]

        switch status {
            case STAT_NAME:
                if p.IsSpace(ch) || ch == Gt || ch == Slash {
                    status = ( ch != Gt && ch != Slash ) ? STAT_ATTR : STAT_END
                }
            case STAT_ATTR:
                p.ParseAttributes('/')
                status = STAT_END
        }

        if status == STAT_END {
            break
        }
    }

    //continue to read
}

func (p *TextParser) ParsePI() {
    status := STAT_NAME
    start := p.current

    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_NAME:
                if p.IsSpace(ch) || ch == Gt {
                    status = ch != Gt ? STAT_ATTR : STAT_END
                }
            case STAT_ATTR:
                p.ParseAttributes('?')
                status = STAT_END
        }

        if status == STAT_END {
            break
        }
    }

    tagName = string(p.buffer[p.current - start])
    fmt.Println(tagName)
}




type HtmlReader struct {
    htmlData string
    buffer []rune
    length int
    prev int
    current int
    next int
    name string
    attrs map[string] string
    eltype int
}

func (p* HtmlReader) IsSpecialCh(r rune) bool {
    return r == Lt || r == Gt
}

func (p *HtmlReader) SetData(data string){
    p.htmlData = data
    p.buffer = []rune(p.htmlData)
    p.length = len(p.buffer)
    p.current = 0
    p.prev = 0
    p.next = 0
} 

func (p *HtmlReader) ClearStatus(){
    p.name = ""
    for k := range p.attrs {
        delete(p.attrs, k)
    }
}

func (p *HtmlReader) Skip(){
    //skip to the blank or tab
    for ;p.next < p.length && (p.buffer[p.next] == Blank || p.buffer[p.next] == Tab); {
        fmt.Println("Skip: ", p.current, p.next)
        p.next = p.next + 1
    }
    
    //if p.next < p.length {
    //    fmt.Println("After Skip:", string(p.buffer[p.current]), string(p.buffer[p.next]))
    //}
}

func (p *HtmlReader) SkipToElementStart () {
    
    next := p.buffer[p.next]
    if next == Slash {
        p.eltype = ElementClose
        p.next = p.next + 1
    }else if next == Exclam {
        thirdCh := p.buffer[p.next + 1]
        if thirdCh == Dash {
            p.eltype = Comment
            p.next = p.next + 3 
        }else{
            p.eltype = CData
            p.next = p.next + 1 
        }
    }else{
        p.eltype = ElementOpen
    }
}

func (p* HtmlReader) Read () rune {
  p.prev = p.current
  p.current = p.next
  if p.current >= p.length {
    return End
  }

  r := p.buffer[p.current]
  if p.current == p.length - 1{
    p.next = p.next + 1
    p.current = p.next
    return r
  }
    
  p.next = p.current + 1
  next := p.buffer[p.next]
  
  if r == RE {
    if next == CR {
       p.next = p.next + 1 
    }
    
    //skip to the blank or tab
    p.Skip() 

    p.current = p.next
    p.next = p.next + 1
    r = p.buffer[p.current]
    if p.next < p.length {
        next = p.buffer[p.next + 1]
    }
  } else if r == CR {
    p.Skip()
    p.current = p.next
    p.next = p.next + 1
    r = p.buffer[p.current]
    
    if p.next < p.length {
        next = p.buffer[p.next]
    }
  }else{
    //do nothing
  }

  if r == Lt {
    p.SkipToElementStart()
    p.current = p.next
    p.next = p.next + 1
    r = p.buffer[p.current]
  }else if r == Slash {
    if next == Gt {
        p.next = p.next + 1
        p.ClearStatus()
    }else{
        //p.current = nextPos 
    }

    r = p.buffer[p.next]
  }else if r == Gt{
    p.ClearStatus()
    if next == Lt {
        //r = p.ReadElementStart(r, next)
        p.SkipToElementStart()
        p.current = p.next
        p.next = p.next + 1
        r = p.buffer[p.current]
    }else{
        r = p.buffer[p.next]
        p.next = p.next + 1
    }
  }else if r == Dash && next == Dash {
    p.next = p.next + 2
    p.current = p.next
    p.next = p.next + 1
    r = p.buffer[p.current]
  }else{
    //p.current = nextPos
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
    p.SetData(htmlData)
   
    var str = ""
    for ; p.current < p.length; {  
        r := p.Read()
        str += string(r)
    }
   
    fmt.Println(str)
    //r := buffer[i]
    //r, n := utf8.DecodeRune(buffer)
    //fmt.Printf("%c", r)
    //buffer = buffer[n:]

    return true
}

func WriteFile(filename string, data string){
    f, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
    }
    n,err := io.WriteString(f, data)
    if err != nil {
        fmt.Println(n, err)
    }
    f.Close()
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
