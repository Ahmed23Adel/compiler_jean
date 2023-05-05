package main

import (
	"fmt"
	"strings"
)

func PrintGraph(root *Node , tokenArray  []TokenStruct) {
	visited := make(map[*Node]bool)
	fmt.Println("Printing tree")
	printNode(root, visited, 0 ,tokenArray)
}

func printNode(node *Node, visited map[*Node]bool, depth int , tokenArray  []TokenStruct) {
	if visited[node] {
		return
	}
	visited[node] = true

	if len(node.adjacent) == 0 && node.start < len(tokenArray){
		fmt.Printf("%s ,Contents :  %s \n", node.name , tokenArray[node.start].Val) 
	}else {
		fmt.Printf("%s \n", node.name ) 
	}
	
	for _, adj := range node.adjacent {
		fmt.Printf("%s%s\n", strings.Repeat(" ", depth*2), "|")
		fmt.Printf("%s%s", strings.Repeat(" ", depth*2), "+-")
		printNode(adj, visited, depth+1 , tokenArray)
	}
}

func printTree(node *Node , tokenArray  []TokenStruct) {
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