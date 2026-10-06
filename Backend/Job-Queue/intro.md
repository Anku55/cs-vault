# Job Queue Fundamentals

A **Job Queue** is a system used to process work asynchronously and reliably.

Instead of forcing an application to execute expensive or time-consuming work immediately, the application creates a **job**, places it into a **queue**, and allows a **worker** to process it asynchronously.

```text
                    Job Queue System

┌──────────────┐
│   Producer   │
│              │
│ Creates Job  │
└──────┬───────┘
       │
       │ Submit Job
       ▼
┌──────────────────────┐
│        Queue         │
│                      │
│ Job A                │
│ Job B                │
│ Job C                │
└──────────┬───────────┘
           │
           │ Claim Job
           ▼
┌──────────────────────┐
│       Consumer       │
│                      │
│ Fetches / Claims Job │
└──────────┬───────────┘
           │
           ▼
┌──────────────────────┐
│        Worker        │
│                      │
│ Executes Job         │
└──────────┬───────────┘
           │
           ▼
      ┌─────────┐
      │ Result  │
      └────┬────┘
           │
      ┌────┴─────┐
      ▼          ▼
    ACK        RETRY
```

---

# 1. What is a Job?

A **job** is a unit of work that needs to be executed.

A job describes **what work needs to be performed** and contains the information required by a worker to perform that work.

### Example

Suppose an application needs to send a welcome email.

Instead of immediately sending the email:

```text
User Registration
       │
       ▼
Send Email
       │
       ▼
Response
```

the application can create a job:

```text
Job
├── ID: job-123
├── Type: send_email
├── Payload: email information
├── Priority: 5
├── Attempts: 0
└── Status: QUEUED
```

The job is then placed into the queue.

```text
Application
     │
     │ Create Job
     ▼
┌────────────────┐
│ Job            │
│                │
│ send_email     │
│ job-123        │
└───────┬────────┘
        │
        ▼
      Queue
```

## Example Job

```json
{
  "id": "job-123",
  "type": "send_email",
  "payload": {
    "to": "user@example.com",
    "subject": "Welcome"
  },
  "priority": 5,
  "attempts": 0,
  "status": "QUEUED"
}
```

## What can a Job represent?

Almost any asynchronous task:

```text
send_email
resize_image
generate_report
process_video
send_notification
deliver_webhook
generate_pdf
process_payment
update_search_index
```

## Job vs Execution

A job is **not the execution itself**.

```text
Job
  ↓
"What needs to be done?"

Worker
  ↓
"Who performs it?"

Handler
  ↓
"How is it performed?"
```

For example:

```text
Job:
    type = resize_image

Worker:
    receives the job

Handler:
    resizeImage(...)
```

---

# 2. Producer

A **Producer** is a component that **creates and submits jobs to the queue**.

Think:

> Producer = "I need this work to be done."

```text
Producer
    │
    │ Create Job
    ▼
  Queue
```

## Example

An e-commerce application receives an order.

```text
User
 │
 ▼
Order Service
 │
 │ Create "send_confirmation_email" job
 ▼
Job Queue
```

Here:

```text
Order Service = Producer
```

The producer does not necessarily execute the job.

## Producer Responsibilities

A producer may:

- Create a job
- Generate a job ID
- Specify job type
- Attach payload
- Set priority
- Set retry configuration
- Set scheduling information
- Submit the job
- Handle submission failures

Example:

```json
{
  "type": "send_email",
  "payload": {
    "to": "user@example.com"
  }
}
```

## Multiple Producers

A queue can have many producers:

```text
Payment Service ──────┐
                      │
Order Service ────────┤
                      │
Notification Service ─┤
                      ├──► Job Queue
                      │
Web API ──────────────┘
```

All of these services can submit jobs.

## Producer Failure

A major distributed-systems problem occurs when the producer doesn't know whether the queue received the job.

```text
Producer
   │
   │ Submit Job
   ▼
Queue
   │
   │ Job accepted
   ▼
Producer
   │
   │ Network failure 💥
   X
```

The producer may not know whether the job was:

```text
Accepted
    OR
Rejected
```

If it retries, duplicate jobs may be created.

```text
Job #123
Job #123
```

This introduces the need for concepts such as:

- Idempotency
- Request IDs
- Deduplication
- Retry policies
- Timeouts

---

# 3. Consumer

A **Consumer** is a component that **receives or claims jobs from the queue**.

Think:

> Consumer = "Give me work that I can process."

