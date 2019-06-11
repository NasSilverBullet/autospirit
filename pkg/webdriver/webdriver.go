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
	appdir     = "/src/github.com/NasSilverBullet/autospirit"
	configfile = "config/config.json"
)

type WebDriver struct {
	Driver   *agouti.WebDriver
	Browser  string `json:"browser"`
	RootURL  string `json:"rooturl"`
	ID       string `json:"id"`
	Password string `json:"password"`
}

func New() (*WebDriver, error) {
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
	case "phantomjs":
		d.Driver = agouti.PhantomJS(
		//agouti.Debug,
		)
	// TODO: 作れたら作る
	//case "selenium":
	//d.Driver = agouti.Selenium(agouti.Browser("chrome"))
	default:
		err = errors.New(fmt.Sprintf("driver '%s' is not compatible, plese use another driver", d.Browser))
	}
	return err
}
