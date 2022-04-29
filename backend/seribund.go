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
	for pv > -1 {
		ins := prog[pp]

		for i := 0; int64(i) < pv; i++ {
			regVal := Registers[ins.Register]
			valVal := ins.Value.Num
			if ins.Value.isReg() {
				valVal = Registers[ins.Value.Reg]
			}
			pv = ins.Operation.Perform(regVal, valVal)
			Registers[ins.Register] = pv
		}

		pp++
		if pp >= len(prog) {
			pp = 0
		}
	}

	return Registers

}
