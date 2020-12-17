package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type CubeState rune

const (
	active   CubeState = '#'
	inactive           = '.'
)

type Boundaries struct {
	maxX, maxY, maxZ, maxW int
	minX, minY, minZ, minW int
}

type WorldState struct {
	turn                  int
	mapCoordStringToState map[string]CubeState
	boundaries            Boundaries
}

func (w *WorldState) addCoordinateValueToState(c Coord, s CubeState) {
	w.mapCoordStringToState[c.toString()] = s
}

func (w *WorldState) NumberOfActiveNeighbours(c Coord) int {
	activeNeighbours := 0
	for _, neighbour := range c.Neighbours() {
		if w.mapCoordStringToState[neighbour.toString()] == active {
			activeNeighbours++
		}
	}
	return activeNeighbours
}

func (w *WorldState) CountActiveCubes() int {
	activeCubes := 0
	for _, cubeState := range w.mapCoordStringToState {
		if cubeState == active {
			activeCubes++
		}
	}
	return activeCubes
}

func (w *WorldState) GetCubeStateOfCoord(c Coord) CubeState {
	if s, exists := w.mapCoordStringToState[c.toString()]; exists {
		return s
	}
	return inactive
}

func (w *WorldState) Print() {
	fmt.Print("Disclaimer: It will only consider W = 0 points\n\n")

	// get all coords by z
	var mapZToCoords = map[int][]Coord{}
	for coordStr := range w.mapCoordStringToState {
		coord := NewCordFromString(coordStr)
		if coord.w != 0 {
			continue
		}
		if _, exists := mapZToCoords[coord.z]; !exists {
			mapZToCoords[coord.z] = []Coord{}
		}
		mapZToCoords[coord.z] = append(mapZToCoords[coord.z], coord)
	}

	for z := w.boundaries.minZ; z <= w.boundaries.maxZ; z++ {
		fmt.Printf("z=%v\n", z)
		for y := w.boundaries.minY; y <= w.boundaries.maxY; y++ {
			for x := w.boundaries.minX; x <= w.boundaries.maxX; x++ {
				fmt.Printf("%s", string(w.mapCoordStringToState[fmt.Sprintf("%v, %v, %v", x, y, z)]))
			}
			fmt.Println()
		}
	}
}

func (w *WorldState) FixBoundaries() {
	// Fix X. Get every point in boundaries of x
	allXBoundariesInactive := true
	// Fix Y. Get every point in boundaries of y
	allYBoundariesInactive := true
	// Fix Z. Get every point in boundaries of z
	allZBoundariesInactive := true

	for coordStr, state := range w.mapCoordStringToState {
		c := NewCordFromString(coordStr)
		if (c.x == w.boundaries.maxX || c.x == w.boundaries.minX) && state == active {
			allXBoundariesInactive = false
		} else if (c.y == w.boundaries.maxY || c.y == w.boundaries.minY) && state == active {
			allYBoundariesInactive = false
		} else if (c.z == w.boundaries.maxZ || c.z == w.boundaries.minZ) && state == active {
			allZBoundariesInactive = false
		}
	}
	if allXBoundariesInactive {
		w.boundaries.minX++
		w.boundaries.maxX--
	}

	if allYBoundariesInactive {
		w.boundaries.minY++
		w.boundaries.maxY--
	}

	if allZBoundariesInactive {
		w.boundaries.minZ++
		w.boundaries.maxZ--
	}
}

type Coord struct {
	x, y, z, w int
}

func NewCordFromString(s string) Coord {
	splitter := strings.Split(s, ", ")
	return Coord{x: pkg.StrToInt(splitter[0]), y: pkg.StrToInt(splitter[1]), z: pkg.StrToInt(splitter[2]), w: pkg.StrToInt(splitter[3])}
}

func (c Coord) toString() string {
	return fmt.Sprintf("%v, %v, %v, %v", c.x, c.y, c.z, c.w)
}

func (c Coord) Neighbours() []Coord {
	var neighbours []Coord
	for ix := c.x - 1; ix <= c.x+1; ix++ {
		for iy := c.y - 1; iy <= c.y+1; iy++ {
			for iz := c.z - 1; iz <= c.z+1; iz++ {
				for iw := c.w - 1; iw <= c.w+1; iw++ {
					neighbourCoord := Coord{ix, iy, iz, iw}
					// self cannot be its neighbour
					if neighbourCoord.toString() == c.toString() {
						continue
					}
					neighbours = append(neighbours, neighbourCoord)
				}
			}
		}
	}

	return neighbours
}

