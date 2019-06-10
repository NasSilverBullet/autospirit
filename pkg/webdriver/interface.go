package webdriver

type WebDriverInterface interface {
	Start(*WebDriver) error
	Login(*WebDriver) error
	Go(*WebDriver) error
	Out(*WebDriver) error
	Stop(*WebDriver) error
}
