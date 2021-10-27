package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"

	"github.com/iamnotjustice/ds-examples/lists/waypoints"
)

func main() {
	//points := list.New()

	points := waypoints.New()

	points.PushBack(&waypoints.Point{
		Lat:  12.23,
		Long: 23.34,
	})

	points.PrintAll()
	//PrintAll(points)

	points.PushBack(&waypoints.Point{
		Lat:  34.45,
		Long: 45.56,
	})
	points.PushBack(&waypoints.Point{
		Lat:  134.45,
		Long: 145.56,
	})

	points.PrintAll()
	//PrintAll(points)

	// try inserting after this new point
	p1 := points.PushFront(&waypoints.Point{
		Lat:  4.15,
		Long: 5.26,
	})
	p2 := points.InsertAfter(&waypoints.Point{
		Lat:  9.45,
		Long: 9.56,
	}, p1)
	if p2 == nil {
		log.Fatal("could not insert after")
	}

	points.PrintAll()
	//PrintAll(points)

	// and then remove both
	points.Remove(p2)
	points.Remove(p1)

	points.PrintAll()
	//PrintAll(points)
}

func PrintAll(l *list.List) error {
	fmt.Printf("Current route is (%d point(s)):\n", l.Len())

	current := l.Front()
	if current == nil {
		return errors.New("WaypointList is empty")
	}

	position := 0

	for e := l.Front(); e != nil; e = e.Next() {
		position++
		fmt.Printf("Point #%d: %v \n", position, e.Value)
	}

	fmt.Printf("=========\n")

	return nil
}
