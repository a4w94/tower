package main

import (
	"fmt"
	te "towerjs/tower/src/tower"
)

func main() {
	var route = "gameset.xlsx"
	//tower.TowerJson(route)

	mosterjson := te.TowerJson(route)

	fmt.Println(string(mosterjson))
}
