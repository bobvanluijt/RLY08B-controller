// Golang implementation of the RLY08B
// https://github.com/bobvanluijt/RLY08B-controller
// MIT

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// global variables
var DEFAULTDEVICE = "/dev/cu.usbmodem00036401"

// execute on the actual devices
func execOnDevice(d string, c int) error {

	// open the device
	f, err := os.OpenFile(d, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	// write the command
	w, err := f.Write([]byte{byte(c)})
	if err != nil {
		return err
	}

	// print succes log
	log.Println("Command ", int(c), " (", w, " bytes) delivered successful.")

	// Sync, flush
	f.Sync()

	// Close the file
	f.Close()

	// success
	return nil
}

// get command
func parseCommand(w http.ResponseWriter, r *http.Request) {

	// validate if POST
	if r.Method == "POST" {

		// check if device is POSTed
		var device string
		if r.Header.Get("device") != "" {
			device = r.Header.Get("device")
		} else {
			device = DEFAULTDEVICE
		}

		// find /command/ at index 0
		index := strings.Index(r.URL.String(), "/command/")
		if int(index) != 0 {
			io.WriteString(w, "No command found!")
		} else {
			// Validate the command
			commandMap := strings.Split(r.URL.String(), "/")
			command, err := strconv.ParseInt(commandMap[2], 0, 8)
			if err != nil {
				io.WriteString(w, "Whoops, wrong command! Check: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm for a list of commands")
			}
			// check if command is in the correct range: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm
			if command >= 100 && command <= 118 {
				err := execOnDevice(device, int(command))
				if err != nil {
					io.WriteString(w, "Whoops, something went wrong. Check the web server log")
				}
			} else {
				io.WriteString(w, "Whoops, wrong command! Check: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm for a list of commands")
			}
		}
	}

}

func main() {

	// get Flags and set custom variables if needed
	rawCOM := flag.Int("command", 100, "Set a command based Based on the commands: http://www.robot-electronics.co.uk/htm/usb_rly08btech.htm")
	rawDevice := flag.String("device", DEFAULTDEVICE, "Set the device.")
	webservice := flag.Bool("webservice", false, "Run as web-service? (true or false)")
	flag.Parse()

	// if the webservice is defined, execute it
	if *webservice == true {

		// run the webservice
		log.Println("Webservice running on localhost:8080")
		http.HandleFunc("/", parseCommand)
		http.ListenAndServe(":8080", nil)

	} else {

		// run command line interface
		err := execOnDevice(*rawDevice, *rawCOM)
		if err != nil {
			log.Panic("Error while executing %v", err)
		}

	}

}
