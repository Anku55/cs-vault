# Go Error Handling

## 1. Error Fundamentals
- error interface
- Error() method
- Returning errors
- Checking errors
- Error as a value
- Multiple return values
- Error propagation
- Error handling patterns

## 2. Creating Errors
- errors.New()
- fmt.Errorf()
- Error messages
- Sentinel errors
- Custom errors

## 3. Error Wrapping
- Error wrapping
- %w
- errors.Unwrap()
- Error chains
- Root cause
- Adding context

## 4. Error Inspection
- errors.Is()
- errors.As()
- Error identity
- Error type matching
- Wrapped error matching
- Type assertions with errors

## 5. Sentinel Errors
- Sentinel error concept
- Package-level errors
- Exported errors
- Unexported errors
- errors.Is()
- When to use sentinel errors
- Sentinel error design

## 6. Custom Error Types
- Custom error types
- Struct-based errors
- Implementing Error()
- Pointer vs value error types
- Error constructors
- Error metadata
- Custom Is() method
- Custom As() method

## 7. Error Context
- Adding operation context
- Adding resource context
- Error wrapping
- Preserving root cause
- Error context design
- Avoiding redundant context

## 8. Error Propagation
- Returning errors
- Propagating errors
- Wrapping errors
- Translating errors
- Handling errors at boundaries
- Error ownership
- Error boundaries

## 9. Error Classification
- Expected errors
- Unexpected errors
- Recoverable errors
- Unrecoverable errors
- Temporary errors
- Permanent errors
- Retryable errors
- Non-retryable errors
- Client errors
- Server errors
- Domain errors
- Infrastructure errors

## 10. nil and Errors
- nil errors
- nil interface
- Typed nil errors
- Comparing errors with nil
- Returning nil errors
- Common nil error bugs

## 11. defer
- defer basics
- Deferred function execution
- LIFO execution
- Deferred cleanup
- Defer with errors
- Defer with named return values
- Defer for resource cleanup

## 12. panic
- panic()
- Runtime panics
- Explicit panics
- Panic propagation
- When to use panic
- When not to use panic
- Panic in goroutines

## 13. recover
- recover()
- Recovering from panic
- Recover inside defer
- Panic recovery middleware
- Recovering HTTP server panics
- Logging recovered panics
- Stack traces

## 14. Resource Cleanup
- Closing files
- Closing HTTP response bodies
- Closing network connections
- Unlocking mutexes
- Database rollback
- Context cancellation
- Resource leaks
- Cleanup after errors

## 15. HTTP Error Handling
- HTTP status codes
- 400 Bad Request
- 401 Unauthorized
- 403 Forbidden
- 404 Not Found
- 409 Conflict
- 422 Unprocessable Content
- 429 Too Many Requests
- 500 Internal Server Error
- 502 Bad Gateway
- 503 Service Unavailable
- 504 Gateway Timeout
- Error response design
- JSON error responses
- HTTP error mapping

## 16. Validation Errors
- Input validation
- Request validation
- Field validation
- Required fields
- Type validation
- Format validation
- Range validation
- Multiple validation errors
- Custom validation errors
- Validation error responses

## 17. Network Errors
- Connection refused
- Connection reset
- Broken pipe
- EOF
- DNS errors
- Timeout errors
- TLS errors
- Connection errors
- Temporary network errors
- Retryable network errors
- Non-retryable network errors

## 18. HTTP Client Errors
- Request creation errors
- Client.Do() errors
- Connection errors
- Timeout errors
- Context cancellation
- HTTP status errors
- Response body errors
- JSON decoding errors
- Retryable HTTP errors
- Non-retryable HTTP errors

## 19. Database Errors
- Connection errors
- Query errors
- Scan errors
- Insert errors
- Update errors
- Delete errors
- Transaction errors
- Commit errors
- Rollback errors
- Duplicate key errors
- Constraint violations
- Deadlock errors
- Timeout errors
- Connection pool errors
- Context cancellation

## 20. Transaction Error Handling
- Begin transaction
- Commit
- Rollback
- Deferred rollback
- Commit failure
- Rollback failure
- Transaction timeout
- Transaction cancellation
- Partial transaction failure

## 21. Concurrency Error Handling
- Goroutine errors
- Error channels
- Result channels
- Multiple goroutine errors
- First-error strategy
- Collecting multiple errors
- Worker errors
- Worker recovery
- Error propagation between goroutines
- Cancellation after errors

