package main

import (
	"fmt"
	"html/template"
	"os"
)

const url = "localhost:3004"

type Film struct {
	Title    string
	Director string
}
type Films []Film

type Films2 map[string][]struct {
	Director string
	Title    struct {
		Uk string
		Fr string
	}
}

var films2 = Films2{
	"Films": {
		{
			Director: "asd",
			Title: struct {
				Uk string
				Fr string
			}{
				Uk: "asd",
				Fr: "123",
			},
		},
	},
}

type ComplexFilmsTitle struct {
	Uk string
	Fr string
}
type ComplexFilms struct {
	Title ComplexFilmsTitle
}

func complexMapOfStructsTemplate() {
	var films = ComplexFilms{
		Title: ComplexFilmsTitle{
			Uk: "asd",
			Fr: "123",
		},
	}

	t, err := template.New("asd").Parse(`{{range .Films}}{{.Title}}{{end}}`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, films)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}

func simpleStructTemplate() {
	fmt.Println("simpleStructTemplate")

	type TData struct {
		Title       string
		ReleaseYear int
	}

	data := TData{Title: "Shawshank Redemption", ReleaseYear: 1991}

	tmpl, err := template.New("test").Parse("{{.Title}} was released on {{.ReleaseYear}}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func mapOfStructsTemplate() {
	films := map[string]Films{
		"Films": {
			{Title: "a", Director: "asd"},
			{Title: "b", Director: "asd"},
		},
	}
	t, err := template.New("asd").Parse(`{{range .Films}}{{.Title}}{{end}}`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, films)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}

func stringArrayTemplate() {
	salutationsSlice := []string{`hi`, `bye`, `cya`}
	templateString := `{{range .}}{{.}}

{{end}}`

	template1, err := template.New("asd").Parse(templateString)
	if err != nil {
		panic(err)
	}
	err = template1.Execute(os.Stdout, salutationsSlice)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}

func simpleMapOfStructsTemplate() {
	tournaments := []struct {
		Place string
		Date  string
	}{
		{Place: "Town1", Date: "01"},
		{Place: "Town2", Date: "02"},
		{Place: "Town3", Date: "03"},
	}
	t, err := template.New("").Parse(`{{range .}}
			{{.Place}}: {{.Date}}
    {{end}}`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, tournaments)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}

func main() {
	if false {
		simpleStructTemplate()
	}

	if false {
		mapOfStructsTemplate()
	}

	if true {
		stringArrayTemplate()
	}

	if false {
		simpleMapOfStructsTemplate()
	}
}
