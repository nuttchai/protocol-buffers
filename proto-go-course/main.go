package main

import (
	"fmt"

	pb "github.com/nuttchai/proto-go-course/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          1,
		IsSimple:    true,
		Name:        "Hello",
		SimpleLists: []int32{1, 2, 3, 4, 5},
	}
}

func main() {
	fmt.Println(doSimple())
}
