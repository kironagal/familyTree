package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// defining a truct to store the person name and the relationship
type Person struct {
	Name         string
	Relationship string
}

var familyTree = make(map[string]Person)

// to add the person and also to check if the person exists
func addPerson(name string) {
	if _, exists := familyTree[name]; exists {
		fmt.Printf("%s already existsin the family tree\n", name)
		return
	}
	familyTree[name] = Person{Name: name}
	fmt.Printf("%s added to the family tree\n", name)
}

// to add relation of the person and alos to check the relation exists
func addRelationship(name, relationship string) {
	if person, exists := familyTree[name]; exists {
		person.Relationship = relationship
		familyTree[name] = person
		fmt.Printf("%s is now markedas %s.\n", name, relationship)
	} else {
		fmt.Printf("%s does not exists in the family tree.Please add the person", name)
	}
}

// connecting two people with the relation between them
func connectRelationship(name1, relationship, name2 string) {
	if person1, exists1 := familyTree[name1]; exists1 {
		if _, exists2 := familyTree[name2]; exists2 {
			if person1.Relationship == "" {
				fmt.Printf("%s does not havea specified relationship.\n", name1)
				return
			}
			if person1.Relationship == relationship {
				fmt.Printf("%s is now the %s of %s.\n", name1, relationship, name2)
			} else {
				fmt.Printf("%s cannot be the %s of %s.\n", name1, relationship, name2)
			}
		} else {
			fmt.Printf("%s does not existsin the family tree. Add the person.\n", name2)
		}
	} else {
		fmt.Printf("%s does not exists in the family tree. Add the person first.\n", name1)
	}
}

// this is to find or get the count of particular relationship
func countRelationship(name, relationship string) int {
	count := 0
	for _, person := range familyTree {
		if person.Relationship == relationship {
			count++
		}
	}
	return count
}

// to check the is the given person a father
func fatherOf(name string) string {
	for _, person := range familyTree {
		if person.Relationship == "father" && person.Name != name {
			return person.Name
		}
	}
	return "unknown"
}

func main() {

	// this is from go package https://pkg.go.dev/github.com/urfave/cli/v2@v2.25.7 for Command Line apps in Go
	app := &cli.App{
		Name:  "Hello",
		Usage: "Welcoming the user",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello")
			return nil
		},
	}
	app.Run(os.Args)

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: family-tree <command> [options]")
		return
	}

	command := args[0]
	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: family-tree add <name> <relationship>")
			return
		}
		addCommand := args[1]
		switch addCommand {
		case "person":
			addPerson(args[2])
		case "relationship":
			addRelationship(args[2], args[3])
		default:
			fmt.Printf("Invalid 'add' command. Use 'person' or 'realtionship'")
		}
	case "connect":
		if len(args) < 7 || args[3] != "as" || args[5] != "of" {
			fmt.Println("Usage: family-tree connect <name1> as <relationship> of <name2>")
			return
		}
		connectRelationship(args[1], args[4], args[6])
	case "count":
		if len(args) < 4 {
			fmt.Println("Usage; family-tree count <relationship> of <name>")
			return
		}
		countCommand := args[1]
		switch countCommand {
		case "sons", "daughters", "wives":
			name := args[3]
			count := countRelationship(name, countCommand)
			fmt.Printf("Number of %s of %s: %d\n", countCommand, name, count)

		case "father":
			name := args[3]
			father := fatherOf(name)
			fmt.Printf("Father of %s: %s\n", name, father)
		default:
			fmt.Println("Invalid 'count' command. Use 'sons','daughters', 'wives', or 'father'.")
		}
	default:
		fmt.Println("Invalid command. Use 'add', 'connect', or 'count'.")
	}
	fmt.Println(args)
}
