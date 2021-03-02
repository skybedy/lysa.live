package routes

import (
	"net/http"

	"lysa.live/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)
}
