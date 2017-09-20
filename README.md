[![Build Status](https://travis-ci.org/israelb/curp.svg?branch=master)](https://travis-ci.org/israelb/curp)
[![Coverage Status](https://coveralls.io/repos/github/israelb/curp/badge.svg?branch=master)](https://coveralls.io/github/israelb/curp?branch=master)

# CURP

This is a little library for getting a CURP.

# Instalation

```shell
# download the go-apex package
go get -v github.com/apex/go-apex
```

# How to use it
You need to use the public method: NewCurp

### Example: 

```go
NewCurp(name, firstLastName, secondLastName, sex, stateCode, birthDate)
```

## How the CURP is building

![CURP](https://github.com/israelb/curp/blob/master/assets/curp-example.png)
