https://www.353solutions.com/c/znga/
https://x.com/jordancurve/status/1108475342468120576

go mod init znga
go build 
go run hw.go
go mod tidy # Removes unnecessary dependencies and adds missing ones.

<!-- 
    Go executable is big in size because 
    -- it is not just your code
    -- go run time
    -- garbage collector 
    -- scheduler 
    -- runtime information
-->
<!-- you can build exec for mac, linux by changing GOARCH GOOS -->

go env GOARCH GOOS
GOOS=linux go build
GOARCH=amd64 GOOS=linux go build

curl -i https://api.github.com/users/parjinderpannu
curl -i -H 'User-Agent: Go' https://api.github.com/ousers/parjinderpannu

cat http.log.gz| gunzip |sha1sum