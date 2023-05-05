package main

import (
	"fmt"
	"os"
)



func main()  {
	
	tokensArray := []TokenStruct{ {Type :  (VAR) , Val : "x" }, {Type :  ASSIGN , Val : "=" }, 
	{Type : VAR , Val : "x" }, {Type : ADD , Val : "+" },  {Type : OPEN_PARAN , Val :  "(" }, {Type : NUMBER , Val : "5" }, {Type : MUL , Val : "*" }, {Type : NUMBER , Val : "3" },
	{Type : CLOSE_PARAN , Val : ")" }, {Type : ADD , Val : "+" }, 
	{Type: OPEN_PARAN , Val : "(" }, {Type : NUMBER , Val : "2" }, 
	{Type : MUL , Val : "*" }, {Type : NUMBER , Val : "4" }, {Type : CLOSE_PARAN , Val : ")" } , {Type: VAR , Val : "y"} , {Type : ASSIGN , Val : "=" }, {Type : NUMBER , Val : "4" }, {Type : ADD , Val : "+" }, {Type : NUMBER , Val : "2" }, {Type : MUL , Val : "*" }, {Type : VAR , Val : "x"} }
	

	fileName := os.Args[1]
	tokensArray = Lexer(fileName) 

	arraysWithoutSep := []TokenStruct{}

	for _ ,  token := range tokensArray {
		if token.Type != SEPARATOR {
			arraysWithoutSep = append(arraysWithoutSep , token)
		}
	}
	// print the tokens
	for _,token := range arraysWithoutSep {
			fmt.Println(token.Type , token.Val)
	}
	parseDocument(arraysWithoutSep)

	// sample_str := readSample("sample.jean")
	// sample_str = sample_str + "\n"
	// tokens := lex_analyzer(sample_str)
	// fmt.Println(tokens)
}