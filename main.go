package main

import (
	"fmt"
	"io/ioutil"
	"log"
	serializer "protobuf-example-go/src/serializer"
	simplepb "protobuf-example-go/src/simple"

	"google.golang.org/protobuf/proto"
)

func main() {
	message := doSimple()
	// readAndWriteDemo(message)
	jsonDemo(message)

}

func jsonDemo(message proto.Message) {
	messageAsString := toJSON(message)
	fmt.Println(messageAsString)

	message2 := &simplepb.SimpleMessage{}
	fromJSON(messageAsString, message2)
	fmt.Println("Successfully created proto struct", message2)
}

func toJSON(pb proto.Message) string {
	out, err := serializer.ProtobufToJSON(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	inBytes := []byte(in)
	err := serializer.JSONToProtobuf(inBytes, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(message proto.Message) {
	writeToFile("simple.bin", message)
	message2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", message2)
	fmt.Println(message2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffer's struct", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	message := simplepb.SimpleMessage{
		Id:         1,
		IsSimple:   true,
		Name:       "Simple Message",
		SampleList: []int32{1, 2, 4},
	}

	fmt.Println(&message)
	return &message
}
