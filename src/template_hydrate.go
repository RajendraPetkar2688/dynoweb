package src

import (
	"fmt"
	"html/template"
)

// Struct data generate and execute template -> input from template_parser.go
// op, err := CreateBundles(fmt.Sprintf("./output/%s.html", data.ID.Hex()))
func HydrateTemplate(tpl *template.Template, data Data) error {
	op, err := CreateBundles(fmt.Sprintf("%s.html", data.ID.Hex()))
	if err != nil {
		return err
	}
	if err := tpl.ExecuteTemplate(op, "index.html", data); err != nil {
		return err
	}

	return nil
}