```text
Queue
  │
  │ Claim / Fetch
  ▼
Consumer
```

## Basic Flow

```text
Queue
  │
  ▼
Consumer
  │
  ▼
Receive Job
  │
  ▼
Process Job
```

## Consumer vs Worker

These terms are often used interchangeably, but they describe different responsibilities conceptually.

| Component | Responsibility |
|---|---|
| Producer | Creates jobs |
| Queue | Stores/manages jobs |
| Consumer | Retrieves/claims jobs |
| Worker | Executes jobs |
| Handler | Contains actual job logic |

A practical implementation can combine Consumer and Worker:

```text
┌─────────────────────────────┐
│           Worker            │
│                             │
│  Claim Job                  │
│      ↓                      │
│  Execute Job                │
│      ↓                      │
│  ACK / Retry                │
└─────────────────────────────┘
```

For your project, this is a good initial design.

---

# 4. Worker

A **Worker** is the component that **actually executes the job**.

Think:

> Worker = "I will perform the work."

```text
Queue
  │
  │ Job
  ▼
Worker
  │
  │ Execute
  ▼
Handler
  │
  ▼
Result
```

## Example

Queue contains:

```text
Job 1 → send_email
Job 2 → resize_image
Job 3 → generate_report
```

Workers execute them:

```text
                 Queue
                   │
        ┌──────────┼──────────┐
        ▼          ▼          ▼
    Worker 1   Worker 2   Worker 3
        │          │          │
        ▼          ▼          ▼
      Email      Image      Report
```

Jobs can therefore be processed concurrently.

---

# Worker Pool

Instead of creating an unlimited number of workers, a system commonly uses a **worker pool**.

```text
                    Queue
                      │
          ┌───────────┼───────────┐
          ▼           ▼           ▼
      Worker 1    Worker 2    Worker 3
          │           │           │
          ▼           ▼           ▼
        Job A       Job B       Job C
```

In Go, workers can be implemented using goroutines:

```go
for i := 0; i < 10; i++ {
    go worker()
}
```

This provides controlled concurrency.

### Why controlled concurrency?

Without a limit:

```text
1 job  → 1 goroutine
10K jobs → 10K goroutines
1M jobs → potentially huge resource usage
```

A worker pool limits concurrent execution:

```text
1,000 jobs
    │
    ▼
Queue
    │
    ▼
10 Workers
    │
    ▼
10 jobs processed concurrently
```

---

# Worker Responsibilities

A production-oriented worker may need to:

- Register with the queue
- Claim jobs
- Execute jobs
- Send heartbeats
- Renew leases
- ACK successful jobs
- Report failures
- Retry jobs
- Handle timeouts
- Handle cancellation
- Gracefully shut down
- Report metrics

---

# Worker Failure

A critical problem:

```text
Queue
  │
  ▼
Worker
  │
  │ Claims Job
  │
  💥 CRASH
```

If the queue permanently removes the job immediately after the worker claims it:

```text
Job
 ↓
Removed
 ↓
Worker crashes
 ↓
❌ Job lost
```

A reliable queue needs a recovery mechanism.

One approach is a **lease / visibility timeout**:

```text
Worker claims Job
       │
       ▼
Job becomes temporarily invisible
       │
       ▼
Worker processes Job
       │
   ┌───┴────┐
   ▼        ▼
  ACK      Crash
   │        │
   ▼        ▼
Complete   Lease expires
              │
              ▼
        Job becomes available
```

---

# 5. Queue

A **Queue** is the intermediary component that **stores and manages jobs between producers and workers**.

Think:

> Queue = "The waiting line for work."

```text
Producer
   │
   │ Enqueue
   ▼
┌───────────────────┐
│       QUEUE       │
│                   │
│ Job A             │
│ Job B             │
│ Job C             │
│ Job D             │
└─────────┬─────────┘
          │
          │ Claim
          ▼
       Worker
```

---

# Why do we need a Queue?

The queue provides **decoupling** between producers and workers.

Without a queue:

```text
Request
   │
   ▼
Application
   │
   ▼
Expensive Work
   │
   ▼
Response
```

The application must wait.

With a queue:

```text
Request
   │
   ▼
Application
   │
   ▼
Queue
   │
   └──────────────► Worker
   │
   ▼
Immediate Response
```

The producer and worker can operate independently.

---

# Queue as a Buffer

Suppose:

```text
Incoming jobs = 10,000/sec
Worker capacity = 2,000/sec
```

The queue temporarily absorbs the difference:

