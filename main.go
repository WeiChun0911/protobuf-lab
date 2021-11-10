package main

import (
	"fmt"
	"log"

	"github.com/WeiChun0911/protobuf-lab/pkg/person"
	"google.golang.org/protobuf/proto"
)

func main() {

	personData1 := &person.Person{
		Name: "Willy",
		Id:   1,
	}

	data, err := proto.Marshal(personData1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	personData2 := &person.Person{}
	err = proto.Unmarshal(data, personData2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(personData2.GetId())
	fmt.Println(personData2.GetName())
}
