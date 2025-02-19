package main

import (
	"fmt"
	"github.com/jwalton/gchalk"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
)

// === RUNE KEY CODES ===

var pixelCharecter = []rune("█")[0]

// === STRUCTS ===

type Point struct {
	X              uint
	Y              uint
	Z              uint
	has3d          bool
	hasBeenMoved   bool
	hasBeenRotated bool
}

type Edge struct {
	pointA        Point
	pointB        Point
	edgeCharecter rune
}

type Face struct {
	edges         []Edge
	fillCharecter string
}

type Shape3d struct {
	faces []Face
}

type CircleOptns struct {
	center    Point
	startAt   float64
	endAt     float64
	radius    uint
	charecter rune
	thickness uint
}

type CameraOptions struct {
	xTranslate   int
	yTranslate   int
	zTranslate   int
	xAngle       float64
	yAngle       float64
	zAngleY      float64
	zAngleX      float64
	screenCenter Point
}

type Screen struct {
	width            uint
	height           uint
	contents         []rune // can't use a string as they dont split nicely with unicode
	originalContents []rune
	items            []Shape3d
	cameraOpts       CameraOptions
	showLogs         bool
	showGuis         bool
	name             string
}

// === HELPERS ===

func addWhitespace(input string, minimumChars uint) string {
	inputLength := len(input)
	if inputLength <= int(minimumChars) { // if the number will fit into the minimum charects
		// then add required whitespace and return
		whitespaceChars := strings.Repeat(" ", int(minimumChars)-inputLength)
		return input + whitespaceChars
	} else {
		return "E"
	}
}

var logLevel = 0

func logSameLevel(showLogs bool, toLog ...any) {
	if showLogs {
		fmt.Printf(strings.Repeat(" | ", logLevel))
		fmt.Println(toLog...)
	}
}

func copysign(a int, b int) int {
	if b >= 0 {
		return a
	} else {
		return -a
	}
}

