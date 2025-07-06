## ✅ **List of Go Concurrency Patterns**

### 🔁 Core Patterns

1. **Fan-Out** – Multiple goroutines consuming from a single channel.
2. **Fan-In** – Multiple goroutines sending results to a single channel.
3. **Pipeline** – Data flows through stages (goroutines), each processing and forwarding.
4. **Worker Pool** – Fixed number of goroutines process jobs from a shared job channel.
5. **Producer-Consumer** – Producers send data to a channel; consumers read and process.

---

### ⏳ Timeout / Cancellation Patterns

6. **Timeout with `select` and `time.After`**
7. **Graceful Shutdown with `done` channel**
8. **Cancellation with `context.Context`** – Most idiomatic and scalable way.


---

### 📦 Coordination Patterns

9. **`sync.WaitGroup` for task synchronization**
10. **`sync.Mutex` / `sync.RWMutex` for mutual exclusion**
11. **`sync.Once` for one-time initialization**
12. **`sync.Cond` for conditional signaling between goroutines**

---

### 📊 Rate Limiting and Scheduling

13. **Ticker Pattern** – Use `time.Ticker` to trigger events periodically.
14. **Rate Limiting with `time.Tick`**
15. **Debounce / Throttle Pattern** – Control bursty events using time windows.

---

### 🧠 Advanced Patterns

16. **Or-Channel Pattern** – Combine multiple done channels to stop on *any* signal.
17. **Future / Promise Pattern** – Return a channel for async result.
18. **Publish-Subscribe (Pub/Sub)** – Simulate event distribution using channels.
19. **Fork-Join Pattern** – Split tasks, run in parallel, and join results.
20. **Event Loop / Dispatcher Pattern** – Central goroutine dispatches tasks or events.

---

### 📚 Utilities and Best Practices

21. **Select with default case** – Non-blocking operations.
22. **Channel Multiplexing** – Use `select` to read from multiple channels.
23. **Closing Channels Correctly** – Only the sender should close channels.
24. **Avoid Goroutine Leaks** – Always ensure goroutines can exit.
25. **Buffered vs Unbuffered Channels** – Use appropriately to manage backpressure.

Videos 
- https://youtube.com/playlist?list=PL7g1jYj15RUNqJStuwE9SCmeOKpgxC0HP&si=55tHEX1pgipSl6Yx
- https://youtube.com/playlist?list=PL5WZs2V9xUA351ALCSHSZ_P5WihHPjHTk&si=CP4OqAHXsT3oTt8P
- https://youtube.com/playlist?list=PL5WZs2V9xUA3hvgGAkVxfjdfy-x9hnfQ6&si=894KxXhXn4cdqFG6
- https://youtube.com/playlist?list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&si=RAT6jdFS44zjfW6P
- https://youtu.be/lbW-KVdIXaY?si=oBwwrZlw6SRccywR
- https://youtube.com/playlist?list=PLq3etM-zISamTauFTO5-G5dqBN07ckzTk&si=igluSkrsePKfoAcv
- https://youtube.com/playlist?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6&si=btSxXw-NazZ2Fn_Q
- https://youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&si=Mx_FbEbm48Hlx2CL
- https://youtube.com/playlist?list=PLKOhSssLmvQQPFv-QqbRZ10rn3UdBLvVC&si=knlWko8qFK0OX9jo
- https://youtube.com/playlist?list=PL8fnAiiuQeFucp_CokHM5rCipBubpyeoz&si=nl_kcUAlt8I_MEM1
- https://www.youtube.com/watch?v=xEmgcGHs3cA
- https://www.youtube.com/watch?v=7EK06n485nk
- https://www.youtube.com/watch?v=PPn_rgdx220
- https://www.youtube.com/watch?v=H7tbjKFSg58
- https://www.youtube.com/watch?v=mm4ztXwysLk
- [ADVANCED TUTORIAL](https://youtube.com/playlist?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&si=8_7OKNcd6j_7uRpB)

Golang Channels 
- https://www.youtube.com/@codeheim
- https://www.youtube.com/@mr_mux408
- https://www.youtube.com/@codeandlearnnow
- https://www.youtube.com/@kantancoding
- https://www.youtube.com/@FloWoelki
- https://www.youtube.com/@anthonygg_
- https://www.youtube.com/@MarioCarrion
- https://youtube.com/@golangcafe
- https://youtube.com/@akhilsharmatech


Sites 
- https://pkg.go.dev/sync
- https://gobyexample.com/
- https://go.dev/learn/
- https://golangforall.com/en/

