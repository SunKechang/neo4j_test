package main

import (
	"fmt"
	"neo4j_1/database/neo4j"
)

func main() {
	world, err := neo4j.HelloWorld("bolt://39.107.25.37:27687", "neo4j", "password")
	fmt.Println(world, err)
}
