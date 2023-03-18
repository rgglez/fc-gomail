D=sendmail
L=sendmail

go:
	cd src/sendmail/src; GOOS=linux CGO_ENABLED=0 go build -o handler *.go
