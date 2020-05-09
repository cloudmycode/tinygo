# TinyGo
One file go web framework. Quickly build a web app by golang use TinyGo framework.

## How to use
Copy tiny folder to your project root, done.

## Samples
### GET
curl "http://localhost:10080/"
```
r := tiny.New()
r.GET("/", func(c *tiny.Context) {
    c.HTML(http.StatusOK, "<h1>Hello TinyGo</h1>")
})
r.Run(":10080")
```
### GET param
curl "http://localhost:10080/hello?name=Bill"

return hello Bill, you're at /hello
```
r := tiny.New()
r.GET("/hello", func(c *tiny.Context) {
    c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
})
r.Run(":10080")
```
