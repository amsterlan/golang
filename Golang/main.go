package main

import "github.com/amsterlan/golang/Golang/internal/entity"

func main() {
	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}
}
