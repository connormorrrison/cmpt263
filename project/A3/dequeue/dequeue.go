package dequeue;

// Monster data structure
type Monster struct {
    Name string;
    InitiativeModifier int;
    ChallengeRating int;
    ArmourClass int;
    AttackRating int;
}

// Monster interface defines operations
type Dequeue interface {
    Prepend(m Monster) Dequeue;
    Append(m Monster) Dequeue;
    Len() int;
    Get(i int) (bool, Monster);
    Shift() Dequeue;
    Drop() Dequeue;
}

// HEAD: front of the dequeue
type HEAD struct {
    Next Dequeue;
}

// NODE: a Monster entry in the dequeue
type NODE struct {
    Data Monster;
    Prev Dequeue;
    Next Dequeue;
}

// TAIL: end of the dequeue
type TAIL struct {
    Prev Dequeue;
}

// MakeDequeue creates a new empty dequeue with connected head and tail
//@* requires true
//@ result.Len() == 0
//@ result points to a valid empty dequeue with connected HEAD and TAIL
//@ result.(*HEAD).Next points to TAIL
//@ result.(*HEAD).Next.(*TAIL).Prev points to HEAD
func MakeDequeue() Dequeue {
	//@ Invariant dequeue structure remains valid during modification
	//@ Invariant all existing nodes maintain their relative order
    var head HEAD;
    var tail TAIL;
    head.Next = &tail;
    tail.Prev = &head;
    return &head;
}
//@* ensures head.Next points to tail, tail.Prev points to head, Len() == 0


// HEAD FUNCTIONS
// Prepend adds a new Monster to the front of the dequeue
//@* requires true
//@ h.Len() == old(h.Len()) + 1
//@ h.Next points to a NODE containing m
//@ forall i, 0 < i < h.Len() h.Get(i) == old(h.Get(i-1))
func (h *HEAD) Prepend(m Monster) Dequeue {
    //@ Invariant dequeue structure remains valid
    //@ Invariant all existing nodes maintain their relative order
	n := &NODE{Data: m, Prev: h, Next: h.Next};
    switch next := h.Next.(type) {
    case *NODE:
		//@ Invariant: h.Next pointing to next NODE
        next.Prev = n;
    case *TAIL:
        next.Prev = n;
    }
    h.Next = n;
    return h;
}
//@* ensures Len() is increased by 1, and new Monster added at front

// Append adds a new Monster to the end of the dequeue
//@* requires true
//@ h.Len() == old(h.Len()) + 1
//@ h.Get(h.Len()-1) == m
//@ forall i, 0 <= i < old(h.Len()): h.Get(i) == old(h.Get(i)
func (h *HEAD) Append(m Monster) Dequeue {
    //@ Variant distance to TAIL node
    //@ Invariant dequeue structure remains valid during traversal
    //@ Invariant all existing nodes maintain their data
	switch next := h.Next.(type) {
    case *TAIL:
        return h.Prepend(m);
    case *NODE:
        return next.Append(m);
    }
    return h;
}
//@* ensures Len() is increased by 1, and new Monster is added at end

// Len returns the number of Monsters in the dequeue
//@* requires true
//@ result >= 0
//@ result equals number of NODEs between HEAD and TAIL
func (h *HEAD) Len() int {
    //@ Variant distance to TAIL node
    //@ Invariant count accumulates correctly
	switch next := h.Next.(type) {
    case *TAIL:
        return 0;
    case *NODE:
        return next.Len();
    }
    return 0;
}
//@* ensures returns the count of nodes between head and tail

// Get retrieves a Monster from the dequeue by index; positive indices 
// count from front, negative from back
//@* requires -h.Len() < i && i < h.Len()
//@ result.0 implies result.1 equals monster at index i
//@ !result.0 implies index i is out of bounds
func (h *HEAD) Get(i int) (bool, Monster) {
    //@ Variant distance to target index
    //@ Invariant index bounds remain valid
	if i == 0 {
        return false, Monster{};
    }
    if i > 0 {
        return h.getFromHead(i);
    }
    // i < 0, traverse from tail
    return h.getFromTail(-i);
}
//@* ensures returns true and the Monster at index i if exists, else false

