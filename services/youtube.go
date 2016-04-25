package services

import (
	"flag"
	"log"
	"net/http"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/googleapi/transport"
	"github.com/toptal/models"
	_"fmt"
	"fmt"
)
const DEVELOPER_KEY = "AIzaSyCYBfp3QD0eMtsenKePuOemVuYqyaQ8l9M"

func GetVideos(search models.Search) (map[string]interface{},error) {
	flag.Parse()
	videos := make (map[string]interface{})
	client := &http.Client{Transport:&transport.APIKey{Key: DEVELOPER_KEY},}

	service, err := youtube.New(client)

	if err !=nil {
		log.Fatalf( "Error creatin new youtube client: %v",err )
		return videos,err
	}
	location := search.Lat+","+search.Lng

	fmt.Println(location)
	call := service.Search.List("id,snippet").Q(search.Term)
	response, err := call.Do()
	fmt.Println(err)
	if err != nil {
		return videos,err
	}
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {

			video :=models.Video{
				VideoId:item.Id.VideoId,
				Description:item.Snippet.Description,
				Title:item.Snippet.Title,
				UrlThumbnails: item.Snippet.Thumbnails.High.Url}
			videos[item.Id.VideoId] = video
		}
	}
	

	return videos,nil
}