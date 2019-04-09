package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	tf := t.Format("Mon, 02 Jan 2006 15:04:05 -0700 (MST)")
	fmt.Println(tf)
}
