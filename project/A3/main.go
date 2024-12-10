// Note: For this assignment I am choosing this option: complete 
// Assignment 3 but not Assignment Four, and have the assignment 
// weight transferred across the first three assignments (each wil
// become 6.67% of the overall course grade)

package main;

import (
    "fmt";
	"A3/dequeue";
)


func main() {
	// Create sample monsters
	m1 := dequeue.Monster{Name: "Godzilla", InitiativeModifier: 10, ChallengeRating: 2, ArmourClass: 15, AttackRating: -10};
    m2 := dequeue.Monster{Name: "Bigfoot", InitiativeModifier: 10, ChallengeRating: 2, ArmourClass: 15, AttackRating: -10};
    m3 := dequeue.Monster{Name: "Dracula", InitiativeModifier: 10, ChallengeRating: 2, ArmourClass: 15, AttackRating: -10};
    m4 := dequeue.Monster{Name: "Werewolf", InitiativeModifier: 10, ChallengeRating: 2, ArmourClass: 15, AttackRating: -10};
    m5 := dequeue.Monster{Name: "Dragon", InitiativeModifier: 10, ChallengeRating: 2, ArmourClass: 15, AttackRating: -10};

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