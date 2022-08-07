package main

func lineSpeedTest (showGuis bool, showLogs bool) {
				screen := newScreen(32,32,"LINE SPEED TEST",nil,nil,showGuis,showLogs)

				for i:=uint(0);i<screen.width;i++ {
								screen.setPix(Piont{X:i,Y:i},"#")
								screen.printMe()
				}
}