```text
10,000 jobs/sec
       │
       ▼
     Queue
       │
       ▼
2,000 jobs/sec
       │
       ▼
    Workers
```

This is called **buffering**.

However, an unlimited queue is dangerous.

Eventually:

```text
Queue grows
   ↓
Memory/storage grows
   ↓
Resource exhaustion
```

Therefore, production systems need **backpressure**.

---

# Queue Operations

A basic queue might provide:

```text
Enqueue()
Dequeue()
```

A reliable distributed queue needs much more:

```text
Enqueue
Claim
ACK
NACK
Retry
Delay
Schedule
Cancel
Dead-letter
```

The exact semantics depend on the architecture.

---

# Queue and Acknowledgement

A worker should generally not cause a job to become permanently completed merely by receiving it.

Instead:

```text
Queue
  │
  ▼
Worker claims Job
  │
  ▼
RUNNING
  │
  ▼
Execute
  │
  ▼
ACK
  │
  ▼
COMPLETED
```

The ACK means:

> "I successfully processed this job."

If processing fails:

```text
Worker
  │
  ▼
Job fails
  │
  ▼
Retry / Failure handling
```

---

# Persistent vs In-Memory Queue

## In-Memory Queue

```text
Producer
   ↓
RAM
   ↓
Worker
```

Fast, but:

```text
Server crashes
     ↓
Jobs disappear
```

Useful for early development.

## Persistent Queue

```text
Producer
   ↓
Queue Server
   ↓
Persistent Storage
   ↓
Worker
```

After a restart:

```text
Server crashes
     ↓
Server restarts
     ↓
Recover jobs
```

This provides much stronger durability.

---

# 6. Job Lifecycle

The **job lifecycle** describes all states and transitions a job goes through from creation until it is completed, cancelled, or permanently failed.

A useful state machine:

```text
                         ┌──────────────┐
                         │    CREATED   │
                         └──────┬───────┘
                                │
                              enqueue
                                │
                                ▼
                         ┌──────────────┐
                         │    QUEUED    │
                         └──────┬───────┘
                                │
                              claim
                                │
                                ▼
                         ┌──────────────┐
                         │   RUNNING    │
                         └──────┬───────┘
                                │
                       ┌────────┴────────┐
                       │                 │
                    success            failure
                       │                 │
                       ▼                 ▼
                 ┌───────────┐     ┌──────────┐
                 │ COMPLETED │     │ RETRYING │
                 └───────────┘     └────┬─────┘
                                        │
                                   retry delay
                                        │
                                        ▼
                                     QUEUED
                                        │
                                      claim
                                        │
                                        ▼
                                    RUNNING
                                        │
                              max attempts exceeded
                                        │
                                        ▼
                                  ┌─────────┐
                                  │  DEAD   │
                                  └────┬────┘
                                       │
                                       ▼
                                      DLQ
```

---

# Job States

A practical system can use:

```text
CREATED
QUEUED
RUNNING
COMPLETED
FAILED
RETRYING
CANCELLED
DEAD
```

Later you can introduce:

```text
SCHEDULED
```

---

## CREATED

The producer has created the job.

```text
Application
    │
    ▼
Create Job
    │
    ▼
CREATED
```

The job may not yet have been successfully submitted to the queue.

---

## QUEUED

The queue has accepted the job.

```text
Producer
   │
   ▼
Queue
   │
   ▼
QUEUED
```

The job is waiting for a worker.

---

## RUNNING

A worker has successfully claimed the job.

```text
QUEUED
   │
   │ claim
   ▼
RUNNING
```

The system should record ownership:

```text
job_id:    job-123
worker_id: worker-01
status:    RUNNING
```

---

## COMPLETED

The worker successfully executes the job and acknowledges it.

```text
RUNNING
   │
   │ success
   ▼
ACK
   │
   ▼
COMPLETED
```

---

## FAILED

The job execution fails.

```text
RUNNING
   │
   │ error
   ▼
FAILED
```

A temporary failure should normally lead to a retry rather than immediate permanent failure.

---

## RETRYING

The system schedules another attempt.

```text
FAILED
   │
   ▼
RETRYING
   │
   │ backoff
   ▼
QUEUED
```

---

# Exponential Backoff

Immediate retries can overload a failing dependency.

Instead:

```text
Attempt 1 → FAIL
              ↓
            wait 1s

Attempt 2 → FAIL
              ↓
            wait 2s

Attempt 3 → FAIL
              ↓
            wait 4s
```

A simplified formula:

```text
delay = base × 2^attempt
```

