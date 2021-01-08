package tempconv

type Celsius float64
type Fahrenheit float64

// Ex 2.1 Kelvin type
type Kelvin float64

const (
	AbsoluteZeroC = -273.15
	FreezingC     = 0.0
	BoilingC      = 100
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CtoK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

func KtoC(k Kelvin) Celsius { return Celsius(k) - AbsoluteZeroC }

func KtoF(k Kelvin) Fahrenheit { return CtoF(KtoC(k)) }

func FtoK(f Fahrenheit) Kelvin { return CtoK(FtoC(f)) }
