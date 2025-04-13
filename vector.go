package vec

import "math"

// Vector3 describes the 3D offset from an origin point
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func FromArray(array []float64) Vector3 {
	return Vector3{X: array[0], Y: array[1], Z: array[2]}
}

func Zero() Vector3 {
	return Vector3{X: 0, Y: 0, Z: 0}
}

// Cpy creates a new Vector3 by taking a copy of v
func Cpy(v Vector3) Vector3 {
	return Vector3{X: v.X, Y: v.Y, Z: v.Z}
}

// Mag returns the magnitude of v
func Mag(v Vector3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Nrm creates a normalised vec from v
func Nrm(v Vector3) Vector3 {
	return Div(v, Mag(v))
}

func Inv(v Vector3) Vector3 {
	return Vector3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func Mul(v Vector3, n float64) Vector3 {
	return Vector3{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func Div(v Vector3, n float64) Vector3 {
	return Vector3{
		X: v.X / n,
		Y: v.Y / n,
		Z: v.Z / n,
	}
}

func Add(a Vector3, b Vector3) Vector3 {
	return Vector3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Sub(a Vector3, b Vector3) Vector3 {
	return Vector3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Dst(a Vector3, b Vector3) float64 {
	return Mag(Sub(a, b))
}
