## Тестирование
```
git clone git@github.com:StarNuik/golang_tenders.git
cd golang_tenders
./test/compose.sh build .
./test/compose.sh up (-d)
go test ./... (-v -count 1)
```

## Запуск (пример)
```
service-tenders:
  build: .
  restart: unless-stopped
  ports:
  - 8080:8080
  environment:
    SERVER_ADDRESS: "0.0.0.0:8080"
    POSTGRES_CONN: "postgres://user:insecure@localhost:5432/"
```
