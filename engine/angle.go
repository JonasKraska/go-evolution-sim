package engine

import "math"

type Rad = float64
type Deg = float64

const (
    factorDegToRad = 180 / math.Pi
    factorRadToDeg = math.Pi / 180
)

type Angle struct {
    rad Rad
    deg Deg
}

func NewAngleRad(rad Rad) Angle {
    return Angle{
        rad: rad,
        deg: RadToDeg(rad),
    }
}

func NewAngleDeg(deg Deg) Angle {
    return Angle{
        rad: DegToRad(deg),
        deg: deg,
    }
}

func (a Angle) GetRad() Rad {
    return a.rad
}

func (a Angle) GetDeg() Deg {
    return a.deg
}

func RadToDeg(rad Rad) Deg {
    return rad * factorDegToRad
}

func DegToRad(deg Deg) Rad {
    return deg * factorRadToDeg
}
