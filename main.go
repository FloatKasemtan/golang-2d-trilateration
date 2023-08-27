package main

import (
	"fmt"
	"trilateration/trilateration/models"
)

func main() {
	pos1 := models.NewPosition(10000, 2000)
	pos2 := models.NewPosition(20000, 2300)
	pos3 := models.NewPosition(14000, 2100)
	pos1.SetRadius(1000)
	pos2.SetRadius(2000)
	pos3.SetRadius(1500)
	sys := models.NewTrilateration([]*models.Position{pos1, pos2, pos3})
	pos, _ := sys.GetPosition()
	fmt.Println(pos)
}
