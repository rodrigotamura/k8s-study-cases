package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", HealthZ)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file", err)
	}

	fmt.Fprintf(w, "My family: %s.", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s. Pass: %s.", user, password)
}

func HealthZ(w http.ResponseWriter, r *http.Request) {

	// calculando tempo que iniciou o fluxo da aplicação
	// com o tempo atual
	duration := time.Since(startedAt)

	// if duration.Seconds() < 10 || duration.Seconds() > 25 {
	if duration.Seconds() < 10 {
		// se a aplicação não alcançar 10 segundos após a criação, estoura erro 500
		// quando passar 25 segundos que a minha aplicação está no ar
		// vamos retornar um header com erro 500
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("ok")))
	}

}
