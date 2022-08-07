package main

import (
	"fmt"
	"time"
)

func spinningLine (showGuis bool, showLogs bool) {
				screen := newScreen(32,32,"SPINNING LINE",nil,nil,showGuis,showLogs)
				
				for true {
								for y := uint(0); (y < screen.height); y++ {
												screen.bresignham3D(
																Piont{X:0             ,Y:y              },
																Piont{X:screen.width-1,Y:screen.height-y-1},
																"#",
												)
												screen.printMe()
												fmt.Println("Press ctrl+c to exit")
												screen.reset()
												time.Sleep(10 * time.Millisecond)

								}
								for x := uint(0); (x < screen.height); x++ {
												screen.bresignham3D(
																Piont{X:x+1             ,Y:screen.height-1},
																Piont{X:screen.width-x-1,Y:0              },
																"#",
												)
												screen.printMe()
												fmt.Println("Press ctrl+c to exit")
												screen.reset()
												time.Sleep(10 * time.Millisecond)
								}

				}
}
