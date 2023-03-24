package api

import (
	"log"

	"github.com/zmb3/spotify"
)

func Recommend(client *spotify.Client, user *Response) []spotify.SimpleTrack {

	artistID := []spotify.ID{}
	trackID := []spotify.ID{}
	genre := []string{}

	artistID = append(artistID, spotify.ID("5a2w2tgpLwv26BYJf2qYwu")) // Sophie msmsmsmsmsmsmsmmsms
	trackID = append(trackID, spotify.ID("3hlksXnvbKogFdPbpO9vel"))   // Super Bass
	genre = append(genre, "pop")

	var seed spotify.Seeds
	seed.Artists = artistID
	seed.Tracks = trackID
	seed.Genres = genre

	ta := spotify.NewTrackAttributes().TargetAcousticness(user.M.Acousticness).
		TargetPopularity(user.M.Popularity).
		TargetDanceability(user.M.Danceability).
		TargetEnergy(user.M.Energy)

	var opt spotify.Options
	var lim int = 100
	var country string = "US"
	var offset int = 0
	var r string = "medium"
	opt.Limit = &lim
	opt.Country = &country
	opt.Offset = &offset
	opt.Timerange = &r

	recs, err := client.GetRecommendations(seed, ta, &opt)
	if err != nil {
		log.Fatalf("Couldn't get recommendation: %v", err)
	}

	return recs.Tracks

}
