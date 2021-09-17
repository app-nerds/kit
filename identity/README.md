# Identity

This package details with common identity-related issues, such as logging in and generating JWT tokens. It helps make the mundane tasks of creating and validating authentication tokens simpler. Here's a sample.

## JWTService

```go
package main

import (
   "github.com/app-nerds/kit/v6/identity"
   "github.com/golang-jwt/jwt"
)

func main() {
   jwtService := identity.NewJWTService(identity.JWTServiceConfig{
      AuthSalt: "salt",
      AuthSecret: "secret",
      Issuer: "issuer://com.some.domain",
      TimeoutInMinutes: 60,
   })

   // Create a token
   token, _ := jwtService.CreateToken(identity.CreateTokenRequest{
      UserID: "user",
      UserName: "My Name",
      AdditionalData: map[string]interface{}{
         "key": "value",
      },
   })

   // token == base64-encoded, encrypted JWT token

   // Parse an incoming token. The result is *jwt.Token and can be 
   // manipulated using the jwt-go library
   parsedToken, _ := jwtService.ParseToken(token)

   userID, userName := jwtService.GetUserFromToken(parsedToken)
   // userID == "user"
   // userName == "My Name"

   additionalData, _ := jwtService.GetAdditionalDataFromToken(parsedToken)
   // additionalData == map[string]interface{}{
   //    "key": "value"
   // }
}
```
