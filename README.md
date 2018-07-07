# linkaggregator

## Usercases

## API

Все удачные ответы имеют структуру: 

```json
{
    "success": true
}
```

или

```json
{
    "success": true,
    "item": {
        "some":"field",
        "another":"field"
    }
}
```

или

```json
{
    "success": true,
    "items": [{
        "some":"field",
        "another":"field"
      },{
        "some2":"field2",
        "another2":"field2"
      }]    
}
```

Если происходит какая-либо ошибка, ответ имеет вид: 

```json
{
    "success": false,
    "message": "Empty field"
}
```

#### Security section
Все запросы имеют базовый префикс **/security**

##### Principal
Метод возвращает базовую информацию о пользователе (такую как id или nickname)

**URL** : `/`

**Method** : `GET`

**Auth required** : YES

**Code** : `200 OK`

**Response**

```json
{
    "id": 1234,
    "nickname": "John Doe",
    "email": "hey@example.com"
}
```

##### Login
В случае успешной авторизации, в куку **X-AUTH-TOKEN** кладется подписанный сервером JWT токен

**URL** : `/login`

**Method** : `POST`

**Auth required** : YES

**Code** : `200 OK`

**Request**

```json
{
    "login":"John",
    "password":"qwerty"
}
```

**Response**

```json
{
    "success":"true"
}
```

##### Reg
Регистрирует пользователя на сайте

**URL** : `/register`

**Method** : `POST`

**Auth required** : YES

**Code** : `200 OK`

**Request**

```json
{
    "login":"John",
    "password":"qwerty",
    "email": "john@example.org"
}
```

**Response**

```json
{
    "success":"true"
}
```

##### Logout
Чистит куку **X-AUTH-TOKEN**

**URL** : `/logout`

**Method** : `GET`

**Auth required** : YES

**Code** : `200 OK`

**Response**

```json
{
    "success":"true"
}
```
## сборка и запуск п
Backend
'''bat

docker build -f  dockerfile -t linkaggregator:<ветка> .
 docker run -rm --name linkaggregator linkaggregator:<ветка>/bin/ash