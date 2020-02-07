package cuda

// This file implements CUDA driver device management

//#include <cuda.h>
import "C"

// CUDA Device number.
type Device int

// Returns the number of devices with compute capability greater than or equal to 1.0 that are available for execution.
func DeviceGetCount() int {
	var count C.int
	err := Result(C.cuDeviceGetCount(&count))
	if err != SUCCESS {
		panic(err)
	}
	return int(count)
}
