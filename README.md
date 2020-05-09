# tinygo
One file go web framework.
Simple to use, Quickly build a web app by golang.

# Start
Copy tiny folder to your project.
Finish.

# How to use

# GET request
    r := tiny.New()
    r.GET("/", func(c *tiny.Context) {
        c.HTML(http.StatusOK, "<h1>Hello TinyGo</h1>")
    })
    r.Run(":10080")
