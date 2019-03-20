// dog is a dog
package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {

    type test struct {
        data int
        answer int
    }

    tests := []test{
        {1,7},
        {2,14},
        {3,21},
        {-1,0},
        {-10,0},
    }

    for _, v := range tests {
	    d := Years(v.data)
	    if d != v.answer {
		    t.Error("Expected: ", d, "Got:", v.answer)
	    }
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
