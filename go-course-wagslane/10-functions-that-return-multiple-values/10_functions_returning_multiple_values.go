package main

import "fmt"

func main() {

    // NOTE: The _ is a way to tell Go to that this variable is unused 

    // The first string will be assigned to "firstName",
    // The second variable "Doe" will be unused 

    firstName, _ := getName()
    fmt.Println("Welcome to Texio,", firstName)

}

// This function return to variables
func getName() (string,string) {
    return "John", "Doe"
}
