/*
Exercise 7.15
Напишите программу, которая читает из стандартного ввода единственное выражение, предлагает
пользователю ввести значения переменных, а затем вычисляет выражение в полученной среде.
Аккуратно обработайте все ошибки.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"GolangBook/chapter7/lesson9/ex7.15/eval"
)

func main() {
	fmt.Print("Enter the expr: ")
	var f string
	_, err := fmt.Scan(&f)
	if err != nil {
		log.Fatal(err)
	}
	expr, err := eval.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	env := Env(expr)
	fmt.Printf("%s = %g\n", f, expr.Eval(env))
}

func Env(expr eval.Expr) eval.Env {
	env := make(eval.Env)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for k := range getVars(expr) {
		fmt.Printf("%s: ", k)
		if !scanner.Scan() {
			log.Fatalf("not enough var!")
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		env[k] = val
	}
	return env
}

func getVars(expr eval.Expr) map[eval.Var]bool {
	vars := make(map[eval.Var]bool, len(expr.Vars()))
	for _, v := range expr.Vars() {
		vars[v] = true
	}
	return vars
}
