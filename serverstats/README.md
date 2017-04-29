# Server Stats

Server Stats provides a structure and Echo handler for tracking server statistics, such as when the server
started, number of requests, and the count of various status codes. Here is an example of using this in
an Echo application.

```
httpServer = echo.New()

/*
 * Server stats middleware
 */
serverStats := serverstats.NewServerStats()
httpServer.Use(serverStats.Middleware)

httpServer.GET("/serverstats", serverStats.Handler)
```
