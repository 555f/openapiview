# Openapi View

Adds the ability to render openapi documentation.

## How to use it?

Example for echo:

```golang
package main

var (
	//go:embed openapi.yaml
    fs embed.FS
)

func main() {
    openapiData, _ := fs.ReadFile("openapi.yaml") // load openapi.yaml file
    ov := openapiview.NewMiddleware("/docs/api", openapiData) // create http middleware

    e := echo.New()
	e.HideBanner = true
	e.HidePort = true

    e.Use(echo.WrapMiddleware(ov.Process))

    if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
```