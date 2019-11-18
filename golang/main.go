package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Passage struct {
	Company string `json:"Company"`
	Money string `json:"Money"` 
}

type Passages []Passage

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func allPassages(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	passages := Passages{
		Passage{Company:"GOL", Money:"1000"},
		Passage{Company:"TAM", Money:"2000"},
		Passage{Company:"AIR", Money:"500"},
	}

	fmt.Println("Todos os dados mostrados", passages)
	json.NewEncoder(w).Encode(passages)
}


func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Pagina inicial")
}

func handleRequests(){
	http.HandleFunc("/", home)
	http.HandleFunc("/passages", allPassages)
	log.Fatal(http.ListenAndServe(":8686", nil))
}

func main(){
	handleRequests()
}