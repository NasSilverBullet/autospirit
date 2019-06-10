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
	r, err := webdriver.New()
	return &webDrriverRepository{r}, err
}

func (r *webDrriverRepository) Start() (*agouti.Page, error) {
	if err := r.driver.Driver.Start(); err != nil {
		return nil, err
	}
	page, err := r.driver.Driver.NewPage()
	if err != nil {
		return page, err
	}
	err = page.Navigate(r.driver.RootURL)
	return page, err
}

func (r *webDrriverRepository) Login(p *agouti.Page) error {
	userName := p.FindByID("username")
	if err := userName.SendKeys(r.driver.ID); err != nil {
		return err
	}
	password := p.FindByID("password")
	if err := password.SendKeys(r.driver.Password); err != nil {
		return err
	}
	err := p.RunScript(click("Login"), nil, nil)
	return err
}

func (r *webDrriverRepository) Go(p *agouti.Page) error {
	if err := p.RunScript(click("btnStInput"), nil, nil); err != nil {
		return err
	}
	return nil
}

func (r *webDrriverRepository) Out(p *agouti.Page) error {
	if err := p.RunScript(click("btnEtInput"), nil, nil); err != nil {
		return err
	}
	return nil
}

func (r *webDrriverRepository) Stop() error {
	time.Sleep(5 * time.Second)
	err := r.driver.Driver.Stop()
	return err
}

// TODO: p.FindByID("Login").Click() のようにする
func click(id string) string {
	return "document.getElementById('" + id + "').click()"
}
