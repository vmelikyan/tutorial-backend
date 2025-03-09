Scale demo

# Tutorial Backend Application

This is a simple Golang API server, that has 3 endpoints

- GET /tasks > returns all tasks from database
- POST /task > creates a new task
- DELETE /task/{id} > soft deletes task with given id

With a dependent database named `tasks` that has one table `tasks` with the below fields

- `id` > primary key
- `task` > description of a task
- `created_at` > created at timestamp
- `deleted` > boolean flag for soft deletion
- `deleted_at` > deleted at timestamp
