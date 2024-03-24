
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

MegaCalc - серверная часть калькулятора выгоды для СберМегаМаркета, отвечает за аутентификацию пользователей в TWA-приложении.

[Frontend repo](https://github.com/malvere/Abuzometer-js)

## Table of Contents

- [Эндпоинты](#endpoints)
- [Запуск](#setup)
- [Зависимости](#dependencies)
- [Использование](#usage)

## Эндпоинты

### Создание инвайт-кодов
- **Method**: GET
- **URL**: `http://localhost:8080/secure/code?code=t3s1c0d3`
- **Description**: Создаёт инвайт код в базе данных, по умолчанию код помечается как активный.
- **Response**:
  ```json
  {
    "code_id": "ed2b3249-3424-47bb-9639-ac5e6dcb3be7",
    "code": "t3s1c0d3",
    "active": true
  }
  ```

### Создание пользователя
- **Method**: GET
- **URL**: `http://localhost:8080/user?tgid=12345678&code=t3s1c0d3`
- **Description**: СОздаёт пользователя и добалвяет к нему активный инвайт-код. Поле `active` у инвайт-кода помечается как `false` в БД.
- **Response**:
  ```json
  {
    "user_id": "b316a917-0652-4f3c-aa7f-375086ead3dd",
    "telegram_id": "12345678",
    "invite_code_id": "ed2b3249-3424-47bb-9639-ac5e6dcb3be7"
  }
  ```

### Удаление пользователя
- **Method**: DELETE
- **URL**: `http://localhost:8080/secure/user`
- **Body**:
  ```json
  {
    "telegram_id": "12345678"
  }
  ```
- **Description**: Удаляет пользователя по `TelegramID`
- **Response**:
  ```json
  "User deleted! Telegram ID: 12345678"
  ```

### Список всех инвайт-кодов
- **Method**: POST
- **URL**: `http://localhost:8080/secure/list-all-codes`
- **Body**:
  ```json
  {
    "page": 0
  }
  ```
- **Description**: Выводит список всех инвайт-кодов по 20 штук на странице.
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

### Промокода

Список активных промокодов
- **Method**: GET
- **URL**: `http://localhost:8080/promo/code?state={Bool}`
- **Description**: Выводит список активных промокодов.
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

Создание нового промокода
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
- **Description**: Создаёт новый промокод в базе.
- **Response**:
  ```json
  {
    "promo_id": "b236da5d-634b-4t8d-993f-f8697gfb52f8",
    "promo_name": "ПРОМИК",
    "promo": "0/0;1000/2000;5000/20000;10000/35000;20000/85000::ПРОМИК",
    "active": true
  }
  ```

## Запуск
1. Клонирование репозитория: `git clone https://github.com/malvere/Megacalc-backend`
2. Установка пакетов: `go mod download`
3. Сборка: `go build`
4. Запуск: `./megacalc-backend`

### Переменные окружения
- `BOT_TOKEN=0123456789:ABCD...` - Токен телеграм бота
- `CHAT_ID=-1000123456789` - ID Чата, участникам которого будет доступен калькулятор (бот должен иметь дотуп к списку участников)
- `CONFIG_PATH=./configs/local.yaml` - Путь к файлу конфигурации
- `DB_URL="host=localhost dbname=megacalc sslmode=disable"` - Строка для подклбчения к БД
- `ORIGIN_ALLOWED=https://abuzometer.js` - Домен где развёрнута фронтовая часть калькулятора
- `PROMO_TOKEN=promik` - Bearer-токен для доступа к промокодам
- `SECRET_TOKEN=0123456789abcd` - Bearer-токен для доступа к инвайт-кодам

## Dependencies
- [Golang](https://golang.org/).


## Использование
1. Создайте промкода через эндпоинт `/secure/code`.
2. Добавить пользователя черещ `/user` с Telegram ID и инвайт-кодом.
3. Список всех инвайт кодов можно получить через `/secure/list-all-codes`.

---
