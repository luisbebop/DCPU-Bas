	SET PUSH, X
	SET PUSH, Y
	SET PUSH, Z
	SET PUSH, I
	SET PUSH, J
	SET A, SP
	SET PUSH, A
	SET Y, 0x7000
	SUB SP, 3 ; Alloc space on stack
	ADD PC, 2
	:c0 DAT "*", 0
	SET A, c0
	BOR A, 0x8000
	SET [0xffff], A
	SET A, 0x1
	SET [0xfffe], A
	SET A, 0x0
	SET [0xfffd], A
	SET A, 0x6
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0x9
	SUB A, 1
	ADD X, A
	ADD PC, 17
	:c1 DAT "0x10c BASIC FTW!", 0
	SET A, c1
	BOR A, 0x8000
	JSR print
	:l0
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x21
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l2
	SET A, 0x1
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, [0xfffe]
	SUB A, 1
	ADD X, A
	SET PC, l3
	:l2
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x2c
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l4
	SET A, 0x2
	SET PUSH, A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x21
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	ADD A, POP
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0x20
	SUB A, 1
	ADD X, A
	SET PC, l3
	:l4
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x4b
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l6
	SET A, 0xc
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0x1f
	SET PUSH, A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x2c
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	SUB A, 1
	ADD X, A
	SET PC, l3
	:l6
	SET A, 0xb
	SET PUSH, A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x4b
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0x1
	SUB A, 1
	ADD X, A
	:l3
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x3
	SET B, POP
	MOD B, A
	SET A, B
	SET PUSH, A
	SET A, [0xfffd]
	SET B, POP
	SET C, 1
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l8
	SET A, 0xe
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, 0x0
	SHL A, 8
	BOR Y, A
	SET PC, l9
	:l8
	SET A, 0x1
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, 0x0
	SHL A, 8
	BOR Y, A
	:l9
	SET A, [0xffff]
	JSR print
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffe], A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x55
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFN C, 0
	SET PC, l10
	SET A, 0x0
	SET [0xfffe], A
	SET A, [0xfffd]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffd], A
	SET A, [0xfffd]
	SET PUSH, A
	SET A, 0x2
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFN C, 0
	SET PC, l10
	SET A, 0x0
	SET [0xfffd], A
	:l10
	SET PC, l0
	:l1
	SET J, POP
	SET I, POP
	SET Z, POP
	SET Y, POP
	SET X, POP
	SET A, POP
	SET SP, A
	SET PC, end
	
	; compiled functions
	:getkey
	SET A, [0x9000]
	SET [0x9000], 0
	SET PC, POP
	:strlen
	SET I, A
	:strlen1
	ADD I, 1
	IFN [I], 0x0
	SET PC, strlen1
	SET A, B
	SET PC, POP
	:printchar
	SET [0x8000+X], A
	BOR [0x8000+X], Y
	ADD X, 1
	IFG X, 0x21f
	SET X, 0
	:pnline
	SET PC, POP
	:printint
	SET I, 0
	:printint1
	SET B, A
	MOD A, 0xa
	ADD A, 0x30
	SET PUSH, A
	SET A, B
	DIV A, 0xa
	ADD I, 1
	IFN A, 0
	SET PC, printint1
	:printint2
	SET A, POP
	JSR printchar
	SUB I, 1
	IFN I, 0
	SET PC, printint2
	SET A, POP
	SET PC, POP
	:printstr
	IFG 0xF000, A
	AND A, 0x7fff
	SET I, A
	:printstr1
	IFE [I], 0
	SET PC, POP
	SET A, [I]
	JSR printchar
	ADD I, 1
	SET PC, printstr1
	:printnl
	DIV X, 32
	ADD X, 1
	MUL X, 32
	SET PC, POP
	:print
	SET B, A
	SHR B, 15
	IFE B, 0
	JSR printint
	IFE B, 1
	JSR printstr
	SET PC, POP
	:end
	IFN SP, 0
	SET PC, POP
	:halt
	SET PC, halt
