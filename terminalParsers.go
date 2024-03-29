package main

import (
	"errors"
)

// numParser: checks if the token at the start index is a number, if it is it appends to the tree
func numParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (NUMBER) || tokenArray[start].Type == (FLOAT) {
		//println("Number parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "numeric-terminal", []*Node{}}
		return end, currentNode, nil
	}
	//println("Number parser failed")
	return -1, nil, errors.New("failed to parse")
}

// varParser: checks if the token at the start index is a variable, if it is it appends to the tree
func varParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (VAR) {
		//println("Variable parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "variable-terminal", []*Node{}}
		return end, currentNode, nil
	}
	//fmt.Println("failed to parse a variable")
	return -1, nil, errors.New("failed to parse")
}

// OpLevel0Parser: checks if the token at the start index is an addition operator, if it is it appends to the tree
func OpLevel0Parser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && (tokenArray[start].Type == (ADD) || tokenArray[start].Type == (SUB) || tokenArray[start].Type == COMP ||tokenArray[start].Type == (GT) || tokenArray[start].Type == (LT) || tokenArray[start].Type == (GTE) || tokenArray[start].Type == (LTE))    {
		//println("Operator parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "addition level operator-terminal", []*Node{}}
		return end, currentNode, nil
	}
	//fmt.Println("failed to parse an operator")
	return -1, nil, errors.New("failed to parse")
}

// multOpParser: checks if the token at the start index is a multiplication operator, if it is it appends to the tree
func OpLevel1Parser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && (tokenArray[start].Type == (MUL) || tokenArray[start].Type == DIV) {
		//println("Operator parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "multiplication level operator-terminal", []*Node{}}
		return end, currentNode, nil
	}
	//fmt.Println("failed to parse an operator")
	return -1, nil, errors.New("failed to parse")

}

// Terminal parser for the assignment operator "="
func assignParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (ASSIGN) {
		//println("Assign parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, ASSIGN_TERMINAL, []*Node{}}
		return end, currentNode, nil
	}
	//fmt.Println("failed to parse an assign")
	return -1, nil, errors.New("failed to parse")
}

func openParanParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (OPEN_PARAN) {
		//println("Left bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, OPEN_PARAN_TERMINAL, []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}

func closedParanParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (CLOSE_PARAN) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, CLOSED_PARAN_TERMINAL, []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")

}

func openCurlyBracketParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (OPEN_CURLY_BRACKET) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "left curly-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")

}

func closeCurlyBracketParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (CLOSE_CURLY_BRACKET) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "right curly-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")

}

func colonParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (COLON) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "colon-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}

func varTypeParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && (tokenArray[start].Type == (INT) || tokenArray[start].Type == (FLT) || tokenArray[start].Type == (STR)) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "dtype-intOrStrOrFlt-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}

func commaParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (COMMA) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "comma-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}

func questionMarkParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (QUESTION_MARK) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, QUESTION_MARK_TERMINAL, []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}

func excMarkParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (EXCLAMATION_MARK) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "exclamation mark-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}

func astParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	if start < len(tokenArray) && tokenArray[start].Type == (MUL) {
		//println("Right bracket parser succeeded")
		end = start + 1
		currentNode = &Node{start, end, "AST-terminal", []*Node{}}
		return end, currentNode, nil
	}
	return -1, nil, errors.New("failed to parse")
}
