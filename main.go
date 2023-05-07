package main

import (
	"fmt"
)

func main() {

	fileName := "test1.jean"
	globalSymbolTable := symbolTable{list: make([]symbolTableRow, 0)}
	globalSymbolTable.pointerToHeader = &globalSymbolTable

	tokensArray := Lexer(fileName)
	//fmt.Println(tokensArray)
	arraysWithoutSep := []TokenStruct{}

	for _, token := range tokensArray {
		if token.Type != SEPARATOR {
			arraysWithoutSep = append(arraysWithoutSep, token)
		}
	}

	// print the tokens
	fmt.Println("Tokens:")
	for _, token := range arraysWithoutSep {
		fmt.Println(token.Type, token.Val)
	}
	fmt.Println("Listed all tokens")
	parseDocument(arraysWithoutSep, &globalSymbolTable)
	fmt.Println(globalSymbolTable)

}
