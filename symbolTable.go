package main

import (
	"errors"
	"fmt"
)

const (
	KIND_FUNC     = "function"
	KIND_PARAM    = "paramter"
	KIND_VARIABLE = "variable"
)

const (
	DTYPE_INT   = "int"
	DTYPE_FLOAT = "float"
	DTYPE_CHAR  = "char"
)

const (
	OTHER_CONST = "const"
	OTHER_FREE  = "free"
)

type symbolTableRowKind string
type symbolTableRowDtype string
type symbolTableRowOther string

// int float, char, if function it will be x int y float,bool
// then in case of function
// name will be x and type will be float
// in case of variable
// name will be float
// type dtype struct{
// 	name string
// 	type
// }

type symbolTableRow struct {
	name              string
	kind              symbolTableRowKind  // function, par, var
	dtype             symbolTableRowDtype // int float, char, if function it will be x int y float,bool
	other             symbolTableRowOther // const?
	pointerToNewTable *[]symbolTable      //{{}}
}

type symbolTable struct {
	list            []symbolTableRow
	pointerToHeader *symbolTable
}

func pushVariableIfPossible(name string, dtype symbolTableRowDtype, other symbolTableRowOther, globalSymbolTable *symbolTable) error {
	if !isVariableExistInSymbolTable(name, globalSymbolTable) {
		for _, row := range globalSymbolTable.list {
			fmt.Println("k1")
			if row.kind == KIND_VARIABLE && row.name == name {
				fmt.Println("k1")
				return errors.New("Variable with name " + name + " already exists in current scope")
			}

		}
		(*globalSymbolTable).list = append((*globalSymbolTable).list, symbolTableRow{name: name, kind: KIND_VARIABLE, dtype: dtype, other: other})
		return nil
	}
	return nil
}

func pushFunctionIfPossible(name string, dtype symbolTableRowDtype, other symbolTableRowOther, globalSymbolTable *symbolTable) error {
	if !isFunctoinExistInSymbolTable(name, globalSymbolTable) {
		for _, row := range globalSymbolTable.list {
			fmt.Println("k1")
			if row.kind == KIND_VARIABLE && row.name == name {
				fmt.Println("k1")
				return errors.New("Variable with name " + name + " already exists in current scope")
			}

		}
		(*globalSymbolTable).list = append((*globalSymbolTable).list, symbolTableRow{name: name, kind: KIND_FUNC, dtype: dtype, other: other})
		return nil
	}
	return nil
}

func pushVariable(name string, dtype symbolTableRowDtype, other symbolTableRowOther, globalSymbolTable *symbolTable) error {
	for _, row := range globalSymbolTable.list {
		fmt.Println("k1")
		if row.kind == KIND_VARIABLE && row.name == name {
			fmt.Println("k1")
			return errors.New("Variable with name " + name + " already exists in current scope")
		}

	}
	(*globalSymbolTable).list = append((*globalSymbolTable).list, symbolTableRow{name: name, kind: KIND_VARIABLE, dtype: dtype, other: other})
	return nil
}

func isVariableExistInSymbolTable(name string, globalSymbolTable *symbolTable) bool {
	for _, row := range globalSymbolTable.list {
		if row.kind == KIND_VARIABLE && row.name == name {
			return true
		}
	}
	return false
}

func isFunctoinExistInSymbolTable(name string, globalSymbolTable *symbolTable) bool {
	for _, row := range globalSymbolTable.list {
		if row.kind == KIND_FUNC && row.name == name {
			return true
		}
	}
	return false
}
