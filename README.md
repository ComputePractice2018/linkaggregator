# linkaggregator

## Usercases

1. Как пользователь я хочу иметь возможность просматривать новости по тематикам.
1. Как пользователь, я хочу иметь возможность добавлять подписки, чтобы просматривать интересные мне новости.
1. Как пользователь, я хочу иметь возможность редактировать подписки, чтобы не просматривать неактуальную информацию.


## API

##запуск и сборка
Backend:

cd backend
docker build -f Dockerfile -t linkaggrergatorbackend:<имя ветки> .
docker run --rm --name inkaggrergator -e NAME=<параметр приложения> inkaggrergatorbackend:<имя ветки>

## REST API

### GET /api/linkaggregator/news

**Ответ**: 200 OK
```json
[{
    "id": 1, 
    "title": "Заголовок",
    "date": "Дата",
    "url": "Ссылка",
    "themes": [ "Тема" ],
    "content": ""
}]
```

### GET /api/linkaggregator/subscriptions

**Ответ**: 200 OK

После запроса принимается массив, в котором перечислены темы из подписок, вида: ['Тема1', 'Тема2', ...]


### GET /api/linkaggregator/news/content/1

**Ответ**: 200 OK

После запроса принимается строка с содержанием новости вида: "Контент"


### PUT /api/linkaggregator/subscriptions/1

**Тело запроса**: передается строка с темой вида: 'Тема'

**Ответ**: 202 Accepted Location /api/linkaggregator/subscriptions/1


### DELETE /api/linkaggregator/subscriptions/1

Ответ: 204 No Content


