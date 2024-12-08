package main;

import (
    "fmt";
	"A3/dequeue";
)


func main() {
	// Create sample monsters
	m1 := dequeue.Monster{Name: "Godzilla", InitiativeModifier: 0, ChallengeRating: 1, ArmourClass: 13, AttackRating: -1};
	m2 := dequeue.Monster{Name: "Bigfoot", InitiativeModifier: 2, ChallengeRating: 2, ArmourClass: 15, AttackRating: 0};
	m3 := dequeue.Monster{Name: "Dracula", InitiativeModifier: 5, ChallengeRating: 0, ArmourClass: 18, AttackRating: 5};

	dq := dequeue.MakeDequeue();

	// TEST Append(), Prepend()
	dq = dq.Append(m1);
	dq = dq.Append(m2);
	dq = dq.Prepend(m3);

	// TEST Len()
	fmt.Println("Len:", dq.Len());


}