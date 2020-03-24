## Installation

Go 는 패키지 매니저가 없으며 VCS에서 직접 받아와서 라이브러리를 사용합니다.

먼저, 다음 페이지에서 window 용 Go 를 설치해줍니다.

[공식 다운로드 사이트](https://golang.org/dl/)

설치가 완료 되었다면. 아래와 같은 구조로 디렉터리를 만들어 줍니다.

```txt
Go/
  bin/
  pkg/
  src/
```

## Install package

```sh
go get github.com/tj/assert
```

## Build

You can build go project with `go build` command.

Also, you can specify runtime os with `GOOS` Environment variable like

`GOOS=windows go build -o main.exe main.go`
`GOOS=linux go build -o main.exe main.go`

## Execute

`go run main.go`

## Test

`go test -v <DIR>/<DIR>`
