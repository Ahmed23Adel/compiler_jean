package main

import (
	"strconv"
)

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

func GetStatement(CFG *Node) (stmt *Node , next *Node) {
	stmt = nil
	if  len(CFG.Children) != 0 {
		stmt = CFG.Children[0]
		next = CFG.Children[1]
	}
	return stmt ,next
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

func EvaluateStatement(stmt *Node , TokenArray []TokenStruct , ) (finalQuads []Quadruple ) {
	finalQuads = []Quadruple{}
	if len(stmt.Children ) >1  && stmt.Children[1].Type == ASSIGN_TERMINAL {  // assignment statement
		variable := TokenArray[stmt.Children[0].Start].Val
		expr := stmt.Children[2] // third is the expression
		binaryTree := ExpressionCFG2BinaryTree(expr, TokenArray)
		binaryTree.Visualize()
		qs , lst := EvaluateExpression(binaryTree)
		finalQuads = append(qs, Quadruple{Op: "=" , Arg1: lst , Arg2: "" , Result: variable})
	} else if len(stmt.Children) >3  && stmt.Children[0].Type == OPEN_PARAN_TERMINAL && stmt.Children[3].Type == QUESTION_MARK_TERMINAL {
		expr := stmt.Children[1] // second is the expression
		codeIfTrue := stmt.Children[5]
		elseStmt := stmt.Children[7]
		
		binaryTree := ExpressionCFG2BinaryTree(expr, TokenArray)
		condition_quads , lst := EvaluateExpression(binaryTree)

		finalQuads = append(finalQuads , condition_quads...)
		true_quads := EvaluateCode(codeIfTrue , TokenArray)
		else_pos :=  len(true_quads) +1
		
		if len(elseStmt.Children) !=0 {
			codeIfFalse := elseStmt.Children[2]
			false_quads := EvaluateCode(codeIfFalse,TokenArray)
			else_pos +=1
			finalQuads = append(finalQuads, Quadruple{Op: "JUMP_IF_NOT" , Arg1: lst , Arg2: "" , Result: strconv.Itoa(else_pos)})
			finalQuads = append(finalQuads, true_quads...)
			finalQuads = append(finalQuads, Quadruple{Op: "JUMP" , Arg1: lst , Arg2: "" , Result: strconv.Itoa( len(false_quads)+1)})
			finalQuads = append(finalQuads, false_quads...)
		}else {
			finalQuads = append(finalQuads, Quadruple{Op: "JUMP_IF_NOT" , Arg1: lst , Arg2: "" , Result: strconv.Itoa(else_pos)})
			finalQuads = append(finalQuads, true_quads...)
		}

		
	} else if len(stmt.Children) >0 && stmt.Children[0].Type == "fun_decl" {

		if len(stmt.Children[0].Children) > 1 && stmt.Children[0].Children[1].Type == "(" {
		  // I should get to point after  {
		  // and write code up till }
		  //fmt.Println("stmtfUNC", stmt)
  
		  finalQuads = append(finalQuads, Quadruple{Op: "Label function" + TokenArray[stmt.Children[0].Children[0].Start].Val, Arg1: "", Arg2: "", Result: ""})
		  // (*(*(*CFG).Children[0]).Children[0]).Children[7].Start = 0
		  // finalQuads = GenerateQuads((*(*(*CFG).Children[0]).Children[0]).Children[7], TokenArray, finalQu 
		  fnCodeQuads := EvaluateCode(stmt.Children[0].Children[7] , TokenArray)
		  
		  finalQuads = append(finalQuads, fnCodeQuads...)
		  finalQuads = append(finalQuads, Quadruple{Op: "MOV", Arg1: "LR", Arg2: "", Result: "PC"})
  
	  }
	} else if len(stmt.Children) > 1 && stmt.Children[0].Type == LOOP_STATEMENT_NON_TERMINAL {
		loop := stmt.Children[0]
		expr := loop.Children[3] 
		codeIfTrue := loop.Children[8]
		binaryTree := ExpressionCFG2BinaryTree(expr, TokenArray)
		condition_quads , lst := EvaluateExpression(binaryTree)
		finalQuads = append(finalQuads , condition_quads...)
		true_quads := EvaluateCode(codeIfTrue , TokenArray)
		finalQuads = append(finalQuads, Quadruple{Op: "JUMP_IF_NOT" , Arg1: lst , Arg2: "" , Result: strconv.Itoa(len(true_quads)+1)})
		finalQuads = append(finalQuads, true_quads...)
		finalQuads = append(finalQuads, Quadruple{Op: "JUMP" , Arg1: lst , Arg2: "" , Result: strconv.Itoa( -len(true_quads))})
	}
	return finalQuads
}

func EvaluateCode(CFG *Node, TokenArray []TokenStruct   ) (finalQuads []Quadruple) {
	finalQuads = []Quadruple{}
	stmt , CFG := GetStatement(CFG)

	// loop over stmts
	for stmt != nil  {
		
		quads := EvaluateStatement(stmt , TokenArray )
		finalQuads = append(finalQuads, quads...)

		stmt , CFG = GetStatement(CFG)
	}
		
	
	return finalQuads
}

func Jump2Goto(quads []Quadruple) []Quadruple {
	for i , quad := range quads {
		if quad.Op == "JUMP" {
			quads[i].Op = "GOTO" 
			converted , _ := strconv.Atoi(quad.Result)
			target := i + converted
			quads[i].Result = strconv.Itoa(target)
		} else if quad.Op == "JUMP_IF_NOT" {
			quads[i].Op = "GOTO_IF_NOT" 
			converted , _ := strconv.Atoi(quad.Result)
			target := i + converted
			quads[i].Result = strconv.Itoa(target)
		}
	}
	return quads
}