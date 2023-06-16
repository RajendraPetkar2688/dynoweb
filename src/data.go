package src

type Data struct {
	PageTitle   string
	HeaderTitle string
	HeaderLinks []HeaderLink
	BlogData    []BlogData
}

type HeaderLink struct {
	Title    string
	URL      string
	IsActive bool
}

type BlogData struct {
	Title       string
	PublishedAt string
	Content     string
	IsActive    bool
}
