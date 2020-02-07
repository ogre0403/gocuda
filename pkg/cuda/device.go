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

// Gets the name of the device.
func DeviceGetName(dev Device) string {
	size := 256
	buf := make([]byte, size)
	cstr := C.CString(string(buf))
	err := Result(C.cuDeviceGetName(cstr, C.int(size), C.CUdevice(dev)))
	if err != SUCCESS {
		panic(err)
	}
	return C.GoString(cstr)
}
