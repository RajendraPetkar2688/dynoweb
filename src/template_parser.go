package src

import "html/template"

var blog_data Data = Data{
	PageTitle:   "Random Page",
	HeaderTitle: "Hello world",
	HeaderLinks: []HeaderLink{
		{
			Title:    "Home",
			URL:      "output.html",
			IsActive: true,
		},
		{
			Title:    "Blog",
			URL:      "Blog.html",
			IsActive: true,
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
			Title:    "Contact Us",
			URL:      "#",
			IsActive: true,
		},
	},

	BlogData: []BlogData{
		{
			Title:       "New Blog1",
			PublishedAt: "Yesterday at 10:00 AM",
			Content:     "Some random content",
			URI:         "Blogs/Blog1.html",
			Image:       "img1.jpg",
			IsActive:    true,
		},
		{
			Title:       "New Blog2",
			PublishedAt: "Yesterday at 10:00 AM",
			Content:     "Some random content",
			URI:         "Blogs/Blog2.html",
			Image:       "assests/img2.jpg",
			IsActive:    true,
		},
		{
			Title:       "New Blog3",
			PublishedAt: "Yesterday at 10:00 AM",
			Content:     "Some random content",
			URI:         "Blogs/Blog2.html",
			Image:       "assests/img2.jpg",
			IsActive:    true,
		},
	},

	ServiceData: []ServiceData{
		{
			STitle:    "Service1",
			SContent:  "Some random content",
			SURI:      "Blogs/Blog1.html",
			SImage:    "img1.jpg",
			SIsActive: true,
		},
		{
			STitle:    "Service2",
			SContent:  "Some random content",
			SURI:      "Blogs/Blog2.html",
			SImage:    "assests/img2.jpg",
			SIsActive: true,
		},
		{
			STitle:    "Service3",
			SContent:  "Some random content",
			SURI:      "Blogs/Blog2.html",
			SImage:    "assests/img2.jpg",
			SIsActive: true,
		},
	},
}

func ParseTemplate() (*template.Template, error) {
	tpl, err := tpl.ParseGlob("index1.html")
	if err != nil {
		return nil, err
	}
	return tpl, nil
}
