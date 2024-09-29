package main

import (
	"encoding/json"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
} 

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type App struct {
	Img string
	Name string
	Path string
}

type AppList struct {
	Apps []App
}

func (a *AppList) Generate() error {
	file, err := os.Open("routes.json")
	if err != nil {
		return err
	}
	
	bytesBuffer, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	
	if err = json.Unmarshal(bytesBuffer, &a.Apps); err!=nil{
		return err
	}

	return nil
}


func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
}


func main() {
	e:= echo.New()
	e.Use(middleware.Logger())

	var data = AppList{}
	data.Generate()
	
	e.Static("/dist", "dist")
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "list", data)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
