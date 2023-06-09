package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
)

func Lexer(filename string) []TokenStruct {
	// Read the file
	fileContent := readSample(filename) + "\n"
	tokenArray := lex_analyzer(fileContent)
	return tokenArray
}

func readSample(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	return fileContent
}

type Position struct {
	line   int
	column int
}

func getPossilbeTerminals() []TokenStruct {
	return []TokenStruct{TokenStruct{Type: RETURN, Val: "return"}, TokenStruct{Type: BREAK, Val: "break"}, TokenStruct{Type: CONINUE, Val: "continue"},
		TokenStruct{Type: FLOAT, Val: "float"}, TokenStruct{Type: INT, Val: "int"}, TokenStruct{Type: CHAR, Val: "char"},
		TokenStruct{Type: AND, Val: "and"}, TokenStruct{Type: OR, Val: "or"}, TokenStruct{Type: NOT, Val: "not"},
		TokenStruct{Type: ABS, Val: "abs"}, TokenStruct{Type: SEPARATOR, Val: "\n"}, TokenStruct{Type: SPACE, Val: " "},
		TokenStruct{Type: COMP, Val: "=="}, TokenStruct{Type: GTE, Val: ">="}, TokenStruct{Type: LTE, Val: "<="},
		TokenStruct{Type: ADD, Val: "+"}, TokenStruct{Type: SUB, Val: "-"}, TokenStruct{Type: MUL, Val: "*"},
		TokenStruct{Type: DIV, Val: "/"}, TokenStruct{Type: POWER, Val: "^"}, TokenStruct{Type: BIT_OR, Val: "|"},
		TokenStruct{Type: BIT_AND, Val: "&"}, TokenStruct{Type: OPEN_PARAN, Val: "("}, TokenStruct{Type: CLOSE_PARAN, Val: ")"},
		TokenStruct{Type: QUESTION_MARK, Val: "?"}, TokenStruct{Type: OPEN_CURLY_BRACKET, Val: "{"}, TokenStruct{Type: CLOSE_CURLY_BRACKET, Val: "}"},
		TokenStruct{Type: EXCLAMATION_MARK, Val: "!"}, TokenStruct{Type: COLON, Val: ":"}, TokenStruct{Type: ASSIGN, Val: "="},
		TokenStruct{Type: GT, Val: ">"}, TokenStruct{Type: LT, Val: "<"}, TokenStruct{Type: COMMA, Val: ","}}
}

func lex_analyzer(input string) []TokenStruct {
	tokens := make([]TokenStruct, 0)
	//
	possibleTerminals := getPossilbeTerminals()
	sort.Slice(possibleTerminals, func(i, j int) bool {
		return len(possibleTerminals[i].Val) > len(possibleTerminals[j].Val)
	})
	maxLen := len(possibleTerminals[0].Val) - 1
	current_pos := Position{line: 0, column: 0}
	for i := 0; i < len(input); i++ {
		found := false
		for lookahead := maxLen; lookahead >= 0; lookahead-- {
			if i+lookahead < len(input) {
				if lookahead > 1 && handleComments(&input, &i, &current_pos, &found) {
					break
				}
				for j := 0; j < len(possibleTerminals); j++ {
					if handleMultiLetterToken(&input, &i, &current_pos, possibleTerminals[j].Val, &tokens, possibleTerminals[j].Type) {
						found = true
						break
					}

				}
			}
			if found {
				break
			}
		}
		if found {
			continue
		}
		if isInt(string(input[i])) {
			handleRealNumbers(&i, &tokens, &input, &current_pos)
			continue

		} else if (string(input[i]) == "'" && string(input[i+2]) == "'") && isChar(string(input[i+1])) {
			handleChar(&current_pos, &input, &i, &tokens)
			continue
		} else {
			handleIdentifier(&i, &input, &tokens, &current_pos)
			continue
		}

	}

	return tokens
}

func handleMultiLetterToken(input *string, idx *int, current_pos *Position, searchable string, tokens *[]TokenStruct, token Token) bool {
	if string((*input)[*idx]) == " " {  // 
		forwardPosOneSpace(current_pos)
		return true
	}
	if *idx+len(searchable) < len(*input)-1 {
		if string((*input)[(*idx):(*idx)+len(searchable)]) == searchable {
			(*tokens) = append((*tokens), TokenStruct{Type: (token), Val: searchable, Pos: *current_pos})
			(*current_pos).column += len(searchable)
			*idx += len(searchable) - 1
			return true
		}

	}
	return false
}

func handleOneLineComment(input *string, idx *int, current_pos *Position) {
	*idx = findFirstRune((*input), (*idx), '\n') - 1
}

func handleBlockComment(input *string, idx *int, current_pos *Position) {
	var numNewLines, distFromLastNewLine int
	*idx, numNewLines, distFromLastNewLine = findFirstStr(*input, *idx, "*/")
	(*current_pos).line += numNewLines
	(*current_pos).column += distFromLastNewLine
}

