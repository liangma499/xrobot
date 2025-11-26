package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Unix(1722441600, 0).UTC().Format("2006-01")
	fmt.Println(date)
}
