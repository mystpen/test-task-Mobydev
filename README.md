# The task for MOBYDEV
## Authorization task

- For register:
    ```http
    POST /signup
    ```
     -  Request in JSON format that needs to be sent:
    ```
    {   
        "username": "user",
        "email": "example@example.com",
        "password": "somepassword"
    }
    ```
- Login:

    ```http
    POST /signin
    ```
     - Request in JSON format:
    ```
    {
        "email": "example@example.com",
        "password": "somepassword"
    }
    ```


## Edit info task

- Get user info:
    ```http
    GET /user/info
    ``` 
- Edit user info:
    ```http
    PUT /user/info
    ```
    - Request:
    ```
    {
        "username": "user",
        "phone": "+77057777777"
    }
    ```
- For get user info by ID:
    ```http
    GET /users/2
    ```
    -  Result:
    ```
    {
        "user_info": {
            "id": 2,
            "username": "user",
            "email": "test@example.com",
            "phone": "+77057777777"
        }
    }
    ```

## Admin permission task

- Created route that only the administrator has access to:
    ```http
    PUT /videos/update/:id
    ```
    - Request example:

    ```
    {
        "title": "some title"
		"type": "type"
		"category": "category"
		"year": 2024
		"description": "some description"
    }
    ```

- Administrator login data:
    ```
    {
        "email": "administrator@example.com",
        "password": "administrator"
    }
    ```
---
### Run project:
- Create a configuration file. *See the ```config.example.yaml``` for an example*:
```
touch config/config.yaml
```
- To run POSTGRES container use docker-compose:
```
docker-compose up -d
``` 
- Run:
```
go run ./cmd/api
```
