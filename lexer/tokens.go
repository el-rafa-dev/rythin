package rythin

type TokenType string

const (
    // Palavras-chave
    TOKEN_MAIN_FUNC    TokenType = "MAIN_FUNC"
    TOKEN_DEF          TokenType = "DEF"
    TOKEN_USING        TokenType = "USING"
    TOKEN_FROM         TokenType = "FROM"
    TOKEN_GET          TokenType = "GET"
    TOKEN_LOOP         TokenType = "LOOP"
    TOKEN_STOP         TokenType = "STOP"
    TOKEN_CONTINUE     TokenType = "CONTINUE"
    TOKEN_IF           TokenType = "IF"
    TOKEN_BUT          TokenType = "BUT"
    TOKEN_TRY          TokenType = "TRY"
    TOKEN_CATCH        TokenType = "CATCH"
    TOKEN_RETURN       TokenType = "RETURN"
    TOKEN_ALLOC        TokenType = "ALLOC"
    TOKEN_FLUSH        TokenType = "FLUSH"
    TOKEN_INPUT        TokenType = "INPUT"
    TOKEN_FWRITE       TokenType = "FWRITE"
    TOKEN_FREAD        TokenType = "FREAD"
    TOKEN_MKDIR	       TokenType = "MKDIR"
    TOKEN_LEN          TokenType = "LEN"
    TOKEN_HAS          TokenType = "HAS"

    // Tipos
    TOKEN_INT          TokenType = "INT"
    TOKEN_BOOL         TokenType = "BOOL"
    TOKEN_STR          TokenType = "STR"
    TOKEN_LONG_INT     TokenType = "LONG_INT"
    TOKEN_VOID         TokenType = "VOID"
    TOKEN_CHAR         TokenType = "CHAR"
    TOKEN_CHARSEQ      TokenType = "CHARSEQ"

    //printers
    TOKEN_PRINT        TokenType = "PRINT"
    TOKEN_PRINT_NEW_LINE TokenType = "PRINT_NEW_LINE"
    TOKEN_PRINT_ERROR    TokenType = "PRINT_ERROR"
    TOKEN_PRINT_ERROR_LOG TokenType = "PRINT_ERROR_LOG"

    // Operadores
    TOKEN_DOT          TokenType = "."
    TOKEN_PLUSPLUS     TokenType = "++"
    TOKEN_ASSIGN       TokenType = "="
    TOKEN_PLUS         TokenType = "+"
    TOKEN_MINUS        TokenType = "-"
    TOKEN_MULTIPLY     TokenType = "*"
    TOKEN_DIVIDE       TokenType = "/"
    TOKEN_NEW_LINE     TokenType = "\\n"
    TOKEN_NOT_EQUAL    TokenType = "NOT_EQUAL"
    TOKEN_EQUAL        TokenType = "EQUAL"
    TOKEN_ARROW_SET    TokenType = "ARROW_SET"          // ->
    TOKEN_COLON        TokenType = "COLON"          // :
    TOKEN_SEMICOLON    TokenType = "SEMICOLON"      // ;
    TOKEN_COMMA        TokenType = "COMMA"
    TOKEN_LBRACKET     TokenType = "LBRACKET"       // [
    TOKEN_RBRACKET     TokenType = "RBRACKET"       // ]
    TOKEN_LPAREN       TokenType = "LPAREN"         // (
    TOKEN_RPAREN       TokenType = "RPAREN"         // )

    // Operadores especiais
    TOKEN_DEREF        TokenType = "DEREF"          // $
    TOKEN_REF          TokenType = "REF"            // &
    TOKEN_COMMENT      TokenType = "COMMENT"        // ;;

    // Literais e identificadores
    TOKEN_IDENTIFIER   TokenType = "IDENTIFIER"
    TOKEN_STRING       TokenType = "STRING"
    TOKEN_NUMBER       TokenType = "NUMBER"
    TOKEN_BOOLEAN      TokenType = "BOOLEAN"
    TOKEN_NULL         TokenType = "NULL"

    // Diversos
    TOKEN_EOF          TokenType = "EOF"
)

type Token struct {
    Type    TokenType
    Literal string
    Line    int
    Column  int
}
