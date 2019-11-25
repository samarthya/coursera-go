package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	Implement the dining philosopher’s problem with the following constraints/modifications.

    1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
    2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
    *
    3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
    *
    4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
    *
    5. The host allows no more than 2 philosophers to eat concurrently.
    *
    6. Each philosopher is numbered, 1 through 5.
    *
    7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
       on a line by itself, where <number> is the number of the philosopher.
    *
    8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
       on a line by itself, where <number> is the number of the philosopher.
    *
*/

// ChopS chopsticks
type ChopS struct{ sync.Mutex }

// Philo Philosophers
type Philo struct {
	// id Condition (6) satisfied.
	id              int
	eligible        int
	leftCS, rightCS *ChopS
}

// Go has two allocation primitives, the built-in functions new and make.
// new(T) allocates zeroed storage for a new item of type T and
// returns its address, a value of type *T. In Go terminology, it returns
// a pointer to a newly allocated zero value of type T.

// NewPhilo constructor
func NewPhilo(name int, leftCS *ChopS, rightCS *ChopS) *Philo {
	ph := new(Philo)
	ph.id = name
	ph.eligible = 0
	ph.leftCS = leftCS
	ph.rightCS = rightCS

	return ph
}

// Each philo can send this as go routine.
func (p Philo) request(wg *sync.WaitGroup, wgHost chan int, wgPermission chan int, respHost chan int) {

	// While it has not eaten three times continue to request.
	for p.eligible < 3 {

		// Requesting the host.
		wgHost <- p.id

		// Waiting for grant
		grant := <-wgPermission

		fmt.Println(grant, " Permission acquired. ")
		p.dinnerTime()

		p.eligible++
		// respHost <- p.id
	}

	wg.Done()

}

// eat Allows the philosopher to eat.
func (p Philo) dinnerTime() {
	rand.Seed(3)
	randomN := rand.Int()

	// fmt.Println(randomN, " : random number ")

	// Condition (3) must be satisfied to pick any random order and use the same order
	// to return
	if randomN%4 == 0 {
		fmt.Printf("[%d] Picked Left Chopstick first.\n", p.id)
		p.leftCS.Lock()
		p.rightCS.Lock()
	} else {
		// fmt.Printf("[%d] Picked Right Chopstick first.\n", p.id)
		p.rightCS.Lock()
		p.leftCS.Lock()
	}

	// Condition (7) and (8)
	fmt.Printf("Starting to eat : [%d]\n", p.id)
	p.eligible++
	// time.Sleep(2 * time.Second)
	fmt.Printf("Finising eating : [%d]\n", p.id)

	if randomN%2 == 0 {
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	} else {
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func hostUp(chHost chan int, chPermission chan int, chResponse chan int) {
	fmt.Println(" ### Host running and listening ###")

	for {

		val, ok := <-chHost

		if ok == false {
			fmt.Println(val, ok, "Time to exit, as all have eaten.")
			break
		} else {
			fmt.Printf("[%d] requested for eating permission.\n", val)
			chPermission <- val
		}
	}
}

/**
 * Main program.
 */
func main() {
	// var once sync.Once
	var wg sync.WaitGroup

	// wgHost (5) Condition - Should allow to allow only two Philosophers to eat
	chHost := make(chan int, 2)

	// wgPermission To grant eligibility
	chPermission := make(chan int)
	chResponse := make(chan int)

	var CSticks []*ChopS
	var philos []*Philo

	CSticks = make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos = make([]*Philo, 5)

	// Condition (1) 5 philosphers
	for i := 0; i < 5; i++ {
		philos[i] = NewPhilo(i, CSticks[i], CSticks[(i+1)%5])
	}

	fmt.Print(" Diners problem\n")

	// Do it only once.
	go hostUp(chHost, chPermission, chResponse)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].request(&wg, chHost, chPermission, chResponse)
	}
	wg.Wait()
	fmt.Println(" Done ")
}
