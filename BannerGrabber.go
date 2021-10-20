package gobg

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type BannerGrabber struct {
	URL string
}

func (bg *BannerGrabber) Grab() {
	fmt.Printf("%s[*] Grabbing banner for: %s...\n%s", ColorYellow, bg.URL, ColorReset)

	res, err := curlGrab(bg.URL)

	if err != nil {
		fmt.Printf("%s[-] E: %s%s\n", ColorRed, err, ColorReset)
	} else {
		fmt.Printf("%s[*] Banner Grab Successful!%s\nBanner:\n%s%s%s", ColorYellow, ColorReset, ColorGreen, res, ColorReset)
	}

}

func (bg *BannerGrabber) Grabs() (string, error) {
	return res, err
}

func curlGrab(address string) (res string, err error) {
	cURL, ok := exec.LookPath("curl")

	if ok != nil {
		return "", errors.New("E: cURL not found!")
	}

	fmt.Printf("%s[*] Using cURL to perform the grab%.\n", ColorYellow, ColorReset)

	curlCommand := exec.Command(cURL, "-s", "-I", address)

	stdout, _ := curlCommand.StdoutPipe()
	stderr, _ := curlCommand.StderrPipe()

	curlCommand.Start()

	outBuf, errBuf := new(bytes.Buffer), new(bytes.Buffer)
	outBuf.ReadFrom(stdout)
	errBuf.ReadFrom(stderr)

	res = outBuf.String()
	errstr := errBuf.String()

	if errstr != "" {
		return "", errors.New(errstr)
	}

	return res, nil
}
