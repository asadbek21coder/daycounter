package handler

import (
	"fmt"
	"net/http"
	"time"
)

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	target := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	remain := int(target.Sub(time.Now()).Hours() / 24)

	temp := `<!DOCTYPE html>
	<html>
	<head>
		<title>Days Remaining</title>
	</head>
	<style>
	html, body, h1, h2 {
	  margin: 0;
	  padding: 0;
	}
	
	.header {
	  background-color: #333;
	  color: #fff;
	  padding: 10px;
	  text-align: center;
	}
	.container {
	  max-width: 300px;
	  margin: 0 auto;
	  padding: 100px;
	}
	
	</style>
	<body>
			<h1 class="header">Days Remaining Until January 1, 2025</h1>
			<h2 class="container">There are <b><i style="font-size: xx-large;">` + fmt.Sprint(remain) + `</i></b> days left.</h2>
	</body>
	</html>`

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, temp)
}
