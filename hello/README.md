# hello
## 패키지
- 패키지명을 폴더명으로 사용
- 폴더가 다르면 패키지도 다름
- package main은 프로그램 시작점인 main()함수가 존재해야 함

## go 모듈
- 1.16이상부터 기본
```
go mod init hello
```

## go 빌드
- goos와 goarch 환경 변수를 조정해서 실행 파일 생성
```
go tool dist list
GOOS=linux GOARCH=amd64 go build 
```