package main

func graphDemo (showGuis bool, showLogs bool) {
				screen := newScreen(32,32,"GRAPH DEMO",nil,nil,showGuis,showLogs)

				screen.bresignham3D(Point{Y: 0,  X: 0 }, Point{Y:16, X:20}, "#")
				screen.bresignham3D(Point{Y: 17, X: 21}, Point{Y:25, X:24}, "*")
				screen.bresignham3D(Point{Y: 24, X: 25}, Point{Y:6, X:31}, "#")
				screen.printMe()
}
