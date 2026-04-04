package main

import (

	"fmt"
	"strings"
)
	


//  var SalesOrder struct {
// 	OrderID     string
// 	CustomerID  string
// 	OrderAmount float64
// 	OrderDate   string
// 	Status      string
// }

var products = map[string]float64 {
	"ProductA": 10.99,
	"ProductB": 20.49,
	"ProductC": 15.75,
};

func calculateTotal(order map[string]int) float64 {
	total := 0.0
	for product, quantity := range order {
		if price, exists := products[product]; exists {
			total += price * float64(quantity)
		}
	}
	return total
}



func processOrder(orderID string, customerID string, order map[string]int) {
	fmt.Printf("Processing order %s for customer %s\n", orderID, customerID)
	totalAmount := calculateTotal(order)
	fmt.Printf("Total amount for order %s: $%.2f\n", orderID, totalAmount)
}


func calculateOrderTotal(itemCode string) (float64, bool) {

	basePrice, exists := products[itemCode]
	if exists {
		return basePrice, true
	} else {
		if strings.HasSuffix(itemCode, "_DISCOUNT") {

			itemCodeWithoutSuffix := strings.TrimSuffix(itemCode, "_DISCOUNT")
			basePrice, exists := products[itemCodeWithoutSuffix]
			if exists {
				basePrice= basePrice * 0.9 // Apply a 10% discount

				fmt.Printf("Applying discount for %s. Original price: $%.2f, Discounted price: $%.2f\n", 
				itemCodeWithoutSuffix, basePrice/0.9, basePrice)

				return basePrice, true
			}


			fmt.Println("Item code with _DISCOUNT suffix does not exist:", itemCode)
			return 0.0, true
		}
	}

	return 0.0, false
}

func main() {
    fmt.Println("----------------Sales order processing system--------------------------")

orderItems:=[]string{"ProductA", "ProductB_DISCOUNT", "ProductC", "ProductC_DISCOUNT"}

var subtotal float64

for _, itemCode := range orderItems {
	price, exists := calculateOrderTotal(itemCode)
	if exists {
		subtotal += price
	} else {
		fmt.Printf("Item code does not exist: %s\n", itemCode)
	}

}
fmt.Printf("Subtotal: $%.2f\n", subtotal)


}