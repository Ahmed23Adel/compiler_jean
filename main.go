package main

import (
	"fmt"
	"os"
)

func main() {

	fileName := os.Args[1]
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
	cfg := parseDocument(arraysWithoutSep)
	if cfg != nil {
		if(semanticCheck(arraysWithoutSep)){
			quads := EvaluateCode(cfg , arraysWithoutSep )
			quads = Jump2Goto(quads)
			PrintQuadruplesToFile(quads , "output.quads")
		}else{
			fmt.Println("Sematic errors.")
		}
	}

}
/////////
