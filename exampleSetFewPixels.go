package main

func setFewPixels (showGuis bool, showLogs bool) {
				screen := newScreen(11,11,"SET FEW PIXELS",nil,nil,showGuis,showLogs)

				screen.setPix(Piont{X: 0,Y: 0}, "#")
				screen.setPix(Piont{X: 0,Y: 10}, "#")
				screen.setPix(Piont{X: 0,Y: 5}, "#")
				screen.setPix(Piont{X: 10,Y: 5}, "#")
				screen.setPix(Piont{X: 10,Y: 0}, "#")
				screen.setPix(Piont{X: 10,Y: 10}, "#")


				screen.printMe()
}
