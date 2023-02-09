/*
Exercise 2-6-1 (2.2)
lesson2.2 prints measurements given on the command line or stdin in various units
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"GolangBook/chapter2/lesson6/lenghtconv"
	"GolangBook/chapter2/lesson6/tempconv"
	"GolangBook/chapter2/lesson6/weightconv"
)

var isLength, isWeight, isTemperature bool

func printMeasurement(arg string) {
	n, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		if arg == "exit" {
			fmt.Print("Shutdown...")
			os.Exit(1)
		}
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	switch {
	case isLength:
		fmt.Println("Selected mode `-l`")
		f := lenghtconv.Feet(n)
		m := lenghtconv.Meter(n)
		fmt.Printf("%s = %s\n%s = %s\n", f, lenghtconv.FToM(f), m, lenghtconv.MToF(m))
	case isWeight:
		fmt.Println("Selected mode `-w`")
		k := weightconv.Kilos(n)
		p := weightconv.Pounds(n)
		fmt.Printf("%s = %s\n%s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
	case isTemperature:
		fmt.Println("Selected mode `-t`")
		c := tempconv.Celsius(n)
		f := tempconv.Fahrenheit(n)
		k := tempconv.Kelvin(n)
		fmt.Printf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
			c, tempconv.CToF(c), tempconv.CToK(c),
			f, tempconv.FToC(f), tempconv.FToK(f),
			k, tempconv.KToC(k), tempconv.KToF(k))
	}
}

func main() {
	flag.BoolVar(&isLength, "l", false, "Converts meters and feet")
	flag.BoolVar(&isWeight, "w", false, "Converts kilograms and pounds")
	flag.BoolVar(&isTemperature, "t", true, "Converts Celsius, Fahrenheit and Kelvin temperatures")
	flag.Parse()
	if len(flag.Args()) > 0 {
		for _, arg := range flag.Args() {
			printMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		fmt.Print("Write a value: ")
		for scan.Scan() {
			printMeasurement(scan.Text())
		}
	}
}
