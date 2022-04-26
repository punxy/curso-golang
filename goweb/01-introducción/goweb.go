package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El método es : " + r.Method)
	fmt.Fprintln(rw, "Hola Mundo!")
}

func pageNT(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Página no encontrada", http.StatusNotFound)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	edad := r.URL.Query().Get("edad")

	fmt.Fprintf(rw, "Hola %s, tu edad es %s años", name, edad)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/notfound", pageNT)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// // Router
	// http.HandleFunc("/", Hola)
	// http.HandleFunc("/notfound", pageNT)
	// http.HandleFunc("/error", Error)
	// http.HandleFunc("/saludar", Saludar)

	// crear servidor
	fmt.Println("El servidor está corriendo en el puerto 9000")
	fmt.Println("Run server: http://localhost:9000")

	//log.Fatal(http.ListenAndServe("localhost:9000", mux))

	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
