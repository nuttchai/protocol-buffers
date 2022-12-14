package main

import (
	"fmt"
	"reflect"

	pb "github.com/nuttchai/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          1,
		IsSimple:    true,
		Name:        "Hello",
		SimpleLists: []int32{1, 2, 3, 4, 5},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 1, Name: "Hello1"},
		MultipleDummies: []*pb.Dummy{
			{Id: 2, Name: "Hello 2"},
			{Id: 3, Name: "Hello 3"},
		},
	}
}

func doEnum() *pb.Enumerations {
	return &pb.Enumerations{
		EyeColor: 1, // or we can write => pb.EyeColor_EYE_COLOR_GREEN instead of 1,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId1": {Id: 41},
			"myId2": {Id: 42},
			"myId3": {Id: 43},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	return jsonString
}

// reflect.Type requires given variable "type"
func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())

	// fmt.Println("This should be Id:")
	// doOneOf(&pb.Result_Id{Id: 42})
	// fmt.Println("This should be Message:")
	// doOneOf(&pb.Result_Message{Message: "a message"})

	// fmt.Println(doMap())

	// doFile(doSimple())

	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	jsonString = doToJSON(doComplex())
	message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	// it will discard the "unknown" field because in message Simple, it doesn't have "unknown" field
	fmt.Println(doFromJSON(`{"id": 32, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))
}
