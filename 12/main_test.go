package main

import (
	"advent-of.code.2020/pkg"
	"reflect"
	"strings"
	"testing"
)

//
// type TestCase struct {
// 	input              string
// 	getNewSeatStateFun NewSeatStateRule
// 	expectedOutput     int
// }

func TestManhattanDistance(t *testing.T) {
	var object Object
	var actualResult, expectedResult int

	object = Object{1, 1}
	expectedResult = 2
	actualResult = ManhattanDistance(object, Object{0, 0})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	object = Object{-1, -1}
	expectedResult = 2
	actualResult = ManhattanDistance(object, Object{0, 0})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	object = Object{10, -1}
	expectedResult = 11
	actualResult = ManhattanDistance(object, Object{0, 0})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestObject_GetDirectionValues(t *testing.T) {
	var object Object
	var actualResult1, actualResult2, expectedResult1, expectedResult2 DirectionValue

	object = Object{1, 1}
	expectedResult1 = DirectionValue{d: North, v: 1}
	expectedResult2 = DirectionValue{d: East, v: 1}
	actualResult1, actualResult2 = object.GetDirectionValues()
	if reflect.DeepEqual(actualResult1, expectedResult1) || reflect.DeepEqual(actualResult2, expectedResult2) {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResult1, actualResult2, expectedResult1, expectedResult2)
	}

	object = Object{0, 0}
	expectedResult1 = DirectionValue{d: North, v: 0}
	expectedResult2 = DirectionValue{d: East, v: 0}
	actualResult1, actualResult2 = object.GetDirectionValues()
	if reflect.DeepEqual(actualResult1, expectedResult1) || reflect.DeepEqual(actualResult2, expectedResult2) {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResult1, actualResult2, expectedResult1, expectedResult2)
	}

	object = Object{1, -1}
	expectedResult1 = DirectionValue{d: South, v: 1}
	expectedResult2 = DirectionValue{d: East, v: 1}
	actualResult1, actualResult2 = object.GetDirectionValues()
	if reflect.DeepEqual(actualResult1, expectedResult1) || reflect.DeepEqual(actualResult2, expectedResult2) {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResult1, actualResult2, expectedResult1, expectedResult2)
	}

	object = Object{-1, -1}
	expectedResult1 = DirectionValue{d: South, v: 1}
	expectedResult2 = DirectionValue{d: West, v: 1}
	actualResult1, actualResult2 = object.GetDirectionValues()
	if reflect.DeepEqual(actualResult1, expectedResult1) || reflect.DeepEqual(actualResult2, expectedResult2) {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResult1, actualResult2, expectedResult1, expectedResult2)
	}
}

func TestObject_Move(t *testing.T) {
	var object Object
	var actualResultX, actualResultY, expectedResultX, expectedResultY int

	object = Object{1, 1}
	object.Move(North, 5)
	expectedResultX = 1
	expectedResultY = 6
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{1, 1}
	object.Move(South, 5)
	expectedResultX = 1
	expectedResultY = -4
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{1, 1}
	object.Move(East, 5)
	expectedResultX = 6
	expectedResultY = 1
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{1, 1}
	object.Move(West, 5)
	expectedResultX = -4
	expectedResultY = 1
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}
}

func TestObject_Turn(t *testing.T) {
	var object Object
	var actualResultX, actualResultY, expectedResultX, expectedResultY int

	object = Object{5, 1}
	object.Turn(Left, 90)
	expectedResultX = -1
	expectedResultY = 5
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{5, 1}
	object.Turn(Left, 180)
	expectedResultX = -5
	expectedResultY = -1
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{5, 1}
	object.Turn(Right, 90)
	expectedResultX = 1
	expectedResultY = -5
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{5, 1}
	object.Turn(Right, 180)
	expectedResultX = -5
	expectedResultY = -1
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	object = Object{5, 1}
	object.Turn(Right, 360)
	expectedResultX = 5
	expectedResultY = 1
	actualResultX = object.x
	actualResultY = object.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}
}

func TestShip_Move(t *testing.T) {
	var ship Ship
	var actualResultX, actualResultY, expectedResultX, expectedResultY int

	ship = Ship{Object{1, 1}, East}
	ship.Move(North, 5)
	expectedResultX = 1
	expectedResultY = 6
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Move(South, 5)
	expectedResultX = 1
	expectedResultY = -4
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Move(East, 5)
	expectedResultX = 6
	expectedResultY = 1
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Move(West, 5)
	expectedResultX = -4
	expectedResultY = 1
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}
}

func TestShip_Turn(t *testing.T) {
	var ship Ship
	var actualResult, expectedResult Direction

	ship = Ship{Object{1, 1}, East}
	ship.Turn(Left, 90)
	expectedResult = North
	actualResult = ship.pointsTo
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Turn(Left, 180)
	expectedResult = West
	actualResult = ship.pointsTo
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Turn(Right, 90)
	expectedResult = South
	actualResult = ship.pointsTo
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Turn(Right, 180)
	expectedResult = West
	actualResult = ship.pointsTo
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	ship = Ship{Object{1, 1}, East}
	ship.Turn(Right, 360)
	expectedResult = East
	actualResult = ship.pointsTo
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestShip_GoForward(t *testing.T) {
	var ship Ship
	var actualResultX, actualResultY, expectedResultX, expectedResultY int

	ship = Ship{Object{1, 1}, East}
	ship.GoForward(5)
	expectedResultX = 6
	expectedResultY = 1
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	ship = Ship{Object{1, 1}, West}
	ship.GoForward(5)
	expectedResultX = -4
	expectedResultY = 1
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	ship = Ship{Object{1, 1}, North}
	ship.GoForward(5)
	expectedResultX = 1
	expectedResultY = 6
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}

	ship = Ship{Object{1, 1}, South}
	ship.GoForward(5)
	expectedResultX = 1
	expectedResultY = -4
	actualResultX = ship.x
	actualResultY = ship.y
	if actualResultX != expectedResultX || actualResultY != expectedResultY {
		t.Errorf("Got (%v, %v); Want (%v, %v)", actualResultX, actualResultY, expectedResultX, expectedResultY)
	}
}

func TestManhattanDistanceOfShip(t *testing.T) {
	input := "F10\nN3\nF7\nR90\nF11"
	expectedOutput := 25

	lines, err := pkg.ReadLines(strings.NewReader(input))
	if err != nil {
		panic(err)
	}

	actualResult := ManhattanDistanceOfShip(lines)
	if actualResult != expectedOutput {
		t.Errorf("Got %v; Want %v", actualResult, expectedOutput)
	}
}

func TestManhattanDistanceOfShipWithWaypoint(t *testing.T) {
	input := "F10\nN3\nF7\nR90\nF11"
	expectedOutput := 286

	lines, err := pkg.ReadLines(strings.NewReader(input))
	if err != nil {
		panic(err)
	}

	actualResult := ManhattanDistanceOfShipWithWaypoint(lines, Object{10, 1})
	if actualResult != expectedOutput {
		t.Errorf("Got %v; Want %v", actualResult, expectedOutput)
	}
}
