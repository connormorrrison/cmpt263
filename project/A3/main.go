package main;

import (
    "fmt";
	"A3/dequeue";
)


func main() {
	// Create sample monsters
	m1 := dequeue.Monster{Name: "Godzilla", InitiativeModifier: 0, ChallengeRating: 1, ArmourClass: 13, AttackRating: -1};
	m2 := dequeue.Monster{Name: "Bigfoot", InitiativeModifier: 2, ChallengeRating: 2, ArmourClass: 15, AttackRating: 0};
	m3 := dequeue.Monster{Name: "Dracula", InitiativeModifier: 5, ChallengeRating: 3, ArmourClass: 18, AttackRating: 5};

	dq := dequeue.MakeDequeue();

	// TEST Append(), Prepend()
	dq = dq.Append(m1);
	dq = dq.Append(m2);
	dq = dq.Prepend(m3);

	// TEST Len()
	fmt.Println("Len:", dq.Len());

	// // TEST Get()
	// if found, monster := dq.Get(1); found {
	// 	fmt.Println("First Monster:", monster)
	// }

	// if found, monster := dq.Get(-1); found {
	// 	fmt.Println("Last Monster:", monster)
	// }

	// // TEST Shift(), Drop()
	// dq = dq.Shift() // Remove first element
	// fmt.Println("Len after Shift:", dq.Len())

	// dq = dq.Drop()
	// fmt.Println("Len after Drop:", dq.Len())


}