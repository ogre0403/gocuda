package cuda

import (
	"fmt"
	"testing"
)

func TestDevice(t *testing.T) {
	count := DeviceGetCount()
	fmt.Println("DeviceGetCount:", count)
	for i := 0; i < count; i++ {
		fmt.Println("Device Name", DeviceGetName(Device(i)))
	}
}
