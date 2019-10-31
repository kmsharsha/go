package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "fmt"
    "log"
)

type resp_format struct {
    Result int `json:"result"`
    Success bool `json:"success"`
}

func main() {
    http.HandleFunc("/add", add_params)
    http.HandleFunc("/sub", sub_params)
    http.ListenAndServe(":9988", nil)
}

func add_params(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    a, _ := strconv.Atoi(params.Get("a"))
    b, _ := strconv.Atoi(params.Get("b"))
    sum_val := a + b

    val := resp_format{Result: sum_val, Success: true}
    resp, err := json.Marshal(val)
    if err != nil {
        fmt.Println("error =", err)
        return
    }

    log.Println(r)
    w.Header().Set("Content-Type", "application/json; char-set: utf-8")
    w.Write(resp)
}


func sub_params(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    a, _ := strconv.Atoi(params.Get("a"))
    b, _ := strconv.Atoi(params.Get("b"))
    sub_val := a - b

    val := resp_format{Result: sub_val, Success: true}
    resp, err := json.Marshal(val)
    if err != nil {
        fmt.Println("error =", err)
        return
    }

    log.Println(r)
    w.Header().Set("Content-Type", "application/json")
    w.Write(resp)
}
