package spotify

import (
	"log"

	"github.com/zmb3/spotify"
)

func Recommend(client *spotify.Client) {

	artistID := []spotify.ID{}
	trackID := []spotify.ID{}
	genre := []string{}

	artistID = append(artistID, spotify.ID("06HL4z0CvFAxyc27GXpf02")) // Sophie msmsmsmsmsmsmsmmsms
	trackID = append(trackID, spotify.ID("0V3wPSX9ygBnCm8psDIegu"))   // Super Bass
	genre = append(genre, "pop")

	var seed spotify.Seeds
	seed.Artists = artistID
	seed.Tracks = trackID
	seed.Genres = genre

	ta := spotify.NewTrackAttributes().TargetAcousticness(0.1).TargetPopularity(40).TargetDanceability(0.8)

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

	log.Println("song recs:", recs.Tracks)

}
