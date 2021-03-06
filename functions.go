/****************************************
  DCPU-Bas - QuickBasic DCPU-16 compiler
      by M4v3R (maciej@trebacz.org)

            Basic functions
 ****************************************/

package main

import (
	"fmt"
)

func FuncStr() {
	Next()
	MatchString("(")
	BoolExpression()
	label := NewLabel()
	EmitLine("ADD PC, 2")
	PostLabel(label)
	EmitLine("DAT 0")
	EmitLine("DAT 0")
	EmitLine(fmt.Sprintf("SET [%s], A", label))
	EmitLine(fmt.Sprintf("SET A, %s", label))
	EmitLine("BOR A, 0x8000")
}

func FuncChr() {
	Next()
	MatchString("(")
	BoolExpression()
	EmitLine("IFG 0xF000, A") // Check if it's not a stack pointer
	EmitLine("AND A, 0x7fff")
	EmitLine("SET PUSH, [A]")
	EmitLine("SET A, POP")
}

func FuncVal() {
	Next()
	MatchString("(")
	BoolExpression()
	Call("atoi")
}

func FuncPeek() {
	Next()
	MatchString("(")
	BoolExpression()
	EmitLine("SET B, [A]")
	EmitLine("SET A, B")
}

func FuncSqr() {
	Next()
	MatchString("(")
	BoolExpression()
	Call("sqrt")
}

func FuncLen() {
	Next()
	MatchString("(")
	BoolExpression()
	Call("strlen")
}
