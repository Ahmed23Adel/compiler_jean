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
	// Constants
	CONST
	// Mathmatical expressions
	ADD   // +
	SUB   // -
	MUL   // *
	DIV   // /
	POWER // ^
	MOD   // %
	ABS   // |
	//logical operators
	AND // and
	OR  // or
	NOT // not
	// if then
	OPEN_PARAN          // (
	CLOSE_PARAN         // )
	QUESTION_MARK       // ?
	OPEN_CURLY_BRACKET  // {
	CLOSE_CURLY_BRACKET // }
	EXCLAMATION_MARK    // !
	//Block structure // DONE
	// Functions
	VAR_TYPES // : f(x: int)
	COLON     // :
	SEPERATOR
)

var tokens = []string{
	// "type(MATHOP, LOGICALOP, VAR, EQUAL, CONST, IFREL--if related--, VARTYPE )"_ EXTRA INFORMATOIN
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	// VARS
	IDENT:  "IDENT", // TODO: REGEX
	EQUAL:  "EQUAL",
	NUMBER: "VAR_NUMBER", // TODO: REGEX
	FLOAT:  "VAR_FLOAT",  // TODO: REGEX
	CHAR:   "VAR_CHAR",   // TODO: REGEX [a-zA-Z]
	// Constants
	CONST: "CONST",
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
	OPEN_CURLY_BRACKET:  "IFREL_OPEN_CURLY_BRACKET",  // {
	CLOSE_CURLY_BRACKET: "IFREL_CLOSE_CURLY_BRACKET", // }
	EXCLAMATION_MARK:    "IFREL_EXCLAMATION_MARK",    // !
	//Block structure // DONE
	// Functions
	VAR_TYPES: "VARTYPE", // : f(x: int)
	COLON:     "COLON",   // :
	SEPERATOR: "SEPERATOR",
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

func lex_analyzer(input string) []string {
	tokens := make([]string, 0)
	current_pos := Position{line: 0, column: 0}
	for i := 0; i < len(input); i++ {
		// do something with i and str[i]

		// Handle Comment
		if i+1 < len(input) {
			if string(input[i])+string(input[i+1]) == "//" {
				current_pos.line += 0
				current_pos.column = 0
				i = findFirstRune(input, i, '\n') - 1
				if len(tokens) > 0 && tokens[len(tokens)-1] != getToken(SEPERATOR) {
					tokens = append(tokens, getToken(SEPERATOR))
				}
				continue
			}
		}

		// handle /**/
		if i+1 < len(input) {
			if string(input[i])+string(input[i+1]) == "/*" {
				current_pos.line += 0
				current_pos.column = 0
				i = findFirstStr(input, i, "*/") - 1

				if len(tokens) > 0 && tokens[len(tokens)-1] != getToken(SEPERATOR) {
					tokens = append(tokens, getToken(SEPERATOR))
				}

				continue
			}
		}
		// space
		if string(input[i]) == " " {
			continue
		}
		if input[i] == '\n' {
			if len(tokens) > 0 && tokens[len(tokens)-1] != getToken(SEPERATOR) {
				tokens = append(tokens, getToken(SEPERATOR))
			}
			fmt.Println("NEW sep")
			continue
		}

		if string(input[i]) == "y" {
			tokens = append(tokens, getToken(IDENT))
		}

		fmt.Println(string(input[i]))
		fmt.Println(tokens)
	}

	return tokens
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

func findFirstStr(input string, startIndex int, searchable string) int {
	for i := startIndex; i < len(input); i++ {
		if i+1 < len(input) {
			if string(input[i])+string(input[i+1]) == searchable {
				return i
			}
		}

	}
	return -1
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
