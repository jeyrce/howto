package main

import (
	"os"
	tt "text/template"
	"time"
)

type Data struct {
	Desc string
	X    int
}

type Param struct {
	Name string
	Time string
	Data Data
}

const (
	Tmpl = `
Name: {{.Name}}
Time: {{.Time}}
Data_X: {{.Data.X}}
Data_Desc: {{.Data.Desc}}
`
)

func main() {
	template, err := tt.New("tmpl").Parse(Tmpl)
	if err != nil {
		panic(err)
	}
	err = template.Execute(os.Stdout, Param{
		Name: "jeyrcelu",
		Time: time.Now().Local().Format(time.DateTime),
		Data: Data{
			Desc: "jeyrcelu test",
			X:    1,
		},
	})
	if err != nil {
		panic(err)
	}
}
