# Tiny Go
TinyGo is a one file web framework for golang.
Simple to use, Quickly build a web app.

# How to use
Copy tiny folder to your project, Done.

# Demo 1 GET
Create main.go file:
    r := tiny.New()
    r.GET("/", func(c *tiny.Context) {
        c.HTML(http.StatusOK, "<h1>Hello TinyGo</h1>")
    })
    r.Run(":10080")
Test: 
curl "http://localhost:10080/"

# Demo 2 POST
