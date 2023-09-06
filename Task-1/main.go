package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	var price, ExpensivePrice int
	for i := 1; i <= n; i++ {
		fmt.Scan(&price)
		if price <= s && price >= ExpensivePrice {
			ExpensivePrice = price
		}
	}
	fmt.Println(ExpensivePrice)
}
