# Base64

Encode decode base64

Simple CLI tool to encode and decode base64
Useful if you use kubernetes a lot ;)

[![codecov](https://codecov.io/gh/samit22/b64/branch/main/graph/badge.svg?token=OCKP3PX85G)](https://codecov.io/gh/samit22/b64)
[![goreport](https://goreportcard.com/badge/github.com/samit22/b64)](https://goreportcard.com/report/github.com/samit22/b64)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_b64&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=samit22_b64)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=samit22_b64&metric=bugs)](https://sonarcloud.io/summary/new_code?id=samit22_b64)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=samit22_b64&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=samit22_b64)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=samit22_b64&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=samit22_b64)

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
