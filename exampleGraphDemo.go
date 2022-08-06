package main

func graphDemo () {
				screen := newScreen(32,32,"GRAPH DEMO",nil,nil)

				screen.bresignham3D(Piont{Y: 0,  X: 0 }, Piont{Y:16, X:20}, "#")
				screen.bresignham3D(Piont{Y: 17, X: 21}, Piont{Y:25, X:24}, "*")
				screen.bresignham3D(Piont{Y: 25, X: 23}, Piont{Y:6, X:31}, "#")
				screen.printMe()
}
