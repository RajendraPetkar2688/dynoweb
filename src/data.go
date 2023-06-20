package src

type Data struct {
	PageTitle   string
	HeaderTitle string
	HeaderLinks []HeaderLink
	ServiceData []ServiceData
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
	URI         string
	Image       string
	IsActive    bool
}

type ServiceData struct {
	STitle    string
	SContent  string
	SURI      string
	SImage    string
	SIsActive bool
}
