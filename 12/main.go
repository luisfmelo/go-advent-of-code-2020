package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Direction string

const (
	North Direction = "north"
	East            = "east"
	South           = "south"
	West            = "west"
	Right           = "right"
	Left            = "left"
)

type DirectionValue struct {
	d Direction
	v int
}

type Object struct {
	x int
	y int
}


func (o *Object) Move(d Direction, value int) {
	switch d {
	case North:
		o.y += value
	case East:
		o.x += value
	case South:
		o.y -= value
	case West:
		o.x -= value
	}
}

func (o *Object) Turn(d Direction, degrees int) {
	turns := degrees / 90
	for turn := 0; turn < turns; turn++ {
		var newX, newY int
		// anti clockwise
		if d == Left {
			newX = o.y * -1 // x' = -y'
			newY = o.x      // y' = x'
		}
		if d == Right {
			newX = o.y      // x' = y'
			newY = o.x * -1 // y' = -x'
		}

		o.x = newX
		o.y = newY
	}
}

func (o *Object) GetDirectionValues() (DirectionValue, DirectionValue) {
	d := make([]DirectionValue, 2)

	if o.x > 0 {
		d[0] = DirectionValue{East, o.x}
	} else {
		d[0] = DirectionValue{West, -1 * o.x}
	}

	if o.y > 0 {
		d[1] = DirectionValue{North, o.y}
	} else {
		d[1] = DirectionValue{South, -1 * o.y}
	}

	return d[0], d[1]
}

type Ship struct {
	Object
	pointsTo Direction
}

func NewShip() *Ship {
	return &Ship{Object{0, 0}, East}
}

func (s *Ship) GoForward(value int) {
	s.Move(s.pointsTo, value)
}

func (s *Ship) Move(d Direction, value int) {
	switch d {
	case North:
		s.y += value
	case East:
		s.x += value
	case South:
		s.y -= value
	case West:
		s.x -= value
	}
}

func (s *Ship) Turn(d Direction, degrees int) {
	clockWise90degreeDirections := map[Direction]Direction{
		North: East,
		East:  South,
		South: West,
		West:  North,
	}

	antiClockWise90degreeDirections := map[Direction]Direction{
		East:  North,
		South: East,
		West:  South,
		North: West,
	}

	turns := degrees / 90
	for turn := 0; turn < turns; turn++ {
		if d == Left {
			s.pointsTo = antiClockWise90degreeDirections[s.pointsTo]
		} else {
			s.pointsTo = clockWise90degreeDirections[s.pointsTo]
		}
	}
}


func ManhattanDistance(o1, o2 Object) int {
	return int(math.Abs(float64(o1.x-o2.x)) + math.Abs(float64(o1.y-o2.y)))
}

func ManhattanDistanceOfShip(actionLines []string) int {
	ship := NewShip()

	for _, actionLine := range actionLines {
		action := rune(actionLine[0])
		value, err := strconv.Atoi(actionLine[1:])
		if err != nil {
			panic(err)
		}

		switch action {
		case 'F':
			ship.GoForward(value)
		case 'N':
			ship.Move(North, value)
		case 'S':
			ship.Move(South, value)
		case 'E':
			ship.Move(East, value)
		case 'W':
			ship.Move(West, value)
		case 'L':
			ship.Turn(Left, value)
		case 'R':
			ship.Turn(Right, value)
		}
	}

	return ManhattanDistance(ship.Object, Object{0,0})
}

func ManhattanDistanceOfShipWithWaypoint(actionLines []string, waypoint Object) int {
	ship := NewShip()

	for _, actionLine := range actionLines {
		action := rune(actionLine[0])
		value, err := strconv.Atoi(actionLine[1:])
		if err != nil {
			panic(err)
		}

		waypointDirectionValue1, waypointDirectionValue2 := waypoint.GetDirectionValues()
		switch action {
		case 'F':
			ship.Move(waypointDirectionValue1.d, value*waypointDirectionValue1.v)
			ship.Move(waypointDirectionValue2.d, value*waypointDirectionValue2.v)
		case 'N':
			waypoint.Move(North, value)
		case 'S':
			waypoint.Move(South, value)
		case 'E':
			waypoint.Move(East, value)
		case 'W':
			waypoint.Move(West, value)
		case 'L':
			waypoint.Turn(Left, value)
		case 'R':
			waypoint.Turn(Right, value)
		}
	}

	return ManhattanDistance(ship.Object, Object{0,0})
}

func main() {
	file, err := os.Open("12/input.txt")
	if err != nil {
		panic(err)
	}

	actions, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(ManhattanDistanceOfShip(actions))
	fmt.Println(ManhattanDistanceOfShipWithWaypoint(actions, Object{10, 1}))
}
