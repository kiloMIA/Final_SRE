Sample curl requests for routes
POST - /tasks

curl -X POST http://localhost:8081/tasks \
     -H "Content-Type: application/json" \
     -d '{
           "description": "Your Task Description",
           "priority": "High",
           "dueDate": "2023-01-01",
           "completed": false
         }'

GET - /tasks
curl -X GET http://localhost:8081/tasks

PUT - /tasks/{taskID}
curl -X PUT http://localhost:8081/tasks/:taskID \
     -H "Content-Type: application/json" \
     -d '{
           "description": "Updated Task Description",
           "priority": "Medium",
           "dueDate": "2023-02-01",
           "completed": true
         }'

DELETE - /tasks/{taskID}
curl -X DELETE http://localhost:8081/tasks/:taskID

utyutiiuyruiyruyruk

