package main

import (
	"fmt"

	"github.com/ttys3/go-build-imports-info/info"

	"github.com/chai2010/gettext-go"
	"github.com/chai2010/gettext-go/examples/hi"
	"github.com/iancoleman/strcase"
	_ "github.com/leonelquinteros/gotext"
)

func init() {
	fmt.Println("=== main.init: default ===")

	// bind app domain
	gettext.BindLocale(gettext.New("hello", "locale"))

	// $(LC_MESSAGES) or $(LANG) or empty
	fmt.Println(gettext.Gettext("Gettext in init."))
	fmt.Println(gettext.PGettext("main.init", "Gettext in init."))
	hi.SayHi()

	// Output(depends on locale environment):
	// ?
	// ?
	// ?
	// ?

	fmt.Println(strcase.ToSnake("GoStruct"))
}

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("go packages this program imports: ")
	versionInfo := info.VersionInfo()

	// tricks learned from https://go-review.googlesource.com/c/tools/+/315852/4/internal/lsp/general.go#99
	// https://github.com/golang/go/issues/45732
	for _, dep := range versionInfo.Deps {
		fmt.Printf("%#v\n", dep)
	}
	fmt.Println("-----------------------------------------")
}
