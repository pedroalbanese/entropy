package main

import (
    "fmt"
    "flag"
    "math"
    "strings"
)

var (
	key = flag.String("k", "", "Passphrase/password.")
)

func main(){
    flag.Parse()
    fmt.Println(H(*key))
}

// for ASCII strings
func H(data string) (entropy float64) {
    if data == "" {
        return 0
    }
    for i := 0; i < 256; i++ {
        px := float64(strings.Count(data, string(byte(i)))) / float64(len(data))
        if px > 0 {
	    entropy += -px * math.Log2(px)
	}
    }
    return entropy
}