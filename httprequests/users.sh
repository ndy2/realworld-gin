# Authentication
curl -X POST http://localhost:8080/api/users/login \
     -d '{
          "user": {
              "email": "jake@jake.jake",
              "password": "jakejake"
          }
     }'

# Registration
curl -X POST http://localhost:8080/api/users \
     -d '{
           "user":{
             "username": "Jacob",
             "email": "jake@jake.jake",
             "password": "jakejake"
           }
         }