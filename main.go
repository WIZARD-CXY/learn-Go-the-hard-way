package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// position is the token position in the source text,
// used for error tracing,line and col counts from 0.
type position struct {
	line int //line number
	col  int //colummn number
}

type TokenType int

//lexical token.
type token struct {
	lit string    //literal value
	typ TokenType //token type
	pos position  //postion
}

//for debug
func (t token) String() string {
	return fmt.Sprintf("<lit:\"%s\",typ:%s,pos line:%d,pos col:%d>", t.lit, tokens[t.typ], t.pos.line, t.pos.col)
}

//stateFn is used for state automation convertion.
type stateFn func(l *lexer) stateFn

//lexical parser.
type lexer struct {
	cur token // current scanned token

	src string // source

	pos   int // current scanning index
	start int // start scanning index
	width int // width of string scanned

	lineNum int // line counter,counts from 0
	colNum  int // column counter,counts from 1

	errors    []string   // errors stack
	state     stateFn    // state function
	tokenChan chan token // token channel
}

const (
	token_begin TokenType = iota
	tNUM                  // number -?digit*.digit*[E|e]-?digit*

	tPLUS  // +
	tMINUS // -

	tMUTIL // *
	tDIV   // /

	tEOF // eof
	token_end
)

const eof = -1

var tokens = map[TokenType]string{
	tMINUS: "[-]",
	tDIV:   "[/]",
	tPLUS:  "[+]",
	tMUTIL: "[*]",
	tEOF:   "[EndOfFile]",
}

//newLexer initiates token channel and go run lexer.run and return lexer.
func newLexer(src string) *lexer {
	l := &lexer{src: src,
		tokenChan: make(chan token),
	}
	go l.run()
	return l
}

//scan next token
func (l *lexer) next() rune {
	if l.pos >= len(l.src) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.src[l.pos:])
	//fmt.Println("haha", l.src[l.pos:], l.pos, w)
	l.width = w
	l.pos += l.width
	l.colNum += l.width
	//fmt.Println("haha", l.src[l.pos:], l.pos, w, l.width)
	fmt.Println(string(r))
	return r
}

//push error message into tracing stack
func (l *lexer) err(e string) {
	l.errors = append(l.errors, e)
}

//format error
func (l *lexer) errf(f string, v ...interface{}) {
	l.errors = append(l.errors, fmt.Sprintf(f, v...))
}

//backup a token
func (l *lexer) backup() {
	l.pos -= l.width
	l.colNum -= l.width
}

//peek a token
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

//ignore a token
func (l *lexer) ignore() {
	l.start = l.pos
}

//emit token to channel tokenChan
func (l *lexer) emit(typ TokenType) {
	var t = token{
		lit: l.src[l.start:l.pos],
		typ: typ,
		pos: position{l.lineNum, l.colNum - (l.pos - l.start)},
	}
	fmt.Println("token is ", t)
	l.tokenChan <- t
	if t.typ == tEOF {
		close(l.tokenChan)
	}
	l.cur = t
	l.start = l.pos

}

//read token one by one
func (l *lexer) token() token {
	token := <-l.tokenChan
	return token
}

//main consumer routine
func (l *lexer) run() {
	for l.state = lexBegin; l.state != nil; {
		l.state = l.state(l)
	}
}

//------------------------------------state function----------------------------------
func lexError(l *lexer) stateFn {
	//premature lexical scanning
	l.emit(tEOF)
	return nil
}

func lexUnkown(l *lexer) stateFn {
	//premature lexical scanning
	l.emit(tEOF)
	return nil
}

//TODO:scan number,and emit the token.
func lexNum(l *lexer) stateFn {
	//unfinished
	l.emit(tNUM)
	return lexBegin
}

//end of scanning
func lexEOF(l *lexer) stateFn {
	l.emit(tEOF)
	return nil
}

//main lex entry
func lexBegin(l *lexer) stateFn {
	switch r := l.next(); {
	case unicode.IsDigit(r) || r == '.' || r == '-':
		//fmt.Printf("%#v\n", l)
		//l.backup()
		if r == '-' && l.cur.typ == tNUM {
			goto L //go to minus
		}
		lexNum(l)
		return lexBegin
	L:
		fallthrough
	case r == '-':
		l.emit(tMINUS)
	case r == '*':
		l.emit(tMUTIL)
	case r == '/':
		l.emit(tDIV)
	case r == '+':
		l.emit(tPLUS)
	case r == ' ':
		l.ignore()
	case r == '\n':
		l.ignore()
		l.lineNum++
		l.colNum = 0
		//l.emit(tNEWLINE),currently not neccesary in parsing.
	case r == eof:
		return lexEOF
	default:
		l.errf("unkown char '%c' at line %d,column %d", r, l.lineNum+1, l.colNum)
		return lexUnkown
	}
	return lexBegin
}

func main() {
	println(`In this task we will focus on a lexer implementation,and it's concurrency part.
lexer is a lexical scanner that consumes source code and produce meaningful tokens. With these tokens we can then 
complete a small calculator.Our simple lexer just need to scan several tokens '+','-','*','/',and numbers.
Lexer is a typical producer-consumer pattern,so we need a channel to send token after lexer initiated and run the scanner in a goroutine.
Instead of switch,we use state function,in order to skip the case statements.
And finally we just need to receive tokens from the channel.
Now edit main.go and finish the task.Utitiles of lexer have been given,you need to write a small regexp engine to complete the 'lexNum' stateFn and pass the test.(Notice don't run 'go test' right now,because lexNum is currently infinite loop.`)
}
