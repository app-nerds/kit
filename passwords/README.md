# Passwords

This package provides basic tools for hashing password.

## Hashing

To create a hashed password:

```go
hashedPassword, err := passwords.HashPassword("password")
```

## Validation

To validate a plaintext password matches your hashed password:

```go
matches := passwords.IsPasswordValid(hashedPassword, "password")
```

## Custom Hashed Password Type

This package also provides a custom type, **HashedPasswordString** which provides a convienent way of working with passwords that you wish to hash and verify. As an example, here is a struct that contains a password field that you may wish to hash prior to saving in a database.

```go
type MyStruct struct {
	Name string                             `json:"name"`
	Password passwords.HashedPasswordString `json:"password"`
}

// Let's pretent we have a user-submitted form that has populated this struct...
u := MyStruct{
	Name: "Adam",
	Password: "password",
}

// Save to imaginary database...
name := u.Name
password := u.Password.Hash()
```
