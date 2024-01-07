package model

type Counter struct {
	Count int
}

// Methods with pointer receivers are often referred to as "mutator methods"
// They have the capability to modify the state of the original instance they are called on.
func (c *Counter) Increment() {
	c.Count++
}

func NewCounterPtr(c int) *Counter {
	return &Counter{Count: c}
}

func NewCounter(c int) Counter {
	return Counter{Count: c}
}
