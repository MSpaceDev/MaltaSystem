rm -rf ../build
mkdir -p ../build
cd ../cmd/server
go build .
mv ./server.exe ../../build
cd ../../build
./server.exe -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password=1234 -db-schema=ye_be
