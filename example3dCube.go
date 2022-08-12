package main

func cube3d (showGuis bool, showLogs bool) {
				screen := newScreen(24,24,"3D CUBE",nil,nil,showGuis,showLogs)

				points := []Point{
								{X: 0,  Y: 5,  Z: 0,  has3d: true},
								{X: 0,  Y: 5,  Z: 10, has3d: true},
								{X: 0,  Y: 15, Z: 0,  has3d: true},
								{X: 0,  Y: 15, Z: 10, has3d: true},
								{X: 10, Y: 5,  Z: 0,  has3d: true},
								{X: 10, Y: 5,  Z: 10, has3d: true},
								{X: 10, Y: 15, Z: 0,  has3d: true},
								{X: 10, Y: 15, Z: 10, has3d: true},
				}
				edges := []Edge{
								//edges for front of cube
								{
												pointA: points[0],
												pointB: points[4],
												edgeCharecter: "#",
								},
								{
												pointA: points[4],
												pointB: points[6],
												edgeCharecter: "#",
								},
								{
												pointA: points[6],
												pointB: points[2],
												edgeCharecter: "#",
								},
								{
												pointA: points[2],
												pointB: points[0],
												edgeCharecter: "#",
								},
								// edges for back of cube
								{
												pointA: points[1],
												pointB: points[5],
												edgeCharecter: "#",
								},
								{
												pointA: points[5],
												pointB: points[7],
												edgeCharecter: "#",
								},
								{
												pointA: points[7],
												pointB: points[3],
												edgeCharecter: "#",
								},
								{
												pointA: points[3],
												pointB: points[1],
												edgeCharecter: "#",
								},
								// edges for linking
								{
												pointA: points[0],
												pointB: points[1],
												edgeCharecter: "#",
								},
								{
												pointA: points[4],
												pointB: points[5],
												edgeCharecter: "#",
								},
								{
												pointA: points[6],
												pointB: points[7],
												edgeCharecter: "#",
								},
								{
												pointA: points[2],
												pointB: points[3],
												edgeCharecter: "#",
								},
				}
				faces := []Face{
								//front face	
								{
												edges: []Edge{
																edges[0],
																edges[1],
																edges[2],
																edges[3],
												},
								},
								// back face
								{
												edges: []Edge{
																edges[4],
																edges[5],
																edges[6],
																edges[7],
												},
								},
								// link them
								{
												edges: []Edge{
																edges[8],
																edges[9],
																edges[10],
																edges[11],
												},
								},
				}
				screen.draw3dShape(
								Shape3d{
												faces: faces,
								},
				)

				screen.printMe()
}
