#!/bin/bash -x

# Run the client
echo Users
go run ../hw15_go_sql/client/ -url "http://localhost" -path "users" -method "get"
go run ../hw15_go_sql/client/ -url "http://localhost" -path "users" -method "put" \
-body '{"id":11,"name":"Tim","email":"tim@example.com","password":"qwerty1"}'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "users" -method "put" \
-body '{"id":12,"name":"Kate","email":"kate@example.com","password":"qwerty"}'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "users" -method "post" \
-body '{"id":12,"name":"Kate","email":"kate@example.com","password":"eFdy453fePt"}'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "users" -method "delete" -body '2'
go run ../hw15_go_sql/client/ -url "http://localhost" -path 'users/stat' -body '1'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "users"

echo Products
go run ../hw15_go_sql/client/ -url "http://localhost" -path "products" -method "get"
go run ../hw15_go_sql/client/ -url "http://localhost" -path "products" -method "put" \
-body '{"id":8,"name":"Juice","price":5.37}'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "products" -method "post" \
-body '{"id":4,"name":"Apple","price":0.8}'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "products" -method "delete" -body '5'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "products" -method "get"

echo Orders
go run ../hw15_go_sql/client/ -url "http://localhost" -path "orders" -method "get" -body '5'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "orders" -method "put" \
-body '{"id":10,"user_id":3,"product_id":8,"amount":3}'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "orders" -method "delete" \
-body '55'
go run ../hw15_go_sql/client/ -url "http://localhost" -path "orders" -method "get" -body '3'