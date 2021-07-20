package textlist

type Option uint32
const(
	NONE Option = 1 << iota
	FILE_TRIMSPACE
	FILE_NEWLINE
	FILE_COMMA
	FILE_PIPE
	FILE_SEMICOLON
	FILE_COLON
)
