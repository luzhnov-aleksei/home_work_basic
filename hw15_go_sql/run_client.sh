#!/bin/bash -x

# Run the client
echo Users
go run ./client/ -path "users" -method "get"
go run ./client/ -path "users" -method "put" \
-body '{"id":11,"name":"Tim","email":"tim@example.com","password":"qwerty1"}'
go run ./client/ -path "users" -method "put" \
-body '{"id":12,"name":"Kate","email":"kate@example.com","password":"qwerty"}'
go run ./client/ -path "users" -method "post" \
-body '{"id":12,"name":"Kate","email":"kate@example.com","password":"eFdy453fePt"}'
go run ./client/ -path "users" -method "delete" -body '2'
go run ./client/ -path 'users/stat' -body '1'
go run ./client/ -path "users"

echo Products
go run ./client/ -path "products" -method "get"
go run ./client/ -path "products" -method "put" \
-body '{"id":8,"name":"Juice","price":5.37}'
go run ./client/ -path "products" -method "post" \
-body '{"id":4,"name":"Apple","price":0.8}'
go run ./client/ -path "products" -method "delete" -body '5'
go run ./client/ -path "products" -method "get"

echo Orders
go run ./client/ -path "orders" -method "get" -body '5'
go run ./client/ -path "orders" -method "put" \
-body '{"id":10,"user_id":3,"product_id":8,"amount":3}'
go run ./client/ -path "orders" -method "delete" \
-body '55'
go run ./client/ -path "orders" -method "get" -body '3'