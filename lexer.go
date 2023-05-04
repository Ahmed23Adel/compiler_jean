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
	// fmt.Println("k0", *idx+len(searchable), len(*input))
	if *idx+len(searchable) < len(*input)-1 {
		// fmt.Println("k1", string((*input)[(*idx):(*idx)+len(searchable)]), searchable)
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
		fmt.Println(string(input[i]))
		if i+1 < len(input) {
			// one line comments
			if string(input[i])+string(input[i+1]) == "//" {
				handleOneLineComment(&input, &i, &current_pos)
				// fmt.Println("//current_pos", (current_pos))

			} else if string(input[i])+string(input[i+1]) == "/*" { // block comments
				handleBlockComment(&input, &i, &current_pos)
				// fmt.Println("/*current_pos", (current_pos))

			} else if string(input[i]) == " " { // space
				forwardPosOneSpace(&current_pos)
				// fmt.Println("spacecurrent_pos", (current_pos))

			} else if input[i] == '\n' { // new line
				handleNewLine(&current_pos, &tokens)
				// fmt.Println("nnncurrent_pos", (current_pos))
			} else if handleOneLetterToken(&input, &i, &current_pos, "+", &tokens, ADD) { // +
				fmt.Println("+current_pos", (current_pos))
			} else if handleOneLetterToken(&input, &i, &current_pos, "-", &tokens, SUB) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "*", &tokens, MUL) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "/", &tokens, DIV) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "^", &tokens, POWER) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "%", &tokens, MOD) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "|", &tokens, ABS) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "(", &tokens, OPEN_PARAN) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, ")", &tokens, CLOSE_PARAN) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "?", &tokens, QUESTION_MARK) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "{", &tokens, OPEN_CURLY_BRACKET) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "}", &tokens, CLOSE_CURLY_BRACKET) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "!", &tokens, EXCLAMATION_MARK) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, ":", &tokens, COLON) { //

			} else if handleOneLetterToken(&input, &i, &current_pos, "=", &tokens, EQUAL) { //

			} else if handleMultiLetterToken(&input, &i, &current_pos, "and", &tokens, AND) { //

			} else if handleMultiLetterToken(&input, &i, &current_pos, "or", &tokens, OR) { //

			} else if handleMultiLetterToken(&input, &i, &current_pos, "not", &tokens, NOT) { //

			} else if (string(input[i]) == "'" && string(input[i+2]) == "'") && isChar(string(input[i+1])) {
				handleChar(&current_pos, &input, &i, &tokens)

			} else if string(input[i]) == "y" || string(input[i]) == "x" || string(input[i]) == "z" || string(input[i]) == "u" { // identifier token
				tokens = append(tokens, tokenStruct{token_type: getToken(IDENT), token_value: string(input[i]), token_pos: current_pos})
				fmt.Println(tokens)
				fmt.Println(current_pos)
				current_pos.column += 1
			} else if isInt(string(input[i])){
				counter1:=i
				counter2:=0
				if counter1+counter2<len(input){
					for (string(input[counter1+counter2])>="0" && string(input[counter1+counter2]) <= "9") || input[counter1+counter2] == '.'{
						counter2 +=1
						//fmt.Println(string(input[counter1:counter1+counter2]))
					}
					fmt.Println(string(input[counter1:counter1+counter2]))
					fmt.Println(counter2, "input: ",string(input[i]))

				}
				i += counter2-1
				if isInt(string(input[counter1:counter1+counter2])){
					tokens = append(tokens,tokenStruct{token_type: getToken(NUMBER), token_value: string(input[counter1:counter1+counter2]), token_pos: current_pos})
					current_pos.column+=counter2
				}
				if isFloat(string(input[counter1:counter1+counter2])){
					tokens = append(tokens,tokenStruct{token_type: getToken(FLOAT), token_value: string(input[counter1:counter1+counter2]), token_pos: current_pos})
					current_pos.column+=counter2
				}


			}

		}

		// "ahmed \" yousef"
		// "ahmed" + "yousef"
		// fmt.Println(string(input[i]))
		// "     ajflkasdjf; \" fsdklasd "
		// break
		// if i > 50 {
		// 	break
		// }
	}
	fmt.Println(tokens)
	return tokens
}

func handleChar(current_pos *Position, input *string, idx *int, tokens *[]tokenStruct) {
	current_pos.column += 1
	(*tokens) = append((*tokens), tokenStruct{token_type: getToken(CHAR), token_value: string((*input)[(*idx)+1]), token_pos: (*current_pos)})
	(*idx) += 3
	(*current_pos).column += 2
}

//  // : f(x: int)
// 	COLON     // :

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
	// sample_str = removeComments(sample_str)
	// fmt.Println(sample_str)
	lex_analyzer(sample_str)
	// fmt.Println(sample_str)
	// fmt.Println(sample_str[0], sample_str[1])
	// fmt.Println(sample_str[0] == byte('x'))

	// fmt.Println(sample_str)
	// tst := "_var_name"
	// fmt.Println(isIdentifier(tst))
	// // tst := "x = 13.0 + 16.5 \n y = 12.6 - 10.4 + .3 + 0.1 + 9 \n"
	// tstReg := regexp.MustCompile(`\d*\.\d+`) // 0.2 .2

	// // Remove single-line comments
	// fmt.Println(tstReg.ReplaceAllString(tst, "FLOAT"))
	// // fmt.Println(reflect.TypeOf(sample_str))
	// fmt.Println(sample_str[0], sample_str[1], sample_str[2], sample_str[4], sample_str[55])
	// age1 = 16.85.5..9(%|and| (| } | :| or space | \n)5
	// age2 = age1/5
	//(age)
	//"any /" "

	// x = - 9
	// x = -y
	// x = - y - -x

	// case operator
}
