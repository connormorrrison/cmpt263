package dequeue;

// Monster data type
type Monster struct {
	Name string; // Non-empty, no whitespace
	InitiativeModifier int; // (-20, 20)
	ChallengeRating int; // > 0
	ArmourClass int; // >= 0
	AttackRating int; // >= -20
}

// Dequeue interface defines the operations on the structure
type Dequeue interface {
	Prepend(m Monster) Dequeue;
	Append(m Monster) Dequeue;
	Len() int;
	Get(i int) (bool, Monster);  // bool if Monster was found
	Shift();
	Drop();
}

// HEAD variant: start of dequeue
type HEAD struct {
	next Dequeue;
}

// NODE variant: contains Monster and links prev and next nodes
type NODE struct {
	data Monster;
	next Dequeue;
	prev Dequeue;
}

// TAIL variant: end of dequeue
type TAIL struct {
	prev Dequeue;
}

func MakeDequeue() Dequeue {
	// Create HEAD and TAIL
	var head HEAD;
	var tail TAIL;

	// Link head and tail together
	head.next = &tail;
	tail.prev = &head;

	// Return the head
	return &head;
}

// Finds head, starting at tail
func backwardTraverse(d Dequeue) *HEAD {
	switch target := d.(type) {
	case *HEAD:
		return target;
	case *NODE:
		return backwardTraverse(target.prev);
	case *TAIL:
		return backwardTraverse(target.prev);
	}
	return nil;
}


// HEAD methods

func (h *HEAD) Prepend(m Monster) Dequeue {
	// TODO
	n := &NODE{data: m, prev: h, next: h.next};
	switch nxt := h.next.(type) {
	// If the next node is a NODE
	case *NODE:
		nxt.prev = n;
	// If the next node is a TAIL
	case *TAIL:
		nxt.prev = n;
	}
	h.next = n;
	return h;
}

func (h *HEAD) Append(m Monster) Dequeue {
	// TODO
	newNext := h.next.Append(m);
	return &HEAD{next: newNext};
}

func (h *HEAD) Len() int {
	// TODO
	switch h.next.(type) {
	case *TAIL:
		return 0;
	default:
		return h.next.Len();
	}
}

func (h *HEAD) Get(i int) (bool, Monster) {
	// TODO
	return h.next.Get(i);
}

func (h *HEAD) Shift() {
	// TODO
}

func (h *HEAD) Drop() {
	// TODO
}


// NODE methods

func (n *NODE) Prepend(m Monster) Dequeue {
	// TODO
	h := backwardTraverse(n);
	return h.Prepend(m);
}

func (n *NODE) Append(m Monster) Dequeue {
	// TODO
	return n;
}

func (n *NODE) Len() int {
	// TODO
	return 1 + n.next.Len(); // example logic (once next is implemented)
}

func (n *NODE) Get(i int) (bool, Monster) {
	// TODO
	return false, Monster{};
}

func (n *NODE) Shift() {
	// TODO
}

func (n *NODE) Drop() {
	// TODO
}


// TAIL methods

// Prepend adds a new Monster to the front of the dequeue
func (t *TAIL) Prepend(m Monster) Dequeue {
	h := backwardTraverse(t);
	return h.Prepend(m);
}

// Append adds a new Monster to the end of the dequeue
func (t *TAIL) Append(m Monster) Dequeue {
	// Insert a NODE before the TAIL
	n := &NODE{data: m, next: t, prev: t.prev};
	switch p := t.prev.(type) {
	case *HEAD:
		p.next = n;
	case *NODE:
		p.next = n;
	}
	t.prev = n;
	return backwardTraverse(t);
}

// Len returns the length of the tail; always 0
func (t *TAIL) Len() int {
	return 0;
}

// Get always returns false and empty monster for TAIL
func (t *TAIL) Get(i int) (bool, Monster) {
	return false, Monster{};
}

// Shift returns the dequeue unchanged for TAIL
func (t *TAIL) Shift() {
	// TODO
}

func (t *TAIL) Drop() {
	// TODO
}



