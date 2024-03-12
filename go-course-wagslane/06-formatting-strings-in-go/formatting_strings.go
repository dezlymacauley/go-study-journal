package main

import "fmt"

func main() {

    // NOTE: %v use this if you are not sure what to use     
    // \n means "print the next statement on a new line"

    fmt.Printf("I am %v years old\n", 30)
    fmt.Printf("%v is here\n", "Dezly Macauley")

    // NOTE: Floats
    // %f is for floats
    // You can also specify the number of decimal points
    // %.2f

    // NOTE: Strings
    // %s for strings

    // NOTE: Integers
    // %d

    // NOTE: What is the difference between "Printf()"
    // "Sprintf()"

    // Printf() will print out the line to standard out.
    // I.e. The print to to console

    // Sprintf() is when you want the value that is being formatted,
    // to be stored in a variable

    const userName = "Dezly Macauley"
    const ticketNumber = 30.5

    openRateMessage := fmt.Sprintf(
        "Hi %s, your ticket number is %.1f", userName, ticketNumber,
    )

    fmt.Println(openRateMessage)

}
