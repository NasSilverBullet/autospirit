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
	p, err := r.driver.Driver.NewPage()
	if err != nil {
		return p, err
	}
	if err = p.Navigate(r.driver.RootURL); err != nil {
		return p, err
	}
	if err = p.Session().SetImplicitWait(30); err != nil {
		return p, err
	}
	return p, err
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
	if err := insertTimestamp(p, "pushStart"); err != nil {
		return err
	}
	if err := p.Screenshot("screenshot/hello.jpg"); err != nil {
		return err
	}
	return nil
}

func (r *webDrriverRepository) Out(p *agouti.Page) error {
	if err := insertTimestamp(p, "pushEnd"); err != nil {
		return err
	}
	if err := p.Screenshot("screenshot/bye.jpg"); err != nil {
		return err
	}
	return nil
}

func (r *webDrriverRepository) Stop() error {
	time.Sleep(10 * time.Second)
	return r.driver.Driver.Stop()
}

func insertTimestamp(p *agouti.Page, id string) error {
	time.Sleep(10 * time.Second)
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
	return nil
}
