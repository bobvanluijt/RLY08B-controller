# USB-RLY08B - 8 relay outputs controller

1. http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm
2. http://www.robot-electronics.co.uk/products/relay-modules/usb-relay/usb-rly08b-8-channel-relay-module.html

Needed lib: `usblib`

# Getting binaries

Don't want to compile? No worries, bins are here: https://github.com/bobvanluijt/RLY08B-controller/tree/master/dist

# Running from command line

`$ ./RLY08B -help` for help

Example:<br>
`$ ./RLY08B -command=100 -device="/dev/cu.usbmodem00036401"`

Commands can be found here: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm

Find your device by running: 

# Running as a webservice

You can also run this package as a webservice.

`$ ./RLY08B -webservice=true`

You can now POST to `http://localhost/command/xxx` where xxx = commands: https://github.com/bobvanluijt/RLY08B-controller/releases

Example: `http://localhost/command/100` (turns on all relays)

To set a custom device add the HEADER: "device" with the path to the device.

Example: `device="/dev/cu.usbmodem00036401"`