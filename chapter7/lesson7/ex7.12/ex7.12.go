/*
Exercise 7.11
Измените обработчик /list так, чтобы его вывод представлял собой таблицу HTML, а не текст.
Вам может пригодиться пакет html/template (раздел 4.6).
*/

package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"socks": 5, "shoes": 50, "hat": 15}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var mu sync.Mutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func getPrice(s string) (dollars, error) {
	p, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, errors.New("you not specify the price")
	} else if p <= 0 {
		return 0, errors.New("invalid price")
	}
	return dollars(p), nil
}

func getName(s string) (string, error) {
	if len(s) > 0 {
		return s, nil
	}
	return "", errors.New("you not specify the item name")
}

type database map[string]dollars

func (db database) isExist(s string) bool { _, ok := db[s]; return ok }

func (db database) list(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "database is empty. Use create")
		return
	}
	w.WriteHeader(http.StatusOK)
	itemList := template.Must(template.ParseFiles("chapter7/lesson7/ex7.12/list.html"))
	if err := itemList.Execute(w, &db); err != nil {
		log.Println(err)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item, err := getName(r.URL.Query().Get("item"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	if !db.isExist(item) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item: %s not found", item)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", db[item])
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	var errs []error
	i, err := getName(r.URL.Query().Get("item"))
	if err != nil {
		errs = append(errs, err)
	}
	p, err := getPrice(r.URL.Query().Get("price"))
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		http.Error(w, fmt.Sprintf("%s", errors.Join(errs...)), http.StatusBadRequest)
		return
	}
	if db.isExist(i) {
		http.Error(w, fmt.Sprintf("%s is exist", i), http.StatusConflict)
		return
	}
	mu.Lock()
	db[i] = p
	mu.Unlock()
	http.Error(w, fmt.Sprintf("item %s with price %s: created sucessfully", i, db[i]), http.StatusOK)
}

func (db database) read(w http.ResponseWriter, r *http.Request) { db.list(w, r) }

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	i, err := getName(r.URL.Query().Get("item"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	if !db.isExist(i) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "the item %s is not exist in DB. Use create", i)
		return
	}
	mu.Lock()
	delete(db, i)
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s deleted sucessfully", i)

}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	var errs []error
	i, err := getName(r.URL.Query().Get("item"))
	if err != nil {
		errs = append(errs, err)
	}
	p, err := getPrice(r.URL.Query().Get("price"))
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error(s) encountered: \n%v", errors.Join(errs...))
		return
	}
	if !db.isExist(i) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "the item %s is not exist in DB. Use create", i)
		return
	}
	if db[i] == p {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "you specify the same price: %s as item %s aready have", p, i)
		return
	}
	mu.Lock()
	oldPrice := db[i]
	db[i] = p
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item %s with %s update. Now it's price is %s", i, oldPrice, db[i])
}
