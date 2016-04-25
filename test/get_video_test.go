package test

import (
	"testing"
	 "github.com/toptal/models"
	"github.com/toptal/services"
	"fmt"
)

func TestSearch(t *testing.T) {
	search := models.Search{Term:"defcon",Lat:"36.1126126",Lng:"-115.1708585"}
	videos , err :=services.GetVideos(search)
	fmt.Println(videos)
	if err !=nil {
		t.Error("GetVideoTest, fail to execute the search" )
		panic(err)
	}

	if len(videos) == 0 {
		t.Error( "GetVideoTest, fails number of elements expected different to zero" )
	}
}

