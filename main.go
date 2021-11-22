package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ejsuarez793/protobuf-example-go/src/complexpb"
	"github.com/ejsuarez793/protobuf-example-go/src/enumpb"
	"github.com/ejsuarez793/protobuf-example-go/src/simplepb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	// read and write bin file Demo
	sm := doSimple()

	writeToFile("simple.bin", sm)

	sm1 := simplepb.SimpleMessage{}
	readFromFile("simple.bin", &sm1)

	sm1.Name = "readed name!"
	fmt.Println(sm1)

	// JSON Demo
	jsonString := toJSON(sm)

	fmt.Println(jsonString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(jsonString, sm2)

	fmt.Println("Successfully created proto struct: ", sm2)
	fmt.Println(sm2.GetSampleList())

	doEnum()

	doComplex()
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:        42,
		DayOfWeek: enumpb.DayOfWeek_THURSDAY,
	}

	fmt.Println(em)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Dummy",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Dummy",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Dummy",
			},
		},
	}

	fmt.Println(cm)
}
func toJSON(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)

	if err != nil {
		log.Fatalln("Cound't unmarshal de JSON into the pb struct", err)
	}
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
	}

	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
	}

	fmt.Println("Data has been written!")

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
