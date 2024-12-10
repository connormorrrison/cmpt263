package dequeue;

// Monster data type
type Monster struct {
    Name string;
    InitiativeModifier int;
    ChallengeRating int;
    ArmourClass int;
    AttackRating int;
}

// Dequeue interface defines the operations on the structure
type Dequeue interface {
	Prepend(m Monster) Dequeue;
	Append(m Monster) Dequeue;
	Len() int;
	Get(i int) (bool, Monster);  // bool if Monster was found
	Shift() Dequeue;
	Drop() Dequeue;
}


// HEAD variant: start of dequeue
type HEAD struct {
	Next Dequeue;
}

// NODE variant: contains Monster and links prev and next nodes
type NODE struct {
	Data Monster;
	Next Dequeue;
	Prev Dequeue;
}

// TAIL variant: end of dequeue
type TAIL struct {
	Prev Dequeue;
}

// MakeDequeue creates a new empty dequeue data structure; returns a dequeue
// with HEAD and TAIL nodes connected
func MakeDequeue() Dequeue {
	var head HEAD;
	var tail TAIL;
	head.next = &tail;
	tail.prev = &head;
	return &head;
}


// // HELPER METHODS

// // Helper method to find the TAIL from HEAD INCOMPLETE
// func (h *HEAD) findTail() *TAIL {
//     current := h.Next;
//     for {
//         switch node := current.(type) {
//         case *NODE:
//             current = node.Next;
//         case *TAIL:
//             return node;
//         default:
//             return nil;
//         }
//     }
// }

// // Helper method to traverse forward for Get INCOMPLETE
// func (h *HEAD) getFromHead(i int) (bool, Monster) {
//     current := h.Next;
//     for current != nil {
//         switch node := current.(type) {
//         case *NODE:
//             if i == 1 {
//                 return true, node.Data;
//             }
//             i--;
//             current = node.Next;
//         case *TAIL:
//             // Reached end without finding the index
//             return false, Monster{};
//         }
//     }
//     return false, Monster{};
// }

// // Helper method to traverse backward for Get INCOMPLETE
// func (h *HEAD) getFromTail(i int) (bool, Monster) {
//     tail := h.findTail();
//     if tail == nil {
//         return false, Monster{};
//     }
//     current := tail.Prev;
//     count := i;
//     for current != nil {
//         switch node := current.(type) {
//         case *NODE:
//             if count == 1 {
//                 return true, node.Data;
//             }
//             count--;
//             current = node.Prev;
//         case *HEAD:
//             // Reached front without finding the index
//             return false, Monster{};
//         }
//     }
//     return false, Monster{};
// }

// HEAD METHODS

// DONE
// Prepend adds a new Monster to the front of the dequeue
//@* requires h is a valid HEAD node, m is a valid Monster
func (h *HEAD) Prepend(m Monster) Dequeue {
	//@ ensures m is the new first Monster in the dequeue
	n := &NODE{Data: m, Prev: h, Next: h.Next};
    switch next := h.Next.(type) {
    case *NODE:
        next.Prev = n;
    case *TAIL:
        next.Prev = n;
    }
    h.Next = n;
    return h;
}
//@* ensures the Monster m is added to the front of the dequeue, Len increases by 1

// DONE
// Appends a new Monster to the end of the dequeue
func (h *HEAD) Append(m Monster) Dequeue {
    switch next := h.Next.(type) {
    case *TAIL:
        return h.Prepend(m);
    case *NODE:
        return next.Append(m);
    }
    return h;
}

// DONE
// Len returns the number of Monsters in the dequeue
func (h *HEAD) Len() int {
    switch next := h.Next.(type) {
    case *TAIL:
        return 0;
    case *NODE:
        return next.Len();
    }
    return 0;
}

// Get retrieves a Monster at the specified index
func (h *HEAD) Get(i int) (bool, Monster) {
	return h.next.Get(i);
}

// Shift removes the first element from the dequeue
func (h *HEAD) Shift() Dequeue {
	// TODO
}


// Drop removes the last element from the dequeue
func (h *HEAD) Drop() Dequeue {
	newNext := h.next.Drop();
	return &HEAD{next: newNext}
}


// NODE METHODS

// DONE
// Prepend adds a new monster to the front of the dequeue
func (n *NODE) Prepend(m Monster) Dequeue {
    // Find HEAD and prepend there
    curr := n.Prev;
    for {
        // Check if the current node is a HEAD
		if head, isHead := curr.(*HEAD); isHead {
            return head.Prepend(m);
        }
        // Check if the current node is a NODE and move to the previous node
        if prevNode, isNode := curr.(*NODE); isNode {
            curr = prevNode.Prev;
        } else {
            break;
        }
    }
    return n;
}

// Append adds a new Monster to the end of the dequeue
func (n *NODE) Append(m Monster) Dequeue {
    switch next := n.Next.(type) {
    case *TAIL:
        newNode := &NODE{Data: m, Prev: n, Next: next};
        next.Prev = newNode;
        n.Next = newNode;
        return n;
    case *NODE:
        return next.Append(m);
    }
    return n;
}

// Len returns the number of Monsters from this node to the end
func (n *NODE) Len() int {
    switch next := n.Next.(type) {
    case *TAIL:
        return 1;
    case *NODE:
        return 1 + next.Len();
    }
    return 1;
}

// Get retrieves a Monster at the specified relative index
func (n *NODE) Get(i int) (bool, Monster) {
	// TODO
	return false, Monster{};
}

// Shift removes the first element from the dequeue
func (n *NODE) Shift() Dequeue {
	// TODO
}

// Drop removes the last element from the dequeue
func (n *NODE) Drop() Dequeue {
	// TODO
}


// TAIL METHODS

// DONE
// Prepend adds a new Monster to the front of the dequeue
func (t *TAIL) Prepend(m Monster) Dequeue {
    // Start traversal from the node before the tail
	curr := t.Prev;
    for {
        // Check if current node is the HEAD
		if head, isHead := curr.(*HEAD); isHead {
            return head.Prepend(m);
        }
        // Move to the next previous node
		curr = curr.(*NODE).Prev;
    }
}

// DONE
// Append adds a new Monster to the end of the dequeue
func (t *TAIL) Append(m Monster) Dequeue {
    curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Append(m);
        }
        curr = curr.(*NODE).Prev;
    }
}

// DONE
// Len returns the length of the tail; always 0
func (t *TAIL) Len() int {
	return 0;
}

// Get always returns false and empty monster for TAIL
func (t *TAIL) Get(i int) (bool, Monster) {
	return false, Monster{};
}

// Shift returns the dequeue unchanged for TAIL
func (t *TAIL) Shift() Dequeue {
	// TODO
	return backwardTraverse(t);
}

// Drop removes the last element from the dequeue
func (t *TAIL) Drop() Dequeue {
	// TODO
}



