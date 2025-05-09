package main

import (
	"fmt"
	"os"
)
import "github.com/el-rafa-dev/rythin/lexer"


func main() {
	data, err := os.ReadFile("/home/el-rafa/VS Code/Rythin/test/soma_test.ry")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	input := string(data)
	lexer := rythin.NewLexer(input)

	for {
		tok := lexer.NextToken()
		fmt.Printf("Token: %-12s Literal: %-20q Linha: %-3d Coluna: %-3d\n", tok.Type, tok.Literal, tok.Line, tok.Column)
		if tok.Type == rythin.TOKEN_EOF {
			break
		}
	}
}
