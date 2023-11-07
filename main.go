package main

import (
	template "html/template"
	os "os"
)

func htmxify(inputStr string, data any) error {
	tmpl, err := template.New("test").Parse(inputStr)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	return nil
}

func simpleStructTemplate() error {
	type TData struct {
		Title       string
		ReleaseYear int
	}

	data := TData{Title: "Shawshank Redemption", ReleaseYear: 1991}
	return htmxify("{{.Title}} was released on {{.ReleaseYear}}", data)
}

func arrayOfStringsTemplate() error {
	salutationsSlice := []string{`hi`, `bye`, `cya`}
	templateString := `{{range .}}{{.}}

{{end}}`

	return htmxify(templateString, salutationsSlice)
}

func structWithinStructTemplate() error {
	type FilmTitles struct {
		Uk string
		Fr string
	}
	type Films struct {
		Title FilmTitles
	}
	var films = Films{
		Title: FilmTitles{
			Uk: "asd",
			Fr: "123",
		},
	}

	const inputStr = `- Uk: {{.Title.Uk}}, Fr: {{.Title.Fr}}
`
	return htmxify(inputStr, films)
}

func arrayOfStructsTemplate() error {
	tournaments := []struct {
		Place string
		Date  string
	}{
		{Place: "Town1", Date: "01"},
		{Place: "Town2", Date: "02"},
		{Place: "Town3", Date: "03"},
	}

	templateString := `{{range .}}
		{{.Place}}: {{.Date}}
	{{end}}`

	return htmxify(templateString, tournaments)
}

func arrayofStructsWithinStructTemplate() error {
	type FilmTitles []struct {
		Uk string
		Fr string
	}
	type Films struct {
		Title FilmTitles
	}
	var films = Films{
		Title: FilmTitles{
			{Uk: "the", Fr: "le"},
			{Uk: "because", Fr: "parceque"},
		},
	}
	const templateString = `{{range .Title}}
- Uk: {{.Uk}}
- Fr: {{.Fr}}
{{end}}
	`

	return htmxify(templateString, films)
}

func structuredPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if false {
		structuredPanic(simpleStructTemplate())
	}

	if false {
		err := arrayOfStringsTemplate()
		if err != nil {
			panic(err)
		}
	}

	if false {
		err := structWithinStructTemplate()
		if err != nil {
			panic(err)
		}
	}

	if false {
		err := arrayOfStructsTemplate()
		if err != nil {
			panic(err)
		}
	}

	if !false {
		err := arrayofStructsWithinStructTemplate()
		if err != nil {
			panic(err)
		}
	}
}
