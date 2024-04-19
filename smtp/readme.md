
## Usage example

```
client, err := smtp.NewSMTP(
    smtp.NewConfig(
        smtp.WithHost("smtp.example.com"),
        smtp.WithUsername("username"),
        smtp.WithPassword("password"),
        smtp.WithPort(587),
    ),
)

if err != nil {
    // TODO
}

err = client.Send(
    "recipient@mail.com",
    "email subject",
    "email content",
    false // is HTML content type
)
```