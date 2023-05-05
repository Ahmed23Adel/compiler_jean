package main

import (
	"errors"
	//"fmt"
)

func codeParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {

	option1 := []parserFunction{stmtParser, codeParser} // code --> stmt code | None
	option2 := []parserFunction{}

	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("Code parser succeeded")
			currentNode.name = "code"
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, code parser")

	return -1, nil, errors.New("failed to parse")

}

func stmtParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{varParser, assignParser, exprParser} // stmt --> var = expr  sep
	option2 := []parserFunction{ifElseParser}                        // stmt --> ifElse
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("Statement parser succeeded")
			currentNode.name = "stmt"
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, stmt parser")
	return -1, nil, errors.New("failed to parse")
}

func exprParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{termParser, addOpParser, exprParser} // expr --> term add_op expr
	option2 := []parserFunction{termParser}                          // expr --> term

	options := [][]parserFunction{option1, option2}
	for _, option := range options {
		//print("expression parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("Expression parser succeeded at option ",i)
			currentNode.name = "expr"
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, expr parser")
	return -1, nil, errors.New("failed to parse")
}

func factorParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{openParanParser, exprParser, closedParanParser} // factor --> ( expr )
	option2 := []parserFunction{varParser}                                      // factor --> var
	option3 := []parserFunction{numParser}                                      // factor --> num

	options := [][]parserFunction{option1, option2, option3}

	for _, option := range options {
		//print("factor parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("factor parser succeeded at option ",i)
			currentNode.name = "factor"
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, factor parser")
	return -1, nil, errors.New("failed to parse")
}

func termParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{factorParser, multOpParser, termParser} // term --> factor mult_op term
	option2 := []parserFunction{factorParser}                           // term --> factor

	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		//print("term parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("term parser succeeded at option ",i)
			currentNode.name = "term"
			return end, currentNode, nil
		}
	}
	//println("All options failed, term parser")
	return -1, nil, errors.New("failed to parse")
}

func ifElseParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{openParanParser, exprParser, closedParanParser, questionMarkParser,
		openCurlyBracketParser, codeParser, closeCurlyBracketParser, elseParser} // ifElse --> ( expr )? {code} else  # if
	options := [][]parserFunction{option1}
	for _, option := range options {
		//print("expression parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {

			currentNode.name = "if stmt"
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, expr parser")
	return -1, nil, errors.New("failed to parse")
}

func elseParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		excMarkParser,
		openParanParser,
		exprParser,
		closedParanParser,
		questionMarkParser,
		openCurlyBracketParser,
		codeParser,
		closeCurlyBracketParser,
		elseParser}
	// else --> !(expr)? {code} else   # else if

	option2 := []parserFunction{
		excMarkParser,
		openCurlyBracketParser,
		codeParser,
		closeCurlyBracketParser} //else --> !{code}               # only else
	option3 := []parserFunction{}
	options := [][]parserFunction{option1, option2, option3}

	for _, option := range options {
		//print("term parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("term parser succeeded at option ",i)
			currentNode.name = "else"
			return end, currentNode, nil
		}
	}
	//println("All options failed, term parser")
	return -1, nil, errors.New("failed to parse")

}