Production systems often add **jitter**:

```text
delay = exponential_backoff + random_jitter
```

This prevents large numbers of jobs from retrying simultaneously.

---

# DEAD / Dead Letter Queue

A job should not retry forever.

Example:

```text
Max attempts = 3

Attempt 1 → FAIL
Attempt 2 → FAIL
Attempt 3 → FAIL
             │
             ▼
            DEAD
             │
             ▼
            DLQ
```

The **Dead Letter Queue (DLQ)** stores jobs that require investigation or manual handling.

Example:

```text
DLQ

Job ID: job-123
Reason: External API unavailable
Attempts: 3
Last error: timeout
```

---

# CANCELLED

A job may be cancelled before or during execution.

```text
QUEUED
   │
   │ cancel
   ▼
CANCELLED
```

A running job may use Go's `context.Context`:

```text
RUNNING
   │
   │ cancellation signal
   ▼
Handler stops
   │
   ▼
CANCELLED
```

---

# SCHEDULED

A scheduled job should execute at a future time.

Example:

```text
CREATED
   │
   ▼
SCHEDULED
   │
   │ wait until scheduled time
   ▼
QUEUED
   │
   ▼
RUNNING
```

Example:

```text
"Generate report at 09:00 tomorrow"
```

---

# Complete Job Lifecycle Example

Consider an image-processing job:

```text
1. User uploads image
          │
          ▼
2. API creates Job
          │
          ▼
3. Job enters QUEUED
          │
          ▼
4. Worker claims Job
          │
          ▼
5. Job becomes RUNNING
          │
          ▼
6. Image processing fails
          │
          ▼
7. RETRYING
          │
          ▼
8. Backoff period
          │
          ▼
9. QUEUED
          │
          ▼
10. Worker claims again
          │
          ▼
11. RUNNING
          │
          ▼
12. Processing succeeds
          │
          ▼
13. ACK
          │
          ▼
14. COMPLETED
```

---

# Common Industry Challenges

Building:

```text
Producer → Queue → Worker
```

is easy.

Building a **reliable distributed job queue** is difficult because of failures and concurrency.

---

## 1. Duplicate Job Execution

Two workers might accidentally claim the same job:

```text
        Job 123
        /     \
       ▼       ▼
 Worker A   Worker B
```

Both execute it.

You need mechanisms such as:

- atomic claiming
- leases
- locking
- idempotent handlers

---

## 2. Worker Crash

```text
Worker
   │
   ▼
Job 123
   │
   💥
```

The system must detect the failure and recover the job.

Typical mechanisms:

```text
Heartbeat
Lease
Visibility timeout
Timeout
Recovery scanner
```

---

## 3. Duplicate Execution After Crash

Consider:

```text
Worker processes job
       │
       ▼
Work succeeds
       │
       ▼
Worker crashes BEFORE ACK
```

The queue may retry the job.

Result:

```text
Job executed twice
```

Therefore, **at-least-once delivery** often requires jobs to be **idempotent**.

---

## 4. Lost Jobs

A job can be lost if:

```text
Job accepted
   ↓
Server crashes
   ↓
Job existed only in RAM
```

Solution:

```text
Durable storage
Write-ahead log
Replication
Recovery
```

---

## 5. Poison Jobs

Some jobs fail every time:

```text
Job
 ↓
FAIL
 ↓
RETRY
 ↓
FAIL
 ↓
RETRY
 ↓
FAIL
 ↓
...
```

Without limits, this can create a retry storm.

Solution:

```text
Maximum retries
Exponential backoff
Dead Letter Queue
```

---

## 6. Backpressure

Producers can overwhelm workers:

```text
Producer rate = 100K jobs/sec
Worker capacity = 10K jobs/sec
```

The queue grows continuously.

Possible solutions:

```text
Rate limiting
Queue limits
Producer throttling
Priority queues
Load shedding
Scaling workers
```

---

## 7. Ordering

Some workloads require:

```text
Job A
  ↓
Job B
  ↓
Job C
```

to execute in that order.

But multiple workers naturally introduce concurrency:

```text
Worker 1 → Job A
Worker 2 → Job B
```

You therefore need to explicitly define your ordering guarantees.

Possible approaches:

```text
Partitioning
Per-key ordering
Single consumer per partition
Sequence numbers
```

---

## 8. Fairness

Suppose one producer submits millions of jobs:

```text
Producer A → 1,000,000 jobs
Producer B → 10 jobs
```

Producer B's jobs could starve.

