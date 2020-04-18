/*
	Julius Olson
	Simple utility for controlling OLED screen brightness
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
	Read current brightness
*/
func getCurrent() (string, error) {
	cmd := exec.Command("bash", "-c", "xrandr --verbose | grep -i brightness")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	curr := strings.Split(string(out), ":")[1]
	return strings.TrimSpace(curr), nil
}

/*
	Set brightness
*/
func set(brightness float64) {
	if brightness > 1.0 || brightness < 0.1 {
		return
	}

	cmd := exec.Command("xrandr", "--output", "eDP1", "--brightness", fmt.Sprintf("%v", brightness))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide argument: current|inc|dec|max")
		return
	}
	arg := os.Args[1]
	switch arg {
	case "current":
		out, err := getCurrent()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Printf("Brightness: %v\n", out)
	case "inc":
		out, err := getCurrent()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		curr, _ := strconv.ParseFloat(out, 64)
		set(curr + 0.1)
	case "dec":
		out, err := getCurrent()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		curr, _ := strconv.ParseFloat(out, 64)
		set(curr - 0.1)
	case "max":
		set(1.0)
	default:
		fmt.Println("Provide argument: current|inc|dec|max")
	}
}
