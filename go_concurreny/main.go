package main

import(
    "fmt"
    "time"
)

func say(s string){
    for i:=0; i<5; i++{
        fmt.Println("Inside for block", time.Now())
        time.Sleep(100*time.Millisecond)
        fmt.Println(time.Now())
        fmt.Println(s)
    }
}


func main(){
    fmt.Println(time.Now())
    go say("world")

    say("hello")
}