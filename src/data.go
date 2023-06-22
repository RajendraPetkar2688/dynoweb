package src

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PageTitle   string             `json:"pageTitle" bson:"pageTitle"`
	HeaderTitle string             `json:"headerTitle" bson:"headerTitle"`
	HeaderLinks []HeaderLink       `json:"headerLinks" bson:"headerLinks"`
	ServiceData []ServiceData      `json:"serviceData" bson:"serviceData"`
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
