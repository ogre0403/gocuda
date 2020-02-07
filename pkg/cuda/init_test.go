package cuda

import (
	"fmt"
)

// needed for all other tests.
func init() {
	Init(0)
	fmt.Println("Init CUDA")
}
