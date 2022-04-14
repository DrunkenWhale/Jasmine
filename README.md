# Jasmine

A distributed cache

its structure like this:

```text
                 /---> node
user -> manager  -----> node  
                 \---> node
```

use __Consistent Hashing__ algorithm to find designated key

## manager

you should config nodes in `nodes.json` like this:

every object has two attributes:
node's `name`
node's `address`

```json

[
  {
    "name": "pigeon1",
    "address": "localhost:8001"
  },
  {
    "name": "pigeon2",
    "address": "localhost:8002"
  },
  {
    "name": "pigeon3",
    "address": "localhost:8003"
  }
]

```

after config

use `manager -port 7070` to run a manager, and you can use `http://localhost:7070/api/?key=114` to get a value

## node

define a `callback` which function signature is `func(key string) ([]byte, error)`

if node can't find key from its cache, it will call `callback(key)` to get value and put it into cache