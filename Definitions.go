package main

type Node struct {
	Start    int
	End      int
	Type     string
	Children []*Node
}

func (n *Node) IsTerminal() bool {
	return len(n.Children) == 0
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
	RETURN              = "Return"
	BREAK               = "Break"
	CONTINUE            = "continue"
	SPACE               = "space"
)

const (
	CODE_NON_TERMINAL                  = "Code"
	STATEMENT_NON_TERMINAL             = "Statement"
	EXPRESSION_NON_TERMINAL            = "Expression"
	FACTOR_NON_TERMINAL                = "Factor"
	TERM_NON_TERMINAL                  = "Term"
	CONDITIONAL_STATEMENT_NON_TERMINAL = "Conditional Statement"
	ELSE_NON_TERMINAL                  = "Else"
	LOOP_STATEMENT_NON_TERMINAL        = "Loop Statement"

	// Terminals
	OPEN_BRACES_TERMINAL   = "{"
	CLOSED_BRACES_TERMINAL = "}"
	OPEN_PARAN_TERMINAL    = "("
	CLOSED_PARAN_TERMINAL  = ")"
	ASSIGN_TERMINAL       = "="
	QUESTION_MARK_TERMINAL = "?"
)

type TokenStruct struct {
	Type Token  //,
	Val  string //x
	Pos  Position
	Usage use
}

type Token string
type use string