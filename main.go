package main

import (
	"fmt"
	"math"
	"strings"
	"time"
	"github.com/jwalton/gchalk"
  "os/exec"
	"os"
)

// === STRUCTS ===

type Piont struct {
  X     uint
  Y     uint
	Z     uint
	has3d bool
}

type Screen struct {
  width                   uint
  height                  uint
  contents                string
	originalContents        string
	perspectiveFunction     func(Piont, map[string]int) Piont
	perspectiveFunctionOpts map[string]int
	showLogs                bool
	showGuis                bool
	name                    string
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
				cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
}

func typeToQuit(message string) {
				fmt.Println(message)
				var i []byte = make([]byte, 1)
				os.Stdin.Read(i)
}

func drawList(header string, subtitle string, items []string, selectedIndex uint) {
				if header != "" {
								fmt.Println(gchalk.Bold(header))
				}

				if subtitle != "" {
								fmt.Println(subtitle)
				}

				for i:=0;i<len(items);i++ {
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
				drawList(header,"  w - up, s -down, d - select",options,selectedIndex)
				for true {
								os.Stdin.Read(i)
								input := string(i)
								if (input == "w" || input == "W") && selectedIndex > 0 {
												selectedIndex -= 1
								} else if (input == "s" || input == "S") && selectedIndex < uint(len(options) - 1) {
												selectedIndex += 1												
								} else if (input == "d" || input == "D") {
												return selectedIndex
								}
								clearScreen()
								drawList(header,"  w - up, s -down, d - select",options,selectedIndex)
				}
				return 0
}

// === FUNCTIONS FOR SCREEN STRUCT ===

func (screen *Screen)bresignham3D(a Piont, b Piont, newChar string) {
				dx := absInt(int(b.X - a.X))
				dy := absInt(int(b.Y - a.Y))
				dz := absInt(int(b.Z - a.Z))
				xs := copysign(1,int(b.X - a.X))
				ys := copysign(1,int(b.Y - a.Y))
				zs := copysign(1,int(b.Z - a.Z))

				if screen.showLogs {
								fmt.Println("LOG: STARTED: drawing bresignham line from X:", a.X, "Y:", a.Y, "Z:", a.Z, "to X:", b.X,"Y:", b.Y, "Z:", b.Z)
								fmt.Println("LOG: deferences: X:", dx, "Y:", dy, "Z", dz)
				}

				screen.setPix(a,newChar)
 
        if (dx >= dy) && (dx >= dz){// Driving axis is X-axis"
								if screen.showLogs {
												fmt.Println("LOG: driving axis for line is X axis")
								}

								p1 := 2 * dy - dx
								p2 := 2 * dz - dx
								for (a.X != b.X){
												a.X = uint(int(a.X) + xs)
												if (p1 >= 0){
																a.Y = uint(int(a.Y) + ys)
																p1 -= 2 * dx
												}
												if (p2 >= 0){
																a.Z = uint(int(a.Z) + zs)
																p2 -= 2 * dx
												}
												p1 += 2 * dy
												p2 += 2 * dz
												screen.setPix(a,newChar)
								}
				} else if (dy >= dx) && (dy >= dz) {// Driving axis is Y axis
								if screen.showLogs {
												fmt.Println("LOG: driving axis for line is Y axis")
								}

								p1 := 2 * dx - dy
								p2 := 2 * dz - dy
								for (a.Y != b.Y) {
												a.Y = uint(int(a.Y) + ys)
												if (p1 >= 0) {
																a.X = uint(int(a.X) + xs)
																p1 -= 2 * dy
												}
												if (p2 >= 0){
																a.Z = uint(int(a.Z) + zs)
																p2 -= 2 * dy
												}
												p1 += 2 * dx
												p2 += 2 * dz
												screen.setPix(a,newChar)
								}
				} else if (dz >= dx) && (dz >= dy) {// Driving axis is Z-axis" 
								if screen.showLogs {
												fmt.Println("LOG: driving axis for line is Z axis")
								}

								p1 := 2 * dy - dz
								p2 := 2 * dx - dz
								for (a.Z != b.Z){
												a.Z = uint(int(a.Z) + zs)
												if (p1 >= 0){
																a.Y += uint(ys)
																p1 -= 2 * dz
												}
												if (p2 >= 0){
																a.X = uint(int(a.X) + xs)
																p2 -= 2 * dz
												}
												p1 += 2 * dy
												p2 += 2 * dx
												screen.setPix(a,newChar)
								}
				} else {
								fmt.Println("Could not find the driving axis")
				}
				if screen.showLogs {
								fmt.Println("LOG: FINISHED: drawing bresignham line ===============================")
				}
}

func newScreen(width uint, height uint, name string, perspectiveFunction func(Piont, map[string]int) Piont, perspectiveFunctionOpts map[string]int, showGuis bool, showLogs bool) Screen {
				if perspectiveFunction == nil { // if the perspective function is undefiened
								perspectiveFunction = pictorial3d // then set it to something
				}
				contents := ""
				for i:=uint(0);i < height;i++ {//for each row
								contents += strings.Repeat("  ", int(width)) + " \n"
				}
				if showLogs {
								fmt.Println("LOG: screen created with", height, "pixels in height and", width, "pixels in width.")
				}

				return Screen{
								width: width,
								height: height,
								contents: contents,
								originalContents: contents,
								perspectiveFunction: perspectiveFunction,
								perspectiveFunctionOpts: perspectiveFunctionOpts,
								showLogs: showLogs,
								showGuis: showGuis,
								name: name,
				}
}

func (s *Screen)setPix(piont Piont, newChar string)  {
  if piont.has3d { // if the piont is 3D
    piont = s.perspectiveFunction(piont, s.perspectiveFunctionOpts) // then make it 2D
	}
	if piont.X < s.width && piont.Y < s.height {
				charToSet := (piont.Y * (s.width + 1)) + piont.X
				charToSet *= 2
				s.contents = s.contents[:charToSet] + strings.Repeat(newChar, 2) + s.contents[charToSet+2:]
				if s.showLogs {
								fmt.Println("LOG: pixel created at Y:", piont.Y, "X:", piont.X, "and the charecter will be:", newChar, "and the charecter to set is", charToSet)
				}
	} else if s.showLogs{
					fmt.Println("LOG: It doesnt appear that the pixel is within the screens bounds")
  }
}

func (screen *Screen)printMe() {
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

func (screen *Screen)reset() {
				screen.contents = screen.originalContents
				if screen.showLogs {
								fmt.Println("LOG: screen reset")
				}
}

// === FUNCTIONS TO HANDLE 3D (take a 3d piont and convert it to a 2d piont that looks 3d) ===

func oblique3d(piont Piont, opts map[string]int) Piont {
				return Piont {
								X: piont.X + piont.Z,
								Y: piont.Y + piont.Z,
				}
}

func isometric3d(piont Piont, opts map[string]int) Piont {
				return Piont {
								X: piont.X + piont.Z,
								Y: piont.Y + piont.Z + uint(math.Round(float64(piont.X) * 0.2)),
				}
}

func pictorial3d(piont Piont, opts map[string]int) Piont {
				// TODO: implement a function for pictorioal 3D
				return isometric3d(piont, opts)
}

func main() {
				// disable input buffering
				exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
				// do not display entered characters on the screen
				exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

				selection := uint(0)
				showLogs := "OFF"
				showGuis := "ON"

				showLogsBool := false
				showGuisBool := true


				mainLoop: for true {
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
																"Show logs " + showLogs,
																"Show guis " + showGuis,
																"Quit",
												},
												selection,
								)

								clearScreen()
								
								start := time.Now()
								switch selection {
												case 0:
																speedTest(showGuisBool, showLogsBool)
												case 1:
																cube3d(showGuisBool,showLogsBool)
												case 2:
																spinningLine(showGuisBool,showLogsBool)
												case 3:
																graphDemo(showGuisBool,showLogsBool)
												case 4:
																lineSpeedTest(showGuisBool,showLogsBool)
												case 5:
																setFewPixels(showGuisBool,showLogsBool)
												case 6:
																if showLogs == "ON" {
																				showLogs = "OFF"
																				showLogsBool = false
																} else {
																				showLogs = "ON"
																				showLogsBool = true
																}
																printQuit, printTook = false, false
												case 7:
																if showGuis == "ON" {
																				showGuis = "OFF"
																				showGuisBool = false
																} else {
																				showGuis = "ON"
																				showGuisBool = true
																}
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
