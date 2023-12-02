package main

import (
	"fmt"
	"ghibli/pkg/config"
	"io"
	"log"
	"net/http"
)

/*
 *   Get Movie IDs from company queries.
     Get Movie Details from Details query.
     Store in DB.
 *
*/
 
func main() {

    config.Init()
    generalConfigs := config.Get()

	url := "https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false&language=JP&page=1&sort_by=popularity.desc&with_companies="+generalConfigs.TMDB.Query.Studio_ghibli_id

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalf(err.Error())
    }

	req.Header.Add("accept", "application/json")
    // Watch out for space before Bearer!
	req.Header.Add("Authorization", "Bearer " + generalConfigs.TMDB.Api_reader_access_token)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatalf(err.Error())
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalf(err.Error())
    }

	fmt.Println(string(body))

}
