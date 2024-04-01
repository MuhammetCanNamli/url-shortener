# Project Title

URL Shortener

## Description

URL shortener coded using Golang

## Getting Started

### Dependencies

* Golang

### Executing program

* First, the Go module needs to be installed.
```
go mod init urlshortener
go get github.com/gorilla/mux
```
* Then, the project can be run via terminal/cmd.
```
go build url_shortener.go
main
```
or
```
go run url_shortener.go
```
* After the localhost server is started, we can shorten the URL we want to shorten via terminal/cmd with the following code:
```
curl -X POST -d "url=link" http://localhost:8080/shorten
```
* Finally, we can copy the URL created on the terminal/cmd screen and use it through the web browser.
</br>

## Project Features

* The project is a simple CLI based URL shortening application.
</br>

## Authors

[Muhammet Can Namli](https://www.linkedin.com/in/muhammet-can-naml%C4%B1-9556311b9/)

## Help
If you have suggestions and recommendations, please contact me.</br>
I would appreciate your support to make the project more modular.
