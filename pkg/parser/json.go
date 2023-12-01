package parser

import (
	"encoding/json"
	"log"
	"os"
)

// JSON parses the file specified with the path variable and parse its content into a structure.
// The mask argument allows JSONfile to accept any kind of structure and decode freely the content into it.
// any and interface{} types are equivalent
func JSON(path_to_file string, mask any) {

    // Try to open file
    if file, err := os.Open(path_to_file); err != nil {

        log.Fatalln("While opening file: " + path_to_file + "\n" + err.Error())
    }else {

        // Try to decode the content of the file into the provided struct
        if err = json.NewDecoder(file).Decode(&mask); err != nil {

            log.Fatalln("While decoding file: " + path_to_file + "\n" + err.Error())
        }
    }
}
