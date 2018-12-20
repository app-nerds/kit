# Server Stats

Server Stats provides a structure and Echo handler for tracking server statistics, such as when the server
started, number of requests, and the count of various status codes. Here is an example of using this in
an Echo application.

```golang
httpServer = echo.New()

/*
 * Server stats middleware
 */
serverStats := serverstats.NewServerStats(nil)
httpServer.Use(serverStats.Middleware)

httpServer.GET("/serverstats", serverStats.Handler)
```

The ServerStats object satifies the IServerStats interface, allowing developers to implement
their own statistics. It also provides the ability to provide a callback method to
add custom stats per application. This is done by passing a method to the NewServerStats
constructor method. Inside this method you can work with a map called **CustomStats**.
This is a map of string keys with interfaces for values.

```golang
httpServer = echo.New()

/*
 * Server stats middleware
 */
serverStats := serverstats.NewServerStats(func(ctx echo.Context, serverStats *ServerStats) {
	var ok bool
	var value int

	if _, ok = serverStats.CustomStats["someCounter"]; !ok {
		serverStats.CustomStats["someCounter"] = 0;
	}

	value = serverStats.CustomStats["someCounter"].(int)
	value++

	serverStats.CustomStats["someCounter"] = value
})
httpServer.Use(serverStats.Middleware)

httpServer.GET("/serverstats", serverStats.Handler)
```
