package main

func setFewPixels(showGuis bool, showLogs bool) {
	screen := newScreen(11, 11, "SET FEW PIXELS", CameraOptions{}, showGuis, showLogs)

	x1 := uint(0)
	x2 := uint(10)
	y1 := uint(0)
	y2 := uint(5)
	y3 := uint(10)

	screen.setPix(Point{X: x1, Y: y1}, pixelCharecter)
	screen.setPix(Point{X: x1, Y: y2}, pixelCharecter)
	screen.setPix(Point{X: x1, Y: y3}, pixelCharecter)
	screen.setPix(Point{X: x2, Y: y1}, pixelCharecter)
	screen.setPix(Point{X: x2, Y: y2}, pixelCharecter)
	screen.setPix(Point{X: x2, Y: y3}, pixelCharecter)

	screen.printContents()
}
