package main

import(
	"fmt"
)



const  (
	VAR  = "var"
	ADD_OP  = "addition operator"
	MULT_OP  = "multiplication operator"
	NUM  = "number"
	ASSIGN  = "="
	LEFT_BRACKET  = "("
	RIGHT_BRACKET  = ")"
)









func main()  {
	// x = x + ( 5 * 3 ) + ( 2  * 4)
	// y = 4 + 2 * x

	
	tokensArray := []Token{ Token{Type :  VAR , Val : "x" }, Token{Type : ASSIGN , Val : "=" }, 
	Token{Type : VAR , Val : "x" }, Token{Type : ADD_OP , Val : "+" },  Token{Type : LEFT_BRACKET , Val :  "(" }, Token{Type : NUM , Val : "5" }, Token{Type : MULT_OP , Val : "*" }, Token{Type : NUM , Val : "3" },
	Token{Type : RIGHT_BRACKET , Val : ")" }, Token{Type : ADD_OP , Val : "+" }, 
	Token{Type: LEFT_BRACKET , Val : "(" }, Token{Type : NUM , Val : "2" }, 
	Token{Type : MULT_OP , Val : "*" }, Token{Type : NUM , Val : "4" }, Token{Type : RIGHT_BRACKET , Val : ")" } , Token{Type: VAR , Val : "y"} , Token{Type : ASSIGN , Val : "=" }, Token{Type : NUM , Val : "4" }, Token{Type : ADD_OP , Val : "+" }, Token{Type : NUM , Val : "2" }, Token{Type : MULT_OP , Val : "*" }, Token{Type : VAR , Val : "x"} }
	
	// print the tokens
	for _,token := range tokensArray {
			fmt.Println(token.Type , token.Val)
	}
	parseDocument(tokensArray)
}