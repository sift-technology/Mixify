package api

import (
	"log"

	"fmt"

	"github.com/zmb3/spotify"
)

func Recommend(client *spotify.Client, user *Response) []spotify.SimpleTrack {

	trackID := []spotify.ID{}

	trackID[0] = user.M.Track1
	trackID[1] = user.M.Track2

	var seed spotify.Seeds
	seed.Tracks = trackID

	ta := spotify.NewTrackAttributes().TargetAcousticness(user.M.Acousticness).
		TargetPopularity(user.M.Popularity).
		TargetDanceability(user.M.Danceability).
		TargetEnergy(user.M.Energy)

	var opt spotify.Options
	var country string = "US"
	var offset int = 0
	var r string = "medium"
	opt.Country = &country
	opt.Offset = &offset
	opt.Timerange = &r

	recs, err := client.GetRecommendations(seed, ta, &opt)
	if err != nil {
		log.Fatalf("Couldn't get recommendation: %v", err)
	}
	for i := 0; i < len(recs.Tracks); i++ {
		trackID[i] = recs.Tracks[i].ID
	}

	playlist, err := client.CreatePlaylistForUser("nateisding", "EEEWEWEWEW", "OKA", true)

	client.AddTracksToPlaylist(playlist.ID, trackID...)
	fmt.Println(playlist.SimplePlaylist.ExternalURLs)

	return recs.Tracks

}
