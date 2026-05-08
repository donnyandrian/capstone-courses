// 214
package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("resources/customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
	}

	encoder.Encode(customer)
}
