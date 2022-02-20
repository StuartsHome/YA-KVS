package main

import (
	"fmt"

	"github.com/StuartsHome/YA-KVS/store"
)

func main() {

	// New store.
	store := store.NewStore()
	fmt.Println(store)

}
