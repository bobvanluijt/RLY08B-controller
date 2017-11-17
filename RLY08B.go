package main

import (
	"flag"
	"log"

	"github.com/google/gousb" // https://godoc.org/github.com/google/gousb
)

func main() {

	// get Flags and set custom variables if needed
	rawCOM := flag.Int64("command", 101, "Set a command based Based on the commands: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm")
	rawVID := flag.Int64("vid", 0x04d8, "Add the VID of the usb device.")
	rawPID := flag.Int64("pid", 0xffee, "Add the PID of the usb device.")
	flag.Parse()

	// get COM (= command)
	COM := gousb.ID(*rawCOM)

	// get VID
	VID := gousb.ID(*rawVID)
	log.Println("Running with VID: ", VID)

	// get PID
	PID := gousb.ID(*rawPID)
	log.Println("Running with PID: ", PID)

	// Based on the commands: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm
	data := make([]byte, 5)
	data[0] = byte(COM) // see command list from url

	// Initialize a new Context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.
	// found via: `$ system_profiler -xml SPUSBDataType``
	dev, err := ctx.OpenDeviceWithVIDPID(VID, PID)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer dev.Close()

	// Switch the configuration to #1.
	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("%s.Config(1): %v", dev, err)
	}
	defer cfg.Close()

	// In the config #1, claim interface #1 with alt setting #0.
	intf, err := cfg.Interface(1, 0)
	if err != nil {
		log.Fatalf("%s.Interface(1, 0): %v", cfg, err)
	}
	defer intf.Close()

	// set writing endpoint
	ep, err := intf.OutEndpoint(0x02)
	if err != nil {
		log.Fatalf("%s.OutEndpoint(2): %v", intf, err)
	}

	// Write data to the USB device.
	numBytes, err := ep.Write(data)

	// check if success
	if numBytes != 5 {
		log.Fatalf("%s.Write([5]): only %d bytes written, returned error is %v", ep, numBytes, err)
	}

	log.Println("Command ", COM, " delivered successful")

}
