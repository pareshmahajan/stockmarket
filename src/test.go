package main

import  (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}


type Exchange struct {
    Symbol string `json:"Symbol"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Exchanges []Exchange





func handleRequests() {
    http.HandleFunc("/stocks", getStock)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	Exchanges = []Exchanges{
        Exchange{Symbol: "NSE", Desc: "NSE Description", Content: "NSE Content"},
        Exchange{Symbol: "NASDAQ", Desc: "NASDAQ Description", Content: "NASDAQ Content"},
    }
    handleRequests()
}


func getStock(w http.ResponseWriter, r *http.Request) {
	stockExchageSymbol := mux.Vars(r)["Symbol"]
	json.NewEncoder(w).Encode("get")
	json.NewEncoder(w).Encode(Exchanges)

	// for _, exchange := range events {
	// 	if exchange.symbol == stockExchageSymbol {
	// 		json.NewEncoder(w).Encode(exchange)
	// 	}
	// }
}