package main

import "fmt"
import "log"

func main(){
    x := 1
    y := 3
    sum, sub := calculate(x, y)
    fmt.Println("sum::w", sum, "sub:", sub)
    log.Println("Sum:", sum)
}

func calculate(x int, y int) (int, int) {
    return x + y, x - y
}
