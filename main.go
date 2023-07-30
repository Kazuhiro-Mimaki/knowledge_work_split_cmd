package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
    var (
        l int
    )

    flag.IntVar(&l, "l", 0, "line_count")

    flag.Parse()

    err := validateL(l)
    if err != nil {
        panic(err)
    }

    fmt.Println("pass")
}

func validateL(l int) (err error) {
    if l < 0 {
        return errors.New("validation error of l")
    }
    return nil
}
