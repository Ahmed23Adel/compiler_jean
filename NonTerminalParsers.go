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
			currentNode.Type = CODE_NON_TERMINAL
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, code parser")

	return -1, nil, errors.New("failed to parse")

}

func stmtParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{funDeclParser} // stmt --> inline_fun_decl | fun_decl
	option2 := []parserFunction{inlineFunDeclParser}
	option3 := []parserFunction{varParser, assignParser, exprParser} // stmt --> var = expr  sep
	option4 := []parserFunction{openParanParser, exprParser, closedParanParser, questionMarkParser,
		openCurlyBracketParser, codeParser, closeCurlyBracketParser, elseParser} // stmt --> ( expr )? {code} else  # if
	option5 := []parserFunction{loopParser}
	option6 := []parserFunction{ifElseParser} // stmt --> ifElse
	options := [][]parserFunction{option1, option2, option3, option4, option5, option6}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("Statement parser succeeded")
			currentNode.Type = STATEMENT_NON_TERMINAL
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, stmt parser")
	return -1, nil, errors.New("failed to parse")
}

func exprParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{termParser, OpLevel0Parser, exprParser} // expr --> term add_op expr
	option2 := []parserFunction{termParser}                          // expr --> term

	options := [][]parserFunction{option1, option2}
	for _, option := range options {
		//print("expression parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("Expression parser succeeded at option ",i)
			currentNode.Type = EXPRESSION_NON_TERMINAL
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, expr parser")
	return -1, nil, errors.New("failed to parse")
}

func factorParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{openParanParser, exprParser, closedParanParser} // factor --> ( expr )
	option2 := []parserFunction{funCallParser}                                  // factor --> funcall
	option3 := []parserFunction{varParser}                                      // factor --> var
	option4 := []parserFunction{numParser}                                      // factor --> num

	options := [][]parserFunction{option1, option2, option3, option4}

	for _, option := range options {
		//print("factor parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("factor parser succeeded at option ",i)
			currentNode.Type = FACTOR_NON_TERMINAL
			return end, currentNode, nil
		}
	}
	//fmt.Println("All options failed, factor parser")
	return -1, nil, errors.New("failed to parse")
}

func termParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{factorParser, OpLevel1Parser, termParser} // term --> factor mult_op term
	option2 := []parserFunction{factorParser}                           // term --> factor

	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		//print("term parser Trying option ",i,"\n")
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			//fmt.Println("term parser succeeded at option ",i)
			currentNode.Type = TERM_NON_TERMINAL
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

			currentNode.Type = CONDITIONAL_STATEMENT_NON_TERMINAL
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
			currentNode.Type = ELSE_NON_TERMINAL
			return end, currentNode, nil
		}
	}
	//println("All options failed, term parser")
	return -1, nil, errors.New("failed to parse")

}

func loopParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		varParser,
		colonParser,
		astParser,
		openParanParser,
		exprParser,
		closedParanParser,
		questionMarkParser,
		openCurlyBracketParser,
		codeParser,
		closeCurlyBracketParser}
	// loop --> var : *(expr)? { code }
	options := [][]parserFunction{option1}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = LOOP_STATEMENT_NON_TERMINAL
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}
func funCallParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		varParser,
		openParanParser,
		closedParanParser}
	option2 := []parserFunction{
		varParser,
		openParanParser,
		argsForCallParser,
		closedParanParser}
	// funcallParser --> var() | var(args_expr)
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "funcall"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func argsParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		arg1Parser,
		arg2Parser}
	option2 := []parserFunction{}
	// args --> none | arg1 arg2
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "args"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func arg1Parser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		varParser,
		colonParser,
		varTypeParser}
	// arg1 --> var : dtype
	options := [][]parserFunction{option1}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "arg1"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func arg2Parser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		commaParser,
		varParser,
		colonParser,
		varTypeParser,
		arg2Parser}
	option2 := []parserFunction{}
	// arg2 --> , var : dtype : arg2 | None
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "arg2"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}
func inlineFunDeclParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		varParser,
		openParanParser,
		argsParser,
		closedParanParser,
		colonParser,
		varTypeParser,
		assignParser,
		exprParser}
	option2 := []parserFunction{varParser,
		openParanParser,
		argsParser,
		closedParanParser,
		assignParser,
		exprParser}
	// inline_fun_decl --> var (args): ret = expr |  var (args) = expr
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "inline_fun_decl"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func funDeclParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		varParser,
		openParanParser,
		argsParser,
		closedParanParser,
		colonParser,
		varTypeParser,
		openCurlyBracketParser,
		codeParser,
		closeCurlyBracketParser}
	option2 := []parserFunction{
		varParser,
		openParanParser,
		argsParser,
		closedParanParser,
		openCurlyBracketParser,
		codeParser,
		closeCurlyBracketParser}
	// fun_decl --> var (args): ret {code} | var (args) {code}
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "fun_decl"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func argsForCallParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		arg1ForCallParser,
		arg2ForCallParser}
	option2 := []parserFunction{}
	// args --> none | arg1_expr_  arg2_expr_
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "args"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func arg1ForCallParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		exprParser}
	// arg1 --> expr
	options := [][]parserFunction{option1}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "arg1"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}

func arg2ForCallParser(start int, tokenArray []TokenStruct) (end int, currentNode *Node, err error) {
	option1 := []parserFunction{
		commaParser,
		exprParser,
		arg2ForCallParser}
	option2 := []parserFunction{}
	// arg2 --> , var arg2 | None
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, err = parseSequential(start, option, tokenArray)
		if err == nil {
			currentNode.Type = "arg2"
			return end, currentNode, nil
		}
	}
	return -1, nil, errors.New("failed to parse")

}
