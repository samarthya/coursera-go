package main

import (
	"fmt"
	"sync"
)

const (
	maxEating     = 3
	fullnessLimit = 3
)

// Chopsticks is an object structure that serves as a tool to eat
type Chopsticks struct {
	mutex  sync.RWMutex
	ID     int
	isUsed bool
}

// NewChopsticks creates a Chopsticks object with the given id label.
func NewChopsticks(id int) *Chopsticks {
	return &Chopsticks{
		mutex:  sync.RWMutex{},
		ID:     id,
		isUsed: false,
	}
}

// Use is to set the Chopsticks status to "In-Use"
func (c *Chopsticks) Use() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.isUsed = true
}

// Free is to set the Chopsticks status to "Free"
func (c *Chopsticks) Free() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.isUsed = false
}

// IsInUse is to check the Chopsticks status, currently "In-Use" or "Free".
func (c *Chopsticks) IsInUse() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.isUsed
}

// Table is the structure serving foods and chopsticks
type Table struct {
	mutex      sync.RWMutex
	isEating   uint
	chopsticks []*Chopsticks
}

// NewTable creates the table object with person quantity.
func NewTable(quantity uint) *Table {
	t := &Table{
		mutex:      sync.RWMutex{},
		isEating:   0,
		chopsticks: []*Chopsticks{},
	}
	for i := 0; i < int(quantity); i++ {
		t.chopsticks = append(t.chopsticks, NewChopsticks(i))
	}

	return t
}

// RequestChopsticks is to allows a customer to eat using an available
// Chopsticks.
func (t *Table) RequestChopsticks() *Chopsticks {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// table is full
	if t.isEating >= maxEating {
		return nil
	}

	// permit to eat. Scan for available chopsticks
	c := t.seekChopsticks()
	c.Use()
	t.isEating++
	return c
}

func (t *Table) seekChopsticks() *Chopsticks {
	// NOTE: here, you can use random. I will use FIFO instead.
	for _, c := range t.chopsticks {
		if !c.IsInUse() {
			return c
		}
	}
	return nil
}

// ReturnChopsticks is to allow a customer to place back chopsticks when
// he/she is done eating
func (t *Table) ReturnChopsticks(c *Chopsticks) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.isEating--
	c.Free()
}

func philosopher(id int, fullness chan int, table *Table) {
	eatCount := fullnessLimit

	for {
		chopsticks := table.RequestChopsticks()
		if chopsticks == nil {
			continue
		}

		// start eating
		fmt.Printf("P%d: START eating with chopstick %d.\n",
			id,
			chopsticks.ID)

		// eating
		eatCount--

		// stop eating
		fmt.Printf("P%d: FINISH eating with chopstick %d.\n",
			id,
			chopsticks.ID)
		table.ReturnChopsticks(chopsticks)

		// check fullness
		if eatCount == 0 {
			fullness <- id
			return
		}

	}
}

func main() {
	headcount := 5

	t := NewTable(uint(headcount))
	c1 := make(chan int)
	fullness := 0

	for i := 0; i < headcount; i++ {
		go philosopher(i, c1, t)
	}

	// Wait for fullness
	for {
		select {
		case person := <-c1:
			fullness++
			fmt.Printf("Philosopher %d is full\n", person)
			if fullness == headcount {
				fmt.Printf("All are full.\n[ ENDED ]\n")
				return
			}
		}
	}
}
