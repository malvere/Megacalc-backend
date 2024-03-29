
```
    8 8          ,ggg, ,ggg,_,ggg,                                        ,gggg,                               
 ad88888ba      dP""Y8dP""Y88P""Y8b                                     ,88"""Y8b,             ,dPYb,          
d8" 8 8 "8b     Yb, `88'  `88'  `88                                    d8"     `Y8             IP'`Yb          
Y8, 8 8          `"  88    88    88                                   d8'   8b  d8             I8  8I          
`Y8a8a8a,            88    88    88                                  ,8I    "Y88P'             I8  8'          
  `"8"8"8b,          88    88    88   ,ggg,     ,gggg,gg    ,gggg,gg I8'             ,gggg,gg  I8 dP    ,gggg, 
    8 8 `8b          88    88    88  i8" "8i   dP"  "Y8I   dP"  "Y8I d8             dP"  "Y8I  I8dP    dP"  "Yb
Y8a 8 8 a8P          88    88    88  I8, ,8I  i8'    ,8I  i8'    ,8I Y8,           i8'    ,8I  I8P    i8'      
 "Y88888P"           88    88    Y8, `YbadP' ,d8,   ,d8I ,d8,   ,d8b,`Yba,,_____, ,d8,   ,d8b,,d8b,_ ,d8,_    _
    8 8              88    88    `Y8888P"Y888P"Y8888P"888P"Y8888P"`Y8  `"Y8888888 P"Y8888P"`Y88P'"Y88P""Y8888PP
                                                    ,d8I'                                                      
                                                  ,dP'8I                                                       
                                                 ,8"  8I                                                       
                                                 I8   8I                                                       
                                                 `8, ,8I                                                       
                                                  `Y8P"                                                       
```
---

# MegaCalc (Backend)
![GoLand](https://img.shields.io/badge/GoLand-0f0f0f?&style=for-the-badge&logo=goland&logoColor=white)
![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

MegaCalc is a Golang backend service designed to facilitate convenient price calculations for users. It manages authentication for a Telegram web app and handles invite codes.

[Frontend repo](https://github.com/malvere/Abuzometer-js)

[RU Docs](.github/README-RU.md)

## Table of Contents

- [Endpoints](#endpoints)
- [Setup](#setup)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Endpoints

### Create Invite Code
- **Method**: GET
- **URL**: `http://localhost:8080/secure/code?code=t3s1c0d3`
- **Description**: Creates an invite code. By default the code is activated.
- **Response**:
  ```json
  {
    "code_id": "ed2b3249-3424-47bb-9639-ac5e6dcb3be7",
    "code": "t3s1c0d3",
    "active": true
  }
  ```

### Create User
- **Method**: GET
- **URL**: `http://localhost:8080/user?tgid=12345678&code=t3s1c0d3`
- **Description**: Creates a user and associates it with an invite code. Invite code is automatically set as `"active": false` in DB.
- **Response**:
  ```json
  {
    "user_id": "b316a917-0652-4f3c-aa7f-375086ead3dd",
    "telegram_id": "12345678",
    "invite_code_id": "ed2b3249-3424-47bb-9639-ac5e6dcb3be7"
  }
  ```

### Delete User
- **Method**: DELETE
- **URL**: `http://localhost:8080/secure/user`
- **Body**:
  ```json
  {
    "telegram_id": "12345678"
  }
  ```
- **Description**: Deletes a user by their Telegram ID.
- **Response**:
  ```json
  "User deleted! Telegram ID: 12345678"
  ```

### List All Codes
- **Method**: POST
- **URL**: `http://localhost:8080/secure/list-all-codes`
- **Body**:
  ```json
  {
    "page": 0
  }
  ```
- **Description**: Retrieves a list of all invite codes. Page contains 20 entries.
- **Response**:
  ```json
  [
    {
      "code_id": "b236f25d-634b-4f8d-993f-f86939fb52f8",
      "code": "test1",
      "active": true
    },
    {
      "code_id": "f22b1c02-7273-4bcc-a02a-af9c8a5b3315",
      "code": "test2",
      "active": true
    }
  ]
  ```

### Promocodes

List all active promocodes
- **Method**: GET
- **URL**: `http://localhost:8080/promo/code?state={Bool}`
- **Description**: Deletes a user by their Telegram ID.
- **Response**:
  ```json
  [
      {
        "promo_id": "e548b625-0387-4995-b065-62280d742456",
        "promo_name": "ДАША",
        "promo": "0/0;1000/6000;5000/30000;10000/55000;20000/110000::ДАША",
        "active": true
      },
      {
        "promo_id": "2ddd69d6-1260-46b6-8955-128f060bade8",
        "promo_name": "ЖАЛКОЧТОЛИ",
        "promo": "0/0;2000/18000;5000/45000;9000/82000;12000/110000::ЖАЛКОЧТОЛИ",
        "active": true
      },
      {
        "promo_id": "4804c5a3-d38a-438e-b5cd-b17bbc2dcb58",
        "promo_name": "ЛАДНОКУПЛЮ",
        "promo": "0/0;1000/11000;2000/18000;3000/27000;4000/40000::ЛАДНОКУПЛЮ",
        "active": true
      }
  ]
  ```

Create new promocode
- **Method**: POST
- **URL**: `http://localhost:8080/promo/code`
- **Body**:
  ```json
  {
    "promo_name": "ПРОМИК",
    "promo_string": "0/0;1000/2000;5000/20000;10000/35000;20000/85000::ПРОМИК",
    "state": true
  }
  ```
- **Description**: Retrieves a list of all invite codes. Page contains 20 entries.
- **Response**:
  ```json
  {
    "promo_id": "b236da5d-634b-4t8d-993f-f8697gfb52f8",
    "promo_name": "ПРОМИК",
    "promo": "0/0;1000/2000;5000/20000;10000/35000;20000/85000::ПРОМИК",
    "active": true
  }
  ```

## Setup
1. Clone the repository: `git clone https://github.com/malvere/Megacalc-backend`
2. Install dependencies: `go mod download`
3. Build the project: `go build`
4. Run the server: `./megacalc-backend`

### Environmental variables
- `BOT_TOKEN=0123456789:ABCD...` - Telegram bot token from @botfather
- `CHAT_ID=-1000123456789` - ChatID, participants will be able to access the app without invite codes (bot must be able to retrieve chat memeber list, e.g. have admin priveleges)
- `CONFIG_PATH=./configs/local.yaml` - Config file path
- `DB_URL="host=localhost dbname=megacalc sslmode=disable"` - Databse connection string
- `ORIGIN_ALLOWED=https://abuzometer.js` - Frontend domain (for CORS policy)
- `PROMO_TOKEN=promik` - Bearer-token for promo-codes management
- `SECRET_TOKEN=0123456789abcd` - Bearer-token for invite-ode management

## Dependencies
- [Golang](https://golang.org/): The programming language used for the backend.


## Usage
1. Create an invite code using the `/secure/code` endpoint.
2. Create a user using the `/user` endpoint with the Telegram ID and invite code.
3. Retrieve a list of all invite codes using the `/secure/list-all-codes` endpoint.

---
