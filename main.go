package main

import(
	"errors"
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


type parserFunction func(start int ,tokenArray  []Token ) (end int,node *Node ,err error)

func parseSequential(start int ,parsers []parserFunction ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {
	startPoint := start


	currentNode = &Node{startPoint,startPoint,"",[]*Node{}}
	for _,parser := range parsers {
		end ,child ,err := parser(startPoint,tokenArray)
		if err != nil {
			return -1 , nil ,errors.New("failed to parse")
		}
		startPoint = end
		currentNode.adjacent = append(currentNode.adjacent,child)
		currentNode.end = end
	}
	return startPoint , currentNode ,nil
}


func parseDocument(tokenArray []Token)  {  
	end ,CFG ,err := codeParser(0,tokenArray)
	if err != nil   {
		fmt.Println("Failed to parse")
	} else if end != len(tokenArray) {
		PrintGraph(CFG, tokenArray)
		println("parser finished before the end of the document")
		println("parser ended at ",end , " and document ended at ",len(tokenArray))
	}else {
		printTree(CFG , tokenArray)
		PrintGraph(CFG, tokenArray)
		println("parser succeeded")
	}

	
}



// -- Terminal parsers --



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
			println(token.Type , token.Val)
	}
	parseDocument(tokensArray)
}