package main

type Node struct{
	start int
	end int
	name  string
	adjacent []*Node
}

type Token struct {
	Type string 
	Val string

}
