package lenghtconv

// MToF converts meters to feet
func MToF(m Meter) Feet { return Feet(m * LengthCoefficient) }

// FToM converts feet to meters
func FToM(f Feet) Meter { return Meter(f / LengthCoefficient) }
