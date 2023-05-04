package main

import(
	"errors"
	"fmt"
)

type Node struct{
	start int
	end int
	name  string
	adjacent []*Node
}


func printTree(node *Node , tokenArray  []Token) {
    // Print the current node
    print("node name: " , node.name, ",Contents: ("  )
	for i := node.start; i < node.end; i++ {
		print( tokenArray[i].Val , " ")
	}
	println(")")
    // Print the adjacent nodes recursively
	println("Children : ")
    for _, adj := range node.adjacent {
        printTree(adj , tokenArray)
    }
	println("End of node ",node.name)
}


const  (
	VAR  = "var"
	ADD_OP  = "addition operator"
	MULT_OP  = "multiplication operator"
	NUM  = "number"
	ASSIGN  = "="
	LEFT_BRACKET  = "("
	RIGHT_BRACKET  = ")"
)

type Token struct {
	Type string 
	Val string

}

type parserFunction func(start int ,tokenArray  []Token ) (end int,node *Node ,err error)

func parseSequential(start int ,parsers []parserFunction ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {
	startPoint := start

	currentNode = &Node{startPoint,startPoint,"",[]*Node{}}
	for _,parser := range parsers {
		end ,child ,err := parser(startPoint,tokenArray)
		if err != nil {
			return -1 , nil ,errors.New("Failed to parse")
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
		println("parser finished before the end of the document")
		println("parser ended at ",end , " and document ended at ",len(tokenArray))
	}else {
		println("parser succeeded")
		printTree(CFG , tokenArray)
	}

	
}

func codeParser(start int ,tokenArray  []Token ) (end int,currentNode *Node , err error) {
	
	option1 := []parserFunction{stmtParser , codeParser}
	option2 := []parserFunction{}

	options := [][]parserFunction{option1,option2}
	
	for _ ,option := range options {
		end ,currentNode ,err = parseSequential(start,option,tokenArray)
			if err == nil {
				fmt.Println("Code parser succeeded")
				currentNode.name = "code"
				return end , currentNode ,nil
			}
	}
	//fmt.Println("All options failed, code parser")
	return -1 , nil,errors.New("Failed to parse")
}

func stmtParser(start int ,tokenArray  []Token ) (end int,currentNode *Node , err error) {  
	option1 := []parserFunction{varParser,assignParser,exprParser} 

	options := [][]parserFunction{option1}

	for _ ,option := range options {
		end ,currentNode ,err = parseSequential(start,option,tokenArray)
		if err == nil {
				fmt.Println("Statement parser succeeded")
				currentNode.name = "stmt"
				return end , currentNode,nil
		}
	}
	//fmt.Println("All options failed, stmt parser")
	return -1 , nil,errors.New("Failed to parse")
}



func exprParser(start int ,tokenArray  []Token ) (end int,currentNode  *Node  , err error) {  
	option1 := []parserFunction{termParser,addOpParser,exprParser}
	option2 := []parserFunction{termParser}

	options := [][]parserFunction{option1,option2}
	for _ ,option := range options {
		//print("expression parser Trying option ",i,"\n")
		end ,currentNode ,err = parseSequential(start,option,tokenArray)
		if err == nil {
			//fmt.Println("Expression parser succeeded at option ",i)
			currentNode.name = "expr"
			return end , currentNode ,nil
		}
	}
	//fmt.Println("All options failed, expr parser")
	return -1 ,nil ,errors.New("Failed to parse")
}

func factorParser(start int ,tokenArray  []Token ) (end int,currentNode  *Node   , err error) {
	option1 := []parserFunction{lbParser,exprParser,rbParser}
	option2 := []parserFunction{varParser}
	option3 := []parserFunction{numParser}

	options := [][]parserFunction{option1,option2,option3}

	for _,option := range options {
		//print("factor parser Trying option ",i,"\n")
		end , currentNode ,err = parseSequential(start,option,tokenArray)
		if err == nil {
			//fmt.Println("factor parser succeeded at option ",i)
			currentNode.name = "factor"
			return end , currentNode,nil
		}
	}
	//fmt.Println("All options failed, factor parser")
	return -1 , nil,errors.New("Failed to parse")
}

func termParser(start int ,tokenArray  []Token ) (end int,currentNode *Node , err error) {  
	option1 := []parserFunction{factorParser,multOpParser,termParser}
	option2 := []parserFunction{factorParser}

	options := [][]parserFunction{option1,option2}

	for  _ ,option := range options {
		//print("term parser Trying option ",i,"\n")
		end , currentNode,err = parseSequential(start,option,tokenArray)
		if err == nil {					
			//fmt.Println("term parser succeeded at option ",i)
			currentNode.name = "term"
			return end ,currentNode ,nil
		}
	}
	//println("All options failed, term parser")
	return -1 , nil,errors.New("Failed to parse")
}

func numParser(start int ,tokenArray  []Token ) (end int, currentNode *Node , err error) {  
	if start < len(tokenArray) && tokenArray[start].Type == NUM {
		//println("Number parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"numeric-terminal",[]*Node{}}
		return end ,currentNode ,nil
	} 
	//println("Number parser failed")
	return -1 ,nil ,errors.New("Failed to parse")
}
func varParser(start int ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {  
	if start < len(tokenArray) && tokenArray[start].Type == VAR {
		//println("Variable parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"variable-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	//fmt.Println("Failed to parse a variable")
	return -1 ,nil , errors.New("Failed to parse")  
}

func addOpParser(start int ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {
	if start < len(tokenArray) && tokenArray[start].Type == ADD_OP  {
		//println("Operator parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"addition operator-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}  
	//fmt.Println("Failed to parse an operator")
	return -1 ,nil , errors.New("Failed to parse")  
}

func multOpParser(start int ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {  
	if start < len(tokenArray) && tokenArray[start].Type == MULT_OP {
		//println("Operator parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"multiplication operator-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	//fmt.Println("Failed to parse an operator")
	return -1 ,nil , errors.New("Failed to parse")  

}

func assignParser(start int ,tokenArray  []Token ) (end int,currentNode *Node ,err error)  {
	if start < len(tokenArray)  && tokenArray[start].Type == ASSIGN {
		//println("Assign parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"assign-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	//fmt.Println("Failed to parse an assign")
	return -1 ,nil , errors.New("Failed to parse")  
}

func lbParser(start int ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {
	if start < len(tokenArray) && tokenArray[start].Type == LEFT_BRACKET {
		//println("Left bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"left bracket-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	return -1 ,nil , errors.New("Failed to parse")  	
}

func rbParser(start int ,tokenArray  []Token ) (end int,currentNode *Node ,err error) {
	if start < len(tokenArray) && tokenArray[start].Type == RIGHT_BRACKET {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"right bracket-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	return -1 ,nil , errors.New("Failed to parse")  
	
}


func main()  {
	// x = x + ( 5 * 3 ) + ( 2  * 4)
	// 0 1 2 3 4 5 6 7 8 9 1011 1213 14
	
	tokensArray := []Token{ Token{Type :  VAR , Val : "x" }, Token{Type : ASSIGN , Val : "=" }, Token{Type : VAR , Val : "x" }, Token{Type : ADD_OP , Val : "+" },  Token{Type : LEFT_BRACKET , Val :  "(" }, Token{Type : NUM , Val : "5" }, Token{Type : MULT_OP , Val : "*" }, Token{Type : NUM , Val : "3" }, Token{Type : RIGHT_BRACKET , Val : ")" }, Token{Type : ADD_OP , Val : "+" }, Token{Type: LEFT_BRACKET , Val : "(" }, Token{Type : NUM , Val : "2" }, Token{Type : MULT_OP , Val : "*" }, Token{Type : NUM , Val : "4" }, Token{Type : RIGHT_BRACKET , Val : ")" } }
	
	// print the tokens
	for _,token := range tokensArray {
			println(token.Type , token.Val)
	}
	parseDocument(tokensArray)
}