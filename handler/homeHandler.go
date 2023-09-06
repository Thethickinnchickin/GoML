package handler

import (
    "net/http"
    "html/template"

)

func HomeHandler(w http.ResponseWriter, r *http.Request) {


	// Render the HTML response
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Home</title>
	</head>
	<body>
		<h1>Home Page</h1>
	</body>
	</html>
	`

	t, err := template.New("registration-success").Parse(tmpl)
	if err != nil {
		http.Error(w, "Failed to render HTML template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to render HTML template", http.StatusInternalServerError)
		return
	}
}
