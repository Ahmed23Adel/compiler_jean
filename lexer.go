package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func readSample(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	return fileContent
}

func removeComments(input string) string {

	singleLineCommentRegex := regexp.MustCompile(`//.*?\n`)
	multiLineCommentRegex := regexp.MustCompile(`(?s)/\*.*?\*/`)
	newlineRegex := regexp.MustCompile(`^\s*$`)

	// Remove single-line comments
	input = singleLineCommentRegex.ReplaceAllString(input, "\n")

	// Remove multi-line comments
	input = multiLineCommentRegex.ReplaceAllString(input, "")

	input = newlineRegex.ReplaceAllString(input, "")

	return input
}

const (
	EOF = iota
	ILLEGAL
	// VARS
	IDENT
	EQUAL
	NUMBER
	FLOAT
	CHAR
	// Mathmatical expressions
	ADD   // +
	SUB   // -
	MUL   // *
	DIV   // /
	POWER // ^
	MOD   // %
	ABS   // |
	//logical operators
	AND // and //TODO
	OR  // or //TODO
	NOT // not //TODO
	// if then
	OPEN_PARAN          // (
	CLOSE_PARAN         // )
	QUESTION_MARK       // ?
	OPEN_CURLY_BRACKET  // {
	CLOSE_CURLY_BRACKET // }
	EXCLAMATION_MARK    // !
	//Block structure // DONE
	// Functions
	COLON // :
	SEPERATOR
)

var tokens = []string{
	// "type(MATHOP, LOGICALOP, VAR, EQUAL, CONST, IFREL--if related--, VARTYPE )"_ EXTRA INFORMATOIN
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	// VARS
	IDENT:  "IDENT",
	EQUAL:  "EQUAL",
	NUMBER: "VAL_NUMBER",
	FLOAT:  "VAL_FLOAT",
	CHAR:   "VAL_CHAR",
	// Mathmatical expressions
	ADD:   "MATHOP_ADD",   // +
	SUB:   "MATHOP_SUB",   // -
	MUL:   "MATHOP_MUL",   // *
	DIV:   "MATHOP_DIV",   // /
	POWER: "MATHOP_POWER", // ^
	MOD:   "MATHOP_MODE",  // %
	ABS:   "MATHOP_ABS",   // |
	//PARANTECIES it works for both mathmatical expressions and if
	OPEN_PARAN:  "PARAN_OPEN",  // (
	CLOSE_PARAN: "PARAN_CLOSE", // )
	//logical operators
	AND: "LOGICALOP_AND", // and
	OR:  "LOGICALOP_OR",  // or
	NOT: "LOGICALOP_NOT", // not
	// if then
	QUESTION_MARK:       "IFREL_QUESTION_MARK",       // ?
	OPEN_CURLY_BRACKET:  "SCOPE_OPEN_CURLY_BRACKET",  // {
	CLOSE_CURLY_BRACKET: "SCOPE_CLOSE_CURLY_BRACKET", // }
	EXCLAMATION_MARK:    "IFREL_EXCLAMATION_MARK",    // !
	//Block structure // DONE
	// Functions
	COLON:     "COLON", // :
	SEPERATOR: "SEPERATOR",
}

type tokenStruct struct {
	token_type  string //, //idenb
	token_value string //x
	token_pos   Position
}

type Token int

func getToken(t Token) string {
	return tokens[t]
}

func isFloat(input string) bool {
	floatRegex := regexp.MustCompile(`^[0-9]*\.[0-9]+$`) // 0.2 .2

	if floatRegex.MatchString(input) {
		return true
	}
	return false

}

