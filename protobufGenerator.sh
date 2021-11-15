protoc --go_out=./pkg ./proto/syncData/*.proto
protoc --go_out=./pkg ./proto/others/*.proto
protoc --go_out=./pkg ./proto/dataExtract/*.proto
protoc --go_out=./pkg ./proto/crazyApiService/*.proto
protoc --go_out=./pkg ./proto/crazyApiService/**/*.proto