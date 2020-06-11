package profiling

//ADD IMPORT PROFILING AND THE FOLLOWING TO YOUR FUNCTION
// defer profiling.TimeTrack(time.Now(), "<funtionName>")

import (
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
