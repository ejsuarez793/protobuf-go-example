package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/ejsuarez793/protobuf-example-go/src/addressbookpb"
	"github.com/ejsuarez793/protobuf-example-go/src/complexpb"
	"github.com/ejsuarez793/protobuf-example-go/src/enumpb"
	"github.com/ejsuarez793/protobuf-example-go/src/simplepb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	// read and write bin file Demo
	sm := doSimple()

	readAndWriteDemo(sm)

	JsonDemo(sm)

	doEnum()

	doComplex()

	addressBookDemo()
}

func printMainMenu() {
	fmt.Println(
		"1 - See Address Book\n " +
			"2 - Add Person\n " +
			"3 - Remove Person\n " +
			"4 - Save and Quit\n " +
			"5 - Quit Discarding Changes\n ")
}

func getInputFromPrompt() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("Error reading prompt input: ", err)
	}

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func clearPrompt() {
	cmdName := "clear"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getPersonFromPrompt() *addressbookpb.Person {
	phones := make([]*addressbookpb.Person_PhoneNumber, 1)
	phone := addressbookpb.Person_PhoneNumber{
		Number: "123",
		Type:   addressbookpb.Person_MOBILE,
	}

	phones = append(phones, &phone)

	fmt.Println("Adding new person info...")
	fmt.Print("Person's Id: -> ")
	person_id_str := getInputFromPrompt()

	person_id, _ := strconv.Atoi(person_id_str)

	fmt.Print("Person's Name: -> ")
	person_name := getInputFromPrompt()

	fmt.Print("Person's Email: -> ")
	person_email := getInputFromPrompt()

	person := addressbookpb.Person{
		Id:     int32(person_id),
		Name:   person_name,
		Email:  person_email,
		Phones: phones,
	}

	return &person
}

func RemoveIndex(s []*addressbookpb.Person, index int) []*addressbookpb.Person {
	return append(s[:index], s[index+1:]...)
}

func DeletePerson(ab *addressbookpb.AddressBook) {
	id_str := getInputFromPrompt()
	id_int, _ := strconv.Atoi(id_str)
	id := int32(id_int)
	for i := 0; i < len(ab.People); i++ {
		if ab.People[i].Id == id {
			ab.People = RemoveIndex(ab.People, i)
			break
		}
	}
}

func addressBookDemo() {

	printMainMenu()
	ab := addressbookpb.AddressBook{}
	err := getAddressBook(&ab)
	if err != nil {
		log.Fatalln("Error exiting")
		return
	}
	continue_run := true
	for continue_run {
		text := getInputFromPrompt()

		clearPrompt()
		fmt.Println(text)
		switch text {
		case "1": // See Address Book
			fmt.Println("See Address Book selected...")
			fmt.Println(ab)
		case "2":
			fmt.Println("Add person selected...")
			person := getPersonFromPrompt()
			ab.People = append(ab.People, person)
		case "3":
			fmt.Println("Delete person selected...")
			DeletePerson(&ab)
		case "4": // Save and Quit
			fmt.Println("Saving...")
			writeToFile("addressbook.bin", &ab)
			continue_run = false
			break
		case "5": // Quit discarding changes
			continue_run = false
			break
		default:
			fmt.Printf("invalid option, please press an option from 1 - 5\n")
		}

		printMainMenu()

	}

	fmt.Println("Exiting execution...")

}

func getAddressBook(pb proto.Message) error {
	// read if file exists otherwise create empty address book
	err := readFromFile("addressbook.bin", pb)

	if !os.IsNotExist(err) { // if error is not of file does not exists
		return err
	}

	// write empty addressbook.bin
	log.Println("address book file does not exists, creating empty addressbook", err)
	err2 := writeToFile("addressbook.bin", pb)
	if err2 != nil {
		log.Fatalln("There was an error getting address book", err2)
	}

	return nil
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

func readAndWriteDemo(pb proto.Message) {
	writeToFile("simple.bin", pb)

	sm1 := simplepb.SimpleMessage{}
	readFromFile("simple.bin", &sm1)

	sm1.Name = "readed name!"
	fmt.Println(sm1)
}

func JsonDemo(pb proto.Message) {
	// JSON Demo
	jsonString := toJSON(pb)

	fmt.Println(jsonString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(jsonString, sm2)

	fmt.Println("Successfully created proto struct: ", sm2)
	fmt.Println(sm2.GetSampleList())
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
		log.Println("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
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
		return err
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
