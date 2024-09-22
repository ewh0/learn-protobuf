package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/ewh0/learn-protobuf-grpc/official-tutorial/go/tutorialpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func promptForAddress(r io.Reader) (*pb.Person, error) {
	rd := bufio.NewReader(r)

	p := &pb.Person{}

	fmt.Print("Input name of the person: ")
	s, err := rd.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	p.Name = strings.TrimSpace(s)

	fmt.Print("Input id number for this person: ")
	s, err = rd.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	id, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatalln(err)
	}
	p.Id = int32(id)

	fmt.Print("Input email of the person: ")
	s, err = rd.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	p.Email = strings.TrimSpace(s)

	for {
		fmt.Print("Input phone number (press enter to exit): ")
		s, err := rd.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		num := strings.TrimSpace(s)
		if len(num) == 0 {
			break
		}

		fmt.Print("Enter phone type: ")
		s, err = rd.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		s = strings.ToLower(strings.TrimSpace(s))
		t := pb.PhoneType_PHONE_TYPE_UNSPECIFIED
		switch s {
		case "mobile":
			t = pb.PhoneType_PHONE_TYPE_MOBILE
		case "home":
			t = pb.PhoneType_PHONE_TYPE_HOME
		case "work":
			t = pb.PhoneType_PHONE_TYPE_WORK
			pb.PhoneType_PHONE_TYPE_HOME.Enum()
		default:
			fmt.Println("unrecognized phone type")
		}

		p.Phones = append(p.Phones, &pb.Person_PhoneNumber{Number: num, Type: t})
	}

	p.LastUpdated = timestamppb.Now()

	return p, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("incorrect number of arguments")
	}

	fname := os.Args[1]
	b, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("failed to read address book from file %s: %s", fname, err.Error())
	}

	book := &pb.AddressBook{}
	if err := proto.Unmarshal(b, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	p, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln(err.Error())
	}

	book.People = append(book.People, p)

	b, err = proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to marshall address book:", err)
	}

	if err := os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalf("failed to read content from file %s: %s", fname, err.Error())
	}
}
