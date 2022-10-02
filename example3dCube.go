package main

func cube3d(showGuis bool, showLogs bool) {
	screen := newScreen(30, 21, "3D CUBE", CameraOptions{}, showGuis, showLogs)

	x1 := uint(0)
	x2 := uint(10)
	y1 := uint(0)
	y2 := uint(10)
	z1 := uint(4)
	z2 := uint(14)

	points := []Point{
		{X: x1, Y: y1, Z: z1, has3d: true},
		{X: x1, Y: y1, Z: z2, has3d: true},
		{X: x1, Y: y2, Z: z1, has3d: true},
		{X: x1, Y: y2, Z: z2, has3d: true},
		{X: x2, Y: y1, Z: z1, has3d: true},
		{X: x2, Y: y1, Z: z2, has3d: true},
		{X: x2, Y: y2, Z: z1, has3d: true},
		{X: x2, Y: y2, Z: z2, has3d: true},
	}
	edges := []Edge{
		//edges for front of cube
		{
			pointA:        points[0],
			pointB:        points[4],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[4],
			pointB:        points[6],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[6],
			pointB:        points[2],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[2],
			pointB:        points[0],
			edgeCharecter: pixelCharecter,
		},
		// edges for back of cube
		{
			pointA:        points[1],
			pointB:        points[5],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[5],
			pointB:        points[7],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[7],
			pointB:        points[3],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[3],
			pointB:        points[1],
			edgeCharecter: pixelCharecter,
		},
		// edges for linking
		{
			pointA:        points[0],
			pointB:        points[1],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[4],
			pointB:        points[5],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[6],
			pointB:        points[7],
			edgeCharecter: pixelCharecter,
		},
		{
			pointA:        points[2],
			pointB:        points[3],
			edgeCharecter: pixelCharecter,
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

	screen.printContents()
}
