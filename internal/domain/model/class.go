package model

import (
	"time"
)

type ClassNumber string

const (
	Class1A ClassNumber = "1А"
	Class1B ClassNumber = "1Б"
	Class1V ClassNumber = "1В"

	Class2A ClassNumber = "2А"
	Class2B ClassNumber = "2Б"
	Class2V ClassNumber = "2В"

	Class3A ClassNumber = "3А"
	Class3B ClassNumber = "3Б"
	Class3V ClassNumber = "3В"

	Class4A ClassNumber = "4А"
	Class4B ClassNumber = "4Б"
	Class4V ClassNumber = "4В"

	Class5A ClassNumber = "5А"
	Class5B ClassNumber = "5Б"
	Class5V ClassNumber = "5В"

	Class6A ClassNumber = "6А"
	Class6B ClassNumber = "6Б"
	Class6V ClassNumber = "6В"

	Class7A ClassNumber = "7А"
	Class7B ClassNumber = "7Б"
	Class7V ClassNumber = "7В"

	Class8A ClassNumber = "8А"
	Class8B ClassNumber = "8Б"
	Class8V ClassNumber = "8В"

	Class9A ClassNumber = "9А"
	Class9B ClassNumber = "9Б"
	Class9V ClassNumber = "9В"

	Class10A ClassNumber = "10А"
	Class10B ClassNumber = "10Б"
	Class10V ClassNumber = "10В"

	Class11A ClassNumber = "11А"
	Class11B ClassNumber = "11Б"
	Class11V ClassNumber = "11В"
)

type Class struct {
	ID        int64
	Name      ClassNumber
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClass(name ClassNumber, now time.Time) *Class {
	return &Class{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
