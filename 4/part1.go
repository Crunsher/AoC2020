package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
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
	return pp.BYR != "" && pp.IYR != "" && pp.EYR != "" && pp.HGT != "" && pp.HGT != "" &&
		pp.HCL != "" && pp.ECL != "" && pp.PID != ""
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
		fmt.Println(scanner.Text())
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
		fmt.Println(pass, pass.IsValid())
		if pass.IsValid() {
			valid++
		}
	}
	fmt.Println("Valid: ", valid)

}