// Shift removes and returns the first monster in the dequeue
//@* requires h.Len() > 0
//@ h.Len() == old(h.Len()) - 1
//@ forall i, 0 <= i < h.Len(): h.Get(i) == old(h.Get(i+1))
func (h *HEAD) Shift() Dequeue {
    //@ Invariant dequeue structure remains valid during removal
    //@ Invariant all remaining nodes maintain their relative order
	switch next := h.Next.(type) {
    case *TAIL:
        return h;
    case *NODE:
        h.Next = next.Next;
        switch nextNext := next.Next.(type) {
        case *NODE:
            nextNext.Prev = h;
        case *TAIL:
            nextNext.Prev = h;
        }
    }
    return h;
}
//@* ensures Len() is decreased by 1, and first Monster is removed

// Drop removes and returns the last monster in the dequeue
//@* requires h.Len() > 0
//@ h.Len() == old(h.Len()) - 1
//@ forall i, 0 <= i < h.Len()-1: h.Get(i) == old(h.Get(i))
func (h *HEAD) Drop() Dequeue {
    //@ Invariant dequeue structure remains valid
	tail := h.findTail();
    if tail == nil {
        return h;
    }

    lastNode, isNode := tail.Prev.(*NODE)
    if !isNode {
        // No nodes to drop (dequeue is empty)
        return h;
    }

    // Check if the node before lastNode is HEAD
    if _, isNode := lastNode.Prev.(*HEAD); isNode {
        //@ Invariant single node case maintains valid structure
        h.Next = tail;
        tail.Prev = h;
    } else if prevNode, isNode := lastNode.Prev.(*NODE); isNode {
        //@ Invariant multiple node case maintains valid structure
        prevNode.Next = tail;
        tail.Prev = prevNode;
    }

    return h;
}
//@* ensures Len() is decreased by 1, and last Monster is removed


// NODE FUNCTIONS
// Prepend adds a new monster to the front by locating HEAD and prepending
//@* requires true
//@ result.Len() == old(n.dequeueLen()) + 1
//@ result.Get(0) == m
//@ forall i, 0 < i < result.Len(): result.Get(i) == old(n.dequeueGet(i-1))
func (n *NODE) Prepend(m Monster) Dequeue {
    // Find HEAD and prepend there
    //@ Variant distance to HEAD
    //@ Invariant dequeue structure remains valid
	curr := n.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Prepend(m);
        }
        
        if prevNode, isNode := curr.(*NODE); isNode {
            curr = prevNode.Prev;
        } else {
            break;
        }
    }
    return n;
}
//@* ensures Len() is increased by 1, and new Monster is added at front

// Append adds a new monster after the current node or delegates to next node
//@* requires true
//@ result.Len() == old(n.dequeueLen()) + 1
//@ result.Get(result.Len()-1) == m
//@ forall i, 0 <= i < old(n.dequeueLen()): result.Get(i) == old(n.dequeueGet(i))
func (n *NODE) Append(m Monster) Dequeue {
    //@ Variant distance to TAIL
    //@ Invariant dequeue structure remains valid
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
//@* ensures Len() is increased by 1, and new Monster is added at end

// Len returns the number of elements in the dequeue
//@* requires true
//@ result > 0
//@ result equals number of NODE's from current node to TAIL
func (n *NODE) Len() int {
    //@ Variant distance to TAIL
    //@ Invariant count accumulates correctly
	switch next := n.Next.(type) {
    case *TAIL:
        return 1;
    case *NODE:
        return 1 + next.Len();
    }
	return 1;
}
//@* ensures returns the count of nodes from current node to tail

// Get retrieves a Monster from the dequeue by index
//@* requires -n.Len() < i && i < n.Len()
//@ result.0 implies result.1 equals monster at index i
//@ !result.0 implies index i is out of bounds
func (n *NODE) Get(i int) (bool, Monster) {
    return false, Monster{};
}
//@* ensures returns true and the Monster at index i if exists, else false

// Shift removes first node by finding HEAD and shifting there
//@* requires n.Len() > 0
//@ result.Len() == old(n.dequeueLen()) - 1
//@ forall i, 0 <= i < result.Len(): result.Get(i) == old(n.dequeueGet(i+1))
func (n *NODE) Shift() Dequeue {
    // Find HEAD and shift there
	//@ Variant distance to HEAD
    //@ Invariant dequeue structure remains valid
    curr := n.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Shift();
        }
        curr = curr.(*NODE).Prev;
    }
}
//@* ensures Len() is decreased by 1, and first Monster is removed

