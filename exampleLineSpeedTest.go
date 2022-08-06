package main

func lineSpeedTest () {
				screen := newScreen(32,32,"LINE SPEED TEST",nil,nil)

				for i:=uint(0);i<screen.width;i++ {
								screen.setPix(Piont{X:i,Y:i},"#")
								screen.printMe()
				}
}
