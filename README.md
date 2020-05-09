# TinyGo
One file go web framework. Quickly build a web app by golang use TinyGo framework.

# How to use
Copy tiny folder to your project root, done.

# Sample 1 GET request
	r := tiny.New()
	r.GET("/", func(c *tiny.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyGo</h1>")
	})
 	r.Run(":10080")
Test:
curl "http://localhost:10080/"
