// dog is a dog
package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(1)
	if v != 7 {
		t.Error("Expected: 7, Got:", v)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}
