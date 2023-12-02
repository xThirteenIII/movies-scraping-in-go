package config

import "ghibli/pkg/parser"


var tmdb tmbd_api_access

var fileNames = map[string]any{
    "tmdb.conf": &tmdb,
}

// The config variable can be accessed by the getter Get
var config projectConfig

func Init() {

    // Iterate through fileNames Map and decode each config file into the correspondig
    // variable
    for file, container := range fileNames {
       parser.JSON("pkg/config/test/"+file, container)
    }

    // Assign that variable to the general projectConfig field
    config = projectConfig{
        TMDB: tmdb,
    }
}

// Config getter
func Get() projectConfig{
    return config
}
