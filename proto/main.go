package main

import (
	pb "ProtoCourse"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "Jed",
		SampleLists: []int32{1, 2, 3, 4, 5},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy1{
			Id:   42,
			Name: "Jedrek",
		},
		MultiDummies: []*pb.Dummy1{{
			Id:   11,
			Name: "Jk",
		},
			{
				Id:   22,
				Name: "edre",
			},
		},
	}
}

func doEnum() *pb.EnumrateEye {
	return &pb.EnumrateEye{
		EyeColor: 2, //pb.Eyecolor_Eye_Color_Blue (Can use enum key)
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id1:
		fmt.Println(message)
		fmt.Println(message.(*pb.Result_Id1).Id1)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("Not string or int32", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId":  {Id: 1},
			"myId2": {Id: 2},
			"myId3": {Id: 3},
		},
	}
}

// go run -v ./proto
func doFile(p proto.Message) {
	path := "proto/simple.bin"

	WriteToFile(path, p)
	message := &pb.Simple{}
	ReadFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	// 	import "google.golang.org/protobuf/encoding/protojson"
	// jsonString2 := protojson.Format(protoMessage)

	jsonString := ToJson(p)
	fmt.Println(jsonString)
	return jsonString
}

// When receieving the json string not sure which type to deserialize in
// If passing the type explicitly
func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)

	return message

}

func main() {
	fmt.Println("Do Simple ----")
	fmt.Printf("%v\n", doSimple())

	fmt.Println("Do Complex ----")
	fmt.Printf("%v\n", doComplex())

	fmt.Println("Do Enum ----")
	fmt.Printf("%v\n", doEnum())

	fmt.Println("Do Oneof ----")
	doOneOf(&pb.Result_Id1{Id1: 32})
	doOneOf(&pb.Result_Message{Message: "Hello!"})

	fmt.Println("Do maps ----")
	fmt.Printf("%v\n", doMap())

	fmt.Println("Do file ----")
	doFile(doSimple())

	fmt.Println("Do file ----")
	doFile(doSimple())

	fmt.Println("Do JSON ----")
	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString, message)

	fmt.Println(doFromJSON(`{"id":42, "unknown":"test"`, reflect.TypeOf(pb.Simple{})))
}
