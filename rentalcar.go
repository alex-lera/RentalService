package main

import (
        "fmt"
        _ "log"
        "net/http"
        "math/rand"
        "github.com/gorilla/mux"
        "encoding/json"
        "io"
        "os"
        "io/ioutil"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "gopkg.in/yaml.v2"
    ) 

type RequestMessage struct {
    Brand string
    Model string
    HorsePow  string
}

type Config struct {
    Server struct {
        Dbname string `yaml:"dbname"`
    } `yaml:"server"`
    Database struct {
        Username string `yaml:"user"`
        Password string `yaml:"pass"`
    } `yaml:"database"`
}

var cfg Config

func main() {

f, err := os.Open("/opt/config.yaml")
    
decoder := yaml.NewDecoder(f)
err = decoder.Decode(&cfg)       

if err != nil {
    return
}

router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/service/v1/cars", newrentalInput)

http.ListenAndServe(":8080", router)
}

func newrentalInput(w http.ResponseWriter, r *http.Request) {
    
    if(r.Method != "POST"){
        w.WriteHeader(405)
        return
    }

    var requestMessage RequestMessage
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    
    if err != nil {
        w.WriteHeader(500)
        return
    }
    
    if err := r.Body.Close(); err != nil {
        w.WriteHeader(500)
        return
    }
    
    if err := json.Unmarshal(body, &requestMessage); err != nil {
        
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        
        if err := json.NewEncoder(w).Encode(err); err != nil {
            w.WriteHeader(500)
            return
        }

    }
        
    conn, err := sql.Open("mysql", cfg.Database.Username+":"+cfg.Database.Password+"@tcp("+cfg.Server.Dbname+")/cars")

    if err != nil {
        w.WriteHeader(500)
        conn.Close()
        return
    }

    statement, err := conn.Prepare("insert into rentals values(?, ?, ?, ?)")

    if err != nil {
        w.WriteHeader(500)
        conn.Close()
        return
    }

    id_db := rand.Intn(10000000)

    rows, err := statement.Exec(id_db, requestMessage.Brand, requestMessage.Model, requestMessage.HorsePow)

    if err != nil && rows != nil{
        w.WriteHeader(500)
        conn.Close()
        return
    }

    w.WriteHeader(200)

    fmt.Fprintln(w, "Hola. ID = ", id_db, ", Brand =", requestMessage.Brand, ", Model =", requestMessage.Model, ", Horse Power =", requestMessage.HorsePow)

    conn.Close()
    
}

func TestDatabase()(int){
	
	conn := cfg.Database.Username+":"+cfg.Database.Password+"@tcp("+cfg.Server.Dbname+")/cars"

    if conn != "dbuser:dbuser@tcp(mymariadb)/cars"{
        return -1
    }

    return 0

}
