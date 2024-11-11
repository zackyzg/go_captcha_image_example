package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mojocn/base64Captcha"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load templates
	t := &Template{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = t

	// Route to serve the login page
	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	})

	// Route to generate captcha
	e.GET("/captcha", func(c echo.Context) error {
    // driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	//driver := base64Captcha.NewDriverString(
	//	60, 160, 40, base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine|base64Captcha.OptionShowSineLine,
	//	4, "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", nil, nil, []string{"DeborahFancyDress.ttf"})
	driver := base64Captcha.NewDriverString(
		60, 160, 20, base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSineLine, 4,
		"1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", nil, nil, []string{"actionj.ttf"})
		cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
		id, b64s, answer, err := cp.Generate()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "captcha generation failed"})
		}
		return c.JSON(http.StatusOK, map[string]string{"id": id, "captcha": b64s, "answer": answer})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
