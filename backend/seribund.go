/*
 *	Seribund interpreter
 *
 *	Syntax follows this kind of structure:
 *
 *		number      ::= [0-9]+
 *		register    ::= [a-z][a-z0-9]*
 *		instruction ::= \s* \( \s* register \s* [+-] \s* (number|register) \s* \) \s* (; .*)?
 *
 */

package backend

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Sam36502/go-seribund/config"
)

type Operation int64

const (
	OP_ADD = Operation(+1)
	OP_SUB = Operation(-1)
)

func (o Operation) Perform(a, b int64) int64 {
	return a + (b * int64(o))
}

type Value struct {
	Reg string
	Num int64
}

func (v Value) isReg() bool {
	return v.Num == -1
}

type Instruction struct {
	Register  string
	Operation Operation
	Value     Value
}

type Program []Instruction

func RunProgram(prog Program) map[string]int64 {

	Registers := make(map[string]int64)

	pp := 0        // Program Pointer
	pv := int64(1) // Previous Value
	runs := 0
	for pv > -1 {
		ins := prog[pp]

		for i := 0; int64(i) < pv; i++ {
			regVal, exists := Registers[ins.Register]
			if !exists {
				regVal = 0
				Registers[ins.Register] = regVal
			}
			valVal := ins.Value.Num
			if ins.Value.isReg() {
				valVal = Registers[ins.Value.Reg]
			}

			Registers[ins.Register] = ins.Operation.Perform(regVal, valVal)
		}
		if pv == 0 {
			pv = 1
		} else {
			pv = Registers[ins.Register]
		}

		pp++
		if pp >= len(prog) {
			pp = 0
		}

		runs++
		if runs >= config.RUNS_LIMIT {
			fmt.Println("Reached run limit; Exiting...")
			break
		}
	}

	return Registers

}

// Orders registers alphabetically and puts their
// ASCII rune into a string. Also ignores negative values
func StringifyRegisters(regs map[string]int64) string {
	index := make([]string, 0)
	for reg := range regs {
		index = append(index, reg)
	}
	sort.Strings(index)

	var out strings.Builder
	for _, reg := range index {
		if regs[reg] < 0 {
			continue
		}
		out.WriteRune(rune(regs[reg]))
	}
	return out.String()
}
