package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	numWords = flag.Int("numwords", 7, "number of words in pass phrase")
	diceFile = flag.String("dicefile", "diceware.wordlist.asc", "path to diceware word list")
)

const (
	ExpectedNumEntries = 6 * 6 * 6 * 6 * 6
	RollWidth          = 5
)

func main() {
	flag.Parse()
	numDice := *numWords * RollWidth
	b := make([]byte, numDice)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	dieRolls := make([]DieRoll, *numWords)
	for i := range dieRolls {
		dieRolls[i] = newDieRoll(b[i*RollWidth : (i+1)*RollWidth])
	}
	wordList, err := parseFile()
	if err != nil {
		log.Fatal(err)
	}
	for i, roll := range dieRolls {
		fmt.Printf("%s", wordList[roll])
		if i < len(dieRolls)-1 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
	for i, roll := range dieRolls {
		fmt.Printf("%s", wordList[roll])
		if i == len(dieRolls)-1 {
			fmt.Printf("\n")
		}
	}
}

func parseFile() (map[DieRoll]string, error) {
	file, err := os.Open(*diceFile)
	if err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	scanner := bufio.NewScanner(file)
	wordList := make(map[DieRoll]string)
	words := make(map[string]bool)
	line := 1
	for scanner.Scan() {
		rollWord := strings.Fields(scanner.Text())
		if len(rollWord) != 2 {
			return nil, fmt.Errorf("parse error: not a pair line %d\n", line)
		}
		if len(rollWord[0]) != RollWidth {
			return nil, fmt.Errorf("parse error: not a die-roll line %d\n", line)
		}
		dieRoll, err := newDieRollASCII(rollWord[0])
      if err != nil {
         return nil, fmt.Errorf("parse error: %v line %d\n", err, line)
      }
		if _, ok := wordList[dieRoll]; ok {
			return nil, fmt.Errorf("parse error: duplicate entry %v line %d\n",
				dieRoll, line)
		} else if _, ok := words[rollWord[1]]; ok {
			return nil, fmt.Errorf("parse error: duplicate word %s line %d\n",
				rollWord[1], line)
		}
		wordList[dieRoll] = rollWord[1]
		line += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if len(wordList) != ExpectedNumEntries {
		return nil, fmt.Errorf("parse error: should be %d entries, but found %d",
			ExpectedNumEntries, len(wordList))
	}
	return wordList, nil
}

type DieRoll [RollWidth]byte

func newDieRoll(bytes []byte) DieRoll {
	return DieRoll{
		bytes[0] % 6,
		bytes[1] % 6,
		bytes[2] % 6,
		bytes[3] % 6,
		bytes[4] % 6,
	}
}

func newDieRollASCII(asciiRoll string) (DieRoll, error) {
	roll := DieRoll{
		asciiRoll[0] - '1',
		asciiRoll[1] - '1',
		asciiRoll[2] - '1',
		asciiRoll[3] - '1',
		asciiRoll[4] - '1',
	}
	if !roll.valid() {
		return roll, fmt.Errorf("newDieRollASCII: invalid roll %v", roll)
	} else {
		return roll, nil
	}
}

func (dr DieRoll) valid() bool {
	return dr[0] >= 0 && dr[0] < 6 &&
		dr[1] >= 0 && dr[1] < 6 &&
		dr[2] >= 0 && dr[2] < 6 &&
		dr[3] >= 0 && dr[3] < 6 &&
		dr[4] >= 0 && dr[4] < 6
}
