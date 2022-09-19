cd ../

 rem CGO_ENABLED=1  GOOS=windows GOARCH=amd64   CC=x86_64-w64-mingw32-gcc go build

rem CGO_ENABLED=1  GOOS=linux GOARCH=amd64   CC=x86_64-w64-mingw32-gcc go build

 set CGO_ENABLED=1
 set GOOS=windows
 set GOARCH=amd64
 set CC=x86_64-w64-mingw32-gcc
 call go build

 set CGO_ENABLED=1 
 set GOOS=linux
 set GOARCH=amd64
 set CC=x86_64-w64-mingw32-gcc
 call go build