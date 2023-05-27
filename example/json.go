package main

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func toJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("Cannot convert to json", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {
	if err := protojson.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Could not unmarshal from json", err)
	}
}