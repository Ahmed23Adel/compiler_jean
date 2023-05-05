package main

type Node struct {
	start    int
	end      int
	name     string
	adjacent []*Node
}

const (
	EOF                 = "end of file"
	ILLEGAL             = "illegal"
	VAR                 = "variable"
	ASSIGN              = "assign (=)"
	NUMBER              = "number (0-9)"
	FLOAT               = "float"
	CHAR                = "char"
	ADD                 = "add (+)"
	SUB                 = "sub (-)"
	MUL                 = "mul (*)"
	DIV                 = "div (/)"
	POWER               = "power (^)"
	MOD                 = "mod (%)"
	ABS                 = "abs"
	AND                 = "and"
	OR                  = "or"
	NOT                 = "not"
	OPEN_PARAN          = "open paran ("
	CLOSE_PARAN         = "close paran )"
	QUESTION_MARK       = "question mark (?)"
	OPEN_CURLY_BRACKET  = "open curly bracket ({)"
	CLOSE_CURLY_BRACKET = "close curly bracket (})"
	EXCLAMATION_MARK    = "exclamation mark (!)"
	COLON               = "colon (:)"
	SEPARATOR           = "separator"
	INT                 = "int(dtype)"
	FLT                 = "float(dtype)"
	CHR                 = "char(dtype)"
	STR                 = "str"
	COMMA               = "comma (,)"
	COMP                = "comparator(==)"
	BIT_AND             = "Bitwise And"
	BIT_OR              = "Bitwise Or(|)"
	GT                  = "Greater than(>)"
	LT                  = "Less than(<)"
	GTE                 = "Greater than or Equal(>=)"
	LTE                 = "Less than or Equal(<=)"
)

type TokenStruct struct {
	Type Token  //,
	Val  string //x
	Pos  Position
}

type Token string
