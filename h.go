package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	fmt.Println(lookupRevRecord("49.181.142.165"))
	fmt.Println(lookupRevRecord("8.8.8.8"))
	fmt.Println(lookupRevRecord("127.0.0.1"))
	fmt.Println(lookupRevRecord("1.2.3.4"))
}

func lookupRevRecord(h string) (string, error) {
	addrs, err := net.LookupAddr(h)
	if err != nil {
		fmt.Println("Reverse Lookup Failed:", err)
	}

	for len(addrs) == 0 {
		return "", errors.New("No Reverse DNS for " + h)
	}

	return addrs[0], nil
}