func handleIdentifier(i *int, input *string, tokens *[]TokenStruct, current_pos *Position) {
	counter1 := *i
	counter2 := 0
	if counter1+counter2 < len((*input)) && isCharaToZ(string((*input)[*i])) {

		for (isCharaToZ(string((*input)[counter1+counter2]))) || (*input)[counter1+counter2] == '_' || ((*input)[counter1+counter2] >= '0') && ((*input)[counter1+counter2] <= '9') {
			counter2 += 1

		}
		*i += counter2 - 1
		(*tokens) = append((*tokens), TokenStruct{Type: (VAR), Val: string((*input)[counter1 : counter1+counter2]), Pos: *current_pos})
		current_pos.column += counter2
	}
}

func isInt(input string) bool {
	intRegex := regexp.MustCompile(`^[0-9]\d*$`) // 0.2 .2

	if intRegex.MatchString(input) {
		return true
	}
	return false

}

func isChar(input string) bool {
	charRegex := regexp.MustCompile(`^[\x00-\x7F]$`)

	if charRegex.MatchString(input) {
		return true
	}
	return false

}

func isCharaToZ(input string) bool {
	charRegex := regexp.MustCompile(`^[a-zA-Z]$`)

	if charRegex.MatchString(input) {
		return true
	}
	return false

}

func isFloat(input string) bool {
	floatRegex := regexp.MustCompile(`^[0-9]*\.[0-9]+$`) // 0.2 .2

	if floatRegex.MatchString(input) {
		return true
	}
	return false

}

func forwardPosOneSpace(current_pos *Position) {
	(*current_pos).column += 1
}

func handleNewLine(current_pos *Position, tokens *[]TokenStruct) {
	if len((*tokens)) > 0 && (*tokens)[len((*tokens))-1].Type != (SEPARATOR) {
		*tokens = append((*tokens), TokenStruct{Type: (SEPARATOR), Val: "\n", Pos: *current_pos})
	}
	current_pos.line += 1
	current_pos.column = 0
}

func handleComments(input *string, i *int, current_pos *Position, found *bool) bool {
	if string((*input)[*i])+string((*input)[(*i)+1]) == "//" {
		handleOneLineComment(input, i, current_pos)
		(*found) = true
		return true

	} else if string((*input)[(*i)])+string((*input)[(*i)+1]) == "/*" { // block comments
		handleBlockComment(input, i, current_pos)
		(*found) = true
	}
	return false
}

func handleRealNumbers(i *int, tokens *[]TokenStruct, input *string, current_pos *Position) {
	var counter1, counter2 int
	counter1, counter2 = (*i), handleNumber(input, i, tokens)
	*i += counter2 - 1
	if isInt(string((*input)[counter1 : counter1+counter2])) {
		(*tokens) = append((*tokens), TokenStruct{Type: (NUMBER), Val: string((*input)[counter1 : counter1+counter2]), Pos: *current_pos})
		current_pos.column += counter2
	}
	if isFloat(string((*input)[counter1 : counter1+counter2])) {
		(*tokens) = append((*tokens), TokenStruct{Type: (FLOAT), Val: string((*input)[counter1 : counter1+counter2]), Pos: *current_pos})
		current_pos.column += counter2
	}
}

func handleChar(current_pos *Position, input *string, idx *int, tokens *[]TokenStruct) {
	current_pos.column += 1
	(*tokens) = append((*tokens), TokenStruct{Type: (CHAR), Val: string((*input)[(*idx)+1]), Pos: (*current_pos)})
	(*idx) += 3
	(*current_pos).column += 2
}

func handleNumber(input *string, i *int, tokens *[]TokenStruct) int {
	counter1 := *i
	counter2 := 0
	if counter1+counter2 < len((*input)) {
		for (string((*input)[counter1+counter2]) >= "0" && string((*input)[counter1+counter2]) <= "9") || (*input)[counter1+counter2] == '.' {
			counter2 += 1

		}

	}
	return (counter2)
}

func findFirstRune(input string, startIndex int, searchable rune) int {
	for i := startIndex; i < len(input); i++ {
		if i+1 < len(input) {
			if input[i] == byte(searchable) {
				return i
			}
		}
	}
	return -1
}

func findFirstStr(input string, startIndex int, searchable string) (int, int, int) {
	numNewLines := 0
	distFromLastNewLine := 0

	for i := startIndex; i < len(input); i++ {
		if i+1 < len(input) {
			if input[i] == '\n' {
				numNewLines += 1
				distFromLastNewLine = 0
			}
			distFromLastNewLine += 1
			if string(input[i])+string(input[i+1]) == searchable {
				return i + 1, numNewLines, distFromLastNewLine
			}
		}

	}
	return -1, -1, -1
}
