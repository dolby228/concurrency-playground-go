# concurrency-playground-go

Мини‑сборник примеров конкуренции в Go: горутины, каналы, fan-in, worker pool, pipeline, sync.Once, sync.Mutex, sync.WaitGroup, atomic.

## Запуск

- go run ./cmd/fanin
- go run ./cmd/workers
- go run ./cmd/onceinit
- go run ./cmd/mutexcounter
- go run ./cmd/wgpipe
- go run ./cmd/atomiccounter

## Тесты

- go test ./...
- go test -race ./...
