package main

import (
	"fmt"
	"time"
)

func onDone(screen Screen) {
				screen.printMe()
				fmt.Println("Press ctrl+c to exit")
				time.Sleep(time.Millisecond * 10)
}
	
func spinningLine (showGuis bool, showLogs bool) {
				screen := newScreen(32,32,"SPINNING LINE",nil,nil,showGuis,showLogs)
			
				for true {
								for y := uint(0); (y < screen.height); y++ {
												screen.bresignham3D(
																Point{X:0             ,Y:y              },
																Point{X:screen.width-1,Y:screen.height-y-1},
																"#",
												)
												onDone(screen)
												screen.reset()//reseting is not in the ondone function since it appears that that function cannot modify the screen
								}
								for x := uint(0); (x < screen.height); x++ {
												screen.bresignham3D(
																Point{X:x+1             ,Y:screen.height-1},
																Point{X:screen.width-x-1,Y:0              },
																"#",
												)
												onDone(screen)
												screen.reset()//reseting is not in the ondone function since it appears that that function cannot modify the screen
								}

				}
}
