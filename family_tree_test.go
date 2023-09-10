package main

import (
	"testing"
)

func TestAddPerson(t *testing.T) {
	familyTree = make(map[string]Person)
	addPerson("Amit Dhakad")
	if len(familyTree) != 1 {
		t.Errorf("Expected 1 person in the family tree, but got %d", len(familyTree))
	}

	// Adding the same person again should not increase the count.
	addPerson("Amit Dhakad")
	if len(familyTree) != 1 {
		t.Errorf("Expected 1 person in the family tree after duplicate addition, but got %d", len(familyTree))
	}
}

func TestAddRelationship(t *testing.T) {
	familyTree = make(map[string]Person)
	addPerson("Amit Dhakad")
	addRelationship("Amit Dhakad", "son")
	if person, exists := familyTree["Amit Dhakad"]; exists {
		if person.Relationship != "son" {
			t.Errorf("Expected relationship 'son' for Amit Dhakad, but got '%s'", person.Relationship)
		}
	} else {
		t.Errorf("Amit Dhakad should exist in the family tree after adding a relationship")
	}

	// Adding a relationship to a non-existent person should print an error message.
	addRelationship("KK Dhakad", "father")
}

func TestConnectRelationship(t *testing.T) {
	familyTree = make(map[string]Person)
	addPerson("Amit Dhakad")
	addPerson("KK Dhakad")
	addRelationship("Amit Dhakad", "son")

	// Valid connection.
	connectRelationship("Amit Dhakad", "child", "KK Dhakad")

	// Invalid connection (mismatched relationship).
	connectRelationship("Amit Dhakad", "father", "KK Dhakad")

	// Non-existent person in the tree.
	connectRelationship("Eve", "daughter", "Amit Dhakad")

	// Person 1 is not in the tree.
	connectRelationship("Aadi", "son", "Amit Dhakad")

	// Person 2 is not in the tree.
	connectRelationship("Amit Dhakad", "sister", "Arun")
}

func TestCountRelationship(t *testing.T) {
	familyTree = make(map[string]Person)
	addPerson("Amit Dhakad")
	addPerson("KK Dhakad")
	addPerson("Aadi")
	addPerson("Arun")

	addRelationship("Amit Dhakad", "son")
	addRelationship("KK Dhakad", "father")
	addRelationship("Aadi", "son")
	addRelationship("Arun", "son")

	// Count sons.
	count := countRelationship("Amit Dhakad", "son")
	if count != 2 {
		t.Errorf("Expected 2 sons for Amit Dhakad, but got %d", count)
	}

	// Count daughters (no daughters in the tree).
	count = countRelationship("Amit Dhakad", "daughter")
	if count != 0 {
		t.Errorf("Expected 0 daughters for Amit Dhakad, but got %d", count)
	}

	// Count father (single father).
	count = countRelationship("Amit Dhakad", "father")
	if count != 1 {
		t.Errorf("Expected 1 father for Amit Dhakad, but got %d", count)
	}

	// Count wives (no wives in the tree).
	count = countRelationship("Amit Dhakad", "wife")
	if count != 0 {
		t.Errorf("Expected 0 wives for Amit Dhakad, but got %d", count)
	}
}

func TestFatherOf(t *testing.T) {
	familyTree = make(map[string]Person)
	addPerson("Amit Dhakad")
	addPerson("KK Dhakad")
	addPerson("Aadi")
	addPerson("Arun")

	addRelationship("Amit Dhakad", "son")
	addRelationship("KK Dhakad", "father")
	addRelationship("Aadi", "son")
	addRelationship("Arun", "son")

	// Father of KK Dhakad.
	father := fatherOf("KK Dhakad")
	if father != "Amit Dhakad" {
		t.Errorf("Expected 'Amit Dhakad' as the father of KK Dhakad, but got '%s'", father)
	}

	// Father of Aadi (no father specified).
	father = fatherOf("Aadi")
	if father != "unknown" {
		t.Errorf("Expected 'unknown' as the father of Aadi, but got '%s'", father)
	}

	// Father of Arun (no father specified).
	father = fatherOf("Arun")
	if father != "unknown" {
		t.Errorf("Expected 'unknown' as the father of Arun, but got '%s'", father)
	}

	// Father of non-existent person.
	father = fatherOf("Suman")
	if father != "unknown" {
		t.Errorf("Expected 'unknown' as the father of Eve, but got '%s'", father)
	}
}
