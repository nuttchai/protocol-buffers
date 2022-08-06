package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	// way to marshall message to (one line) json
	// out, err := protojson.Marshal(pb)

	// way to marshall message to multi line json
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("Cannot convert to JSON", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {
	// if err := protojson.Unmarshal([]byte(in), pb); err != nil {
	// 	log.Fatalln("Cannot unmarshal from JSON", err)
	// }

	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Cannot unmarshal from JSON", err)
	}
}
