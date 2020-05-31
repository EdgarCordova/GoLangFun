### Endpoints:

| Method | Endpoint | Description |
| --- | --- | --- |
| GET | /Todo(/) | Gets all todo items |
| GET | /Todo/{id} | Gets todo item by id |
| POST | /Todo(/) | Creates a todo item. (requests body with description) | 
| POST | /Todo/{id} | Update the description on the todo item | 
| DELETE | /Todo/{id} | Delete a todo item |
### To run locally. 

- run `go run main.go`
- `localhost:8080/` should return `"homepage endpoint Hit"`
