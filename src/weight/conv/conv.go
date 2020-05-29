package weight

// KToP converts Kilograms to Pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * Conversion) }

// PToK converts Pounds to Kilograms.
func PToK(p Pounds) Kilograms { return Kilograms(p / Conversion) }
