package main

import (
	"errors"
	//"fmt"
)

func codeParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {

	option1 := []parserFunction{stmtParser, codeParser} // code --> stmt code | None
	option2 := []parserFunction{}

	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errSemantic != nil {
			return -1, nil, nil, errSemantic
		}
		if errParsing == nil {
			//fmt.Println("Code parser succeeded")
			currentNode.name = "code"
			return end, currentNode, errParsing, errSemantic
		}
	}
	//fmt.Println("All options failed, code parser")

	return -1, nil, errors.New("failed to parse"), nil

}

func stmtParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{funDeclParser} // stmt --> inline_fun_decl | fun_decl
	option2 := []parserFunction{inlineFunDeclParser}
	option3 := []parserFunction{assignmentStmtParser}     // stmt --> var = expr  sep
	option4 := []parserFunction{functionHeaderBodyParser} // stmt --> ( expr )? {code} else  # if
	option5 := []parserFunction{loopParser}
	option6 := []parserFunction{ifElseParser} // stmt --> ifElse
	options := [][]parserFunction{option1, option2, option3, option4, option5, option6}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errSemantic != nil {
			return -1, nil, nil, errSemantic
		}
		if errParsing == nil {
			//fmt.Println("Statement parser succeeded")
			currentNode.name = "stmt"
			return end, currentNode, nil, nil
		}
	}
	//fmt.Println("All options failed, stmt parser")
	return -1, nil, errParsing, errSemantic
}

func assignmentStmtParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{varParser, assignParser, exprParser} // stmt --> var = expr  sep                                                                                       // expr --> term
	// errorString := ""
	options := [][]parserFunction{option1}
	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "assignmentStmt"
			// TODO: y= expr you must calc output tpye of this expression
			err2 := pushVariableIfPossible(tokenArray[start].Val, dtypeStruct{dtype: string(tokenArray[end-1].Type)}, OTHER_FREE, globalSymbolTable) // z = x + y // z if not in symbol table insert if it's is then do nothing // x, y must be in symbol table
			for i := end - 1; i >= start+2; i-- {
				if tokenArray[i].Type == VAR && !isVariableExistInSymbolTable(tokenArray[i].Val, globalSymbolTable) {
					return -1, nil, nil, errors.New("variable " + tokenArray[i].Val + " not found in symbol table. Used before declaration")
				}
			}
			if err2 == nil {
				// errorString = err2.Error()
				return end, currentNode, nil, nil
			}

			return -1, nil, nil, err2

		}
	}
	return -1, nil, errors.New("failed to parse"), nil
}

func functionHeaderBodyParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{openParanParser, exprParser, closedParanParser, questionMarkParser,
		openCurlyBracketParser, codeParser, closeCurlyBracketParser, elseParser} // stmt --> ( expr )? {code} else  # if                                                                                  // expr --> term
	// errorString := ""
	options := [][]parserFunction{option1}
	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "assignmentStmt"
			// err2 := pushFunctionIfPossible(tokenArray[start].Val, symbolTableRowDtype(tokenArray[start].Type), OTHER_FREE, globalSymbolTable) // z = x + y // z if not in symbol table insert if it's is then do nothing // x, y must be in symbol table
			// for i := end - 1; i >= start+2; i-- {
			// 	if tokenArray[i].Type == VAR && !isVariableExistInSymbolTable(tokenArray[i].Val, globalSymbolTable) {
			// 		return -1, nil, nil, errors.New("variable " + tokenArray[i].Val + " not found in symbol table. Used before declaration")
			// 	}
			// }
			// if err2 == nil {
			// 	// errorString = err2.Error()
			// 	return end, currentNode, nil, nil
			// }

			// return -1, nil, nil, err2

		}
	}
	return -1, nil, errors.New("failed to parse"), nil
}

func exprParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{termParser, addOpParser, exprParser} // expr --> term add_op expr
	option2 := []parserFunction{termParser}                          // expr --> term

	options := [][]parserFunction{option1, option2}
	for _, option := range options {
		//print("expression parser Trying option ",i,"\n")
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			//fmt.Println("Expression parser succeeded at option ",i)
			currentNode.name = "expr"
			return end, currentNode, nil, nil
		}
	}
	//fmt.Println("All options failed, expr parser")
	return -1, nil, errors.New("failed to parse"), nil
}

func factorParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{openParanParser, exprParser, closedParanParser} // factor --> ( expr )
	option2 := []parserFunction{funCallParser}                                  // factor --> funcall
	option3 := []parserFunction{varParser}                                      // factor --> var
	option4 := []parserFunction{numParser}                                      // factor --> num

	options := [][]parserFunction{option1, option2, option3, option4}

	for _, option := range options {
		//print("factor parser Trying option ",i,"\n")
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			//fmt.Println("factor parser succeeded at option ",i)
			currentNode.name = "factor"
			return end, currentNode, nil, nil
		}
	}
	//fmt.Println("All options failed, factor parser")
	return -1, nil, errors.New("failed to parse"), nil
}

func termParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{factorParser, multOpParser, termParser} // term --> factor mult_op term
	option2 := []parserFunction{factorParser}                           // term --> factor

	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		//print("term parser Trying option ",i,"\n")
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			//fmt.Println("term parser succeeded at option ",i)
			currentNode.name = "term"
			return end, currentNode, nil, nil
		}
	}
	//println("All options failed, term parser")
	return -1, nil, errors.New("failed to parse"), nil
}

func ifElseParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{openParanParser, exprParser, closedParanParser, questionMarkParser,
		openCurlyBracketParser, codeParser, closeCurlyBracketParser, elseParser} // ifElse --> ( expr )? {code} else  # if
	options := [][]parserFunction{option1}
	for _, option := range options {
		//print("expression parser Trying option ",i,"\n")
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {

			currentNode.name = "if stmt"
			return end, currentNode, nil, nil
		}
	}
	//fmt.Println("All options failed, expr parser")
	return -1, nil, errors.New("failed to parse"), nil
}

func elseParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
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
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			//fmt.Println("term parser succeeded at option ",i)
			currentNode.name = "else"
			return end, currentNode, nil, nil
		}
	}
	//println("All options failed, term parser")
	return -1, nil, errors.New("failed to parse"), nil

}

func loopParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
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
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "loop"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}
func funCallParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
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
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "funcall"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}

func argsParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{
		arg1Parser,
		arg2Parser}
	option2 := []parserFunction{}
	// args --> none | arg1 arg2
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "args"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}

func arg1Parser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{
		varParser,
		colonParser,
		varTypeParser}
	// arg1 --> var : dtype
	options := [][]parserFunction{option1}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "arg1"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}

func arg2Parser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
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
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "arg2"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}
func inlineFunDeclParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{varParser, openParanParser, argsParser, closedParanParser, colonParser, varTypeParser, assignParser, exprParser}
	option2 := []parserFunction{varParser, openParanParser, argsParser, closedParanParser, assignParser, exprParser}
	// inline_fun_decl --> var (args): ret = expr |  var (args) = expr
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "inline_fun_decl"
			argsEnd := false
			var args dtypeStructList
			idxReturn := -1
			for i := start + 2; i < end; i++ {
				// TODO don't put it in dtype but put in the new referenced symbol table and make dtype nil for function
				if !argsEnd {
					// I will put return at end of array
					// if name and val are nil they it return nothing
					if tokenArray[i].Type == VAR {
						args.list = append(args.list, dtypeStruct{name: tokenArray[i].Val, dtype: tokenArray[i+2].Val})
						i += 1
					}
					if tokenArray[i].Type == CLOSE_PARAN {
						argsEnd = true
						idxReturn = i + 2
						break

					}
				}
			}
			args.list = append(args.list, dtypeStruct{dtype: tokenArray[idxReturn].Val})
			err2 := pushFunctionIfPossible(tokenArray[start].Val, args, OTHER_FREE, globalSymbolTable) // z = x + y // z if not in symbol table insert if it's is then do nothing // x, y must be in symbol table
			if err2 == nil {
				// errorString = err2.Error()
				return end, currentNode, nil, nil
			}

			return -1, nil, nil, err2

		}

	}
	return -1, nil, errors.New("failed to parse"), nil

}

func funDeclParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
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
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "fun_decl"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}

func argsForCallParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{
		arg1ForCallParser,
		arg2ForCallParser}
	option2 := []parserFunction{}
	// args --> none | arg1_expr_  arg2_expr_
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "args"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}

func arg1ForCallParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{
		exprParser}
	// arg1 --> expr
	options := [][]parserFunction{option1}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "arg1"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}

func arg2ForCallParser(start int, tokenArray []TokenStruct, globalSymbolTable *symbolTable) (end int, currentNode *Node, errParsing error, errSemantic error) {
	option1 := []parserFunction{
		commaParser,
		exprParser,
		arg2ForCallParser}
	option2 := []parserFunction{}
	// arg2 --> , var arg2 | None
	options := [][]parserFunction{option1, option2}

	for _, option := range options {
		end, currentNode, errParsing, errSemantic = parseSequential(start, option, tokenArray, globalSymbolTable)
		if errParsing == nil && errSemantic == nil {
			currentNode.name = "arg2"
			return end, currentNode, nil, nil
		}
	}
	return -1, nil, errors.New("failed to parse"), nil

}
