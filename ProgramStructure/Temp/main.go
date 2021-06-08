package main

import (
	"TheGoProgLangBook/ProgramStructure/Temp/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, args := range os.Args[1:] {
		t, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Println(err)
		}
		f := tempconv.Farenheit(t)
		c := tempconv.Celcius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
