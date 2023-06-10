package main

import "fmt"

func semanticCheck(arraysWithoutSep []TokenStruct) bool{
	varArray := []TokenStruct{}
    funArray := []TokenStruct{}
	//fmt.Println(len(arraysWithoutSep))
	/*for index, token := range arraysWithoutSep {
		//fmt.Println(index)
		if arraysWithoutSep[index].Type == VAR {
			temb := index+1
			if index == len(arraysWithoutSep)-1{
				temb = index 
			}
			if arraysWithoutSep[temb].Type == ASSIGN {
				token.Usage = "initialised"
				varArray = append(varArray, token)
			}else if arraysWithoutSep[temb].Type != ASSIGN && arraysWithoutSep[temb].Type != OPEN_PARAN && arraysWithoutSep[temb].Type != COLON{
				token.Usage = "used"
				varArray = append(varArray, token)
			}
			
		}
		if token.Type == OPEN_CURLY_BRACKET {
			varArray = append(varArray, token)
		}
		if token.Type == CLOSE_CURLY_BRACKET {
			varArray = append(varArray, token)
		}
	}
	
	for index, token := range arraysWithoutSep {
		//fmt.Println(index)
		if arraysWithoutSep[index].Type == VAR {
			temb := index+1
			if index == len(arraysWithoutSep)-1{
				temb = index 
			}
			if index != len(arraysWithoutSep)-1 && arraysWithoutSep[temb].Type == OPEN_PARAN {
				if arraysWithoutSep[temb+2].Type == OPEN_CURLY_BRACKET || arraysWithoutSep[temb+2].Type == COLON {
					token.Usage = "initialised"
				}else if arraysWithoutSep[temb+2].Type == COMMA{
					token.Usage = "used"
				}
				funArray = append(funArray, token)
			}
			
		}
		if token.Type == OPEN_CURLY_BRACKET {
			funArray = append(funArray, token)
		}
		if token.Type == CLOSE_CURLY_BRACKET {
			funArray = append(funArray, token)
		}
	}*/
	
	flag := 0
	depth := 0
	for index, token := range arraysWithoutSep {
		//fmt.Println(index)
		//flag = 0
	    //depth = 0
		if arraysWithoutSep[index].Type == VAR {
			temb := index + 1
			if index == len(arraysWithoutSep)-1 {
				temb = index
			}
			if index != len(arraysWithoutSep)-1 && arraysWithoutSep[temb].Type == OPEN_PARAN {
				if arraysWithoutSep[temb+2].Type == OPEN_CURLY_BRACKET || arraysWithoutSep[temb+2].Type == COLON {
					token.Usage = "initialised"
					flag = 1
				} else if arraysWithoutSep[temb+2].Type == COMMA {
					token.Usage = "used"
				}
				funArray = append(funArray, token)
			}
			if arraysWithoutSep[temb].Type == ASSIGN && flag == 0 {
				token.Usage = "initialised"
				varArray = append(varArray, token)
			} else if arraysWithoutSep[temb].Type != ASSIGN && arraysWithoutSep[temb].Type != OPEN_PARAN && arraysWithoutSep[temb].Type != COLON && flag == 0 {
				token.Usage = "used"
				varArray = append(varArray, token)
			}

		}
		if token.Type == OPEN_CURLY_BRACKET {
			if flag == 1 {
				depth += 1
			}
			varArray = append(varArray, token)
			funArray = append(funArray, token)
		}
		if token.Type == CLOSE_CURLY_BRACKET {
			if flag == 1 {
				depth -= 1
			}
			if depth == 0 {
				flag = 0
			}
			varArray = append(varArray, token)
			funArray = append(funArray, token)
		}
	}

	fmt.Println(".............................................")
	fmt.Println("VarTokens:")
	for _, token := range varArray {
		fmt.Println(token.Type, token.Val, token.Usage)
	}
	fmt.Println("Listed all var tokens and {}")
	fmt.Println(".............................................")

	fmt.Println(".............................................")
	fmt.Println("FunTokens:")
	for _, token := range funArray {
		fmt.Println(token.Type, token.Val, token.Usage)
	}
	fmt.Println("Listed all var tokens and {}")
	fmt.Println(".............................................")

	fmt.Println(len(varArray))
	x := usageBeforeInitialised(varArray)
	unusage(varArray)

	y := usageBeforeInitialised(funArray)
	return (x && y)

}

func usageBeforeInitialised(varArray []TokenStruct) bool {
    no_sem_error := true
	for index, token := range varArray {
		if token.Type == VAR {
			if token.Usage == "used" {
				//fmt.Println(">>>>",token.Val)
				//fmt.Println(token.Val, token.Usage)
				count := 0
				err_flag := 1
				for i := index; i >= 0; i-- {
					if varArray[i].Type == CLOSE_CURLY_BRACKET {
						count += 1
					}
					if varArray[i].Type == OPEN_CURLY_BRACKET {
						tmb := count - 1
						if tmb >= 0 {
							count = tmb
						} else {
							count = 0
						}
					}
					//fmt.Println(varArray[i].Type, varArray[i].Usage, count, varArray[i].Val)
					if varArray[i].Type == VAR && varArray[i].Usage == "initialised" && count == 0 && varArray[i].Val == token.Val {
						err_flag = 0
						//fmt.Println(varArray[i].Usage)
					}
				}
				if err_flag == 1 {
					fmt.Println("Variable ",token.Val," used before being initialised!")
					no_sem_error = false
				}
			}
		}
	}
	return no_sem_error

}

func unusage(varArray []TokenStruct) {

	for index, token := range varArray {
		if token.Type == VAR {
			if token.Usage == "initialised" {
				count := 0
				err_flag := 1
				//flag_for_the_most_inner_scope := 0
				for i := index; i < len(varArray); i++ {
					if varArray[i].Type == CLOSE_CURLY_BRACKET && count == 0 {
						break
						//i = len(varArray)
					}
					if varArray[i].Type == OPEN_CURLY_BRACKET {
						//flag_for_the_most_inner_scope = 1
						count += 1
					}
					if varArray[i].Type == CLOSE_CURLY_BRACKET && count != 0 {
						//flag_for_the_most_inner_scope = 0
						tmb := count - 1
						if tmb >= 0 {
							count = tmb
						} else {
							break
							//i = len(varArray)
						}
					}
					//fmt.Println(varArray[i].Type, varArray[i].Usage, count, varArray[i].Val)
					if varArray[i].Type == VAR && varArray[i].Usage == "used" && count >= 0 && varArray[i].Val == token.Val {
						err_flag = 0
					}
				}
				//fmt.Println(token.Type, token.Usage, token.Val)
				if err_flag == 1 {
					fmt.Println("Unused variable of ",token.Val)
				}

			}
		}
	}

}


func eachFuncDic(varArray []TokenStruct) {

	

}
