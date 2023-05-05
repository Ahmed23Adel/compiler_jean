package main

import (
	"errors"
	"fmt"
)

type parserFunction func(start int ,tokenArray  []tokenStruct ) (end int,node *Node ,err error)

func parseSequential(start int ,parsers []parserFunction ,tokenArray  []tokenStruct ) (end int,currentNode *Node ,err error) {
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


func parseDocument(tokenArray []tokenStruct)  {  
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