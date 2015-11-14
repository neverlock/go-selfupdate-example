# go-selfupdate-example

## At your server
- mkdir /var/www/auto-update/go-selfupdate-example

## At local
- install go-selfupdate from https://github.com/sanbornm/go-selfupdate
- go build go-selfupdate-example.go
- go-selfupdate go-selfupdate-example 1.0
- scp -r public/* user@yourdomain.com:/var/www/auto-update/go-selfupdate-example
- cp go-selfupdate-example version 1.0 run in other location
- edit your code 
- recompile your code with go build go-selfupdate-example.go
- go-selfupdate go-selfupdate-example 2.0
- scp -r public/* user@yourdomain.com:/var/www/auto-update/go-selfupdate-example
- stop your go-selfupdate-example version 1.0 that running and run it again
