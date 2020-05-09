# TinyGo
One file go web framework. Quickly build a web app by golang use TinyGo framework.

# How to use
Copy tiny folder to your project root, done.

# Samples
- GET
curl "http://localhost:10080/"
    r := tiny.New()
    r.GET("/", func(c *tiny.Context) {
        c.HTML(http.StatusOK, "<h1>Hello TinyGo</h1>")
    }
    r.Run(":10080")
 
