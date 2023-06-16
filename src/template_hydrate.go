package src

import "html/template"

//Struct data generate and execute template -> input from template_parser.go

func HydrateTemplate(tpl *template.Template) error {
	op, err := CreateBundles()
	if err != nil {
		return err
	}
	if err := tpl.ExecuteTemplate(op, "index.html", blog_data); err != nil {
		return err
	}

	return nil
}
