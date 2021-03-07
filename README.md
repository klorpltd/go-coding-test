# go-coding-test

### Create an HTTP handler that creates a new user
   1. User must contain name
   2. Returned with uuid and name
   3. Stored in in-memory store

`REST POST /user`
```json
{
  "name": "bob"
}
```
Returns
```json
{
  "id": "9bff75a6-bb3b-4ff0-8ac5-e2aa1f62220d",
  "name": "bob"
}
```
### Start async user processing

1. Every user creation starts an async "data processing" for user
2. Only 3 users maybe be processing at one time
3. Service must not terminate while processing user