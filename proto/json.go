package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Pb to JSON
func ToJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}

	file, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return string(file)
}

func fromJSON(in string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	//Unmarshall into byte from PB
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Could not unmarshal from JSON")
	}

}

// func fromJSON2(in string, pb proto.Message) {
// 	req := proto.pb{}
// 	err := protojson.Unmarshal([]byte(in), req)
// }
