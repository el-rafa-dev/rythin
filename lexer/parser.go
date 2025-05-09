package rythin

import "github.com/el-rafa-dev/rythin/lexer/tokens"
//import "../tokens" // para testar localmente (não recomendado para produção)


type Parser struct {
	tokens  []tokens.Token
	current int
	curToken tokens.Token // Adicionando a referência do token atual
}

func NewParser(tokens []tokens.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) parseFunctionDef() FunctionDef {
	// Consome o 'def' e avança
	p.consume(tokens.DEF, "Esperado 'def'")

	// Verifica se há ponteiro '*'
	//isPointer := p.match(tokens.STAR)

	// Consome o identificador (nome da função)
	name := p.consume(tokens.IDENTIFIER, "Esperado nome da função")

	// Consome o ':' que define o tipo de retorno
	p.consume(tokens.COLON, "Esperado ':'")
	typ := p.consume(tokens.VOID, "Esperado tipo de retorno (ex: 'void')").Literal

	// Consome o '->' (seta de retorno)
	p.consume(tokens.ARROW_SET, "Esperado '->'")

	// Consome o '[' que indica início do corpo da função
	p.consume(tokens.LBRACKET, "Esperado '[' para início do corpo")

	// Começa a ler as declarações do corpo
	var body []Statement
	for !p.check(tokens.RBRACKET) && !p.isAtEnd() {
		body = append(body, p.parseStatement())
	}

	// Consome o ']' no final do corpo
	p.consume(tokens.RBRACKET, "Esperado ']' no final da função")

	// Retorna a definição da função
	return FunctionDef{
		Name:       name.Literal,
		ReturnType: typ,
		Body:       body,
	}
}

// Função principal de parsing
func (p *Parser) Parse() Program {
	var functions []FunctionDef

	// Parse todas as funções no código
	for !p.isAtEnd() {
		fn := p.parseFunctionDef()
		functions = append(functions, fn)
	}

	// Retorna o programa com todas as funções parseadas
	return Program{Functions: functions}
}

// Consome o token atual e avança para o próximo token
func (p *Parser) consume(expectedType tokens.TokenType, errorMsg string) tokens.Token {
	if p.curToken.Type == expectedType {
		tok := p.curToken
		p.nextToken() // Avança para o próximo token
		return tok
	}
	// Lança um erro ou pode retornar um erro caso o tipo não seja o esperado
	return panic(errorMsg + " encontrado: " + p.curToken.Literal)
}

// Avança para o próximo token
func (p *Parser) nextToken() {
	if p.isAtEnd() {
		return
	}
	p.curToken = p.tokens[p.current]
	p.current++
}

// Verifica se o próximo token é do tipo esperado
func (p *Parser) match(expected tokens.TokenType) bool {
	if p.curToken.Type == expected {
		p.nextToken()
		return true
	}
	return false
}

// Verifica se o token atual é do tipo esperado
func (p *Parser) check(expected tokens.TokenType) bool {
	return p.curToken.Type == expected
}

// Verifica se chegou ao fim da lista de tokens
func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens)
}

// Exemplo de um Statement que poderia ser parseado
func (p *Parser) parseStatement() Statement {
	// Aqui você implementaria a lógica para processar uma declaração (statement)
	// Dependeria de como suas declarações estão estruturadas no código da linguagem
	// Exemplo: retorno de uma expressão, atribuições, etc.
	//return Statement{}
	return nil
}
