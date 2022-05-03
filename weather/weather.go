package weather

import "fmt"

// Prediction describes a weather prediction.
type Prediction uint8

// The supported weather prediction types.
const (
	Sunny Prediction = iota
	Rain
	Overcast
	Snow
	Unknown
)

// Locator is implemented by objects that can represent a location as a
// pair of GPS coordinates.
type Locator interface {
	Coords() (float64, float64, error)
}

// Predict the weather at the specified location.
func Predict(loc Locator) (Prediction, error) {
	coord1, coord2, err := loc.Coords()

	if err != nil {
		fmt.Println(coord1)
		fmt.Println(coord2)
		return Unknown, err
	}
	// ...

	return Prediction(5), nil
}

func GetVersion() string {
	return "v1.0.0"	
}
