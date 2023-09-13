package floatUtil

import (
	"fmt"
	"strconv"
)

func ToFixed2(val float64) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	return v
}

type Float64WithPrecision2 float64

func (f Float64WithPrecision2) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', 2, 64)), nil
}
func (f Float64WithPrecision2) String() string {
	return strconv.FormatFloat(float64(f), 'f', 2, 64)
}

type Float64WithPrecision3 float64

func (f Float64WithPrecision3) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), 'f', 3, 64)), nil
}
func (f Float64WithPrecision3) String() string {
	return strconv.FormatFloat(float64(f), 'f', 3, 64)
}
