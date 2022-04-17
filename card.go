package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	Value int
	Suit  string
}

var suits = []string{"♧", "♤", "♢", "♡"}

func (c Card) ValueToString() string {
	var value string
	switch c.Value {
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	case 14:
		value = "A"
	default:
		value = strconv.Itoa(c.Value)
	}
	return value
}

func StringToFullString(v string) string {
	var value string
	switch v {
	case "J":
		value = "Jack"
	case "Q":
		value = "Queen"
	case "K":
		value = "King"
	case "A":
		value = "Ace"
	default:
		value = v
	}
	return value
}

func StringToValue(s string) int {
	var value int
	switch s {
	case "J":
		value = 11
	case "Q":
		value = 12
	case "K":
		value = 13
	case "A":
		value = 14
	default:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		value = int(v)
	}
	return value
}

func (c Card) String() string {
	return fmt.Sprintf("%v%-4v", c.ValueToString(), c.Suit)
}
