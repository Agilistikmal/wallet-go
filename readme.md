### Wallet GO
Simple RESTful API using GoLang.

---
#### GET /api/user
Get all users data.
###### Example response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 1,
      "name": "Agil Ghani Istikmal",
      "email": "email@gmail.com",
      "phone": "+628123456789",
      "wallet_amount": 15000
    }
  ]
}
```


---
#### GET /api/user/:userId
Get user by id.
###### Example response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "Agil Ghani Istikmal",
    "email": "email@gmail.com",
    "phone": "+628123456789",
    "wallet_amount": 15000
  }
}
```


---
#### POST /api/user
Create user data.
###### Example request:
```json
{
  "name": "Agil Ghani Istikmal",
  "email": "email@gmail.com",
  "phone": "+628123456789",
}
```
###### Example response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "Agil Ghani Istikmal",
    "email": "email@gmail.com",
    "phone": "+628123456789",
    "wallet_amount": 0
  }
}
```


---
#### PUT /api/user/:userId
Update user data.
###### Example request:

```json
{
  "name": "Agil Ghani Istikmal",
  "email": "newemail@gmail.com",
  "phone": "+628123456789",
}
```
###### Example response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "Agil Ghani Istikmal",
    "email": "newemail@gmail.com",
    "phone": "+628123456789",
    "wallet_amount": 0
  }
}
```


---
#### PUT /api/wallet/:userId
Update user wallet amount.
###### Example request:

```json
{
  "amount": 500
}
```
###### Example response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "user_id": 1,
    "wallet_amount": 500
  }
}
```


---
#### DELETE /api/wallet/:userId
Delete user data.
###### Example response:
```json
{
  "code": 200,
  "status": "OK",
  "data": nil
}
```