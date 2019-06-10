package teamspirit

import (
	"time"

	"github.com/NasSilverBullet/autospirit/pkg/webdriver"
)

type webDriverInterface struct{}

func NewWedDriverInteface() webdriver.WebDriverInterface {
	return &webDriverInterface{}
}

func (i *webDriverInterface) Start(d *webdriver.WebDriver) error {
	if err := d.Driver.Start(); err != nil {
		return err
	}
	page, err := d.Driver.NewPage()
	if err != nil {
		return err
	}
	err = page.Navigate(d.RootURL)
	time.Sleep(3 * time.Second)
	return err
}

func (i *webDriverInterface) Login(d *webdriver.WebDriver) error {
	return nil
}

func (i *webDriverInterface) Go(d *webdriver.WebDriver) error {
	return nil
}

func (i *webDriverInterface) Out(d *webdriver.WebDriver) error {
	return nil
}

func (i *webDriverInterface) Stop(d *webdriver.WebDriver) error {
	return nil
}
