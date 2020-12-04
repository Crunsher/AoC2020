package main

import (
	"bufio"
	"container/list"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	BYR string
	IYR string
	EYR string
	HGT string
	HCL string
	ECL string
	PID string
	CID string
}

func (pp Passport) IsValid() bool {
	if !(pp.BYR != "" && pp.IYR != "" && pp.EYR != "" && pp.HGT != "" && pp.HGT != "" &&
		pp.HCL != "" && pp.ECL != "" && pp.PID != "") {
		return false
	}

	byr, _ := strconv.Atoi(pp.BYR)
	if byr > 2002 || byr < 1920 {
		return false
	}

	iyr, _ := strconv.Atoi(pp.IYR)
	if byr > 2020 || iyr < 2010 {
		return false
	}

	eyr, _ := strconv.Atoi(pp.EYR)
	if eyr > 2030 || eyr < 2020 {
		return false
	}

	if pp.HGT[len(pp.HGT)-2:] == "cm" {
		hgt, _ := strconv.Atoi(pp.HGT[:len(pp.HGT)-2])
		if hgt > 193 || hgt < 150 {
			return false
		}
	} else if pp.HGT[len(pp.HGT)-2:] == "in" {
		hgt, _ := strconv.Atoi(pp.HGT[:len(pp.HGT)-2])
		if hgt > 76 || hgt < 59 {
			return false
		}
	} else {
		return false
	}

	if pp.HCL[0] != '#' || len(pp.HCL) != 7{
		return false
	}

	_, err := hex.DecodeString(pp.HCL[1:])
	if err != nil {
		return false
	}

	eclFound := false
	for _, ecl := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if pp.ECL == ecl {
			eclFound = true
			break
		}
	}
	if !eclFound {
		return false
	}

	_, err = strconv.Atoi(pp.PID)
	if len(pp.PID) != 9 || err != nil {
		return false
	}

	return true
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ll := list.New()

	var pp Passport
	for scanner.Scan() {
		if scanner.Text() == "" {
			ll.PushBack(pp)
			pp = Passport{}
			continue
		}
		kvs := strings.Fields(scanner.Text())
		for _, kv := range kvs {
			value := strings.Split(kv, ":")
			switch value[0] {
			case "byr":
				pp.BYR = value[1]
			case "iyr":
				pp.IYR = value[1]
			case "eyr":
				pp.EYR = value[1]
			case "hgt":
				pp.HGT = value[1]
			case "hcl":
				pp.HCL = value[1]
			case "ecl":
				pp.ECL = value[1]
			case "pid":
				pp.PID = value[1]
			case "cid":
				pp.CID = value[1]
			}
		}
	}
	ll.PushBack(pp)
	valid := 0
	for e := ll.Front(); e != nil; e = e.Next() {
		pass := e.Value.(Passport)
		if pass.IsValid() {
			valid++
		}
	}
	fmt.Println("Valid: ", valid)

}
