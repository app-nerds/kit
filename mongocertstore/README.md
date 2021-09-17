# Mongo Cert Store

This package allows you to store autocerts (like Let's Encrypt) in a MongoDB database.

```go
certStore := mongocertstore.NewCertCache(db, "certs")
// Provide this to your favorite HTTP library for storing certificates
```
