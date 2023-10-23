## Key value pair db like redis implementation in golang

## Completly inspired from ONE2N golang playbook

```
https://playbook.one2n.in/key-value-db-redis-exercise#714257b1224c462aa936497aa2261671
```

- implementing a small subset of Redis commands. We’ll focus on the following features:
  - It will be reachable over TCP
  - It can handle concurrent clients. That is, it can accept connections from multiple clients and responds to them regardless of the order in which they connect and send requests. For instance, client C1 connects, client C2 connects, C2 can send a request and get a response no matter what C1 is doing, whether staying idle, disconnecting or sending requests as well.
  - The following will be implemented:
    - `GET`: Accepts a string, and return the value stored for that key, if any
    - `SET`: Accepts two strings, a key and a value, and sets the value for the key, overriding any values that may have been present
    - `DEL`: Accepts a string and deletes the value that may have been there
    - `INCR`: Accepts a single argument and increments the existing value. If the value is not an integer, it’s an error, if there are no values, it gets initialized to `1`, resulting in an identical outcome as calling `SET <key> 1`.