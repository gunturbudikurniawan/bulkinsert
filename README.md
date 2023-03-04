user mengirimkan request dengan 1000 record
logical process di service a membutuhkan 10ms/record
user harus mendapatkan response time dalam 1s

request payload json :
{
"request_id":123456,
"data":[{
"id":123456,
"customer":"John Doe",
"quantity":1,
"price":10.00,
"timestamps":"2022-01-01 22:10:4444"
},...,...
]
}

seluruh data harus berhasil dimasukkan dalam database
service dibuat menggunakan golang
data transaksi disimpan ke dalam database
jika diperlukan boleh membuat lebih dari 1 service
jika diperlukan boleh menggunakan lebih dari 1 database

docker.for.mac.localhost
docker-compose -f mysqlCore.yml up
docker container start mysqlCore
docker container stop mysqlCore

Run the below command to get the benchmark for speed and memory performance:
go test -benchmem -run=^$ bpjs -bench BenchmarkOrmCreate

Run the test function with the command below:
go test -benchmem -run=^$ bpjs -bench BenchmarkCreate

Run the test function with the command below:
go test -benchmem -run=^$ bpjs -bench BenchmarkBulkCreate

Test benchmarks with a different batch size by running the command below:
go test -benchmem -run=^$ bpjs -bench BenchmarkBulkCreateSize

type Response struct {
RequestID int `json:"request_id"`
Data []struct {
ID int `json:"id"`
Customer string `json:"customer"`
Quantity int `json:"quantity"`
Price float64 `json:"price"`
Timestamps string `json:"timestamps"`
} `json:"data"`
}

first
docker-compose -f rabbit.yml up
docker-compose -f mysqlCore.yml up
docker-compose -f adminerCore.yml up

if you have existing
docker container start mysqlCore
