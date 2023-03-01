package engine

import (
	"fmt"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"math"
)

const Epsilon = 0.00001

type Vector struct {
	X, Y float64
}

func Dot(ihs, rhs Vector) float64 {
	return ihs.X*rhs.X + ihs.Y*rhs.Y
}

func Lerp(a, b Vector, t float64) Vector {
	return NewVector(
		a.X+(b.X-a.X)*t,
		a.Y+(b.Y-a.Y)*t,
	)
}

func Distance(a, b Vector) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func Reflect(ihs, rhs Vector) Vector {
	factor := -2.0 * Dot(ihs, rhs)
	return NewVector(
		factor*ihs.X+rhs.X,
		factor*ihs.Y+rhs.Y,
	)
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Copy() Vector {
	return NewVector(v.X, v.Y)
}

func RandomVectorOnUnitCircle() Vector {
	randomRad := random.FloatBetween(0, 360) * math.Pi

	return NewVector(
		math.Cos(randomRad),
		math.Sin(randomRad),
	).Normalize()
}

func (v Vector) Set(x, y float64) Vector {
	v.X = x
	v.Y = y
	return v
}

func (v Vector) Add(other Vector) Vector {
	return NewVector(v.X+other.X, v.Y+other.Y)
}

func (v Vector) AddScalar(scalar float64) Vector {
	return NewVector(v.X+scalar, v.Y+scalar)
}

func (v Vector) AddScalars(x, y float64) Vector {
	return NewVector(v.X+x, v.Y+y)
}

func (v Vector) Sub(other Vector) Vector {
	return NewVector(v.X-other.X, v.Y-other.Y)
}

func (v Vector) SubScalar(scalar float64) Vector {
	return NewVector(v.X-scalar, v.Y-scalar)
}

func (v Vector) SubScalars(x, y float64) Vector {
	return NewVector(v.X-x, v.Y-y)
}

func (v Vector) Mul(other Vector) Vector {
	return NewVector(v.X*other.X, v.Y*other.Y)
}

func (v Vector) MulScalar(scalar float64) Vector {
	return NewVector(v.X*scalar, v.Y*scalar)
}

func (v Vector) MulScalars(x, y float64) Vector {
	return NewVector(v.X*x, v.Y*y)
}

func (v Vector) Div(other Vector) Vector {
	return NewVector(v.X/other.X, v.Y/other.Y)
}

func (v Vector) DivScalar(scalar float64) Vector {
	return NewVector(v.X/scalar, v.Y/scalar)
}

func (v Vector) DivScalars(x, y float64) Vector {
	return NewVector(v.X/x, v.Y/y)
}

func (v Vector) Distance(other Vector) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector) Lerp(other Vector, t float64) Vector {
	return NewVector(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
	)
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Normalize() Vector {
	m := v.Magnitude()

	if m > Epsilon {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v Vector) Reflect(other Vector) Vector {
	factor := -2.0 * v.Dot(other)
	return NewVector(
		factor*v.X+other.X,
		factor*v.Y+other.Y,
	)
}

func (v Vector) Equals(other Vector) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector) ToString() string {
	return fmt.Sprintf("Vector(%f, %f)", v.X, v.Y)
}

func (v Vector) Rotate(angle Angle) Vector {
	return Vector{
        X: (math.Cos(angle.rad) * v.X) - (math.Sin(angle.rad) * v.Y),
        Y: (math.Sin(angle.rad) * v.X) + (math.Cos(angle.rad) * v.Y),
	}
}

func (v Vector) AngleBetween(other Vector) Angle {
    return NewAngleRad(math.Atan2(
        other.Y * v.X - other.X * v.Y,
        other.X * v.X - other.Y * v.Y,
    ))
}

func (v Vector) ToPoint() Point {
	return Point{
		X: int(math.Floor(v.X)),
		Y: int(math.Floor(v.Y)),
	}
}
