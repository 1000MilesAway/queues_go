package main

import (
	"fmt"
	"math"
	"github.com/1000MilesAway/queues_go/tree/main/tracker"
	geom "github.com/twpayne/go-geom"
	"gocv.io/x/gocv"
)

func MinX(person [][]float64) float64 {
	var min float64 = person[0][0]
	for _, xy := range person {
		if xy[0] < min {
			min = xy[0]
		}
	}
	return min
}

func MaxX(person [][]float64) float64 {
	var max float64 = person[0][0]
	for _, xy := range person {
		if xy[0] > max {
			max = xy[0]
		}
	}
	return max
}

func MinY(person [][]float64) float64 {
	var min float64 = person[0][1]
	for _, xy := range person {
		if xy[1] < min {
			min = xy[1]
		}
	}
	return min
}

func MaxY(person [][]float64) float64 {
	var max float64 = person[0][1]
	for _, xy := range person {
		if xy[1] > max {
			max = xy[1]
		}
	}
	return max
}

func PersonArea(person [][]float64) float64 {
	var coordinates = make([]geom.Coord, 0)
	for _, xy := range person {
		coordinates = append(coordinates, geom.Coord(xy))
	}
	fmt.Println(coordinates)
	kal := [][]geom.Coord{coordinates}
	personPolygon := geom.NewPolygon(geom.XY).MustSetCoords(kal)
	return math.Abs(personPolygon.Area())
}

func main() {
	webcam, _ := gocv.VideoCaptureFile("video.mp4")
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()
	kal := tracker.KalmanFilter{Score: 0.8}
	fmt.Println(kal)
	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
		h, w := img.Size()[0], img.Size()[1]
		persons := [][][]float64{{{50, 50}, {50, 100}, {100, 100}, {100, 50}}}
		var boxes = make([][4]int, 0)
		var scores = make([]float64, 0)

		for _, person := range persons {
			if len(person) <= 3 {
				continue
			}
			x1 := MinX(person)
			x2 := MaxX(person)
			y1 := MinY(person)
			y2 := MaxY(person)
			area := PersonArea(person)
			scores = append(scores, area / ((x2 - x1) * (y2 - y1)))
			boxes = append(boxes, [4]int{int(x1 * float64(w)), int(y1 * float64(h)), int(x2 * float64(w)), int(y2 * float64(h))})
		}

		fmt.Println(boxes)
		fmt.Println(scores)
	}
}
