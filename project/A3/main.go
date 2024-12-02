package main;

import "fmt";

// Each node contains a MONSTER data type
type Monster struct {
	Name string;
	InitiativeModifier int;
	ChallengeRating int;
	ArmourClass int;
	AttackRating int;
}

// Node variants
type Head struct {
	next *Node;
}

type Node struct {
	data Monster;
	next *Node;
	prev *Node;
}

type Tail struct {
	prev *Node;
}

type NodeOperations interface {
	Prepend(m Monster);  // Adds the data element to the beginning of the dequeue
	Append(m Monster);  // Adds the data element to the end of the dequeue
	Len() int;  // Returns an int telling us how many elements are in the dequeue
	Get(i int) (bool, Monster);  // Returns the Monster at position i
	Shift();  // To remove at the front
	Drop();  // To remove at the back
}