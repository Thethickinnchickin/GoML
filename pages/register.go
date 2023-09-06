package pages

import (
    "net/http"
    "html/template"
	//"fmt"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
    
	// Render the HTML registration form
	//fmt.Print("Hey Matt")
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Registration</title>
	</head>
	<body>
		<h1>Registration</h1>
		<form method="post" action="http://localhost:8080/register-user">
			<label for="username">Username:</label>
			<input type="text" id="username" name="username" required><br>
			<label for="password">Password:</label>
			<input type="password" id="password" name="password" required><br>
			<input type="submit" value="Register">
		</form>
	</body>
	</html>
	`

	t, err := template.New("registration-form").Parse(tmpl)
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
