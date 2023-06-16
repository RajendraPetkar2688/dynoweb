package main

import "github.com/cleverextechnology/site-generator/src"

func main() {
	tpl, err := src.ParseTemplate()
	if err != nil {
		panic(err)
	}
	if err := src.HydrateTemplate(tpl); err != nil {
		panic(err)
	}
}
