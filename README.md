# Rest Api Generator

**rest-api-generator** generates standard structure rest api application using module [flt-go-router](https://github.com/fajarardiyanto/flt-go-router).

## Getting started

### Flags
|  flag   |                     description                     |
|:-------:|:---------------------------------------------------:|
|   pkg   | name package (example: github.com/testing/pkg-name) |
| appName |      name application (example: testing-name)       |
| modules |           embed modules (example: mysql)            |

### Modules Register
| Name  |
|:-----:|
| mysql |
| redis |
| mongo |

### Usage
move file **rest-api-generator** to your folder project or folder /bin

```shell
./rest-api-generator --pkg testing --appName testing-name --modules mysql,redis,mongo
```
or
```shell
./rest-api-generator --pkg testing --appName testing-name --modules mysql --modules redis --modules mongo
```
<br />

```shell
cd testing && make run
```

or see Makefile

### Todo
- [ ] Select Framework (gin, echo, chi, native-mux)