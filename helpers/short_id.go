package helpers

import (
	"github.com/teris-io/shortid"
	"strings"
)

func NewShortID() string {
	var id string
	for {
		sid, err := shortid.New(0, shortid.DefaultABC, 1)
		if err == nil {
			if id, err = sid.Generate(); err == nil && !(strings.Contains(id, "-") || strings.Contains(id, "_")) {
				break
			}
		}
	}
	return id
}
