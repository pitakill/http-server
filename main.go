package main

import (
	"fmt"
	"log"
	"net/http"
)

const html = `<!doctype html>
<html>
	<head>
		<title>Hola</title>
	</head>
	<body>
		<h1>Hola</h1>
	</body>
	<script>
		console.log("hola mundooooo!!!1");
		const h2 = document.createElement("h2");
		h2.innerText = "chí esto es dinámico";

		document.getElementsByTagName("body")[0].appendChild(h2);
	</script>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprint(w, html)
			return
		}

		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		for key, value := range r.Form {
			for i, name := range value {
				fmt.Fprintf(w, "Enviaste la siguiente información %s=%s (%d)\n", key, name, i)
			}
		}
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