Possible solutions:

```text
Fair scheduling
Per-tenant queues
Weighted scheduling
Rate limits
Priority policies
```

---

## 9. Large Payloads

Putting huge payloads directly into the queue can cause:

```text
High memory usage
Large network transfers
Slow serialization
Storage pressure
```

A common architecture is:

```text
Job
 │
 ├── metadata
 └── object reference
          │
          ▼
       Object Storage
```

Instead of storing a 500 MB file inside the job itself.

---

## 10. Poison / Malformed Jobs

A producer could submit invalid data:

```json
{
  "type": "send_email",
  "payload": "INVALID"
}
```

Workers need:

```text
Validation
Schema checking
Payload limits
Safe deserialization
```

---

## 11. Graceful Shutdown

A worker shouldn't simply disappear:

```text
Worker
  │
  ├── Job A running
  ├── Job B queued locally
  │
  💥 shutdown
```

Instead:

```text
SIGTERM
   │
   ▼
Stop claiming new jobs
   │
   ▼
Finish current jobs
   │
   ▼
ACK
   │
   ▼
Shutdown
```

---

## 12. Observability

You need to know what your system is doing.

Important metrics:

```text
jobs_submitted_total
jobs_completed_total
jobs_failed_total
jobs_retried_total

queue_depth
processing_latency
job_wait_time
worker_count
worker_failures
retry_rate
```

Logs should contain useful identifiers:

```json
{
  "event": "job_failed",
  "job_id": "job-123",
  "worker_id": "worker-7",
  "attempt": 3
}
```

---

# End-to-End Architecture

Putting everything together:

```text
                         PRODUCERS
                             │
                ┌────────────┼────────────┐
                │            │            │
                ▼            ▼            ▼
             Service A   Service B    Service C
                │            │            │
                └────────────┼────────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │   Queue Server  │
                    │                 │
                    │ Job Management │
                    │ Scheduling     │
                    │ Retry Manager   │
                    │ Routing        │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │     Storage     │
                    │                 │
                    │ Jobs            │
                    │ Attempts        │
                    │ Workers         │
                    └─────────────────┘
                             │
              ┌──────────────┼──────────────┐
              ▼              ▼              ▼
         ┌─────────┐    ┌─────────┐    ┌─────────┐
         │ Worker  │    │ Worker  │    │ Worker  │
         │ Node 1  │    │ Node 2  │    │ Node 3  │
         └────┬────┘    └────┬────┘    └────┬────┘
              │              │              │
              ▼              ▼              ▼
           Handler        Handler        Handler
              │              │              │
              └──────────────┼──────────────┘
                             │
                       ACK / RETRY
                             │
                    ┌────────┴────────┐
                    ▼                 ▼
               COMPLETED             DLQ
```

---

# Core Mental Model

The entire system can be remembered as:

```text
Producer
   │
   │ "I need this work done."
   ▼
Queue
   │
   │ "I'll hold and manage this work."
   ▼
Consumer
   │
   │ "Give me work."
   ▼
Worker
   │
   │ "I'll execute it."
   ▼
Handler
   │
   │ "Here's the actual implementation."
   ▼
Result
   │
   ├──────────────► ACK → COMPLETED
   │
   └──────────────► FAILURE → RETRY → QUEUED
                              │
                              ▼
                            DLQ
```

---

# Key Concepts to Master

Before implementing your Job Queue, you should understand:

- [ ] Job
- [ ] Producer
- [ ] Consumer
- [ ] Worker
- [ ] Queue
- [ ] Worker Pool
- [ ] Job Lifecycle
- [ ] Job States
- [ ] Acknowledgement (ACK)
- [ ] Retry
- [ ] Exponential Backoff
- [ ] Jitter
- [ ] Dead Letter Queue
- [ ] Visibility Timeout
- [ ] Lease
- [ ] Heartbeat
- [ ] Graceful Shutdown
- [ ] Idempotency
- [ ] At-most-once delivery
- [ ] At-least-once delivery
- [ ] Backpressure
- [ ] Concurrency
- [ ] Durability
- [ ] Persistence
- [ ] Failure Recovery
- [ ] Scheduling
- [ ] Priority
- [ ] Ordering
- [ ] Observability
- [ ] Load Balancing

> **Core principle:** A production-grade job queue is not simply `Producer → Queue → Worker`. The difficult part is guaranteeing correct behavior when jobs fail, workers crash, networks fail, producers retry, servers restart, and thousands or millions of jobs arrive concurrently.