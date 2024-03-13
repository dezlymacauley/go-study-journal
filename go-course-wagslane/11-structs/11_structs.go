package main

import "fmt"

type messageToSend struct {
    phoneNumber int
    message string
}

// The variable msg is of type messageToSend,
// which is a struct that has two fields
func test(msg messageToSend) {
    fmt.Printf("Sending message: '%s' to: %v\n", msg.message, msg.phoneNumber)
    fmt.Println("========================")
}

func main() {
    
    test(messageToSend{
        phoneNumber: 2729927,
        message: "Thanks for signing up",
    })

}
