package main

import (
	"fmt"
	"strings"
	"os"
)



func PrintQuadruplesToFile(quads []Quadruple, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Print the header line
	fmt.Fprintf(file, "%-8s%-8s%-8s%-8s%-8s\n", "Addr", "Op", "Arg1", "Arg2", "Result")

	// Print the quadruples
	for i, quad := range quads {
		fmt.Fprintf(file, "%-8d%-8s%-8s%-8s%-8s\n", i+1, quad.Op, quad.Arg1, quad.Arg2, quad.Result)
	}

	return nil
}

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

	if len(node.Children) == 0 && node.Start < len(tokenArray){
		if node.Start == -1 {
			fmt.Printf("%s ,Contents :  %s \n", node.Type , "None") 
		} else {
			fmt.Printf("%s ,type: %s ,Contents :  '%s' \n", node.Type ,tokenArray[node.Start].Type ,tokenArray[node.Start].Val)
		}
	}else {
		fmt.Printf("%s \n", node.Type ) 
	}
	
	for _, adj := range node.Children {
		fmt.Printf("%s%s\n", strings.Repeat(" ", depth*2), "│")  // old "|"
		fmt.Printf("%s%s", strings.Repeat(" ", depth*2), "├─")  // old "+-"
		printNode(adj, visited, depth+1 , tokenArray)
	}
}


func (root *BinaryNode) Visualize() {

	fmt.Println("Printing binary tree")
	printBinaryNode(root, 0 )
}

func  printBinaryNode(node *BinaryNode, depth int ) {


	
	fmt.Printf("%s \n", node.Value ) 
	
	if node.Left != nil {
		fmt.Printf("%s%s\n", strings.Repeat(" ", depth*2), "│")  // old "|"
		fmt.Printf("%s%s", strings.Repeat(" ", depth*2), "├─")  // old "+-"
		printBinaryNode(node.Left , depth+1 )
	}

	if node.Right != nil {
		fmt.Printf("%s%s\n", strings.Repeat(" ", depth*2), "│")  // old "|"
		fmt.Printf("%s%s", strings.Repeat(" ", depth*2), "├─")  // old "+-"
		printBinaryNode(node.Right , depth+1 )
	}

}

// func printTree(node *Node , tokenArray  []TokenStruct) {
//     // Print the current node
//     print("node name: " , node.name, ",Contents: ("  )
// 	for i := node.start; i < node.end; i++ {
// 		print( tokenArray[i].Val , " ")
// 	}
// 	println(")")
//     // Print the adjacent nodes recursively
// 	println("Children : ")
//     for _, adj := range node.adjacent {
//         printTree(adj , tokenArray)
//     }
// 	println("End of node ",node.name)
// }