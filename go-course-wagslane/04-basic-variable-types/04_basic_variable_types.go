package main

import "fmt"

func main() {

    // NOTE: bool (True or False)

    // bool

    // NOTE: string (Text)
    
    // string
    
    // NOTE: Signed integers (Whole numbers that can be positive or negative)

    // int, int8, int16, int32, int64

    // NOTE: rune (same as int32)

    // A run is an alias for int32
    // It represents a unicode code point


    // NOTE: Unsigned integers 
    
    // uint, uint8, uint16, uint32, uint64, uintptr

    // NOTE: byte (same as uint8)
    
    // A byte is an alias for uint8

    // NOTE: float
    
    // float32, float64

    // NOTE: Complex Numbers
    
    // complex64, complex128

//-----------------------------------------------------------------------------

// NOTE: What does the size of a type indicate?

// E.g. float32
// The 32 at the end means 32 bits
// a byte is 8 bits
// a nibble is 4 bits

//-----------------------------------------------------------------------------

// NOTE: How to declare a variable in Go

// In Go, variables are declared using the keyword "var"

var smsSendingLimit int
var costPerSMS float64
var hasPermission bool
var username string

fmt.Printf(
    "%v %f %v %q\n",
    smsSendingLimit,
    costPerSMS,
    hasPermission,
    username,

    // %v is a placeholder for a value of any type. 
    // %f is a placeholder for a floating-point number. 
    // %q is a placeholder for a quoted string. 

    // This will print 0 0.000000 false ""
)

//-----------------------------------------------------------------------------

// NOTE: How to use the short assignment operator in Go
// := 
// When you use this operator: 
// You do not need to add "var"
// Go will infer the type of the variable
// In Go, most of the time, you will see variables declared,
// using the short assignment operator

// E.g. var username string
// is the same as username := string

congrats := "happy birthday"
fmt.Println(congrats)


//-----------------------------------------------------------------------------

// NOTE: How to declare constants in Go
// You can't use the short assignment indicator with constants in Go

// Constants are declared using the const keyword

const currencyName = "Euro"

//-----------------------------------------------------------------------------

}
