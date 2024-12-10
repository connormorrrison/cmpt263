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
func MakeDequeue() Dequeue {
    var head HEAD;
    var tail TAIL;
    head.Next = &tail;
    tail.Prev = &head;
    return &head;
}


// HEAD FUNCTIONS
// Prepend adds a new Monster to the front of the dequeue
func (h *HEAD) Prepend(m Monster) Dequeue {
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

// Append adds a new Monster to the end of the dequeue
func (h *HEAD) Append(m Monster) Dequeue {
    switch next := h.Next.(type) {
    case *TAIL:
        return h.Prepend(m);
    case *NODE:
        return next.Append(m);
    }
    return h;
}

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

// Get retrieves a Monster from the dequeue by index; positive indices 
// count from front, negative from back
func (h *HEAD) Get(i int) (bool, Monster) {
    if i == 0 {
        return false, Monster{};
    }
    if i > 0 {
        return h.getFromHead(i);
    }
    // i < 0, traverse from tail
    return h.getFromTail(-i);
}

// Shift removes and returns the first monster in the dequeue
func (h *HEAD) Shift() Dequeue {
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

// Drop removes and returns the last monster in the dequeue
func (h *HEAD) Drop() Dequeue {
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
        // Only one node exists; remove it by linking HEAD directly to TAIL
        h.Next = tail;
        tail.Prev = h;
    } else if prevNode, isNode := lastNode.Prev.(*NODE); isNode {
        // More than one node exists; bypass lastNode
        prevNode.Next = tail;
        tail.Prev = prevNode;
    }

    return h;
}


// NODE FUNCTIONS
// Prepend adds a new monster to the front by locating HEAD and prepending
func (n *NODE) Prepend(m Monster) Dequeue {
    // Find HEAD and prepend there
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

// Append adds a new monster after the current node or delegates to next node
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

// Len returns the number of elements in the dequeue
func (n *NODE) Len() int {
    switch next := n.Next.(type) {
    case *TAIL:
        return 1;
    case *NODE:
        return 1 + next.Len();
    }
    return 1;
}

// Get retrieves a Monster from the dequeue by index
func (n *NODE) Get(i int) (bool, Monster) {
    return false, Monster{};
}

// Shift removes first node by finding HEAD and shifting there
func (n *NODE) Shift() Dequeue {
    // Find HEAD and shift there
    curr := n.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Shift();
        }
        curr = curr.(*NODE).Prev;
    }
}

// Drop removes this node from the dequeue and updates links
func (n *NODE) Drop() Dequeue {
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


// TAIL FUNCTIONS
// Prepend adds a monster to front by finding HEAD and prepending there
func (t *TAIL) Prepend(m Monster) Dequeue {
    curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Prepend(m);
        }
        curr = curr.(*NODE).Prev;
    }
}

// Append adds a monster to end by finding HEAD and appending there
func (t *TAIL) Append(m Monster) Dequeue {
    curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Append(m);
        }
        curr = curr.(*NODE).Prev;
    }
}

// Len returns 0 as TAIL is not a data node
func (t *TAIL) Len() int {
    return 0;
}

// Get retrieves a Monster from the dequeue by index, this function returns false
func (t *TAIL) Get(i int) (bool, Monster) {
    return false, Monster{};
}

// Shift removes first node by finding HEAD and shifting there
func (t *TAIL) Shift() Dequeue {
    curr := t.Prev;
    for {
        if head, isHead := curr.(*HEAD); isHead {
            return head.Shift();
        }
        curr = curr.(*NODE).Prev;
    }
}

// Drop removes the last item; it has no effect on TAIL
func (t *TAIL) Drop() Dequeue {
    return t;
}


// HELPER FUNCTIONS

// Helper method to find the TAIL from HEAD
func (h *HEAD) findTail() *TAIL {
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

// Helper method to traverse forward for Get
func (h *HEAD) getFromHead(i int) (bool, Monster) {
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

// Helper method to traverse backward for Get
func (h *HEAD) getFromTail(i int) (bool, Monster) {
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