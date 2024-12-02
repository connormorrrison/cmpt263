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

// Define NodeType for const expression
type NodeType int;

const (
	HEAD NodeType = iota
	NODE
	TAIL
)