## 22. errgroup
- errgroup.Group
- Group.Go()
- Group.Wait()
- Error propagation
- Context cancellation
- Fail-fast behavior
- Parallel error handling
- Bounded concurrency

## 23. Retry Handling
- Retryable errors
- Non-retryable errors
- Fixed retry
- Retry delay
- Exponential backoff
- Exponential backoff with jitter
- Maximum retry count
- Maximum retry duration
- Retry budget
- Retry storms
- Thundering herd problem
- Retry-after
- Idempotency

## 24. Idempotency
- Idempotent operations
- Non-idempotent operations
- Idempotency keys
- Duplicate requests
- Duplicate jobs
- Duplicate events
- Preventing duplicate processing
- Database uniqueness
- Exactly-once processing concept

## 25. Worker Error Handling
- Worker failures
- Job failures
- Job retry
- Maximum retries
- Job timeout
- Job cancellation
- Failed job state
- Dead-letter queue
- Poison messages
- Worker recovery
- Graceful worker shutdown

## 26. Event Streaming Error Handling
- Producer errors
- Consumer errors
- Broker errors
- Connection errors
- Serialization errors
- Deserialization errors
- Invalid messages
- Partition errors
- Offset errors
- Offset commit failures
- Storage errors
- Disk errors
- Consumer failures
- Producer retries
- Consumer recovery
- Dead-letter events
- Poison events

## 27. Distributed System Errors
- Partial failures
- Network partitions
- Service unavailable
- Timeouts
- Retries
- Backoff
- Idempotency
- Duplicate requests
- Duplicate messages
- Cascading failures
- Circuit breakers
- Rate limiting
- Graceful degradation
- Failure detection

## 28. Circuit Breaker
- Circuit breaker concept
- Closed state
- Open state
- Half-open state
- Failure threshold
- Recovery timeout
- Failure counting
- Preventing cascading failures

## 29. Timeouts
- Request timeout
- Connection timeout
- Database timeout
- Worker timeout
- Job timeout
- Context timeout
- Context deadline
- Context cancellation
- Timeout propagation

## 30. Error Logging
- Error logging
- Structured logging
- Log levels
- Debug
- Info
- Warn
- Error
- Fatal
- Error context
- Request ID
- Correlation ID
- Stack traces
- Avoiding duplicate logs
- Avoiding sensitive data in logs

## 31. Error Codes
- Error code design
- Domain error codes
- API error codes
- Internal error codes
- Machine-readable errors
- Human-readable messages
- Stable error codes
- Error code documentation

## 32. Error Translation
- Internal errors
- External errors
- Domain errors
- Infrastructure errors
- Database error translation
- HTTP error translation
- Client-safe error messages
- Preventing internal error leakage

## 33. Error Boundaries
- Error ownership
- Error boundaries
- Repository error handling
- Service error handling
- Handler error handling
- Error translation at boundaries
- Error propagation across layers

## 34. Error Testing
- Testing expected errors
- Testing unexpected errors
- Testing errors.Is()
- Testing errors.As()
- Testing wrapped errors
- Testing custom errors
- Testing validation errors
- Testing HTTP errors
- Testing database errors
- Testing network errors
- Testing timeout errors
- Testing retry behavior
- Testing panic recovery
- Testing concurrent errors

## 35. Production Error Handling
- Structured errors
- Error classification
- Error wrapping
- Error codes
- Logging
- Metrics
- Tracing
- Retries
- Timeouts
- Circuit breakers
- Rate limiting
- Graceful degradation
- Alerting
- Error monitoring

## 36. Error Observability
- Error rate
- Error count
- Error type
- Error code
- Endpoint-level errors
- Service-level errors
- Retry count
- Timeout count
- Failed jobs
- Dead-letter messages
- Consumer failures
- Error metrics
- Distributed tracing

## 37. Go Standard Library Error APIs
- errors.New()
- errors.Is()
- errors.As()
- errors.Unwrap()
- fmt.Errorf()
- fmt.Errorf("%w")
- io.EOF
- io.ErrUnexpectedEOF
- os.ErrNotExist
- os.ErrPermission
- context.Canceled
- context.DeadlineExceeded

## 38. Common Error Handling Mistakes
- Ignoring errors
- Using panic for normal errors
- Losing the root cause
- Incorrect error wrapping
- Over-wrapping errors
- Logging the same error multiple times
- Comparing error strings
- Retrying every error
- Retrying non-idempotent operations
- Exposing internal errors
- Creating too many custom error types
- Using errors as control flow unnecessarily
- Ignoring cleanup errors
- Ignoring context cancellation