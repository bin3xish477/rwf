windows:
	GOOS=windows GOARCH=amd64 go build -o ./rwf.exe ./main.go
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./rwf.exe ./main.go
	tar -czvf rwf-windows-amd64.tar.gz ./rwf.exe ./README.md
	sha256sum rwf-windows-amd64.tar.gz > rwf-windows-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./rwf.exe ./main.go
	tar -czvf rwf-windows-386.tar.gz ./rwf.exe ./README.md
	sha256sum rwf-windows-386.tar.gz > rwf-windows-386.tar.gz.sha256sum

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./rwf ./main.go
darwin-release:
	GOOS=darwin GOARCH=amd64 go build -o ./rwf ./main.go
	tar -czvf rwf-darwin-amd64.tar.gz ./rwf ./README.md
	sha256sum rwf-darwin-amd64.tar.gz > rwf-darwin-amd64.tar.gz.sha256sum

linux:
	GOOS=linux GOARCH=amd64 go build -o ./rwf ./main.go
linux-release:
	GOOS=linux GOARCH=amd64 go build -o ./rwf ./main.go
	tar -czvf rwf-linux-amd64.tar.gz ./rwf ./README.md
	sha256sum rwf-linux-amd64.tar.gz > rwf-linux-amd64.tar.gz.sha256sum
	GOOS=linux GOARCH=386 go build -o ./rwf ./main.go
	tar -czvf rwf-linux-386.tar.gz ./rwf ./README.md
	sha256sum rwf-linux-386.tar.gz > rwf-linux-386.tar.gz.sha256sum

releases: linux-release windows-release darwin-release

clean:
	@if [ -f rwf ]; then rm rwf; fi
	@if [ -f rwf.exe ]; then rm rwf.exe; fi

