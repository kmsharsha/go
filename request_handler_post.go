package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "fmt"
    "log"
)

type resp_format struct {
    Result int `json:"result"`
    Success bool `json:"success"`
}

type req_params struct {
    A int `json:"a"`
    B int `json:"b"`
}

func main() {
    http.HandleFunc("/add", add_params)
    http.HandleFunc("/sub", sub_params)
    http.ListenAndServe(":9988", nil)
}

func add_params(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Println(err)
    }

    var params req_params
    err = json.Unmarshal(body, &params)

    if err != nil {
        log.Println(err)
    }

    a := params.A
    b := params.B
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
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Println(err)
    }

    var params req_params
    err = json.Unmarshal(body, &params)
    if err != nil {
        log.Println(err)
    }

    a := params.A
    b := params.B
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
