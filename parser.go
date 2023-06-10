package main

import (
	"errors"
	"fmt"
)

type parserFunction func(start int ,tokenArray  []TokenStruct ) (end int,node *Node ,err error)

func parseSequential(start int ,parsers []parserFunction ,tokenArray  []TokenStruct ) (end int,currentNode *Node ,err error) {
	startPoint := start

	if len(parsers) == 0 {
		currentNode = &Node{-1,-1,"",[]*Node{}}
	} else {
		currentNode = &Node{startPoint,startPoint,"",[]*Node{}}
		for _,parser := range parsers {
			end ,child ,err := parser(startPoint,tokenArray)
			if err != nil {
				return -1 , nil ,errors.New("failed to parse")
			}
			startPoint = end
			currentNode.Children = append(currentNode.Children,child)
			currentNode.End = end
		}
	}
	return startPoint , currentNode ,nil
}


func parseDocument(tokenArray []TokenStruct)  {  
	end ,CFG ,err := codeParser(0,tokenArray)
	if err != nil   {
		fmt.Println("Failed to parse")
		fmt.Println("Returned with error")
	} else if end != len(tokenArray) {
		PrintGraph(CFG, tokenArray)
		fmt.Println("Failed to parse")
		println("parser finished before the end of the document")
		println("parser ended at ",end , " and document ended at ",len(tokenArray))
	}else {
		//printTree(CFG , tokenArray)
		PrintGraph(CFG, tokenArray)
		println("parser succeeded")
	}

	
}