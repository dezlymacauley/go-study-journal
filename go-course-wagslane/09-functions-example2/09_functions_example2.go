package main

import "fmt"

func main() {

    sendsSoFar := 430
    const sendsToAdd = 25

    // NOTE: I am updating the value of sendsSoFar in the main function,
    // with the value that is returned by the incrementSends function

    sendsSoFar = incrementSends(sendsSoFar, sendsToAdd) 
    fmt.Println("You've sent", sendsSoFar, "messages")

}

// Remember that this function does not modify the value of sendsSoFar,
// in the main function of the variable!

func incrementSends(sendsSoFar int, sendsToAdd int) int {
    sendsSoFar = sendsSoFar + sendsToAdd
    return sendsSoFar
}
