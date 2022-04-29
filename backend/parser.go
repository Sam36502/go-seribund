package backend

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Sam36502/go-seribund/config"
)

func ParseProgram(program string) Program {
	lines := strings.Split(program, config.LINE_ENDING)

	parsed := make(Program, 0)
	problems := make([]string, 0)
	regRegex, err := regexp.Compile(config.REGEX_REG)
	if err != nil {
		fmt.Println("[ERROR] Failed to compile RegEx '" + config.REGEX_REG + "'.")
		os.Exit(1)
	}
	for li, line := range lines {

		if len(line) == 0 {
			continue
		}

		// Remove Comments
		inner := line
		if ci := strings.IndexRune(inner, config.COMMENT_RUNE); ci != -1 {
			inner = strings.TrimSpace(line[:ci])
		}
		if len(inner) == 0 {
			continue
		}

		// Get contents of parentheses
		startParen := strings.IndexRune(inner, config.PAREN_L_RUNE)
		endParen := strings.IndexRune(inner, config.PAREN_R_RUNE)
		if startParen == -1 || endParen == -1 {
			problems = append(problems, SyntaxError(li, "Missing Parentheses", line))
			continue
		}
		inner = strings.TrimSpace(inner[startParen+1 : endParen])
		if len(inner) == 0 {
			problems = append(problems, SyntaxError(li, "Empty Instruction", line))
			continue
		}

		// Parse Operation
		var ins Instruction
		opi := strings.IndexRune(inner, config.ADD_RUNE)
		if opi > -1 {
			ins.Operation = OP_ADD
		} else {
			opi = strings.IndexRune(inner, config.SUB_RUNE)
			if opi > -1 {
				if ins.Operation != 0 {
					problems = append(problems, SyntaxError(li, "Too many Operators", line))
					continue
				}
				ins.Operation = OP_SUB
			} else {
				problems = append(problems, SyntaxError(li, "Missing Operator", line))
				continue
			}
		}

		// Parse Register
		ins.Register = strings.TrimSpace(inner[:opi])
		if len(ins.Register) == 0 {
			problems = append(problems, SyntaxError(li, "Missing Register", line))
			continue
		}
		isReg := regRegex.MatchString(ins.Register)
		if !isReg {
			problems = append(problems, SyntaxError(li, "Invalid Register", line))
		}

		// Parse Value
		val := strings.TrimSpace(inner[opi+1:])
		if len(val) == 0 {
			problems = append(problems, SyntaxError(li, "Missing Value", line))
			continue
		}
		isReg = regRegex.MatchString(val)
		if isReg {
			ins.Value.Reg = val
			ins.Value.Num = -1
		} else {
			ins.Value.Reg = ""
			ins.Value.Num, err = strconv.ParseInt(val, config.VALUE_BASE, 64)
			if err != nil {
				problems = append(problems, SyntaxError(li, "Invalid Value", line))
				continue
			}
		}

		// Add instruction to program
		parsed = append(parsed, ins)

	}

	if len(problems) > 0 {
		for _, p := range problems {
			fmt.Print(p)
		}
		fmt.Println("Parsing failed; Exiting...")
		os.Exit(1)
	}

	return parsed
}

func SyntaxError(li int, msg, line string) string {
	return fmt.Sprintf("Syntax Error on line %d: %s\n    %s\n", li+1, msg, line)
}
