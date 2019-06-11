package teamspirit

import (
	"time"

	"github.com/NasSilverBullet/autospirit/pkg/webdriver"
	"github.com/sclevine/agouti"
)

type webDrriverRepository struct {
	driver *webdriver.WebDriver
}

func NewWebDrriverRepository() (*webDrriverRepository, error) {
	d, err := webdriver.New()
	return &webDrriverRepository{d}, err
}

func (r *webDrriverRepository) Start() (*agouti.Page, error) {
	if err := r.driver.Driver.Start(); err != nil {
		return nil, err
	}
	page, err := r.driver.Driver.NewPage()
	if err != nil {
		return page, err
	}
	if err = page.Navigate(r.driver.RootURL); err != nil {
		return page, err
	}
	return page, err
}

func (r *webDrriverRepository) Login(p *agouti.Page) error {
	if err := p.FindByID("username").Fill(r.driver.ID); err != nil {
		return err
	}
	if err := p.FindByID("password").Fill(r.driver.Password); err != nil {
		return err
	}
	if err := p.FindByID("Login").Click(); err != nil {
		return err
	}
	return nil
}

func (r *webDrriverRepository) Go(p *agouti.Page) error {
	err := insertTimestamp(p, "pushStart")
	return err
}

func (r *webDrriverRepository) Out(p *agouti.Page) error {
	err := insertTimestamp(p, "pushEnd")
	return err
}

func (r *webDrriverRepository) Stop() error {
	time.Sleep(10 * time.Second)
	return r.driver.Driver.Stop()
}

func insertTimestamp(p *agouti.Page, id string) error {
	if err := p.FindByID("publisherAttach09D2800000A981a").Click(); err != nil {
		return err
	}
	if err := p.FindByID("09D2800000A981a_06628000006IAtS").SwitchToFrame(); err != nil {
		return err
	}
	if err := p.FindByID(id).Click(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	if err := p.Screenshot("test.jpg"); err != nil {
		return err
	}
	return nil
}
