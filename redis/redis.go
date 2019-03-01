package main

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

func main() {

	var client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	valstr1, err := client.Get("value1").Result()
	if err == nil {
		var val1 int
		fmt.Println(val1)
		val1, err = strconv.Atoi(valstr1)
		// fmt.Println(val1)
		if err != nil {
			fmt.Println("value1 is not a valid integer")
			return
		}

	} else if err != redis.Nil {
		fmt.Println("redis access error reason:" + err.Error())
		return
	}

}
