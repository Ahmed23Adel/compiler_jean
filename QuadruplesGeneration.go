package main

import "strconv"

type Quadruple struct {
	Op     string
	Arg1   string
	Arg2   string
	Result string
}

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value string
}

func (n BinaryNode) IsTerminal() bool {
	return n.Left == nil && n.Right == nil
}

func (node *BinaryNode) GetDeepestNonTerminal() *BinaryNode {
	if node == nil || node.IsTerminal() {
		return nil
	}

	queue := []*BinaryNode{node}
	deepest := node

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil && !node.Left.IsTerminal() {
			deepest = node.Left
			queue = append(queue, deepest)
		}
		if node.Right != nil && !node.Right.IsTerminal() {
			deepest = node.Right
			queue = append(queue, deepest)
		}
	}
	return deepest
}

func ExpressionCFG2BinaryTree(CFG *Node, TokenArray []TokenStruct) *BinaryNode {
	parent := &BinaryNode{}
	println(CFG.Type)
	if CFG.IsTerminal() {
		parent.Value = TokenArray[CFG.Start].Val
	} else if len(CFG.Children) == 1 {
		parent = ExpressionCFG2BinaryTree(CFG.Children[0], TokenArray)
	} else if CFG.Children[0].Type == OPEN_PARAN_TERMINAL { // ( expr )

		parent = ExpressionCFG2BinaryTree(CFG.Children[1], TokenArray)
	} else { // term op term .... or some shit
		print(CFG.Children[0].Type , " ")
		print(CFG.Children[1].Type, " ")
		println(CFG.Children[2].Type)
		parent.Value = TokenArray[CFG.Children[1].Start].Val
		parent.Left = ExpressionCFG2BinaryTree(CFG.Children[0], TokenArray)
		parent.Right = ExpressionCFG2BinaryTree(CFG.Children[2], TokenArray)
	}

	return parent
}

func GetStatements(CFG *Node) []*Node {
	stmts := []*Node{CFG.Children[0]}
	parent := CFG.Children[1]
	for len(parent.Children) != 0 {
		stmts = append(stmts, parent.Children[0])
		parent = parent.Children[1]
	}
	return stmts
}

func EvaluateExpression(CFG *BinaryNode) (quads []Quadruple ,lastVar string ) {
	// get statements from CFG
	quads = []Quadruple{}
	if CFG.IsTerminal() {
		lastVar = CFG.Value
	} else {
		tempNum := 0
		// get the deepest non terminal 
		parent := CFG
		
		for parent.GetDeepestNonTerminal() != nil{
			CFG.Visualize()
			deepest := parent.GetDeepestNonTerminal()
			lastVar = "t" + strconv.Itoa(tempNum)
			tempNum ++
			newQuad := Quadruple{Op: deepest.Value , Arg1: deepest.Left.Value , Arg2: deepest.Right.Value , Result: lastVar}
			quads = append(quads, newQuad)

			// replace deepest Node with the new temp
			deepest.Value = lastVar
			deepest.Left = nil
			deepest.Right = nil
		}
		CFG.Visualize()
	}
	return quads , lastVar
}

func GenerateQuads(CFG *Node, TokenArray []TokenStruct) (finalQuads []Quadruple) {
	finalQuads = []Quadruple{}
	stmts := GetStatements(CFG)
	print("Length of statements is")
	println(len(stmts))
	// loop over stmts
	for _, stmt := range stmts {
		// assume all statements are assignments
		// for _, child := range stmt.Children {
		// 	println(child.Type)
		// }
		if len(stmt.Children ) >1  && stmt.Children[1].Type == ASSIGN_TERMINAL {
			variable := TokenArray[stmt.Children[0].Start].Val
			expr := stmt.Children[2] // third is the expression
			binaryTree := ExpressionCFG2BinaryTree(expr, TokenArray)
			binaryTree.Visualize()
			quads , lst := EvaluateExpression(binaryTree)
			finalQuads = append(finalQuads, quads...)
			finalQuads = append(finalQuads, Quadruple{Op: "=" , Arg1: lst , Arg2: "" , Result: variable})
			println("Last var" , lst)
			for _, quad := range quads {
				println("Op:", quad.Op ,"Arg1:", quad.Arg1 ,"Arg2:", quad.Arg2 ,"Result:", quad.Result)
		}
		}
		
	}
	return finalQuads
}