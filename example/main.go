package main

import (
	pb "example.com/m/proto/generated"
	"fmt"
	"google.golang.org/protobuf/proto"
	"reflect"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "A name",
		SampleList: []int32{1, 2, 3, 4, 5},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   42,
			Name: "Uzair Tariq",
		},
		MultipleDummies: []*pb.Dummy{
			{Id: 1, Name: "One"},
			{Id: 2, Name: "Two"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		_ = fmt.Errorf("message has unexpected type: %v", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId": {
				Id: 52,
			},
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
	jsonString := toJson(p)
	fmt.Println(jsonString)
	return jsonString
}

func doFromJson(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnum())
	/*doOneOf(&pb.Result_Id{
		Id: 1,
	})
	doOneOf(&pb.Result_Message{
		Message: "This is a testing message",
	})*/
	//fmt.Println(doMap())
	//doFile(doSimple())

	jsonString := doToJSON(doSimple())
	message := doFromJson(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)
}