package main

func graphDemo(showGuis bool, showLogs bool) {
	screen := newScreen(32, 32, "GRAPH DEMO", CameraOptions{}, showGuis, showLogs)

	screen.bresignham3D(Point{Y: 00, X: 00}, Point{Y: 16, X: 20}, pixelCharecter)
	screen.bresignham3D(Point{Y: 17, X: 21}, Point{Y: 25, X: 24}, pixelCharecter)
	screen.bresignham3D(Point{Y: 24, X: 25}, Point{Y: 06, X: 31}, pixelCharecter)
	screen.printContents()
}
