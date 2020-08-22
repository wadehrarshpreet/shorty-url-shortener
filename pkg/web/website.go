package web

import (
	"html/template"
	"io"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
)

// Template custom type of Echo Renderer
type Template struct {
	templates *template.Template
}

// Render Custom Echo Rendered function
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// InitWebsite to initialize Web Routes and output html
func InitWebsite(e *echo.Echo) {

	e.Logger.Info("Initialize webview routes")
	// Initialize SEO data
	seoData, err := initSEO(path.Join(path.Dir("pkg/web/"), "SEOData.json"))
	if err != nil {
		e.Logger.Errorf("Failed to initialize SEO data %v", err)
	}

	t := &Template{
		templates: template.Must(template.ParseFiles("web/index.html")),
	}

	e.Renderer = t

	// Register HTML route SPA
	e.GET("/*", func(c echo.Context) error {
		requestURLPath := c.Request().URL.Path
		pageSEOData, ok := seoData[requestURLPath]
		if !ok {
			pageSEOData = seoData["/"]
		}

		return c.Render(http.StatusOK, "index.html", echo.Map{
			"SEO": pageSEOData,
		})
	})

}
