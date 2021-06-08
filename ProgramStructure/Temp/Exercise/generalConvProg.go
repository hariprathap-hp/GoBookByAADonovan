package main

import (
	"TheGoProgLangBook/ProgramStructure/Temp/tempconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//check if command line arguments are present
	if len(os.Args[1:]) > 0 {
		for _, args := range os.Args[1:] {
			v, err := strconv.ParseFloat(args, 64)
			if err != nil {
				fmt.Println(err)
			}
			m := tempconv.Metres(v)
			ft := tempconv.Feet(v)

			fmt.Printf("%s = %s, %s = %s", m, tempconv.MToF(m), ft, tempconv.FToM(ft))
		}
	} else {
		//If command Line arguments are not present, get it via standard io console
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			n := (scanner.Text())
			if n == "end" {
				break
			} else {
				v, err := strconv.ParseFloat(n, 64)
				if err != nil {
					fmt.Println(err)
				}
				m := tempconv.Metres(v)
				ft := tempconv.Feet(v)

				fmt.Printf("%s = %s, %s = %s\n", m, tempconv.MToF(m), ft, tempconv.FToM(ft))
			}
		}

	}
}
