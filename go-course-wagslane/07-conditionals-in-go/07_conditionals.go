package main

import "fmt"

func main() {

    messageLength := 10
    maxMessageLength := 20

    // FIX: Learn how to use Println to display this

//     fmt.Println("Trying to send a message of length:",
//     messageLength, "and a max length of:", maxMessageLength)  
// )

// This should print:
// "Trying to send a message of length: 10 and max length of: 20"

    if messageLength <= maxMessageLength {
        fmt.Println("Message Sent")
    } else {
        fmt.Println("Message not sent")
    }


}
