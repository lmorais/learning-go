package main

import "fmt"

type Rectangle struct {
    length, width int
    name string
}

func main() {

    //initialize values in order they are defined in struct
    r1 := Rectangle{2, 1, "my_r1"}
    fmt.Println("Rectangle r1 is: ", r1)

    //initialize values by variable name in any order
    r2 := Rectangle{width:3, name:"my_r2", length:4}
    fmt.Println("Rectangle r2 is: ", r2)

    //get pointer to an instance with new keyword
    pr := new (Rectangle)

    //set value using . notation by dereferencing pointer.
    (*pr).width = 6

     //set value using . notation - same as previous.  There is no -> operator like in c++. Go automatically converts
    pr.length = 8
    pr.name = "ptr_to_rectangle"

    //Go performs default printing of structs
    fmt.Println("Rectangle pr as address is: ", pr)

    //address and value are differentiated with an & symbol
    fmt.Println("Rectangle pr as value is: ", *pr)
}