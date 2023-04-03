/*
Exercise 7.16
Напишите программу-калькулятор для веб.
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"GolangBook/chapter7/lesson9/ex7.16/eval"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("chapter7/lesson9/ex7.16/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
		return
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("chapter7/lesson9/ex7.16/index.html"))

	s := r.PostFormValue("expr")
	if s == "" {
		fmt.Fprintf(w, "empty expr")
		return
	}
	expr, err := eval.Parse(s)
	if err != nil {
		fmt.Errorf("%w", err)
		return
	}
	env, err := parseEnv(r.PostFormValue("env"))
	if err != nil {
		fmt.Errorf("%w", err)
		return
	}
	if err := tmpl.Execute(w, expr.Eval(env)); err != nil {
		log.Fatal(err)
		return
	}
}

func parseEnv(s string) (eval.Env, error) {
	env := make(eval.Env)
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(`:=[]{},\"`, r) || unicode.IsSpace(r)
	})
	for i := 0; i < len(fields); i += 2 {
		k := strings.TrimSpace(fields[i])
		v := strings.TrimSpace(fields[i+1])
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		env[eval.Var(k)] = val
	}
	return env, nil
}
