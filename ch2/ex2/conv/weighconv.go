package conv

import "fmt"

type Kilogram float64
type Pond float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pond) String() string     { return fmt.Sprintf("%glb", p) }

func KilogramToPond(k Kilogram) Pond { return Pond(k / 0.453592) }
func PondToKilogram(p Pond) Kilogram { return Kilogram(p * 0.453592) }
