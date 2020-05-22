package doggy

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	total := 70
	if n != total {
		t.Error("Expected", total, "Got", n)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(20)
	if n != 140 {
		t.Error("Expected", 140, "Got", n)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	// Output:
	// 140
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}
