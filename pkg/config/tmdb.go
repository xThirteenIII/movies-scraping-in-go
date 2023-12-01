package config

// struct fields must be Capitalized to be accessed elsewhere
type tmbd_api_access struct {
    Api_key string  `json:"api_key"` 
    Api_reader_access_token string `json:"Api_reader_access_token"`
    Query []query_fields `json:"query"`
}

type query_fields struct {
    Walt_disney_production_id string `json:"Walt_disney_production_id"`
    Studio_ghibli_id string `json:"studio_ghibli_id"`
}
