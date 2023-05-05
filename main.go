package main

import (
	"fmt"
)




func main()  {

	
	// x = x + ( 5 * 3 ) + ( 2  * 4)
	// y = 4 + 2 * x

	
	tokensArray := []tokenStruct{ {Type :  VAR , Val : "x" }, {Type : ASSIGN , Val : "=" }, 
	{Type : VAR , Val : "x" }, {Type : ADD_OP , Val : "+" },  {Type : LEFT_BRACKET , Val :  "(" }, {Type : NUM , Val : "5" }, {Type : MULT_OP , Val : "*" }, {Type : NUM , Val : "3" },
	{Type : RIGHT_BRACKET , Val : ")" }, {Type : ADD_OP , Val : "+" }, 
	{Type: LEFT_BRACKET , Val : "(" }, {Type : NUM , Val : "2" }, 
	{Type : MULT_OP , Val : "*" }, {Type : NUM , Val : "4" }, {Type : RIGHT_BRACKET , Val : ")" } , {Type: VAR , Val : "y"} , {Type : ASSIGN , Val : "=" }, {Type : NUM , Val : "4" }, {Type : ADD_OP , Val : "+" }, {Type : NUM , Val : "2" }, {Type : MULT_OP , Val : "*" }, {Type : VAR , Val : "x"} }
	
	// print the tokens
	for _,token := range tokensArray {
			fmt.Println(token.Type , token.Val)
	}
	parseDocument(tokensArray)
}