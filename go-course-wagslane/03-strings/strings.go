package main

import "fmt"

func main() {
    var username string = "dezlymacauley"
    var password string = "argentina828628"

    fmt.Println("Authorization: Basic", username+":"+password)
    
    // This will print:
    // Authorization: Basic dezlymacauley:argentina828628

}
