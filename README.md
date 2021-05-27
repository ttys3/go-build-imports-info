# go-build-imports-info


tricks learned from <https://go-review.googlesource.com/c/tools/+/315852/4/internal/lsp/general.go#99>

see also <https://github.com/golang/go/issues/45732>


## demo output

### Build info

```shell
github.com/ttys3/go-build-imports-info master
    github.com/ttys3/go-build-imports-info@(devel)
    github.com/chai2010/gettext-go@v1.0.2 h1:1Lwwip6Q2QGsAdl/ZKPCwTe9fe0CjlUbqj5bFNSjIRk=
    github.com/iancoleman/strcase@v0.1.3 h1:dJBk1m2/qjL1twPLf68JND55vvivMupZ4wIzE8CTdBw=
    github.com/leonelquinteros/gotext@v1.5.0 h1:ODY7LzLpZWWSJdAHnzhreOr6cwLXTAmc914FOauSkBM=
    golang.org/x/text@v0.3.0 h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=
```
