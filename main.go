package main

import (
	"fmt"
	"trilateration/trilateration/models"
)

func main() {
	pos1 := models.NewPositionWithHeight(10000, 2000, 100)
	pos2 := models.NewPositionWithHeight(20000, 2300, 100)
	pos3 := models.NewPositionWithHeight(14000, 2100, 150)
	if err := pos1.SetRadius(1000); err != nil {
		panic(err)
	}
	if err := pos2.SetRadius(2000); err != nil {
		panic(err)
	}
	if err := pos3.SetRadius(1500); err != nil {
		panic(err)
	}
	sys := models.NewTrilateration([]*models.Position{pos1, pos2, pos3})
	pos, _ := sys.GetPosition()
	fmt.Println(pos)
}
