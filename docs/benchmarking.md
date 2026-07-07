# Benchmarking Guide

You can run benchmarks locally to measure task scheduling and callback delivery performance.

## Prerequisites
* Install Go 1.25+
* Install standard HTTP load generators (e.g. `wrk`, `hey`, or `vegeta`).

## Load Testing HTTP Server
Run the `atimer` binary on a target port:
```bash
./atimer -heaps 4 -workers 2 -port 8080
```

Use `wrk` to send POST requests simulating high concurrency:
```bash
wrk -t4 -c100 -d30s -s benchmark.lua http://localhost:8080/api
```

## Running Benchmarks locally
You can run automated Go benchmark tests targeting the memory heap directly:
```bash
go test -bench=. ./internals/timer/...
```
