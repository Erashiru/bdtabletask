package main

import (
	"fmt"
	"log"
	"strings"
)

func (app *application) Manager(page []string) error {
	err := app.storage.Insert()
	if err != nil {
		return err
	}
	fmt.Printf("=+=+=+=\nСтраница сборки заказа %s\n\n", page)

	for _, order := range page {
		shelves, err := app.storage.Get(order)
		if err != nil {
			log.Fatal(err)
		}

		for _, shelf := range shelves {
			fmt.Printf("===Стеллаж %s\n", shelf.Name)

			fmt.Printf("%s (id=%d)\nзаказ %d, %d шт\n", shelf.ProductName, shelf.ProductID, shelf.OrderNumber, shelf.Quantity)

			if len(shelf.Additional) > 0 {
				fmt.Printf("доп стеллаж: %s\n\n", strings.Join(shelf.Additional, ","))
			} else {
				fmt.Println()
			}
		}
	}

	return nil
}
