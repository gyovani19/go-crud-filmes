package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Filme struct {
	ID      string   `json:"id"`
	ISBN    string   `json:"isbn"`
	Titulo  string   `json:"title"`
	Diretor *Diretor `json:"diretor"`
}

type Diretor struct {
	PrimeiroNome string `json:"primeiroNome"`
	UltimoNome   string `json:"ultimoNome"`
}

var filmes []Filme

func getFilmes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filmes)
}

func deleteFilme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range filmes {
		if item.ID == params["id"] {
			filmes = append(filmes[:index], filmes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(filmes)
}

func getFilme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range filmes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createFilme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filme Filme
	_ = json.NewDecoder(r.Body).Decode(&filme)
	filme.ID = strconv.Itoa(rand.Intn(1000000))
	filmes = append(filmes, filme)
	json.NewEncoder(w).Encode(filme)
}

func updateFilme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range filmes {
		if item.ID == params["id"] {
			filmes = append(filmes[:index], filmes[index+1:]...)
			var filme Filme
			_ = json.NewDecoder(r.Body).Decode(&filme)
			filme.ID = params["id"]
			filmes = append(filmes, filme)
			json.NewEncoder(w).Encode(filme)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	filmes = append(filmes, Filme{ID: "1", ISBN: "12345", Titulo: "Era Uma Vez em... Hollywood", Diretor: &Diretor{PrimeiroNome: "Quentin", UltimoNome: "Tarantino"}})
	filmes = append(filmes, Filme{ID: "2", ISBN: "47890", Titulo: "O Irlandês", Diretor: &Diretor{PrimeiroNome: "Martin", UltimoNome: "Scorsese"}})
	r.HandleFunc("/filmes", getFilmes).Methods("GET")
	r.HandleFunc("/filmes/{id}", getFilme).Methods("GET")
	r.HandleFunc("/filmes", createFilme).Methods("POST")
	r.HandleFunc("/filmes/{id}", updateFilme).Methods("PUT")
	r.HandleFunc("/filmes/{id}", deleteFilme).Methods("DELETE")

	fmt.Printf("Servidor rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
