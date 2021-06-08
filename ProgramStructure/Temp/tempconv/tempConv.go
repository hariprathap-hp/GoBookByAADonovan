package tempconv

func CToF(c Celcius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

func FToC(f Farenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func CToK(c Celcius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KToC(k Kelvin) Celcius {
	return Celcius(k - FreezingK)
}

func FToM(ft Feet) Metres {
	return Metres(ft / 3.28)
}

func MToF(m Metres) Feet {
	return Feet(m * 3.28)
}

func KToP(kg Kilogram) Pounds {
	return Pounds(kg * 2.2046)
}

func PToK(p Pounds) Kilogram {
	return Kilogram(p / 2.2046)
}
