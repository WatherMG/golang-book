package weightconv

// KToP Converts kilograms to pounds
func KToP(k Kilos) Pounds { return Pounds(k / KiloPoundCoefficient) }

// PToK converts pounds into kilograms
func PToK(p Pounds) Kilos { return Kilos(p * KiloPoundCoefficient) }
