package main

import (
	json2 "github.com/angao/practice/20180318/json"
)

type Person struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Birthday json2.Time `json:"birthday"`
}

func main()  {
	//src := `{"id":1,"name":"alen","birthday":"2018-03-18 16:05:51"}`
	//
	//p := new(Person)
	//err := json.Unmarshal([]byte(src), p)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(p)
	//
	//ps , _ := json.Marshal(p)
	//log.Println(string(ps))
}
