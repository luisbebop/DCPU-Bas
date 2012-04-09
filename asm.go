/****************************************
  DCPU-Bas - QuickBasic DCPU-16 compiler
      by M4v3R (maciej@trebacz.org)

      Functions that print assembly
 ****************************************/

package main

import (
	"fmt"
	"strconv"
)

func Negate() {
	Push()
	EmitLine("SET A, 0")
	EmitLine("SUB A, POP")
}

func Clear() {
	EmitLine("SET A, 0")
}

func Not() {
	EmitLine("XOR A, -1")
}

func LoadConst(s string) {
	val, _ := strconv.Atoi(s)
	EmitLine(fmt.Sprintf("SET A, %#x", val))
}

func LoadConstString(s string) {
	label := NewConst()
	EmitLine(fmt.Sprintf(":%s DAT \"%s\", 0", label, s))
	EmitLine(fmt.Sprintf("SET A, %s", label))
	EmitLine(fmt.Sprintf("BOR A, 0x8000"))
}

func LoadVar(s string) {
	if !InTable(s) {
		Undefined(s)
	}
	symbol := Symbols[Locate(GetSymbol(s))]
	EmitLine(fmt.Sprintf("SET A, [%#x]", (0xffff + symbol.l)))
}

func Push() {
	StackDepth++
	EmitLine("SET PUSH, A")
}

func PopAdd() {
	StackDepth--
	EmitLine("ADD A, POP")
}

func PopSub() {
	StackDepth--
	EmitLine("SUB A, POP")
	Negate()
}

func PopMul() {
	StackDepth--
	EmitLine("MUL A, POP")
}

func PopDiv() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("DIV B, A")
	EmitLine("SET A, B")
	EmitLine("SET B, 0")
}

func PopAnd() {
	StackDepth--
	EmitLine("AND A, POP")
}

func PopOr() {
	StackDepth--
	EmitLine("BOR A, POP")
}

func PopXor() {
	StackDepth--
	EmitLine("XOR A, POP")
}

func PopCompare() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("SET C, 1")
}

func SetEqual() {
	EmitLine("IFE A, B")
	EmitLine("SET C, 0")
}

func SetNotEqual() {
	EmitLine("IFN A, B")
	EmitLine("SET C, 0")
}

func SetGreater() {
	EmitLine("IFG B, A")
	EmitLine("SET C, 0")
}

func SetLess() {
	EmitLine("IFG A, B")
	EmitLine("SET C, 0")
}

func SetGreaterOrEqual() {
	SetGreater()
	SetEqual()
}

func SetLessOrEqual() {
	SetLess()
	SetEqual()
}

func Store(s string) {
	symbol := Symbols[Locate(GetSymbol(s))]
	EmitLine(fmt.Sprintf("SET [%#x], A", (0xffff + symbol.l)))
}

func Branch(s string) {
	EmitLine(fmt.Sprintf("SET PC, %s", s))
}

func BranchFalse(s string) {
	EmitLine("IFN C, 0")
	Branch(s)
}

func Prolog() {
	EmitLine("SET PC, begin")
	EmitLine("")
	Lib()
	EmitLine("")
	PostLabel("begin")
}

func Ret() {
	EmitLine("SET PC, POP")
}

func Cls() {
	EmitLine("SET A, 0x200")
	l := NewLabel()
	PostLabel(l)
	EmitLine("SET B, 0x8000")
	EmitLine("ADD B, A")
	EmitLine("SET [B], 0")
	EmitLine("SUB A, 1")
	BranchFalse(l)
	EmitLine("SET [0x8000], 0")
	Next()
}

func Loc() {
	Next()
	BoolExpression()
	EmitLine("SUB A, 1")
	EmitLine("SET PUSH, 0x20")
	EmitLine("MUL A, POP")
	EmitLine("SET X, A")
	if Token == ',' {
		Next()
		BoolExpression()
		EmitLine("SUB A, 1")
		EmitLine("ADD X, A")
	}
}

func Lib() {
	PostLabel("printchar") // Print char
	EmitLine("SET B, X") // Get current cursor position
	EmitLine("ADD B, 0x8000") // Add video mem address
	EmitLine("BOR A, Y") // Apply color code
	EmitLine("SET [B], A") // Set video memory byte to show char
	EmitLine("ADD X, 1") // Increment cursor position
	EmitLine("IFN X, 0x160") // Check if we should do next line (X > 32)
	EmitLine("SET PC, pnline")
	EmitLine("SET X, 0") // First row, first column
	PostLabel("pnline")
	Ret()

	PostLabel("printint") // Print integer
	EmitLine("SET I, 0") // Loop counter
	PostLabel("printint1") // Loop: divide A by 10 until 0 is left
	EmitLine("SET B, A") // Store A (number) for later
	EmitLine("MOD A, 0xa") // Get remainder from division by 10
	EmitLine("ADD A, 0x30") // Add 0x30 to the remainder to get ASCII code
	EmitLine("SET PUSH, A") // Store the remainder (digit) on the stack
	EmitLine("SET A, B") // Get A (number) back
	EmitLine("DIV A, 0xa") // Divide the number by 10
	EmitLine("ADD I, 1") // Increment loop counter
	EmitLine("IFN A, 0") // A > 10: jump
	EmitLine("SET PC, printint1")
	PostLabel("printint2") // Loop: print character by character
	EmitLine("SET A, POP") // Get digit from stack
	EmitLine("JSR printchar") // Print character
	EmitLine("SUB I, 1") // Decrement loop counter
	EmitLine("IFN I, 0")
	EmitLine("SET PC, printint2") // Jump back if there are more chars
	EmitLine("SET A, POP")
	Ret()

	PostLabel("printstr") // Print string
	EmitLine("AND A, 0x7fff")
	EmitLine("SET I, A") // Get string address
	PostLabel("printstr1")
	EmitLine("IFE [I], 0") // Return if we've reached end of string
	Ret()
	EmitLine("SET A, [I]") // Set A to address of next char
	EmitLine("JSR printchar") // Print char
	EmitLine("ADD I, 1") // Increment char index
	EmitLine("SET PC, printstr1") // Loop

	PostLabel("print")
	EmitLine("SET B, A") // Check variable type
	EmitLine("SHR B, 15")
	EmitLine("IFE B, 0") // Integer
	EmitLine("JSR printint")
	EmitLine("IFE B, 1") // String
	EmitLine("JSR printstr")
	Ret()
}

func Epilog() {
	PostLabel("crash")
	EmitLine("SET PC, crash")
}