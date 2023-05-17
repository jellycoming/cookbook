package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func toJsonString(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("toJsonString(): %v, %s", v, err.Error())
	}
	return string(data)
}

func main() {
	s1 := []uint64{1, 2, 3, 4}
	fmt.Println(toJsonString(s1))
	s2 := map[string]string{"hello": "world", "foo": "bar"}
	fmt.Println(toJsonString(s2))
}
