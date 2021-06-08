package tempconv

import "fmt"

//declare the variables that will be exported
type Celcius float64
type Farenheit float64
type Kelvin float64

type Kilogram float64
type Pounds float64

type Feet float64
type Metres float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilngC       Celcius = 100
	FreezingK     Kelvin  = 273.15
)

func (c Celcius) String() string {
	return fmt.Sprintf("%g°c", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%g°f", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%g KGms", kg)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%g pounds", p)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%g feet", ft)
}

func (mts Metres) String() string {
	return fmt.Sprintf("%g metres", mts)
}
