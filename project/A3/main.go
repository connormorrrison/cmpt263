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

	// Test empty dequeue
    fmt.Println("Initial Len:", dq.Len());
    if found, _ := dq.Get(1); found {
        fmt.Println("Error: Found monster in empty dequeue");
    }

	// TEST Append(), Prepend()
    fmt.Println("\nTesting Append and Prepend:");
    dq = dq.Append(m1);
    dq = dq.Append(m2);
    dq = dq.Prepend(m3);
    dq = dq.Append(m4);
    dq = dq.Prepend(m5);

	// TEST Len()
    fmt.Println("Len:", dq.Len());

    // TEST Get() with positive indices
    fmt.Println("\nTesting forward traversal:");
    for i := 1; i <= dq.Len(); i++ {
        if found, monster := dq.Get(i); found {
            fmt.Printf("Monster at position %d: %s\n", i, monster.Name);
        }
    }

    // TEST Get() with negative indices
    fmt.Println("\nTesting backward traversal:");
    for i := -1; i >= -dq.Len(); i-- {
        if found, monster := dq.Get(i); found {
            fmt.Printf("Monster at position %d: %s\n", i, monster.Name);
        }
    }

    // TEST Shift()
    fmt.Println("\nTesting Shift (remove at the front):");
    dq = dq.Shift(); // Remove first element
    fmt.Println("Len after Shift:", dq.Len());
    if found, monster := dq.Get(1); found {
        fmt.Println("New first monster:", monster.Name);
    }

	// TEST Shift()
    fmt.Println("\nTesting Shift (remove at the front):");
    dq = dq.Shift(); // Remove first element
    fmt.Println("Len after Shift:", dq.Len());
    if found, monster := dq.Get(1); found {
        fmt.Println("New first monster:", monster.Name);
    }

    // TEST Drop()
    fmt.Println("\nTesting Drop (remove at the back):");
    dq = dq.Drop(); // Remove last element
    fmt.Println("Len after Drop:", dq.Len());
    if found, monster := dq.Get(-1); found {
        fmt.Println("New last monster:", monster.Name);
    }

	// TEST Get() with positive indices
    fmt.Println("\nTesting forward traversal:");
    for i := 1; i <= dq.Len(); i++ {
        if found, monster := dq.Get(i); found {
            fmt.Printf("Monster at position %d: %s\n", i, monster.Name);
        }
    }
}