func absInt(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux only
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func typeToQuit(message string) {
	fmt.Println(message)
	var i []byte = make([]byte, 1)
	os.Stdin.Read(i)
}

func toggleBtn(option string, toggled bool) string {
	if toggled {
		return option + " ON"
	} else {
		return option + " OFF"
	}
}

func drawList(header string, subtitle string, items []string, selectedIndex uint) {
	if header != "" {
		fmt.Println(gchalk.Bold(header))
	}

	if subtitle != "" {
		fmt.Println(subtitle)
	}

	for i := 0; i < len(items); i++ {
		text := " - " + items[i]
		if i == int(selectedIndex) {
			fmt.Println(gchalk.Bold(text))
		} else {
			fmt.Println(text)
		}
	}
}

func selectTui(header string, options []string, selectedIndex uint) uint {
	var i []byte = make([]byte, 1)
	clearScreen()
	drawList(header, "  w - up, s -down, d - select", options, selectedIndex)
	for true {
		os.Stdin.Read(i)
		input := string(i)
		if (input == "w" || input == "W") && selectedIndex > 0 {
			selectedIndex -= 1
		} else if (input == "s" || input == "S") && selectedIndex < uint(len(options)-1) {
			selectedIndex += 1
		} else if input == "d" || input == "D" {
			return selectedIndex
		}
		clearScreen()
		drawList(header, "  w - up, s -down, d - select", options, selectedIndex)
	}
	// Just to keep the go code checking happy we must return a number even if ln. 142 does it for us
	return 0
}

// === FUNCTIONS FOR SCREEN STRUCT ===

func (screen *Screen) bresignham3D(a Point, b Point, newChar rune) {
	a = transformAndRotate(a, screen.cameraOpts)
	b = transformAndRotate(b, screen.cameraOpts)

	dx := absInt(int(b.X - a.X))
	dy := absInt(int(b.Y - a.Y))
	dz := absInt(int(b.Z - a.Z))
	xs := copysign(1, int(b.X-a.X))
	ys := copysign(1, int(b.Y-a.Y))
	zs := copysign(1, int(b.Z-a.Z))

	logSameLevel(screen.showLogs, "Drawing bresignham line from X:", a.X, "Y:", a.Y, "Z:", a.Z, "to X:", b.X, "Y:", b.Y, "Z:", b.Z, "deferences X:", dx, "Y:", dy, "Z:", dz)
	logLevel += 1

	if (dx >= dy) && (dx >= dz) { // Driving axis is X-axis"
		logSameLevel(screen.showLogs, "Driving axis for line is X axis")
		screen.setPix(a, newChar)

		p1 := 2*dy - dx
		p2 := 2*dz - dx
		for a.X != b.X {
			a.X = uint(int(a.X) + xs)
			if p1 >= 0 {
				a.Y = uint(int(a.Y) + ys)
				p1 -= 2 * dx
			}
			if p2 >= 0 {
				a.Z = uint(int(a.Z) + zs)
				p2 -= 2 * dx
			}
			p1 += 2 * dy
			p2 += 2 * dz
			screen.setPix(a, newChar)
		}
	} else if (dy >= dx) && (dy >= dz) { // Driving axis is Y axis
		logSameLevel(screen.showLogs, "Driving axis for line is Y axis")
		screen.setPix(a, newChar)

		p1 := 2*dx - dy
		p2 := 2*dz - dy
		for a.Y != b.Y {
			a.Y = uint(int(a.Y) + ys)
			if p1 >= 0 {
				a.X = uint(int(a.X) + xs)
				p1 -= 2 * dy
			}
			if p2 >= 0 {
				a.Z = uint(int(a.Z) + zs)
				p2 -= 2 * dy
			}
			p1 += 2 * dx
			p2 += 2 * dz
			screen.setPix(a, newChar)
		}
	} else if (dz >= dx) && (dz >= dy) { // Driving axis is Z-axis"
		logSameLevel(screen.showLogs, "Driving axis for line is Z axis")
		screen.setPix(a, newChar)

		p1 := 2*dy - dz
		p2 := 2*dx - dz
		for a.Z != b.Z {
			a.Z = uint(int(a.Z) + zs)
			if p1 >= 0 {
				a.Y += uint(ys)
				p1 -= 2 * dz
			}
			if p2 >= 0 {
				a.X = uint(int(a.X) + xs)
				p2 -= 2 * dz
			}
			p1 += 2 * dy
			p2 += 2 * dx
			screen.setPix(a, newChar)
		}
	} else {
		fmt.Println("Could not find the driving axis")
	}
	logLevel -= 1
}

func (screen *Screen) raycastToContents() {
	// TODO make raycasting
}

func newScreen(width uint, height uint, name string, cameraOptions CameraOptions, showGuis bool, showLogs bool) Screen {
	if (cameraOptions == CameraOptions{}) { // if no camera options are present
		cameraOptions = CameraOptions{ // then generate them
			xTranslate: 0,
			yTranslate: 0,
			zTranslate: 0,
			xAngle:     -0.3, // adds to Y depending on X * xAngle
			yAngle:     0,    // adds to X depending on Y * yAngle
			zAngleY:    0.7,  // adds to Y depending on Z * zAngleY
			zAngleX:    1,    // adds to X depending on Z * zAngleX
			screenCenter: Point{
				X: width / 2,
				Y: height / 2,
			},
		}
	}
	lineContents := strings.Repeat("  ", int(width)) + " \n"
	contents := strings.Repeat(lineContents, int(height))
	runeContents := []rune(contents)
	logSameLevel(showLogs, "Screen created with", height, "pixels in height and", width, "pixels in width.")
	logLevel += 1

	return Screen{
		width:            width,
		height:           height,
		contents:         runeContents,
		originalContents: runeContents,
		cameraOpts:       cameraOptions,
		showLogs:         showLogs,
		showGuis:         showGuis,
		name:             name,
	}
}

func (screen *Screen) setPix(point Point, newChar rune) {
	point = transformAndRotate(point, screen.cameraOpts)
	if point.X < screen.width && point.Y < screen.height { // if the point fits in the screen
		charToSet := (point.Y * (screen.width + 1)) + point.X
		charToSet *= 2
		logSameLevel(
			screen.showLogs,
			"Pixel created at",
			"Y:", addWhitespace(fmt.Sprint(point.Y), 8),
			"X:", addWhitespace(fmt.Sprint(point.X), 8),
			"and the charecter will be:", newChar,
			"and the previous charecters were:", screen.contents[charToSet:charToSet+2],
			"and the charecter to set is", charToSet,
		)
		screen.contents[charToSet] = newChar
		screen.contents[charToSet+1] = newChar
	} else {
		logSameLevel(screen.showLogs, "It doesnt appear that the pixel is within the screens bounds")
	}
}

func (screen *Screen) drawEdge(edge Edge) {
	screen.bresignham3D(edge.pointA, edge.pointB, edge.edgeCharecter)
}

func drawCirclePixel(screen *Screen, origin Point, xc uint, yc uint, charecter rune) {
	// draws 8 pixels of a circle from 1 pixel see https://lectureloops.com/wp-content/uploads/2021/01/image-5.png explaining this process
	screen.setPix(Point{X: origin.X + xc, Y: origin.Y + yc}, charecter)
	screen.setPix(Point{X: origin.X + xc, Y: origin.Y - yc}, charecter)
	screen.setPix(Point{X: origin.X - xc, Y: origin.Y + yc}, charecter)
	screen.setPix(Point{X: origin.X - xc, Y: origin.Y - yc}, charecter)
	screen.setPix(Point{X: origin.X + yc, Y: origin.Y + xc}, charecter)
	screen.setPix(Point{X: origin.X + yc, Y: origin.Y - xc}, charecter)
	screen.setPix(Point{X: origin.X - yc, Y: origin.Y + xc}, charecter)
	screen.setPix(Point{X: origin.X - yc, Y: origin.Y - xc}, charecter)
}

func (screen *Screen) drawCircle(circle CircleOptns) {
	logSameLevel(screen.showLogs, "Drawing circle with optns", circle)
	logLevel += 1

	x := uint(0)
	y := circle.radius
	d := 3 - 2*int(circle.radius)
	drawCirclePixel(screen, circle.center, x, y, circle.charecter)
	for y >= x {
		// for each pixel we will
		// draw all eight pixels

		x++

		// check for decision parameter
		// and correspondingly
		// update d, x, y
		if d > 0 {
			y--
			d = d + 4*(int(x)-int(y)) + 10
		} else {
			d = d + 4*int(x) + 6
		}
		for i := uint(0); i < circle.thickness; i++ {
			drawCirclePixel(screen, circle.center, x, y-i, circle.charecter)
		}
	}

	logLevel -= 1
}

func (screen *Screen) drawFace(face Face) {
	logSameLevel(screen.showLogs, "Drawing face")
	logLevel += 1

	// draw the borders
	for edge := 0; edge < len(face.edges); edge++ {
		screen.drawEdge(face.edges[edge])
	}

	// fill it in
	// TODO: add code to fill faces in

	logLevel -= 1
}

func (screen *Screen) draw3dShape(shape Shape3d) {
	logSameLevel(screen.showLogs, "Drawing shape")
	logLevel += 1

	for face := 0; face < len(shape.faces); face++ {
		screen.drawFace(shape.faces[face])
	}
}

func (screen *Screen) printContents() {
	if screen.showGuis {
		if screen.name != "" {
			fmt.Println(screen.name)
		}
		fmt.Print(string(screen.contents))
	}
	logSameLevel(screen.showLogs, "Screen printed")
}

func (screen *Screen) reset() {
	screen.contents = screen.originalContents
	logSameLevel(screen.showLogs, "Screen reset")
}

// === FUNCTIONS TO HANDLE 3D (take a 3d point and convert it to a 2d point that looks 3d) ===

func move(point Point, opts CameraOptions) Point {
	//fmt.Println("LOG: started transform X:", point.X, "Y: ", point.Y, "Z:", point.Z)
	point = Point{
		X:              uint(int(point.X) + opts.xTranslate),
		Y:              uint(int(point.Y) + opts.yTranslate),
		Z:              uint(int(point.Z) + opts.zTranslate),
		has3d:          point.has3d,
		hasBeenMoved:   true,
		hasBeenRotated: point.hasBeenRotated,
	}
	//fmt.Println("LOG: finished transform X:", point.X, "Y: ", point.Y, "Z:", point.Z)
	return point
}

func rotate(point Point, opts CameraOptions) Point {
	return Point{
		X: uint(math.Round(
			(float64(point.X)) +
				(float64(point.Y) * float64(opts.yAngle)) +
				(float64(point.Z) * float64(opts.zAngleX)),
		)),
		Y: uint(math.Round(
			(float64(point.Y)) +
				(float64(point.X) * float64(opts.xAngle)) +
				(float64(point.Z) * float64(opts.zAngleY)),
		)),
		Z:              point.Z,
		has3d:          point.has3d,
		hasBeenRotated: true,
		hasBeenMoved:   point.hasBeenMoved,
	}
}

func transformAndRotate(point Point, opts CameraOptions) Point {
	if !point.hasBeenMoved {
		point = move(point, opts)
	}
	if !point.hasBeenRotated && point.has3d {
		point = rotate(point, opts)
	}
	return point
}

// === MAIN ===

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	selection := uint(0)
	showLogs := false
	showGuis := true

mainLoop:
	for true {
		logLevel = 0 // reset log level
		printTook := true
		printQuit := true
		selection = selectTui(
			"ﳑ GO ENGINE - select an option",
			[]string{
				"Run a speed test",
				"Show a 3d cube example",
				"Show a spining line",
				"Show a circle demo",
				"Show a graph demo",
				"Show a line to test performance",
				"Show some pixels to test code",
				toggleBtn("Show logs", showLogs),
				toggleBtn("Show guis", showGuis),
				"Quit",
			},
			selection,
		)

		clearScreen()

		start := time.Now()
		switch selection {
		case 0:
			speedTest(showGuis, showLogs)
		case 1:
			cube3d(showGuis, showLogs)
		case 2:
			spinningLine(showGuis, showLogs)
		case 3:
			circleDemo(showGuis, showLogs)
		case 4:
			graphDemo(showGuis, showLogs)
		case 5:
			lineSpeedTest(showGuis, showLogs)
		case 6:
			setFewPixels(showGuis, showLogs)
		case 7:
			showLogs = !showLogs
			printQuit, printTook = false, false
		case 8:
			showGuis = !showGuis
			printQuit, printTook = false, false
		case 9:
			break mainLoop
		}
		if printTook {
			fmt.Println("Took:", time.Since(start))
		}
		if printQuit {
			typeToQuit("Press any key to exit")
		}
	}

	fmt.Println("Bye bye 👋")
}
