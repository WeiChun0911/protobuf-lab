package main

import (
	"fmt"
	"log"

	"github.com/WeiChun0911/protobuf-lab/pkg/crazyApiService/apiResponse"
	"google.golang.org/protobuf/proto"
)

func main() {

	// personData1 := &others.Person{
	// 	Name: "Willy",
	// 	Id:   1,
	// }

	// data, err := proto.Marshal(personData1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(data)

	// personData2 := &others.Person{}
	// err = proto.Unmarshal(data, personData2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(personData2.GetId())
	// fmt.Println(personData2.GetName())

	personData3 := &apiResponse.PlatformMonitor{}

	data, err := proto.Marshal(personData3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
