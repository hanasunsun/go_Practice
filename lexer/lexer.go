// lexer/lexer.go

package lexer

import "github.com/hanasunsun/monkey/token"

type Lexer struct {
    input string
    position int        // 入力における現在の位置
    readPosition int    // これから読み込む位置
    ch byte             // 現在検査中の文字
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input){
        l.ch = 0
    }else{
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

// 変数や関数の名前はLetter型になるがこれを全てIDENTで扱うために用意する
func (l *Lexer) readIdentifier() string {
    position := l.position;
    for isLetter(l.ch){
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token{
    var tok token.Token
    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch){
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
    return token.Token{Type: tokenType, Literal: string(ch)}
}

// 変数や関数の名前等に使うことができる文字をこの関数で定義している
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}
