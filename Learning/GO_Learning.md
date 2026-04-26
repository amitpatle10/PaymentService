# GO_Learning — My Personal Go Reference

> Built concept by concept as I build the PaymentService.
> Every entry here is something I understood, ran, and verified myself.

---

## How to Read This File

Each concept has:

- **What it is** — plain English
- **Real-world analogy** — something familiar
- **Code snippet** — from this project
- **First seen in** — which phase / file

New entries are appended at the bottom as we learn them.

---

<!-- Concepts will be appended here as they are learned -->

---

## Concept 3 — Go Modules & go.mod

**What it is:**
A Go module is a self-contained project with a unique name, a Go version, and a list of dependencies. It is defined by a `go.mod` file at the root of the project.

**Real-world analogy:**
The module is your apartment. `go.mod` is the address plate on the door — it tells the world what this place is called and what version of Go it runs on.

**Code snippet:**
```
module github.com/amitpatle/paymentservice

go 1.26.2
```

**How to create it:**
```bash
go mod init github.com/amitpatle/paymentservice
```

**First seen in:** Phase 1 — Project Setup

---

## Concept 4 — Packages

**What it is:**
A package is a folder of `.go` files that belong together. Every `.go` file must declare which package it belongs to at the top with `package <name>`.

`package main` is special — it is the entry point of any runnable Go program. It must contain a `func main()`.

**Real-world analogy:**
If the module is the apartment, packages are the rooms — each room has a specific purpose (kitchen, bedroom). `package main` is the front door — execution always starts there.

**Code snippet:**
```go
package main

import "fmt"

func main() {
    fmt.Println("PaymentService starting...")
}
```

**First seen in:** Phase 1 — main.go

---

## Concept 5 — Structs

**What it is:**
A struct is a way to group related data together under one name. It is Go's replacement for classes (Go has no classes).

**Real-world analogy:**
A payment receipt — it has multiple fields (amount, date, merchant, status) all on one piece of paper. A struct is that receipt as a Go type.

**Code snippet:**
```go
type Payment struct {
    ID        string
    UserID    string
    Amount    float64
    Currency  string
    Provider  string
    Status    PaymentStatus
    CreatedAt time.Time
}
```

Creating a struct (struct literal):
```go
p := models.Payment{
    ID:     "pay_001",
    Amount: 999.99,
    Status: models.StatusPending,
}
```

Any field not set gets a **zero value** — `""` for string, `0` for numbers, `false` for bool. Go never leaves memory uninitialized.

**First seen in:** Phase 2 — models/models.go

---

## Concept 6 — Custom Types & Constants

**What it is:**
You can create a new named type based on an existing type. Combined with `const`, this replaces magic strings with safe, readable names.

**Real-world analogy:**
Instead of writing "red", "green", "yellow" as raw strings on traffic signals everywhere, you define named constants — `Red`, `Green`, `Yellow`. No typos, no guessing.

**Code snippet:**
```go
type PaymentStatus string   // new type based on string

const (
    StatusPending PaymentStatus = "pending"
    StatusSuccess PaymentStatus = "success"
    StatusFailed  PaymentStatus = "failed"
)
```

**First seen in:** Phase 2 — models/models.go

---

## Concept 7 — Short Variable Declaration (:=)

**What it is:**
`:=` creates a new variable and infers its type automatically from the right-hand side. No need to declare the type explicitly.

**Real-world analogy:**
Instead of saying "I have a box of type `string`, and I'm putting `hello` in it" — you just say "here's `hello`, figure out the box type yourself."

**Code snippet:**
```go
p := models.Payment{...}   // Go infers type as models.Payment
x := 42                    // Go infers type as int
```

vs the long form (rarely used):
```go
var p models.Payment = models.Payment{...}
```

**First seen in:** Phase 2 — main.go

---

## Concept 1 — How Go Compiles (vs Java)

**What it is:**
Go compiles your source code directly into a native machine code binary. The CPU understands it directly — no middleman needed at runtime.

Java takes a two-step path: `javac` compiles your `.java` files into **bytecode** (an intermediate form, not machine code), and then the **JVM** reads and runs that bytecode at runtime.

