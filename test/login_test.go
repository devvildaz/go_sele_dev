package seleni_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/tebeka/selenium"
	"github.com/ysmood/got"
)

const (
	chromeDriverPath = "../tools/chromedriver"
	port             = 8888
)

func generateSelenium() (*selenium.Service, selenium.WebDriver) {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)
	fmt.Print("INIT DEBUG SELENIUM")
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()
	fmt.Print("INIT SERVICE SELENIUM")
	caps := selenium.Capabilities{"browserName": "chrome"}
	// https://sum.unmsm.edu.pe
	wd, err := selenium.NewRemote(caps, "http://localhost:"+strconv.Itoa(port)+"/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()
	fmt.Print("INIT REMOTE SELENIUM")
	return service, wd
}

func TestSeleniAdvancedSuite(t *testing.T) {

	got.Each(t, func(t *testing.T) SeleniAdvancedSuite {
		g := got.New(t)

		//g.Parallel()

		g.PanicAfter(time.Second * 10)

		return SeleniAdvancedSuite{g}
	})
}

type SeleniAdvancedSuite struct {
	got.G
}

func (g SeleniAdvancedSuite) SeleniTest() {
	service, driver := generateSelenium()
	defer service.Stop()

	if err := driver.Get("https://sum.unmsm.edu.pe/"); err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	g.Eq("10", "10")
}

// func (g SeleniAdvancedSuite) SeleniTest2() {
// 	service, driver := generateSelenium()
// 	defer service.Stop()
// 	time.Sleep(5 * time.Second)
// 	if err := driver.Get("https://www.google.com/"); err != nil {
// 		panic(err)
// 	}
// 	time.Sleep(5 * time.Second)
// 	g.Eq("10", "10")
// }
