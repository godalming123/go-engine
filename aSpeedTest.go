package main

import (
	"fmt"
	"time"
)

func speedTest(showGuis bool, showLogs bool) {
	// test time to create 100 screens
	fmt.Println("Testing time to create screens")
	start := time.Now()

	var screen Screen
	for i := 0; i < 100; i++ {
		screen = newScreen(901, 541, "test", CameraOptions{}, showGuis, showLogs)
	}

	fmt.Println("Took:", time.Since(start))

	// test time to draw 100 lines
	fmt.Println("Testing time to draw lines")
	start = time.Now()

	for i := 0; i < 100; i++ {
		screen.bresignham3D(Point{X: 0, Y: 0}, Point{X: 250, Y: 319}, pixelCharecter)
	}

	fmt.Println("Took:", time.Since(start))

	// test time to draw 100 circles
	fmt.Println("Testing time to draw circles")
	start = time.Now()

	for i := 0; i < 100; i++ {
		screen.drawCircle(CircleOptns{center: Point{X: 270, Y: 270}, radius: 270, thickness: 5, charecter: pixelCharecter})
	}

	fmt.Println("Took:", time.Since(start))

	screen.printContents()
}
