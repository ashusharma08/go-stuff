package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	hydrogen chan struct{} // Controls hydrogen atom release
	oxygen   chan struct{} // Controls oxygen atom release
	barrier  *sync.WaitGroup
}

func NewH2O() *H2O {
	return &H2O{
		hydrogen: make(chan struct{}, 2), // Allows two hydrogen atoms at a time
		oxygen:   make(chan struct{}, 1), // Allows one oxygen atom at a time
		barrier:  &sync.WaitGroup{},
	}
}

func (h2o *H2O) ghydrogen(releaseHydrogen func()) {
	h2o.hydrogen <- struct{}{} // Add a hydrogen atom
	h2o.barrier.Wait()         // Wait for other atoms to arrive
	releaseHydrogen()          // Print 'H'
}

func (h2o *H2O) goxygen(releaseOxygen func()) {
	h2o.oxygen <- struct{}{} // Add an oxygen atom
	h2o.barrier.Add(2)       // Ensure 2 hydrogen atoms are present
	h2o.barrier.Wait()       // Wait for other atoms to arrive
	releaseOxygen()          // Print 'O'

	// Reset for the next molecule
	<-h2o.hydrogen
	<-h2o.hydrogen
	<-h2o.oxygen
	h2o.barrier = &sync.WaitGroup{} // Reset barrier
}

func main() {
	h2o := NewH2O()

	var wg sync.WaitGroup
	releaseHydrogen := func() { fmt.Print("H") }
	releaseOxygen := func() { fmt.Print("O") }

	// Simulate multiple threads forming water
	for i := 0; i < 10; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			h2o.ghydrogen(releaseHydrogen)
		}()
		go func() {
			defer wg.Done()
			h2o.ghydrogen(releaseHydrogen)
		}()
		go func() {
			defer wg.Done()
			h2o.goxygen(releaseOxygen)
		}()
	}

	wg.Wait()
}
