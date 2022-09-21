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
	edgeCharecter string
}

type Face struct {
	edges         []Edge
	fillCharecter string
}

type Shape3d struct {
	faces []Face
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
	contents         string
	originalContents string
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
	return 0
}

// === FUNCTIONS FOR SCREEN STRUCT ===

func (screen *Screen) bresignham3D(a Point, b Point, newChar string) {
	a = transformAndRotate(a, screen.cameraOpts)
	b = transformAndRotate(b, screen.cameraOpts)

	dx := absInt(int(b.X - a.X))
	dy := absInt(int(b.Y - a.Y))
	dz := absInt(int(b.Z - a.Z))
	xs := copysign(1, int(b.X-a.X))
	ys := copysign(1, int(b.Y-a.Y))
	zs := copysign(1, int(b.Z-a.Z))

	if screen.showLogs {
		fmt.Println("LOG: STARTED: drawing bresignham line from X:", a.X, "Y:", a.Y, "Z:", a.Z, "to X:", b.X, "Y:", b.Y, "Z:", b.Z)
		fmt.Println("LOG: deferences: X:", dx, "Y:", dy, "Z", dz)
	}

	screen.setPix(a, newChar)

	if (dx >= dy) && (dx >= dz) { // Driving axis is X-axis"
		if screen.showLogs {
			fmt.Println("LOG: driving axis for line is X axis")
		}

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
		if screen.showLogs {
			fmt.Println("LOG: driving axis for line is Y axis")
		}

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
		if screen.showLogs {
			fmt.Println("LOG: driving axis for line is Z axis")
		}

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
	if screen.showLogs {
		fmt.Println("LOG: FINISHED: drawing bresignham line ===============================")
	}
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
	contents := ""
	for i := uint(0); i < height; i++ { //for each row
		contents += strings.Repeat("  ", int(width)) + " \n" // add the blankspace
	}
	if showLogs {
		fmt.Println("LOG: screen created with", height, "pixels in height and", width, "pixels in width.")
	}

	return Screen{
		width:            width,
		height:           height,
		contents:         contents,
		originalContents: contents,
		cameraOpts:       cameraOptions,
		showLogs:         showLogs,
		showGuis:         showGuis,
		name:             name,
	}
}

func (screen *Screen) setPix(point Point, newChar string) {
	point = transformAndRotate(point, screen.cameraOpts)
	if point.X < screen.width && point.Y < screen.height { // if the point fits in the screen
		charToSet := (point.Y * (screen.width + 1)) + point.X
		charToSet *= 2
		screen.contents = screen.contents[:charToSet] + strings.Repeat(newChar, 2) + screen.contents[charToSet+2:]
		if screen.showLogs {
			fmt.Println("LOG: pixel created at Y:", point.Y, "X:", point.X, "and the charecter will be:", newChar, "and the charecter to set is", charToSet)
		}
	} else if screen.showLogs {
		fmt.Println("LOG: It doesnt appear that the pixel is within the screens bounds")
	}
}

func (screen *Screen) drawEdge(edge Edge) {
	screen.bresignham3D(edge.pointA, edge.pointB, edge.edgeCharecter)
}

func (screen *Screen) drawFace(face Face) {
	if screen.showLogs {
		fmt.Println("LOG: started to draw face")
	}

	// draw the borders
	for edge := 0; edge < len(face.edges); edge++ {
		screen.drawEdge(face.edges[edge])
	}

	// fill it in
	// TODO: add code to fill faces in
}

func (screen *Screen) draw3dShape(shape Shape3d) {
	if screen.showLogs {
		fmt.Println("LOG: started to draw shape")
	}
	for face := 0; face < len(shape.faces); face++ {
		screen.drawFace(shape.faces[face])
	}
}

func (screen *Screen) printContents() {
	if screen.showGuis {
		if screen.name != "" {
			fmt.Println(screen.name)
		}
		fmt.Print(screen.contents)
	}
	if screen.showLogs {
		fmt.Println("LOG: screen printed")
	}
}

func (screen *Screen) reset() {
	screen.contents = screen.originalContents
	if screen.showLogs {
		fmt.Println("LOG: screen reset")
	}
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
		printTook := true
		printQuit := true
		selection = selectTui(
			"ﳑ GO ENGINE - select an option",
			[]string{
				"Run a speed test",
				"Show a 3d cube example",
				"Show a spining line",
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
			graphDemo(showGuis, showLogs)
		case 4:
			lineSpeedTest(showGuis, showLogs)
		case 5:
			setFewPixels(showGuis, showLogs)
		case 6:
			showLogs = !showLogs
			printQuit, printTook = false, false
		case 7:
			showGuis = !showGuis
			printQuit, printTook = false, false
		case 8:
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
