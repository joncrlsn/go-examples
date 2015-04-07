package main

import (
	"os"
	templ "text/template"
)

type Context struct {
	People []Person
}

type Person struct {
	Name   string //exported field since it begins with a capital letter
	Senior bool
}

func main() {

	// With example 1
	t := templ.New("'With' Example 1")
	templ.Must(t.Parse(`
With inline Variable:{{with $3 := "hello"}} {{$3}} {{end}}, with an inline variable 
`))
	t.Execute(os.Stdout, nil)

	// With example 2
	t2 := templ.New("'With' Example 2")
	templ.Must(t2.Parse(`
Current Value:{{with "mom"}} {{.}} {{end}}, with the current value
`))
	t2.Execute(os.Stdout, nil)

	// If-Else example
	tIfElse := templ.New("If Else Example")
	ctx := Person{Name: "Mary", Senior: false}
	tIfElse = templ.Must(
		tIfElse.Parse(`
If-Else example:{{if eq $.Senior true }} Yes, {{.Name}} is a senior {{else}} No, {{.Name}} is not a senior {{end}}
`))
	tIfElse.Execute(os.Stdout, ctx)

	// Range example
	tRange := templ.New("Range Example")
	ctx2 := Context{People: []Person{Person{Name: "Mary", Senior: false}, Person{Name: "Joseph", Senior: true}}}
	tRange = templ.Must(
		tRange.Parse(`
Range example:
{{range $i, $x := $.People}} Name={{$x.Name}} Senior={{$x.Senior}}  
{{end}}

`))
	tRange.Execute(os.Stdout, ctx2)

}