**Real-world analogy:**
- Java is like writing a recipe in French, then hiring a translator (JVM) to read it aloud in English every time you cook.
- Go is like writing the recipe directly in English — no translator needed, just cook.

**The flow:**

```
Java:
  .java → javac → .class (bytecode) → JVM interprets/JIT-compiles → runs on hardware
  (JVM must be installed on the machine)

Go:
  .go → go build → native binary (.exe on Windows, no extension on Mac/Linux) → runs directly on hardware
  (nothing needs to be installed — binary is self-contained)
```

**Key difference:**
| | Java | Go |
|---|---|---|
| Compiled to | Bytecode (needs JVM) | Machine code (runs directly) |
| Runtime dependency | JVM must be installed | Nothing — binary includes everything |
| Portability | Same bytecode, any OS with JVM | Separate binary per OS, zero dependency |

**First seen in:** Pre-Phase 1 discussion (before project setup)

---

## Concept 2 — Garbage Collection in Go

**What it is:**
Garbage collection (GC) is the automatic process of finding memory your program no longer uses and freeing it — so you don't run out of RAM.

Without GC, you'd have to manually free every variable you create. Forget once → memory leak → program crashes or slows down.

**Real-world analogy:**
A janitor in an office. You leave empty coffee cups around. The janitor figures out which cups are abandoned (no one is going to drink from them) and throws them away. You never have to call the janitor — they just work in the background.

**How Go's GC works — Mark and Sweep:**

```
Step 1 — MARK
  Start from "roots" (variables currently in use — global vars, active function vars)
  Follow every pointer, mark everything reachable as ALIVE

Step 2 — SWEEP
  Scan all memory
  Anything NOT marked = unreachable = garbage → free it
```

**The Tri-color system:**
Go assigns every object a color during the mark phase:

| Color | Meaning |
|---|---|
| White | Not yet visited — potential garbage |
| Grey | Visited, but children not yet checked |
| Black | Visited + all children checked — definitely alive |

At the end: everything still **white** is freed.

**The important part — Go's GC runs concurrently:**
Java's older GCs would stop your entire program to collect garbage (stop-the-world pause). Go's GC runs *alongside* your program — very short pauses (microseconds), designed for low-latency systems.

```
Your program:   ──────────────────────────────────→
Go's GC:              ~~~~~~~~~~~~~~~~~~~~~~~~~~~
                      (runs in background, mostly)
```

**Where does Go's GC live?**
Unlike Java (where GC is part of the separately-installed JVM), Go's GC is **compiled into your binary**. When you ship the `.exe`, the GC ships with it. No external dependency.

**Java GC vs Go GC:**
| | Java | Go |
|---|---|---|
| Where it lives | Inside JVM (separate install) | Embedded in your binary |
| Pause time | Can be significant (tunable) | Sub-millisecond, concurrent |
| Control | Many flags and GC algorithms | Minimal — mostly automatic |
| Moving GC? | Yes (G1 relocates objects) | No — simpler but can fragment memory |

**First seen in:** Pre-Phase 1 discussion (before project setup)

---

## Concept 3 — Go Modules & go.mod

**What it is:**
A Go module is a self-contained project with a unique name, a Go version, and a list of dependencies. It is defined by a `go.mod` file at the root of the project.

**Real-world analogy:**
The module is your apartment. `go.mod` is the address plate on the door — it tells the world what this place is called and what version of Go it runs on.

**Code snippet:**
```
module github.com/amitpatle/paymentservice

go 1.26.2
```

**How to create it:**
```bash
go mod init github.com/amitpatle/paymentservice
```

**First seen in:** Phase 1 — Project Setup

---

## Concept 4 — Packages

**What it is:**
A package is a folder of `.go` files that belong together. Every `.go` file must declare which package it belongs to at the top with `package <name>`.

`package main` is special — it is the entry point of any runnable Go program. It must contain a `func main()`.

**Real-world analogy:**
If the module is the apartment, packages are the rooms — each room has a specific purpose. `package main` is the front door — execution always starts there.

**Code snippet:**
```go
package main

import "fmt"

func main() {
    fmt.Println("PaymentService starting...")
}
```

**First seen in:** Phase 1 — main.go

---

## Concept 5 — Structs

**What it is:**
A struct is a way to group related data together under one name. It is Go's replacement for classes (Go has no classes).