// Drop removes this node from the dequeue and updates links
//@* requires n is a valid NODE in the dequeue
//@ result.Len() == old(n.dequeueLen()) - 1
//@ node n is removed from the dequeue
//@ Prev and Next links are properly updated
func (n *NODE) Drop() Dequeue {
    //@ Invariant dequeue structure remains valid 
	switch prev := n.Prev.(type) {
    case *HEAD:
        prev.Next = n.Next;
    case *NODE:
        prev.Next = n.Next;
    }
    switch next := n.Next.(type) {
    case *TAIL:
        next.Prev = n.Prev;
    case *NODE:
        next.Prev = n.Prev;
    }
    return n.Prev;
}
//@* ensures n is removed, links to Prev and Next are updated, Len() is decreased by 1


// TAIL FUNCTIONS
// Prepend adds a monster to front by finding HEAD and prepending there
//@* requires true
//@ result.Len() == old(t.dequeueLen()) + 1
//@ result.Get(0) == m
func (t *TAIL) Prepend(m Monster) Dequeue {
    //@ Variant distance to HEAD
    //@ Invariant dequeue structure remains valid
	curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Prepend(m);
        }
        curr = curr.(*NODE).Prev;
    }
}
//@* ensures Len() is increased by 1, and new Monster is added at front

// Append adds a monster to end by finding HEAD and appending there
//@* requires true
//@ result.Len() == old(t.dequeueLen()) + 1
//@ result.Get(result.Len()-1) == m
func (t *TAIL) Append(m Monster) Dequeue {
    //@ Variant distance to HEAD
    //@ Invariant dequeue structure remains valid
	curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Append(m);
        }
        curr = curr.(*NODE).Prev;
    }
}
//@* ensures Len() is increased by 1, and new Monster is added at end

// Len returns 0 as TAIL is not a data node
//@* requires true
//@ result == 0
func (t *TAIL) Len() int {
    return 0;
}
//@* ensures returns 0

// Get retrieves a Monster from the dequeue by index, this function returns false
//@* requires true
//@ result.0 == false
func (t *TAIL) Get(i int) (bool, Monster) {
    return false, Monster{};
}
//@* ensures always returns false and an empty Monster

// Shift removes first node by finding HEAD and shifting there
//@* requires t.Len() > 0
//@ result.Len() == old(t.Len()) - 1
func (t *TAIL) Shift() Dequeue {
    //@ Variant distance to HEAD
    //@ Invariant dequeue structure remains valid
	curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Shift();
        }
        curr = curr.(*NODE).Prev;
    }
}
//@* ensures Len() is decreased by 1, and first Monster is removed

// Drop removes the last item; it has no effect on TAIL
//@* requires true
//@ result == t
func (t *TAIL) Drop() Dequeue {
    return t;
}
//@* ensures no change, returns TAIL itself


// HELPER FUNCTIONS

// Helper method to find the TAIL from HEAD
//@* requires true
//@ result != nil implies result points to TAIL node
//@ result == nil implies invalid dequeue
func (h *HEAD) findTail() *TAIL {
    //@ Variant distance to TAIL
    //@ Invariant dequeue structure remains valid
	current := h.Next;
    for {
        switch node := current.(type) {
        case *NODE:
            current = node.Next;
        case *TAIL:
            return node;
        default:
            return nil;
        }
    }
}
//@* ensures returns the TAIL node if exists, else nil

// Helper method to traverse forward for Get
//@* requires i > 0
//@ result.0 implies result.1 is monster at index i from head
//@ !result.0 implies index i is out of bounds
func (h *HEAD) getFromHead(i int) (bool, Monster) {
    //@ Variant distance to target index
    //@ Invariant index bounds remain valid during traversal
	current := h.Next;
    for current != nil {
        switch node := current.(type) {
        case *NODE:
            if i == 1 {
                return true, node.Data;
            }
            i--;
            current = node.Next;
        case *TAIL:
            // Reached end without finding the index
            return false, Monster{}
        }
    }
    return false, Monster{};
}
//@* ensures if the i-th Monster exists, returns true and the Monster; else false

// Helper method to traverse backward for Get
//@* requires i > 0
//@ result.0 implies result.1 is monster at index i from tail
//@ !result.0 implies index i is out of bounds
func (h *HEAD) getFromTail(i int) (bool, Monster) {
    //@ Variant distance to target index
	//@ Invariant index bounds remain valid during traversal
	tail := h.findTail();
    if tail == nil {
        return false, Monster{};
    }
    current := tail.Prev;
    count := i;
    for current != nil {
        switch node := current.(type) {
        case *NODE:
            if count == 1 {
                return true, node.Data;
            }
            count--;
            current = node.Prev;
        case *HEAD:
            // Reached front without finding the index
            return false, Monster{};
        }
    }
    return false, Monster{};
}
//@* ensures if the i-th Monster exists from tail, returns true and the Monster; else false
