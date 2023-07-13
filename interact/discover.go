package interact

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/mdlayher/hdhomerun"
)

func discover(timeout int) ([]*hdhomerun.DiscoveredDevice, error) {
	d, err := hdhomerun.NewDiscoverer(
		hdhomerun.DiscoverDeviceType(hdhomerun.DeviceTypeTuner),
	)
	if err != nil {
		return nil, fmt.Errorf("error starting discovery: %v", err)
	}

	// Discover devices for up to 2 seconds or until canceled.
	var devices []*hdhomerun.DiscoveredDevice
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

outer:
	for {
		device, err := d.Discover(ctx)
		switch err {
		case nil:
			// Found a device.
			// If only one was expected, could invoke cancel here.
			devices = append(devices, device)

		case io.EOF:
			// No more devices to be found.
			break outer
		default:
			return nil, fmt.Errorf("error during discovery: %v", err)
		}
	}

	return devices, nil
}
