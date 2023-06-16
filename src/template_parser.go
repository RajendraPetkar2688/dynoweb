package src

import "html/template"

var blog_data Data = Data{
	PageTitle:   "Random Page",
	HeaderTitle: "Hello world",
	HeaderLinks: []HeaderLink{
		{
			Title:    "Home",
			URL:      "#",
			IsActive: true,
		},
		{
			Title:    "Blog",
			URL:      "Blog.html",
			IsActive: false,
		},
		{
			Title:    "Services",
			URL:      "#",
			IsActive: true,
		},
		{
			Title:    "About",
			URL:      "#",
			IsActive: true,
		},
		{
			Title:    "Contact",
			URL:      "#",
			IsActive: true,
		},
	},
	BlogData: []BlogData{
		{
			Title:       "New Blog1",
			PublishedAt: "Yesterday at 10:00 AM",
			Content:     "Some random content",
			IsActive:    true,
		},
		{
			Title:       "New Blog2",
			PublishedAt: "Yesterday at 10:00 AM",
			Content:     "Some random content",
			IsActive:    true,
		},
	},
}

func ParseTemplate() (*template.Template, error) {
	tpl, err := tpl.ParseGlob("index.html")
	if err != nil {
		return nil, err
	}
	return tpl, nil
}
