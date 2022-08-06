package main

import (
				"time"
)

func spinningLine () {
				screen := newScreen(32,32,"SPINNING LINE",nil,nil)
				
				for true {
								for y := uint(0); y < screen.height; y++ {
												screen.bresignham3D(
																Piont{X:0             ,Y:y              },
																Piont{X:screen.width-1,Y:screen.height-y-1},
																"*",
												)
												screen.printMe()
												screen.reset()
												time.Sleep(100 * time.Millisecond)
												clearScreen()
								}
								for x := uint(0); x < screen.height; x++ {
												screen.bresignham3D(
																Piont{X:x+1             ,Y:screen.height-1},
																Piont{X:screen.width-x-1,Y:0              },
																"*",
												)
												screen.printMe()
												time.Sleep(100 * time.Millisecond)
												clearScreen()
								}
				}
}
