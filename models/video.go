package models


type Video struct {
	VideoId string  	`json:"videoId"`
	Title string		`json:"title"`
	UrlThumbnails string 	`json:"urlThumbnails"`
	Description string 	`json:"description"`
}