func isInt(input string) bool {
	intRegex := regexp.MustCompile(`^[1-9]\d*$`) // 0.2 .2

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

func isIdentifier(input string) bool {
	idenRegex := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

	if idenRegex.MatchString(input) {
		return true
	}
	return false
}

type Position struct {
	line   int
	column int
}

func posGoNextLine(pos *Position) {
	(*pos).line = pos.line + 1
	(*pos).column = 0
}

func appendSeperator(tokens *[]string) {
	if len((*tokens)) > 0 && (*tokens)[len((*tokens))-1] != getToken(SEPERATOR) {
		*tokens = append((*tokens), getToken(SEPERATOR))
	}
}

func handleOneLineComment(input *string, idx *int, current_pos *Position) {
	// posGoNextLine(current_pos)
	*idx = findFirstRune((*input), (*idx), '\n') - 1
	appendSeperator(&tokens)
}

func handleBlockComment(input *string, idx *int, current_pos *Position) {
	var numNewLines, distFromLastNewLine int
	*idx, numNewLines, distFromLastNewLine = findFirstStr(*input, *idx, "*/")
	(*current_pos).line += numNewLines
	(*current_pos).column += distFromLastNewLine

	if len(tokens) > 0 && tokens[len(tokens)-1] != getToken(SEPERATOR) {
		tokens = append(tokens, getToken(SEPERATOR))
	}
}

func forwardPosOneSpace(current_pos *Position) {
	(*current_pos).column += 1
}

func handleNewLine(current_pos *Position, tokens *[]tokenStruct) {
	if len((*tokens)) > 0 && (*tokens)[len((*tokens))-1].token_type != getToken(SEPERATOR) {
		*tokens = append((*tokens), tokenStruct{token_type: getToken(SEPERATOR), token_value: "\n", token_pos: *current_pos})
	}

	current_pos.line += 1
	current_pos.column = 0
}

func handleOneLetterToken(input *string, idx *int, current_pos *Position, searchable string, tokens *[]tokenStruct, token Token) bool {
	if string((*input)[(*idx)]) == searchable {
		(*tokens) = append((*tokens), tokenStruct{token_type: getToken(token), token_value: searchable, token_pos: *current_pos})
		(*current_pos).column += 1
		return true
	}
	return false
}

func handleMultiLetterToken(input *string, idx *int, current_pos *Position, searchable string, tokens *[]tokenStruct, token Token) bool {
	if *idx+len(searchable) < len(*input)-1 {
		if string((*input)[(*idx):(*idx)+len(searchable)]) == searchable {
			(*tokens) = append((*tokens), tokenStruct{token_type: getToken(token), token_value: searchable, token_pos: *current_pos})
			(*current_pos).column += len(searchable)
			*idx += len(searchable) - 1
			return true
		}

	}
	return false
}

func lex_analyzer(input string) []tokenStruct {
	tokens := make([]tokenStruct, 0)
	current_pos := Position{line: 0, column: 0}
	for i := 0; i < len(input); i++ {
		// if i > 50 {
		// 	break
		// }

		if i+2 < len(input) { // handle 3 characters
			if handleMultiLetterToken(&input, &i, &current_pos, "and", &tokens, AND) { //
				continue

			} else if handleMultiLetterToken(&input, &i, &current_pos, "or", &tokens, OR) { //
				continue

			} else if handleMultiLetterToken(&input, &i, &current_pos, "not", &tokens, NOT) { //
				continue

			} else if (string(input[i]) == "'" && string(input[i+2]) == "'") && isChar(string(input[i+1])) {
				handleChar(&current_pos, &input, &i, &tokens)
				continue
			}
		}
		if i+1 < len(input) { // handle two characters
			if string(input[i])+string(input[i+1]) == "//" {
				handleOneLineComment(&input, &i, &current_pos)
				continue

			} else if string(input[i])+string(input[i+1]) == "/*" { // block comments
				handleBlockComment(&input, &i, &current_pos)
				continue
			}
		}
		if i < len(input) { // handle one character
			if string(input[i]) == " " { // space
				forwardPosOneSpace(&current_pos)
				continue
			} else if input[i] == '\n' { // new line
				handleNewLine(&current_pos, &tokens)
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "+", &tokens, ADD) { // +
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "-", &tokens, SUB) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "*", &tokens, MUL) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "/", &tokens, DIV) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "^", &tokens, POWER) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "%", &tokens, MOD) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "|", &tokens, ABS) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "(", &tokens, OPEN_PARAN) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, ")", &tokens, CLOSE_PARAN) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "?", &tokens, QUESTION_MARK) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "{", &tokens, OPEN_CURLY_BRACKET) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "}", &tokens, CLOSE_CURLY_BRACKET) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "!", &tokens, EXCLAMATION_MARK) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, ":", &tokens, COLON) { //
				continue

			} else if handleOneLetterToken(&input, &i, &current_pos, "=", &tokens, EQUAL) { //
				continue

			}
		}

		if isInt(string(input[i])) {
			handleRealNumbers(&i, &tokens, &input, &current_pos)
			continue

		} else {
			handleIdentifier(&i, &input, &tokens, &current_pos)
			continue
		}

	}
	return tokens
}

func handleIdentifier(i *int, input *string, tokens *[]tokenStruct, current_pos *Position) {
	counter1 := *i
	counter2 := 0
	if counter1+counter2 < len((*input)) && isCharaToZ(string((*input)[*i])) {

		for (isCharaToZ(string((*input)[counter1+counter2]))) || (*input)[counter1+counter2] == '_' || ((*input)[counter1+counter2] >= '0') && ((*input)[counter1+counter2] <= '9') {
			counter2 += 1

		}
		*i += counter2 - 1
		(*tokens) = append((*tokens), tokenStruct{token_type: getToken(IDENT), token_value: string((*input)[counter1 : counter1+counter2]), token_pos: *current_pos})
		current_pos.column += counter2
	}
}

func handleRealNumbers(i *int, tokens *[]tokenStruct, input *string, current_pos *Position) {
	var counter1, counter2 int
	counter1, counter2 = (*i), handleNumber(input, i, tokens)
	*i += counter2 - 1
	if isInt(string((*input)[counter1 : counter1+counter2])) {
		(*tokens) = append((*tokens), tokenStruct{token_type: getToken(NUMBER), token_value: string((*input)[counter1 : counter1+counter2]), token_pos: *current_pos})
		current_pos.column += counter2
	}
	if isFloat(string((*input)[counter1 : counter1+counter2])) {
		(*tokens) = append((*tokens), tokenStruct{token_type: getToken(FLOAT), token_value: string((*input)[counter1 : counter1+counter2]), token_pos: *current_pos})
		current_pos.column += counter2
	}
}

func handleChar(current_pos *Position, input *string, idx *int, tokens *[]tokenStruct) {
	current_pos.column += 1
	(*tokens) = append((*tokens), tokenStruct{token_type: getToken(CHAR), token_value: string((*input)[(*idx)+1]), token_pos: (*current_pos)})
	(*idx) += 3
	(*current_pos).column += 2
}

func handleNumber(input *string, i *int, tokens *[]tokenStruct) int {
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

func main() {
	sample_str := readSample("sample.jean")
	sample_str = sample_str + "\n"
	tokens := lex_analyzer(sample_str)
	fmt.Println(tokens)
}
