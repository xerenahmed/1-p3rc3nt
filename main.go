package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	start int
	realization float64

	breakpointsRaw string
	breakpoints = make(map[int]bool)

	high int
)

func init() {
	flag.IntVar(&start, "s", 100, "start")
	flag.Float64Var(&realization, "r", 1, "realization")
	flag.StringVar(&breakpointsRaw, "b", "", "breakpoints")
	flag.Parse()

	s := strings.Split(breakpointsRaw, ",")
	if breakpointsRaw == "" || len(s) == 0 {
		panic("you must specify breakpoints")
	}

	for _, b := range s {
		bi, err := strconv.Atoi(b)
		if err != nil {
			panic(fmt.Errorf("invalid breakpoint: %v", err))
		}

		breakpoints[bi] = true
		if high < bi {
			high = bi
		}
	}
}

func main() {
	balance := float64(start)
	for i := 0; i <= high; i++ {
		daily := (balance * realization) / 100
		balance += daily

		if breakpoints[i] {
			fmt.Printf("Stage %d\n- Balance: %.2f$\n- Daily Earn: %.2f$\n- Away Start: %%%.2f \n\n",
				i, balance, daily, (balance / float64(start) * 100) - 100)
		}
	}
}
