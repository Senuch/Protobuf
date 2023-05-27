package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func writeToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Cannot serialize to bytes ", err)
		return
	}

	if err = os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cannot write to file ", err)
		return
	}

	fmt.Println("Data has been written")
}

func readFromFile(fname string, pb proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalln("Cannot read file ", err)
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Could not unmarshal", err)
		return
	}
}