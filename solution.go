package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Num int

const SidesTriangle Num = 3
const SidesSquare Num = 4
const SidesCircle Num = 0

func CalcSquare(sideLen float64, sidesNum Num) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3.0/4.0)
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0.0
	}
}
