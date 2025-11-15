package main

// https://grugbrain.dev/

import (
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func GetCommandOutput(command []string) string {

	var prog string = command[0]
	var rest []string = command[1:]

	cmd := exec.Command(prog, strings.Join(rest, " "))

	// Execute the command and capture its standard output
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	return string(output)

}

func GetHostIp() string {
	switch runtime.GOOS {
	case "windows":
		log.Fatal("I haven't made this support Windows yet")
	case "linux":
		// TODO: gotta grep or something
		// since ip a is suuuuuuper verbose??
		return GetCommandOutput(strings.Split("ip a", " "))
	}

	return "{something went wrong}"
}

func GetHostname() string {
	switch runtime.GOOS {
	case "windows":
		return GetCommandOutput(strings.Split("hostname", " "))
	case "linux":
		return strings.TrimSpace(GetCommandOutput(strings.Split("cat /etc/hostname", " ")))
	}

	return "{something went wrong}"
}

// No need for get os family, since runtime.GOOS does that already!

func GetOSSubtype() string {
	switch runtime.GOOS {
	case "windows":
		var blorp string = GetCommandOutput(strings.Split("systeminfo | findstr /B /C:\"OS Name\" /C:\"OS Version\"", ""))
		var re = regexp.MustCompile(`(?m)\d\d \w\w\w`)
		var res = re.FindString(blorp)

		return res

	case "linux":
		var glorp string = GetCommandOutput(strings.Split("cat /etc/os-release", " "))
		var redistro = regexp.MustCompile(`(?m)\bNAME=\"\w*`)
		var distro = strings.Replace(redistro.FindString(glorp), "NAME=\"", "", 1)
		var rerel = regexp.MustCompile(`(?m)\bBUILD_ID=\w*`)
		var release = strings.Replace(rerel.FindString(glorp), "BUILD_ID=", "", 1)

		return distro + " " + release
	}

	return "{confused in this empty place}"
}
