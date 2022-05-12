package config

// Source Parsing Config
const (
	LINE_ENDING  = "\n"
	COMMENT_RUNE = ';'
	PAREN_L_RUNE = '('
	PAREN_R_RUNE = ')'
	ADD_RUNE     = '+'
	SUB_RUNE     = '-'
	VALUE_BASE   = 10
	REGEX_REG    = "^[a-z][a-z0-9]*$"
	RUNS_LIMIT   = 1000
)

// Command Flags
const (
	FL_VALUES  = "values"
	FLS_VALUES = "v" // Shorthand
	FL_STEP    = "step"
	FLS_STEP   = "s" // Shorthand
)
