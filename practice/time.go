package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now Time is: ", now)

	t := time.Date(2023, time.January, 21, 19, 28, 0, 0, time.UTC)
	fmt.Println("Date is: ", t)
	fmt.Println("Parsed Date is: ", t.Format(time.ANSIC))
}
