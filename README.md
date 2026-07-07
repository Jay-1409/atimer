# atimer

> A highly performant, sharded clock node capable of receiving multiple timer requests and firing asynchronous callbacks when they expire.

![License](https://img.shields.io/badge/license-Apache2.0-blue)
![Language](https://img.shields.io/badge/language-Go1.25-blue)

---

## What's in there for you
- **Lock-Free Routing**: Uses atomic CPU instructions (`sync/atomic`) to distribute incoming tasks to multiple timer heaps using round-robin scheduling, minimizing lock contention.
- **CPU-Efficient Polling**: Avoids CPU-intensive busy waiting by using Go's timer channel scheduler. The idle compute usage is close to `0%`.
- **Dedicated Heap Worker Pools**: Each sharded heap runs its own dedicated and configurable worker pool (`TimerEventHandler`) to execute asynchronous HTTP POST callbacks concurrently.
- **Simple HTTP API**: A lightweight HTTP API endpoint `/api` that accepts URL form-encoded inputs and schedules tasks instantly.

---

## Features

* **Sharded Min-Heaps**: Divides the scheduling load across multiple, isolated min-heaps to prevent thread bottlenecking.
* **Smart Sleep & Wake**: Timers sleep efficiently and wake up instantly if a higher-priority task (with a closer execution time) is added to the queue.
* **Robust Workers**: Handles high-concurrency event dispatching with a configurable worker pool size. If callbacks fail, the server logs and proceeds rather than crashing.

---

## Quick Start Guide

### 1. Build and Run
Build the project using standard Go compiler tools:

```bash
go build ./cmd/atimer
```

Start the server using CLI flags to configure the port, number of heaps, and worker threads per heap:

```bash
./atimer -port 8080 -heaps 4 -workers 2
```

### 2. Schedule a Task
Send a `POST` request with form values to the `/api` endpoint:

```bash
curl -X POST http://localhost:8080/api \
  -d "id=task123" \
  -d "timer_time=10" \
  -d "callback_url=http://example.com/callback"
```

---

## Documentation

- API Endpoint reference [here](docs/endpoints.md)
- Timer Architecture details [here](docs/architecture.md)
- Bruno collection [here](bruno/)
- System whiteboard drawings (excalidraw) [here](docs/whiteboard.excalidraw)
- Just a journal about this project [here](docs/journal.md)
---

## 📄 License

This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for details.
