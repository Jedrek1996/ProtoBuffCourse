package main

import (
	pb "ProtoCourse"
	"fmt"
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
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("Not string or int32", x)
	}
}

func main() {
	fmt.Println("Do Simple ----")
	fmt.Printf("%v\n", doSimple())

	fmt.Println("Do Complex ----")
	fmt.Printf("%v\n", doComplex())

	fmt.Println("Do Enum ----")
	fmt.Printf("%v\n", doEnum())

	fmt.Println("Do Oneof ----")
	doOneOf(&pb.Result_Id{Id: 32})
	doOneOf(&pb.Result_Message{Message: "Hello!"})

}