**Real-world analogy:**
A payment receipt — it has multiple fields (amount, date, merchant, status) all on one piece of paper. A struct is that receipt as a Go type.

**Code snippet:**
```go
type Payment struct {
    ID        string
    UserID    string
    Amount    float64
    Currency  string
    Provider  string
    Status    PaymentStatus
    CreatedAt time.Time
}
```

Creating a struct (struct literal):
```go
p := models.Payment{
    ID:     "pay_001",
    Amount: 999.99,
    Status: models.StatusPending,
}
```

Any field not set gets a **zero value** — `""` for string, `0` for numbers, `false` for bool. Go never leaves memory uninitialized.

**First seen in:** Phase 2 — models/models.go

---

## Concept 6 — Custom Types & Constants

**What it is:**
You can create a new named type based on an existing type. Combined with `const`, this replaces magic strings with safe, readable names.

**Real-world analogy:**
Instead of writing "red", "green", "yellow" as raw strings on traffic signals everywhere, you define named constants — `Red`, `Green`, `Yellow`. No typos, no guessing.

**Code snippet:**
```go
type PaymentStatus string   // new type based on string

const (
    StatusPending PaymentStatus = "pending"
    StatusSuccess PaymentStatus = "success"
    StatusFailed  PaymentStatus = "failed"
)
```

**First seen in:** Phase 2 — models/models.go

---

## Concept 7 — Short Variable Declaration (:=)

**What it is:**
`:=` creates a new variable and infers its type automatically from the right-hand side. No need to declare the type explicitly.

**Real-world analogy:**
Instead of saying "I have a box of type string, and I'm putting hello in it" — you just say "here's hello, figure out the box type yourself."

**Code snippet:**
```go
p := models.Payment{...}   // Go infers type as models.Payment
x := 42                    // Go infers type as int
```

vs the long form (rarely used):
```go
var p models.Payment = models.Payment{...}
```

**First seen in:** Phase 2 — main.go

---

## Concept 8 — Interfaces

**What it is:**
An interface is a contract — a list of method signatures a type must implement. Go interfaces are implicit: if your struct has all the required methods, it automatically satisfies the interface. No `implements` keyword needed.

**Real-world analogy:**
A power socket is an interface. Contract: "give me two pins, 220V." Whether it's a phone charger or a laptop adapter — doesn't matter. As long as it follows the contract, it plugs in.

**Java vs Go:**
```java
// Java — explicit
class Stripe implements PaymentProvider { ... }
```
```go
// Go — implicit
type StripeProvider struct { ... }
// If it has Charge() and Refund(), it IS a PaymentProvider. No declaration needed.
```

**Code snippet:**
```go
type PaymentProvider interface {
    Charge(req ChargeRequest) error
    Refund(paymentID string) error
}
```

**First seen in:** Phase 3 — provider/provider.go

---

## Concept 9 — Methods & Receivers

**What it is:**
A method is a function attached to a struct. The receiver `(s StripeProvider)` is Go's way of saying "this method belongs to this struct" — like `this` in Java, but written explicitly.

**Code snippet:**
```go
func (s StripeProvider) Charge(req provider.ChargeRequest) error {
    // s gives access to StripeProvider's fields, e.g. s.APIKey
    fmt.Printf("[STRIPE] Charging %.2f %s\n", req.Payment.Amount, req.Payment.Currency)
    return nil
}
```

**First seen in:** Phase 4 — adapter/stripe.go

---

## Concept 10 — Error Handling

**What it is:**
Go has no exceptions. Functions that can fail return an `error` as the last return value. The caller checks `if err != nil`. `nil` means success.

**Real-world analogy:**
Like a delivery status. The courier either hands you the package (`nil` — no problem) or gives you a slip saying what went wrong (`error` with a message). You decide what to do with either outcome.

**Code snippet:**
```go
// returning an error
return fmt.Errorf("stripe charge failed: %s", reason)

// returning success
return nil

// handling it
err := provider.Charge(req)
if err != nil {
    fmt.Printf("Payment failed: %s\n", err)
    return
}
fmt.Println("Payment successful!")
```

**First seen in:** Phase 4 — adapter/mock.go, adapter/stripe.go
