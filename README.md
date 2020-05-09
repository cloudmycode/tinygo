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
```
r := tiny.New()
r.GET("/hello", func(c *tiny.Context) {
    c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
})
r.Run(":10080")
```
### GET param in URL path
curl "http://localhost:10080/hello/Bill"
```
r := tiny.New()
r.GET("/hello/:name", func(c *tiny.Context) {
    c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
})
r.Run(":10080")
```
### GET param with * match
curl "http://localhost:10080/assets/path1/path2/filename"
```
r := tiny.New()
r.GET("/assets/*filepath", func(c *tiny.Context) {
    c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
})
r.Run(":10080")
```
### POST param
curl -X POST -d "username=Bill&password=1234" "http://localhost:10080/login"
```
r := tiny.New()
r.POST("/login", func(c *tiny.Context) {
    c.JSON(http.StatusOK, tiny.H{
        "username": c.PostForm("username"),
        "password": c.PostForm("password"),
    })
})
r.Run(":10080")
```
### Static assets
curl "http://localhost:10080/images/image.jpeg"
```
r := tiny.New()
r.Static("/images", "./static")
r.Run(":10080")
```
### Path Group
```
r := tiny.New()
v1 := r.Group("/v1")
{
	// curl "http://localhost:10080/v1/"
	v1.GET("/", func(c *tiny.Context) {
        c.HTML(http.StatusOK, "<h1>Hello Tiny Group v1</h1>")
    })

    // curl "http://localhost:10080/v1/hello"
    v1.GET("/hello", func(c *tiny.Context) {
        c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
    })
}

v2 := r.Group("/v2")
{
    // curl "http://localhost:10080/v2/hello/Bill"
    v2.GET("/hello/:name", func(c *tiny.Context) {
        c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
    })

    // curl -X POST -d "username=Bill&password=1234" "http://localhost:10080/v2/login"
    v2.POST("/login", func(c *tiny.Context) {
        c.JSON(http.StatusOK, tiny.H{
        "username": c.PostForm("username"),
        "password": c.PostForm("password"),
        })
	})
}
r.Run(":10080")
```
### Plugin
curl "http://localhost:10080/v2/hello/Bill"
```
r := tiny.New()
r.AddPlugin(traceTime())
r.Run(":10080")

// Print request time
func traceTime() tiny.HandlerFunc {
	return func(c *tiny.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
```
