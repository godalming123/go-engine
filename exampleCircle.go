package main

func circleDemo(showGuis bool, showLogs bool) {
	screen := newScreen(901, 541, "CIRCLE DEMO", CameraOptions{}, showGuis, showLogs)

	screen.drawCircle(CircleOptns{center: Point{X: 270, Y: 270}, radius: 270, charecter: pixelCharecter, thickness: 5})
	// screen.drawCircle(CircleOptns{center: Point{X: 270, Y: 270}, radius: 269, charecter: pixelCharecter})
	// screen.drawCircle(CircleOptns{center: Point{X: 270, Y: 270}, radius: 268, charecter: pixelCharecter})
	screen.printContents()
}
