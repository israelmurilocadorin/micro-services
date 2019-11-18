package main

import (
	"fat"
	"log"
	"net/http"
)

type Passage struct{
	Company string 'json:"Company"'
	Money string 'json:"Money"'
}

type Passages []Passage

func allPassages(w http.ResponseWriter, r *http.Request){
	passages != Passages{
		Passage{
			Companhia: "Test 1",
			Valor: "Test 2"
		},
	}


	fat.Println("Todos os dados mostrados")
	json.NewEncoder(w).Encode(passages)
}


func home(w http.ResponseWriter, r *http.Request){
	fat.Fprintf(w, "Pagina inicial")
}

func handleRequests(){
	http.HandleFunc("/", home)
	http.HandleFunc("/Passages", allPassages)

	log.Fatal(http.ListenAndServer(":8686", nil))
}

func main(){
	handleRequests()
}