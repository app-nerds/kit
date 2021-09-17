# MailService

MailService provides the basic ability to send email.

### Example

```go
service := email.NewMailService(email.Config{
	Host: "mail.something.com",
	Password: "password",
	Port: 25,
	UserName: "user",
})

if err = service.Connect(); err != nil {
	// Handle error
}

mail := email.Mail{
	Body: "This is an example",
	From: email.Person{
		Name: "Adam",
		EmailAddress: "test@test.com",
	},
	Subject: "This is a sample",
	To: []Person{
		{
			Name: "Bob Hope",
			EmailAddress: "address1@test.com",
		},
		{
			Name: "Elvis Presley",
			EmailAddress: "address2@test.com",
		},
	},
}

if err = service.Send(mail); err != nil {
	// Handle error
}
```

### Validating Email Address

```go
isValid = email.IsValidEmailAddress("whatever")
// isValid == false
```
