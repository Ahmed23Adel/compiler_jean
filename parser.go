package main

import (
	"fmt"
)

type parserFunction func(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, node *Node, errParsing error, errSemantic error)

func parseSequential(start int, parsers []parserFunction, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	startPoint := start

	if len(parsers) == 0 {
		currentNode = &Node{-1, -1, "", []*Node{}}
	} else {
		currentNode = &Node{startPoint, startPoint, "", []*Node{}}
		for _, parser := range parsers {
			end, child, errParsing, errSemantic := parser(startPoint, tokenArray, globalSymbolTable)
			if errParsing != nil || errSemantic != nil {
				return -1, nil, errParsing, errSemantic
			}
			startPoint = end
			currentNode.adjacent = append(currentNode.adjacent, child)
			currentNode.end = end
		}
	}
	return startPoint, currentNode, nil, nil
}

func parseDocument(tokenArray []TokenStruct, globalSymbolTable *symbolTable) {
	end, CFG, errParsing, errSemantic := codeParser(0, tokenArray, globalSymbolTable)
	if errParsing != nil || errSemantic != nil {
		fmt.Println("Failed to parse")
		fmt.Println("Returned with error")
		// fmt.Println("errParsing:", errParsing.Error())
		fmt.Println("errSemantic:", errSemantic.Error())
	} else if end != len(tokenArray) {
		PrintGraph(CFG, tokenArray)
		fmt.Println("Failed to parse")
		println("parser finished before the end of the document")
		println("parser ended at ", end, " and document ended at ", len(tokenArray))
	} else {
		//printTree(CFG , tokenArray)
		PrintGraph(CFG, tokenArray)
		println("parser succeeded")
	}

}
