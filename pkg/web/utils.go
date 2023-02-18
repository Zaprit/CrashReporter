package web

import (
    "fmt"
    "html/template"
)


func RenderTemplate(templateFile string) (*template.Template, error){
    tmpl, err := template.ParseFiles("static/template/base.gohtml", fmt.Sprintf("static/template/%s.gohtml",templateFile))

    return tmpl, err
}