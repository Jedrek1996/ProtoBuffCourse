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

func main() {
	fmt.Println("Do Simple ----")
	fmt.Println(doSimple())

	fmt.Println("Do Complex ----")
	fmt.Println(doComplex())
}
