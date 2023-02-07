package main

import (
	"fmt"
	"github.com/guisfits/fullcycle/applicatiion/route"
)


func main() {
	route := route.Route{
		ID:	"1",
		ClientID: "1",
	}

	route.LoadPositions()
	result, _ := route.ExportJsonPositions()

	fmt.Println(result[1])
}
