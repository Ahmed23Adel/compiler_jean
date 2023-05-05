package main

type Node struct{
	start int
	end int
	name  string
	adjacent []*Node
}

const  (
	VAR  = "var"
	ADD_OP  = "addition operator"
	MULT_OP  = "multiplication operator"
	NUM  = "number"
	ASSIGN  = "="
	LEFT_BRACKET  = "("
	RIGHT_BRACKET  = ")"
)

type tokenStruct struct {
	Type string 
	Val string
}
