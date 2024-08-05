package main

import (
	"fmt"
)

func main() {
	machine := NewMachine(4)
	machine.Randomize()

	program := []Instruction{
		inst(0, 1, STAY, 1, 0, RIGHT, 0),
		{Null: true},
	}

	currentInst := 0
	for {
		fmt.Printf("%+v %d\n", machine.Data, machine.Head)

		if program[currentInst].Null {
			fmt.Println("program has ended")
			break
		}

		if machine.Head == machine.Size {
			fmt.Println("head outside tape")
			break
		}
		currentInst = machine.Execute(&program[currentInst])
	}
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
		Expected:    ex,
		WriteYes:    wy,
		DirYes:      dy,
		NextInstYes: ny,
		WriteNo:     wn,
		DirNo:       dn,
		NextInstNo:  nn,
	}
}
