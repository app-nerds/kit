# ENV Package

The **env** package provides methods for working with the OS environment. Go already provides a
rich set of tools for this purpose, and these methods simply add some convience.

## Package Methods

Below is a reference of exported methods in this package.

### Getenv(key, defaultValue string) string
```Getenv``` takes a key name, and a default value and returns a string. If there is an
environment variable in the OS with this key name then that value is returned. Otherwise
the default value is returned.
