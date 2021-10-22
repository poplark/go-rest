package main

import (
	"fmt"
	"poplark/rest-blog/dbs"
	"poplark/rest-blog/router"
)

func main() {
	fmt.Println("Hello, World!")
	dbs.Count(false)
	dbs.CreateUser("pop", "pop@pop.com", "pop")
	dbs.Find(5, 5, false)
	dbs.FindOneById(2)
	router.Handler()
}
