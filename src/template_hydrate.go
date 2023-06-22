package src

import (
	"fmt"
	"html/template"
)

//Struct data generate and execute template -> input from template_parser.go

func HydrateTemplate(tpl *template.Template, data Data) error {
	//fmt.Sprintf("%s/Output/%s.html", data.ID.Hex(), data.ID.Hex())
	op, err := CreateBundles(fmt.Sprintf("Output/%s.html", data.ID.Hex()))
	if err != nil {
		return err
	}
	if err := tpl.ExecuteTemplate(op, "index1.html", data); err != nil {
		return err
	}

	return nil
}
