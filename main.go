package main

import (
	"fmt"
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

	url := "https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false&language=JP&page=1&sort_by=popularity.desc&with_companies="

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalf(err.Error())
    }

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIwYTFhN2VkMGQ4N2Y3YzAwZjM4OTNkOGU4YmZjMGZmMyIsInN1YiI6IjY1NjljM2Q2Y2Y0OGExMDEzYjk1NzBiZSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.j2dG9FMEd9braI7coiyM1XMyL3AO3NDRxH-VHsvnmSE")

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
