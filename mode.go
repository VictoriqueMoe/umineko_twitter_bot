package main

import "fmt"

type Mode string

const (
	ModeRandom Mode = "random"
	ModeErika  Mode = "erika"
)

func (m *Mode) String() string {
	return string(*m)
}

func (m *Mode) Set(v string) error {
	switch Mode(v) {
	case ModeRandom, ModeErika:
		*m = Mode(v)
		return nil
	default:
		return fmt.Errorf("must be one of: random, erika")
	}
}
