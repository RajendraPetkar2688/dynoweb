package src

import "html/template"

func ParseTemplate() (*template.Template, error) {
	tpl, err := tpl.ParseGlob("index1.html")
	if err != nil {
		return nil, err
	}
	return tpl, nil
}
