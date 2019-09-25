package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
)


type Beers struct {
  ID string `json:"id"`
  name string `json:"name"`
  brewery string `json:"Brewery"`
  country string `json:"Country"`
  price float64 `json:"Price"`
  currency string `json:"Currency"`
}

type BeerBox struct {
  PriceTotal float64 `json:"Price Total"`
  
}



var db *sql.DB
var err error
func main() {
db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/Bender")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  router := mux.NewRouter()

  router.HandleFunc("/Beers", getBeers).Methods("GET")

  router.HandleFunc("/Beers", createBeer).Methods("POST")
  router.HandleFunc("/Beers/{id}", getBeer).Methods("GET")
  
  router.HandleFunc("/Beers/{id}/boxprice", getBoxprice).Methods("GET")
  



  router.HandleFunc("/Beers/{id}", updateBeer).Methods("PUT")
  router.HandleFunc("/Beers/{id}", deleteBeer).Methods("DELETE")


  http.ListenAndServe(":8000", router)
}



func getBeers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var beers []Beers
  result, err := db.Query("SELECT id, name, brewery, country, price, currency from beer")
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var beer Beers
    err := result.Scan(&beer.ID, &beer.name  , &beer.brewery, &beer.country, &beer.price, &beer.currency )
    if err != nil {
      panic(err.Error())
    }
    beers = append(beers, beer)
  }
  json.NewEncoder(w).Encode(beers)
}



func createBeer(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  stmt, err := db.Prepare("INSERT INTO beer( name) VALUES(?)")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  name := keyVal["name"]
  _, err = stmt.Exec(name)
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "New Beer was created")
}





func getBeer(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  result, err := db.Query("SELECT id, name FROM beer WHERE id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  var beer Beers
  for result.Next() {
    err := result.Scan(&beer.ID, &beer.name)
    if err != nil {
      panic(err.Error())
    }
  }
  json.NewEncoder(w).Encode(beer)
}





func updateBeer(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("UPDATE beer SET name = ? WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  newName := keyVal["name"]
  _, err = stmt.Exec(newName, params["id"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Beer with ID = %s was updated", params["id"])
}






func deleteBeer(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("DELETE FROM beer WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec(params["id"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Beer with ID = %s was deleted", params["id"])
}





func getBoxprice(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  result, err := db.Query("SELECT id, name , price , quantity FROM beer WHERE id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }


 


  defer result.Close()
  var beer Beers
  for result.Next() {
    err := result.Scan(&beer.ID, &beer.name)
    if err != nil {
      panic(err.Error())
    }
  }
  json.NewEncoder(w).Encode(beer)
}




