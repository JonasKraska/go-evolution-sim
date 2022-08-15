package engine

type Gameable struct {
}

type Gamer interface {
    GetDimensions() Size
}
