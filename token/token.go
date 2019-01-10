// token/token.go
package token

type TokenType string

type Token struct{
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOL"

    //識別子リテラル
    IDENT = "IDENT"
    INT = "INT"

    // 演算子
    ASSIGN = "="
    PLUS = "+"

    // デリミタ
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // キーワード
    FUNCTION = "FUNCTION"
    LET = "LET"
)

var keywords = map[string] TokenType {
    "fn" : FUNCTION,
    "let" : LET,
}

func  LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}
