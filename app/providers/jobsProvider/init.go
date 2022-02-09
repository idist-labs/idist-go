package jobsProvider

import (
	"fmt"
	"github.com/bamzi/jobrunner"
)

func Init() {
	fmt.Println("------------------------------------------------------------")
	jobrunner.Start()
}
