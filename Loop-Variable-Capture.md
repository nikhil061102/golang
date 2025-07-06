Ah, the classic Go confusion!
You asked:

> **Why do we need `ch := ch` inside the loop in `fanIn`?**
> Isn't `ch` already defined in the `for _, ch := range inputChans`?

---

## 🔍 The Reason: **Loop Variable Capture in Goroutines**

### TL;DR:

`ch := ch` inside the loop creates a **new local copy** of `ch`, which **captures the correct channel** for each goroutine.

Without it, **all goroutines might close over the same `ch`** — the **last one** in the loop — causing wrong behavior.

---

## 📌 What Happens Without `ch := ch`?

Let's say `inputChans = [chan1, chan2, chan3]`.

You write:

```go
for _, ch := range inputChans {
	go func() {
		fmt.Println(<-ch)
	}()
}
```

**Wrong behavior**: All goroutines might read from `chan3` (the final value of `ch`)
This is because `ch` is **shared** across iterations, and the goroutines capture **its pointer**, not its value at each point in time.

---

## ✅ Correct Way — Shadowing the Variable

```go
for _, ch := range inputChans {
	ch := ch // ✅ create a new scoped copy
	go func() {
		fmt.Println(<-ch)
	}()
}
```

Now:

* Each goroutine closes over its **own unique `ch`**.
* Each one reads from the **correct input channel**.

---

## 🧠 Deeper Insight: Closures in Go

* Loop variables in Go are **reused** for each iteration.
* Goroutines run asynchronously — so they often execute **after** the loop has moved on.
* Capturing the loop variable directly means you’re capturing its **future value**, not its current one.

---

## ✅ Summary

| Code                     | Behavior                           |
| ------------------------ | ---------------------------------- |
| `for _, ch := range ...` | Iterates normally                  |
| `go func() { ... }()`    | Captures `ch` by reference         |
| `ch := ch`               | Shadows `ch` to capture by value ✅ |

This trick is essential in Go whenever you launch goroutines inside a loop.

---

Let me know if you'd like a visual example or to test this live in the Go playground!
