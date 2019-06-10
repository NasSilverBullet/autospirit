package webdriver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sclevine/agouti"
)

const (
	appdir     = "/src/github.com/NasSilverBullet/auto-insert_timestamps-on-teamsprit"
	configfile = "config.json"
)

type WebDriver struct {
	Driver   *agouti.WebDriver
	Browser  string `json:"browser"`
	RootURL  string `json:"rooturl"`
	ID       string `json:"id"`
	Password string `json:"password"`
}

func NewWebDriver() (*WebDriver, error) {
	b, err := openConfigFile()
	if err != nil {
		return nil, err
	}
	d := &WebDriver{}
	err = json.Unmarshal(b, d)
	if err != nil {
		return nil, err
	}
	err = setDriver(d)
	return d, err
}

func openConfigFile() ([]byte, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return nil, errors.New(fmt.Sprintf("\"%s\" is not defined, please define", gopath))
	}
	dir := gopath + appdir
	if err := os.Chdir(dir); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(configfile, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()
	b, err := ioutil.ReadAll(f)
	return b, err
}

func setDriver(d *WebDriver) error {
	var err error
	switch d.Browser {
	case "chrome":
		d.Driver = agouti.ChromeDriver(agouti.Browser("chrome"))
	case "firefox":
	//	TODO: 作れたら作る
	// Edgeとか他のはいいよね
	default:
		err = errors.New(fmt.Sprintf("driver '%s' is not compatible, plese use another driver", d.Browser))
	}
	return err
}
