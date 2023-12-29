Simple jwt authentication with mongoDB and Gin Framework Golang

Sample register: curl --location 'http://localhost:8000/users/signup'
--header 'Content-Type: application/json'
--header 'Authorization: Bearer <>'
--data-raw '{ "first_name":"ngo", "last_name":"nguyen", "email":"tanngontn@gmail.com", "password":"123456", "phone":"123456789", "user_type":"ADMIN" }'

Sample login curl --location 'http://localhost:8000/users/login'
--header 'Content-Type: application/json'
--header 'Authorization: Bearer <>'
--data-raw '{ "email":"tanngontn@gmail.com", "password":"123456" }'
