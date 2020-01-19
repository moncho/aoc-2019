package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

const (
	radiansToDegrees = 180 / math.Pi
)

type angle float64

// degrees returns the angle in degrees between 0-360.
// where 0 is the angle to any point to the left with the same y.
// Rotation is counter clockwise.
func (a angle) degrees() float64 {
	return 180 - (radiansToDegrees * a.radians())
}
func (a angle) radians() float64 { return float64(a) }

type xy struct {
	x int
	y int
}

func (xy xy) sortByAngleTo(others []xy) {
	sort.Slice(others, func(x, y int) bool {
		return xy.angleTo(others[x]) < xy.angleTo(others[y])
	})
}

func (xy xy) angleTo(other xy) angle {
	y := other.y - xy.y
	x := other.x - xy.x
	return angle(math.Atan2(float64(y), float64(x)))
}

func (xy xy) distanceTo(other xy) int {
	return abs(xy.y-other.y) + abs(xy.x-other.x)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	asteroids := asteroids(file)

	max, count := asteroidWithMoreVisibility(asteroids)

	fmt.Printf("How many other asteroids can be detected from that location? %d count at %v \n", count, max)

	angles, grouped := groupAndSort(max, asteroids)

	var destroyed []xy
	cont := true
	for cont {
		cont = false
		for _, angleTo := range angles {
			if asteroids := grouped[angleTo]; len(asteroids) > 0 {
				destroyed = append(destroyed, asteroids[0])
				grouped[angleTo] = asteroids[1:]
				cont = true
			}
		}
	}

	fmt.Printf("What do you get if you multiply the X coordinate by 100 and then add its Y coordinate of 200th asteroid to be vaporized? %d\n",
		destroyed[199].x*100+destroyed[199].y)
}

func asteroids(r io.Reader) []xy {
	scanner := bufio.NewScanner(r)

	var m []xy
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, s := range strings.Split(line, "") {
			if s == "#" {
				m = append(m, xy{x, y})
			}
		}
		y++
	}

	return m
}

// groups the given asteroids by their angle to the given pos
// and sorts those in the same angle by their distance to the pos,
// the first element being the closest asteroid.
func groupAndSort(a xy, asteroids []xy) ([]float64, map[float64][]xy) {

	groupBy := make(map[float64][]xy)
	var angles []float64

	for _, asteroid := range asteroids {
		if asteroid == a {
			continue
		}
		angle := a.angleTo(asteroid).degrees()

		if g := groupBy[angle]; g != nil {
			groupBy[angle] = append(groupBy[angle], asteroid)
		} else {
			groupBy[angle] = []xy{asteroid}
			angles = append(angles, angle)
		}
	}

	// Sort backwards
	sort.Slice(angles, func(i, j int) bool {
		return angles[i] > angles[j]
	})
	// Take the index of the 270 angle and start the slice of angles
	// in that pos
	norhtIndex := sort.Search(len(angles), func(n int) bool {
		return angles[n] <= 270
	})

	angles = append(angles[norhtIndex:], angles[:norhtIndex]...)

	//Sort asteroids by distance to a
	for angle, asteroids := range groupBy {
		sort.Slice(asteroids, func(i, j int) bool {
			return a.distanceTo(asteroids[i]) < a.distanceTo(asteroids[j])
		})
		groupBy[angle] = asteroids
	}

	return angles, groupBy
}

func asteroidWithMoreVisibility(asteroids []xy) (xy, int) {
	abl := asteroidsByLocation(asteroids)

	var max xy
	var maxCount int
	for xy, lines := range abl {
		if len(lines) > maxCount {
			maxCount = len(lines)
			max = xy
		}
	}

	return max, maxCount
}

func asteroidsByLocation(asteroids []xy) map[xy][]float64 {
	res := make(map[xy][]float64)

	for _, a := range asteroids {
		for _, b := range asteroids {
			if a == b {
				continue
			}
			angleTo := float64(a.angleTo(b))
			if containsf(res[a], angleTo) {
				continue
			}
			res[a] = append(res[a], angleTo)
		}
	}

	return res
}

func containsf(ff []float64, f float64) bool {
	for _, e := range ff {
		if e == f {
			return true
		}
	}
	return false
}

func abs(f int) int {
	if f < 0 {
		return -f
	}
	return f
}
