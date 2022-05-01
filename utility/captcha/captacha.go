package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var Captcha *base64Captcha.Captcha

func Setup() {
	driver := &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: 2,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"chromohv.ttf"},
	}
	driver = driver.ConvertFonts()
	var store base64Captcha.Store
	//if red == nil {
	//	store = base64Captcha.DefaultMemStore
	//} else {
	//	store = NewRedisStore(g.Redis())
	//}
	store = base64Captcha.DefaultMemStore
	//store = NewRedisStore()

	Captcha = base64Captcha.NewCaptcha(driver, store)
}
