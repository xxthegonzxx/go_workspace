// Package weight converts pounds to kilograms.
package weight

import "fmt"

// Kilograms holds value in float64 format.
type Kilograms float64

// Pounds holds value in float64 format.
type Pounds float64

const (
	// Conversion formula for converting Pounds to Kilograms and vice versa.
	Conversion = 2.205
)

func (p Pounds) String() string    { return fmt.Sprintf("%g lbs", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%g kilograms", k) }
