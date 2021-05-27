# go-build-imports-info


tricks learned from <https://go-review.googlesource.com/c/tools/+/315852/4/internal/lsp/general.go#99>

see also <https://github.com/golang/go/issues/45732>


## demo output

```shell
go run .
-----------------------------------------
go packages this program imports:
&info.Module{ModuleVersion:info.ModuleVersion{Path:"github.com/chai2010/gettext-go", Version:"v1.0.2", Sum:"h1:1Lwwip6Q2QGsAdl/ZKPCwTe9fe0CjlUbqj5bFNSjIRk="}, Replace:(*info.ModuleVersion)(nil)}
&info.Module{ModuleVersion:info.ModuleVersion{Path:"github.com/iancoleman/strcase", Version:"v0.1.3", Sum:"h1:dJBk1m2/qjL1twPLf68JND55vvivMupZ4wIzE8CTdBw="}, Replace:(*info.ModuleVersion)(nil)}
&info.Module{ModuleVersion:info.ModuleVersion{Path:"github.com/leonelquinteros/gotext", Version:"v1.5.0", Sum:"h1:ODY7LzLpZWWSJdAHnzhreOr6cwLXTAmc914FOauSkBM="}, Replace:(*info.ModuleVersion)(nil)}
&info.Module{ModuleVersion:info.ModuleVersion{Path:"golang.org/x/text", Version:"v0.3.0", Sum:"h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg="}, Replace:(*info.ModuleVersion)(nil)}
-----------------------------------------
```
