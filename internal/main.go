package main

import (
	"os"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath     = "tools/selenium-server-4.3.0.jar"
	chromeDriverPath = "tools/chromedriver"
	port             = 8080
)

//"os"
// "github.com/tebeka/selenium"

func main() {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	// https://sum.unmsm.edu.pe
	wd, err := selenium.NewRemote(caps, "http://localhost:"+strconv.Itoa(port)+"/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	wd.Refresh()
	wd.Get("https://sum.unmsm.edu.pe/")
	time.Sleep(time.Second * 5)
}
