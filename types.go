package main

import (
	"math/rand"
)

type Machine struct {
	Data []uint

	Head int
}

func NewMachine() *Machine {
	m := Machine{}
	m.Data = []uint{}
	return &m
}

func (m *Machine) Randomize(size int) {
	m.Data = []uint{}
	for range size {
		m.Data = append(m.Data, uint(rand.Intn(2)))
	}
}

func (m *Machine) Execute(inst *Instruction) int {
	if m.Data[m.Head] == inst.Expected {
		m.Data[m.Head] = inst.Yes.Write
		m.Head += int(inst.Yes.Dir)
		return inst.Yes.NextInst
	} else {
		m.Data[m.Head] = inst.No.Write
		m.Head += int(inst.No.Dir)
		return inst.No.NextInst
	}
}

type Direction int

const (
	LEFT  Direction = -1
	STAY  Direction = 0
	RIGHT Direction = 1
)

type Instruction struct {
	Expected uint

	Yes YesOrNo
	No  YesOrNo

	Null bool
}
type YesOrNo struct {
	Write    uint
	Dir      Direction
	NextInst int
}
