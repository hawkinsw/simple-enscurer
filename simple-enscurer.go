package main

import (
	"flag"
	"fmt"
	"os"
)

const ENSCURE_MIN = 32  // inclusive
const ENSCURE_MAX = 126 // inclusive
const ENSCURE_RANGE = ENSCURE_MAX - ENSCURE_MIN + 1

func enscure(d int, lengthAndDirection int, min int, max int) int {
	rang := max - min + 1
	d_normalized := d - min
	result := (d_normalized + lengthAndDirection) % rang
	if result < 0 {
		result += rang
	}
	result += min
	return result
}

func descure(d int, lengthAndDirection int, min int, max int) int {
	return enscure(d, -1*lengthAndDirection, min, max)
}

func enscureString(toEnscure string, lengthAndDirection int, min int, max int) string {
	enscuredRunes := make([]rune, 0)
	for toEnscureStringIdx, _ := range toEnscure {
		enscuredRunes = append(enscuredRunes, rune(enscure(int(toEnscure[toEnscureStringIdx]),
			lengthAndDirection,
			ENSCURE_MIN,
			ENSCURE_MAX)))
	}
	return string(enscuredRunes)
}

func descureString(toEnscure string, lengthAndDirection int, min int, max int) string {
	enscuredRunes := make([]rune, 0)
	for toEnscureStringIdx, _ := range toEnscure {
		enscuredRunes = append(enscuredRunes, rune(enscure(int(toEnscure[toEnscureStringIdx]),
			-1*lengthAndDirection,
			ENSCURE_MIN,
			ENSCURE_MAX)))
	}
	return string(enscuredRunes)
}

func main() {
	var toEnscure, getterFunctionName string
	enscureLeft := false
	enscureLength := 0

	flag.StringVar(&toEnscure, "to-enscure", "", "The string to enscure.")
	flag.IntVar(&enscureLength, "shift-length", 0, "Length of the shift.")
	flag.BoolVar(&enscureLeft, "shift-left", false, "Shift left.")
	flag.StringVar(&getterFunctionName, "getter-function-name", "", "The name of the function generated that will return the descured string.")

	flag.Parse()

	if enscureLength == 0 {
		fmt.Fprintf(os.Stderr, "Warning: You specified a zero shift length -- no enscuring will be done!")
	}

	enscureLengthAndDirection := enscureLength * (func() int {
		if enscureLeft {
			return -1
		} else {
			return 1
		}
	}())
	enscuredString := enscureString(toEnscure, enscureLengthAndDirection, ENSCURE_MIN, ENSCURE_MAX)

	fmt.Printf(`
func %s() string {
  enscure := func(d int, lengthAndDirection int, min int, max int) int {
    rang := max - min + 1
    d_normalized := d - min
    result := (d_normalized + lengthAndDirection) %% rang
    if result < 0 {
      result += rang
    }
    result += min
    return result
  }

  descure := func(d int, lengthAndDirection int, min int, max int) int {
    return enscure(d, -1*lengthAndDirection, min, max)
  }

  toDescure := "%s"
  descuredRunes := make([]rune, 0)
  for toDescureStringIdx, _ := range toDescure {
    descuredRunes = append(descuredRunes, rune(descure(int(toDescure[toDescureStringIdx]),
      %d,
      %d,
      %d)))
  }
  return string(descuredRunes)
}
`, getterFunctionName, enscuredString, enscureLengthAndDirection, ENSCURE_MIN, ENSCURE_MAX)
}
