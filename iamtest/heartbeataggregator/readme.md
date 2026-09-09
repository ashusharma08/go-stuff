Day 2: The High-Scale "Heartbeat" Aggregator
We’ve solved Entitlement (the "Can I watch?" part). Now we need to solve Concurrency Tracking (the "Are they currently watching?" part).

The Challenge:
Every video player sends a "Heartbeat" event every 30 seconds to say "I'm still here."

Scale: 100,000 events per second.

Problem: If you increment Redis for every single heartbeat, you'll hit a bottleneck.

Goal: Build a Batch Aggregator in Go that:

Receives heartbeats on a channel.

Buffers them in memory (e.g., "User X has 2 active heartbeats").

Flushes the counts to the database in one big "Write" every 5 seconds.

How would you start the design for this? Would you use a single global channel with multiple workers, or a "Sharded" approach to avoid channel contention?