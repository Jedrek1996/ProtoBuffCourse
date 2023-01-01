package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

// Pb is like interface
func WriteToFile(fname string, pb proto.Message) {
	file, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Cannot serialize to byte", err)
		return
	}

	if err = ioutil.WriteFile(fname, file, 0644); err != nil {
		log.Fatalln("Can't write error to file", err)
	}

	fmt.Println("Data written")
}

func ReadFromFile(fname string, pb proto.Message) {
	file, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't READ file", err)
	}

	if err = proto.Unmarshal(file, pb); err != nil {
		log.Fatalln("Cannot ummarshal", err)
		return
	}
}
