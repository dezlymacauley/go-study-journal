package main

import "fmt"

func main() {

    messagesFromDoris := []string{
        "Hello",
        "Want to get pizza?",
        "The movie is starting",
        "How is Alex?",
    }

    numMessages := float64(len(messagesFromDoris))
    costPerMessage := 0.2

    totalCost := costPerMessage * numMessages
    
    fmt.Printf("Doris spent $%.2f on text messages today\n", totalCost)

}
