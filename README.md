# libs

A collection of modules other projects may be dependent upon.
Mainly wrappers to third-party projects to allow for configuration via files,
configuration references and some sensible default configurations.

Implementations are built for necessity rather than completion and may not have all of the expected functions available.

## Modules

### concurrent

Thread-safe objects to protect against potential data races.

### filesystem

Convenience functions for filesystem operations.

### metrics

A wrapper to rcrowley/go-metrics with configuration options for file and HTTP outupt.
A base set of runtime stats is provided as well:
uptime, memory usage (alloc, sys, total), and threads (goroutines)

### out

A logger with archiving (rolling, compression) functionality and optional asynchronous output.

### random

Convenience functions for randomized output.

### types

Convenience functions for casting values between types.

### websocket

Websocket provides management capabilities for websocket connections.
