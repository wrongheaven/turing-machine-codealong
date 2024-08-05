package main

import (
	"fmt"
)

func main() {
	machine := NewMachine()
	machine.Randomize(8)

	program := []Instruction{
		inst(0, 1, STAY, 1, 0, RIGHT, 0),
		{Null: true},
	}

	currentInst := 0
	for {
		if program[currentInst].Null {
			break
		}

		if machine.Head == len(machine.Data) {
			machine.Data = append(machine.Data, 0)
		}
		fmt.Printf("%+v %d\n", machine.Data, machine.Head)
		currentInst = machine.Execute(&program[currentInst])
	}
	fmt.Printf("%+v %d\n", machine.Data, machine.Head)
}

func inst(
	ex uint,
	wy uint,
	dy Direction,
	ny int,
	wn uint,
	dn Direction,
	nn int,
) Instruction {
	return Instruction{
		Expected: ex,
		Yes: YesOrNo{
			Write:    wy,
			Dir:      dy,
			NextInst: ny,
		},
		No: YesOrNo{
			Write:    wn,
			Dir:      dn,
			NextInst: nn,
		},
	}
}
