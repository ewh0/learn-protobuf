#!/usr/bin/env python3

import argparse
import pathlib

from tutorialpb import addressbook_pb2


def prompt_for_new_person() -> addressbook_pb2.Person | None:
    try:
        person = addressbook_pb2.Person()

        name = input("Name of the person: ").strip()
        id = int(input("Id of the person: ").strip())
        email = input("Email of the person: ").strip()

        person.name = name
        person.id = id
        person.email = email

        while True:
            num = input("Phone number of the person (type empty to exit): ").strip()
            if not num:
                break

            t = addressbook_pb2.PHONE_TYPE_UNSPECIFIED
            match input("Type of this phone: ").strip().lower():
                case "home":
                    t = addressbook_pb2.PHONE_TYPE_HOME
                case "mobile":
                    t = addressbook_pb2.PHONE_TYPE_MOBILE
                case "work":
                    t = addressbook_pb2.PHONE_TYPE_WORK
                case _:
                    print("Unrecognized phone type")

            person.phones.append(addressbook_pb2.Person.PhoneNumber(number=num, type=t))

        print(f"New person added: \n{person}")
        return person
    except EOFError:
        print("\nBye!")
        return None


def main():
    parser = argparse.ArgumentParser(
        prog="add_person", description="Add person to the address book"
    )
    parser.add_argument(
        "file", help="Path to the protobuf message file", type=pathlib.Path
    )
    args = parser.parse_args()

    p: pathlib.Path = args.file
    with p.open("r+b") as f:
        a = addressbook_pb2.AddressBook()
        a.ParseFromString(f.read())
        print(f"Address book:\n{a}")

        if person := prompt_for_new_person():
            a.people.append(person)

        f.seek(0)
        f.write(a.SerializeToString())


if __name__ == "__main__":
    main()
