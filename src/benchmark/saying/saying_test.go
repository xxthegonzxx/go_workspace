package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	x := "Welcome my dear James"
	if s != x {
		t.Error("Expected", s, "Got", x)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
