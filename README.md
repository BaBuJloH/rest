# REST API

## ENDPOINTS

GET /users - list of users -- 200, 404, 500
GET /user/:id - user by id -- 200, 404, 500
POST /user - create user -- 204, 4хх, Header Location: url
PUT  /user/:id - fully update user -- 204/200, 400, 404, 500
PATCH /user/:id - partially update user -- 204/200, 400, 404, 500
DELETE /user/:id - delete user -- 204, 404, 400