package main
import "fmt"
import "math"

type Piont struct {
  X uint
  Y uint
	Z uint
}

type Screen struct {
  width               uint
  height              uint
  contents            string
	headerContents      string
	originalContents    string
	perspectiveFunction func(Piont) Piont
}

func (s *Screen)printMe()  {
    fmt.Println(s.headerContents)
    fmt.Println(s.contents)
}

func (s *Screen)setPix()  {
    fmt.Println(s.headerContents)
    fmt.Println(s.contents)
}

// functions to handle 3d (take a 3d piont and convert it to a 2d piont that looks 3d)
func oblique3d(piont Piont) Piont {
				return Piont {
								X: piont.X + piont.Z,
								Y: piont.Y + piont.Z,
				}
}

func isometric3d(piont Piont) Piont {
				return Piont {
								X: piont.X + piont.Z,
								Y: piont.Y + piont.Z + uint(math.Round(float64(piont.X) * 0.2)),
				}
}

func pictorial3d(piont Piont) Piont {
				return isometric3d(piont)
}



func newScreen(width uint, height uint, perspectiveFunction func(Piont) Piont) Screen {
    if (perspectiveFunction == nil ){
        perspectiveFunction = pictorial3d
    }
    return Screen{
						width: width,
						height: height,
						contents: "someContents",
						headerContents: "someHeaderContents",
						originalContents: "some original contents",
						perspectiveFunction: perspectiveFunction,
		}
}

func main() {
				ourScreen := newScreen(32,32,nil)
	ourScreen.printMe()
}