func PrintCycles(worldStates []*WorldState) {
	for _, worldState := range worldStates {
		if worldState.turn == 0 {
			fmt.Println("Before any cycles:")
		} else {
			fmt.Printf("After %d cycle(s):\n", worldState.turn)
		}
		fmt.Println()
		worldState.Print()
	}
}

func CountActiveCubes(matrix [][]rune, lastTurn int, dimensions int, printCycles bool) int {
	var worldStates []*WorldState

	initialBoundaries := Boundaries{
		maxX: len(matrix[len(matrix)-1]),
		maxY: len(matrix),
		maxZ: 0,
		minX: 0,
		minY: 0,
		minZ: 0,
	}

	// Parse State of the world
	worldStates = append(worldStates, &WorldState{turn: 0, mapCoordStringToState: map[string]CubeState{}, boundaries: initialBoundaries})
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			worldStates[0].addCoordinateValueToState(Coord{x, y, 0, 0}, CubeState(matrix[y][x]))
		}
	}

	// Process world until lastTurn
	for turn := 0; turn < lastTurn; turn++ {
		// Get World State
		worldState := worldStates[turn]

		// Prepare new world state
		newWorldState := &WorldState{
			turn:                  turn + 1,
			mapCoordStringToState: map[string]CubeState{},
			boundaries: Boundaries{
				maxX: worldState.boundaries.maxX + 1,
				maxY: worldState.boundaries.maxY + 1,
				maxZ: worldState.boundaries.maxZ + 1,
				minX: worldState.boundaries.minX - 1,
				minY: worldState.boundaries.minY - 1,
				minZ: worldState.boundaries.minZ - 1,
			},
		}
		if dimensions == 4 {
			newWorldState.boundaries.minW = worldState.boundaries.minW - 1
			newWorldState.boundaries.maxW = worldState.boundaries.maxW + 1
		}

		// prepare new State
		for x := newWorldState.boundaries.minX; x <= newWorldState.boundaries.maxX; x++ {
			for y := newWorldState.boundaries.minY; y <= newWorldState.boundaries.maxY; y++ {
				for z := newWorldState.boundaries.minZ; z <= newWorldState.boundaries.maxZ; z++ {
					for w := newWorldState.boundaries.minW; w <= newWorldState.boundaries.maxW; w++ {
						coord := Coord{x, y, z, w}
						newWorldState.mapCoordStringToState[coord.toString()] = worldState.GetCubeStateOfCoord(coord)
						activeNeighbours := worldState.NumberOfActiveNeighbours(coord)
						switch worldState.GetCubeStateOfCoord(coord) {
						case active:
							// Cube is active? if 2 or 3 neighbors are also active -> remains active. Otherwise, becomes inactive.
							if !(activeNeighbours == 2 || activeNeighbours == 3) {
								newWorldState.mapCoordStringToState[coord.toString()] = inactive
							}
						case inactive:
							// Cube is inactive? 3 of its neighbors are active -> becomes active. Otherwise, remains inactive.
							if activeNeighbours == 3 {
								newWorldState.mapCoordStringToState[coord.toString()] = active
							}
						}
					}
				}
			}
		}
		worldState.FixBoundaries()

		// Assign New State
		worldStates = append(worldStates, newWorldState)
	}

	if printCycles {
		PrintCycles(worldStates)
	}

	return worldStates[len(worldStates)-1].CountActiveCubes()
}

func main() {
	file, err := os.Open("17/input.txt")
	if err != nil {
		panic(err)
	}

	matrix, err := pkg.ReadMatrix(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	start := time.Now()
	log.Printf("1st Part result: %v\n", CountActiveCubes(matrix, 6, 3, false))
	log.Printf("1st Part took: %s", time.Since(start))

	start = time.Now()
	log.Printf("2nd Part result: %v\n", CountActiveCubes(matrix, 6, 4, false))
	log.Printf("2nd Part took: %s", time.Since(start))
}
