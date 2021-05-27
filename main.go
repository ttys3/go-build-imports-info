package main

import (
	"context"
	"fmt"
	"os"

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
	info.PrintVersionInfo(context.Background(), os.Stdout, true, info.Markdown)
}
