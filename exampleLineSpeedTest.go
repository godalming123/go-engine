package main

func lineSpeedTest (showGuis bool, showLogs bool) {
				screen := newScreen(32,32,"LINE SPEED TEST",CameraOptions{},showGuis,showLogs)

				for i:=uint(0);i<screen.width;i++ {
								screen.setPix(Point{X:i,Y:i},"#")
								screen.printContents()
				}
}
