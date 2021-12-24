package app

import (
	"github.com/osdifera/romano/controllers/lists"
	"github.com/osdifera/romano/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/lista", lists.GetList)
}
