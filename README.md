# Base64

Encode decode base64

Simple CLI tool to encode and decode base64
Useful if you use kubernetes a lot ;)

[![codecov](https://codecov.io/gh/samit22/calendarN/branch/main/graph/badge.svg?token=A5XND1948Y)](https://codecov.io/gh/samit22/calendarN)
[![goreport](https://goreportcard.com/badge/github.com/samit22/calendarN)](https://goreportcard.com/report/github.com/samit22/calendarN)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=bugs)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_calendarN&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=samit22_calendarN)

### Installation

```
go install github.com/samit22/b64@latest
```

### Commands

- To encode a string to base64 format

  ```
  b64 encode samit
  ```

- To decode base64 string

  ```
  b64 decode c2FtaXQ=
  ```

  The generated string is copies to the clipboard, ready to paste!

### Requirement

- Go 1.18+

### Contributing
