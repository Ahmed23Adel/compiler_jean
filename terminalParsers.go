package main

import (
	"errors"
)

// numParser: checks if the token at the start index is a number, if it is it appends to the tree
func numParser(start int ,tokenArray  []TokenStruct ) (end int, currentNode *Node , err error) {  
	if start < len(tokenArray) && tokenArray[start].Type ==  (NUMBER)  || tokenArray[start].Type ==  (FLOAT)   {
		//println("Number parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"numeric-terminal",[]*Node{}}
		return end ,currentNode ,nil
	} 
	//println("Number parser failed")
	return -1 ,nil ,errors.New("failed to parse")
}

// varParser: checks if the token at the start index is a variable, if it is it appends to the tree
func varParser(start int ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error) {  
	if start < len(tokenArray) && tokenArray[start].Type ==  (VAR) {
		//println("Variable parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"variable-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	//fmt.Println("failed to parse a variable")
	return -1 ,nil , errors.New("failed to parse")  
}


// addOpParser: checks if the token at the start index is an addition operator, if it is it appends to the tree
func addOpParser(start int ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error) {
	if start < len(tokenArray) &&( tokenArray[start].Type ==  (ADD) ||  tokenArray[start].Type ==  (SUB)) {  
		//println("Operator parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"addition operator-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}  
	//fmt.Println("failed to parse an operator")
	return -1 ,nil , errors.New("failed to parse")  
}

// multOpParser: checks if the token at the start index is a multiplication operator, if it is it appends to the tree
func multOpParser(start int ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error) {  
	if start < len(tokenArray) && (tokenArray[start].Type ==  (MUL) ||  tokenArray[start].Type ==  DIV  ){
		//println("Operator parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"multiplication operator-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	//fmt.Println("failed to parse an operator")
	return -1 ,nil , errors.New("failed to parse")  

}

func assignParser(start int ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error)  {
	if start < len(tokenArray)  && tokenArray[start].Type ==  (ASSIGN) {
		//println("Assign parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"assign-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	//fmt.Println("failed to parse an assign")
	return -1 ,nil , errors.New("failed to parse")  
}

func openParanParser(start int ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error) {
	if start < len(tokenArray) && tokenArray[start].Type ==  (OPEN_PARAN) {
		//println("Left bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"left bracket-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	return -1 ,nil , errors.New("failed to parse")  	
}

func closedParanParser(start int ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error) {
	if start < len(tokenArray) && tokenArray[start].Type ==  (CLOSE_PARAN) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start,end,"right bracket-terminal",[]*Node{}}
		return end ,currentNode ,nil
	}
	return -1 ,nil , errors.New("failed to parse")  
	
}