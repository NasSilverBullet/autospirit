package teamspirit

import (
	"time"

	"github.com/NasSilverBullet/autospirit/pkg/webdriver"
)

type webDrriverRepository struct {
	d *webdriver.WebDriver
}

func NewWebDrriverRepository() (*webDrriverRepository, error) {
	r, err := webdriver.New()
	return &webDrriverRepository{r}, err
}

func (r *webDrriverRepository) Start() error {
	if err := r.d.Driver.Start(); err != nil {
		return err
	}
	page, err := r.d.Driver.NewPage()
	if err != nil {
		return err
	}
	err = page.Navigate(r.d.RootURL)
	time.Sleep(3 * time.Second)
	return err
}

func (r *webDrriverRepository) Login() error {
	return nil
}

func (r *webDrriverRepository) Go() error {
	return nil
}

func (r *webDrriverRepository) Out() error {
	return nil
}

func (r *webDrriverRepository) Stop() error {
	return nil
}
