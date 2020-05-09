# TinyGo
One file go web framework. Quickly build a web app by golang use TinyGo framework.

# How to use
Copy tiny folder to your project root, done.

# Samples
    func main() {
	r := tiny.New()

	// 1. Test GET
	// curl "http://localhost:10080/"
	r.GET("/", func(c *tiny.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyGo</h1>")
	})

	// 2. Test GET param
	// curl "http://localhost:10080/hello?name=Bill"
	r.GET("/hello", func(c *tiny.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	// 3. Test param in URL
	// curl "http://localhost:10080/hello/Bill"
	r.GET("/hello/:name", func(c *tiny.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	// 4. Test param with * match
	// curl "http://localhost:10080/assets/path1/path2/filename"
	r.GET("/assets/*filepath", func(c *tiny.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	// 5. Test POST param
	// curl -X POST -d "username=Bill&password=1234" "http://localhost:10080/login"
	r.POST("/login", func(c *tiny.Context) {
		c.JSON(http.StatusOK, tiny.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 6. Test static assets
	// curl "http://localhost:10080/images/image.jpeg"
	r.Static("/images", "./static")

	// 7. Test Group
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

	// Test plugin, print request time
	// curl "http://localhost:10080/v2/hello/Bill"
	r.AddPlugin(traceTime())

	r.Run(":10080")
    }

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
