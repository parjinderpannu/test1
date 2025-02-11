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