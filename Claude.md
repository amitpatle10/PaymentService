# PaymentService — Build & Learn Guide

> **Language:** Go (learning from scratch)
> **Goal:** Build a production-grade, loosely coupled, plug-and-play Payment Service
> **Scale Target:** Millions of users — SDE2 interview-ready

---

## What We Are Building

A payment service that:

- Can plug in any payment provider (Stripe, PayPal, Razorpay) without changing core logic
- Handles failures gracefully (retries, fallback, circuit breaker)
- Never double-charges a user (idempotency)
- Keeps a reliable record of every rupee/dollar moved (ledger)
- Scales to millions of transactions

---

## Learning Philosophy

- **Go is new for the learner.** Every Go concept is introduced as we need it — never dumped in advance.
- **Technical terms always explained in simple brackets.** Example: _interface (a contract — list of rules a struct must follow)_.
- **One concept at a time.** Finish understanding one thing completely before moving to the next.
- **Simple language first.** Real-world analogies preferred over jargon.
- **Discussion always before code.** No file is created until the concept is discussed and agreed upon.

---

## Principles

- **Discuss first (mandatory).** Structure, tradeoffs, alternatives — all discussed before any code is written.
- **Explain every Go keyword/syntax** the first time it appears. Never assume the learner knows it.
- **Name every design pattern** we use — why this pattern, what problem it solves, what alternative exists.
- **Tradeoff analysis at every decision.** No choice is made without discussing at least two alternatives.
- **Technical terms in brackets.** Every time a technical term appears, add a simple explanation in brackets right after it.
- **One topic at a time.** Never introduce two new concepts in one response.
- **Test immediately.** After every piece of code, run it and verify it works before moving forward.
- **SDE2 interview questions surfaced at each topic.** After each concept, list the interview questions that concept answers.
- **Reflect after each phase.** Summarize what was learned — the pattern, tradeoff, or Go concept.

---

## Approach (Every Session)

1. **Discuss** — what are we building next, why, what are the tradeoffs
2. **Go concept** — introduce the Go feature needed, explain in simple terms with analogy
3. **Build** — write code step by step, every line explained
4. **Run** — execute and verify
5. **Interview questions** — what SDE2 questions does this concept answer
6. **Reflect** — what did we learn today

---

## Go Concepts — Introduced in Order as We Build

| Phase | What We Build         | Go Concept Introduced                          |
| ----- | --------------------- | ---------------------------------------------- |
| 1     | Project setup         | Modules, packages, go.mod, how Go compiles     |
| 2     | Core models           | Structs, basic types, zero values              |
| 3     | Provider interface    | Interfaces (the core of plug-and-play)         |
| 4     | Mock + Stripe adapter | Implementing interfaces, error handling        |
| 5     | Orchestrator          | Methods on structs, pointer vs value receivers |
| 6     | Idempotency + DB      | database/sql, transactions, context            |
| 7     | API layer             | net/http, handlers, middleware                 |
| 8     | Async events          | Goroutines, channels, select                   |
| 9     | Circuit breaker       | sync.Mutex, atomic operations                  |

---

## Payment Service Concepts — SDE2 Must-Know

| Concept              | Simple Explanation                                          | Interview Question It Answers                 |
| -------------------- | ----------------------------------------------------------- | --------------------------------------------- |
| Idempotency          | Same request = same result, no matter how many times sent   | "How do you prevent double charges?"          |
| Saga Pattern         | Break big transaction into small steps, each reversible     | "How do you handle distributed transactions?" |
| Outbox Pattern       | Save event to DB same time as data, then publish            | "How do you prevent dual-write failures?"     |
| Circuit Breaker      | Stop calling a failing service before it takes you down too | "How do you handle provider downtime?"        |
| Double-entry Ledger  | Every debit has a matching credit — money never disappears  | "How do you ensure financial accuracy?"       |
| Provider Abstraction | One common interface, multiple provider implementations     | "How do you make the system extensible?"      |

---

## Interview Readiness Checklist

By the end of this project, you should be able to answer:

- [ ] How does your payment service handle a provider going down?
- [ ] What happens if the network drops after charge but before your DB is updated?
- [ ] How do you prevent a user from being charged twice?
- [ ] How would you add a new payment provider without touching existing code?
- [ ] How does your system scale to 1 million transactions per day?
- [ ] What database would you choose and why?
- [ ] How do you ensure exactly-once payment processing?
- [ ] Walk me through the flow of a single payment — end to end
- [ ] What is a saga? When would you use it over a traditional DB transaction?
- [ ] How do you handle reconciliation when your DB and provider are out of sync?

---

## Ground Rules

- No code before discussion is complete
- Every new Go keyword explained the first time it appears
- Simple analogies before technical definitions
- If a term is used without a bracket explanation, ask for it
- One topic per session — go deep, not wide

---

## GO_Learning.md — Living Concept Log

A file `GO_Learning.md` lives at the root of this project. It is a running log of every Go concept learned during this project.

**Rules for maintaining it:**

- Every time a new Go concept is introduced and understood, append it to `GO_Learning.md`
- Each entry must include: the concept name, a simple plain-English explanation, a real-world analogy, a code snippet showing it in use, and which phase/file it was first encountered in
- Update it immediately after the concept is discussed and verified — not at the end of a session
- Never rewrite or restructure old entries — only append new ones
- The file is the learner's personal Go reference, built from their own project
