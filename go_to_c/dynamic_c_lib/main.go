package main

import (
	"fmt"
	"runtime"
)

func main() {
	solveLinearProgram()

	// Garbage collecting at the end of the program isn't necessary, but this
	// will illustrate how the Go wrapper finalizer frees the C structure.
	runtime.GC()
	runtime.GC()
	runtime.GC()
}

func solveLinearProgram() {
	lp := NewLP(0, 2)
	lp.AddConstraint([]float64{0.0, 110.0, 30.0}, LE, 4000.0)
	lp.AddConstraint([]float64{0.0, 1.0, 1.0}, LE, 75.0)
	lp.SetObjFn([]float64{0.0, 143.0, 60.0})
	lp.SetMaximize()

	fmt.Println("Go says: About to solve lineary program ...")
	lp.Solve()

	vars := lp.Variables()
	fmt.Println("\nGo says results are: ")
	fmt.Printf("Plant %.3f acres of barley\n", vars[0])
	fmt.Printf("And  %.3f acres of wheat\n", vars[1])
	fmt.Printf("For optimal profit of $%.2f\n\n", lp.Objective())
}
