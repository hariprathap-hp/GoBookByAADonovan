package main

import (
	"flag"
	"fmt"
	"strings"
)

type MyFlag struct {
	s []string
}

func (mf *MyFlag) String() string {
	var str strings.Builder
	for _, v := range mf.s {
		str.WriteString(v)
	}
	return str.String()
}

func (mf *MyFlag) Set(s string) error {
	r := strings.Split(s, ",")
	mf.s = r
	return nil
}

func ParseMyFlag(name string, value string, usage string) *[]string {
	s := strings.Split(value, ",")
	f := MyFlag{s}
	flag.CommandLine.Var(&f, name, usage)
	return &f.s
}

var test = ParseMyFlag("flg", "harry,pp", "parse string")

func main() {
	//To give the flag a default value, check the function above "ParseMyFlag" which provides default value

	//The below commented section is enough to declare an user defined flag with user defined values
	/*var f MyFlag
	flag.Var(&f, "flg", "usage")
	flag.Parse()
	fmt.Println(f.s)*/

	flag.Parse()
	fmt.Println(*test)
}
