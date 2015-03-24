package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	fmt.Println(v)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
	k := reflect.TypeOf(*r)
	v := reflect.ValueOf(*r)
	for i := 0; i < v.NumField(); i++ {
		fmt.Fprintf(w, "%v : %v\n", k.Field(i), v.Field(i).Interface())
	}
}
