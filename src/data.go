package src

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PageTitle     string             `json:"pageTitle" bson:"pageTitle"`
	HeaderTitle   string             `json:"headerTitle" bson:"headerTitle"`
	WebsiteLogo   string             `json:"websiteLogo" bson:"websiteLogo"`
	HomePageImage string             `json:"homePageImage" bson:"homePageImage"`
	HeaderLinks   []HeaderLink       `json:"headerLinks" bson:"headerLinks"`
	ServiceData   []ServiceData      `json:"serviceData" bson:"serviceData"`
	BlogData      []BlogData         `json:"blogData" bson:"blogData"`
}

type HeaderLink struct {
	Title    string `json:"title" bson:"title"`
	URL      string `json:"url" bson:"url"`
	IsActive bool   `json:"isActive" bson:"isActive"`
}

type ServiceData struct {
	Title    string `json:"title" bson:"title"`
	Content  string `json:"content" bson:"content"`
	URI      string `json:"uri" bson:"uri"`
	Image    string `json:"image" bson:"image"`
	IsActive bool   `json:"isActive" bson:"isActive"`
}

type BlogData struct {
	Title      string `json:"title" bson:"title"`
	Published  string `json:"" bson:"published"`
	Content    string `json:"content" bson:"content"`
	AuthorName string `json:"authorName" bson:"authorName"`
	Image      string `json:"image" bson:"image"`
	IsActive   bool   `json:"isActive" bson:"isActive"`
}
