package main

import (
	"fmt"
	"errors"
)

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

	return -1 , nil,errors.New("failed to parse")
	
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
	return -1 , nil,errors.New("failed to parse")
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
	return -1 ,nil ,errors.New("failed to parse")
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
	return -1 , nil,errors.New("failed to parse")
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
	return -1 , nil,errors.New("failed to parse")
}