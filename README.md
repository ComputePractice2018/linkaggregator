# linkaggregator

## Usercases

1. Как пользователь я хочу иметь возможность искать новости по нужным мне критериям.
1. Как пользователь, я хочу иметь возможность добавлять подписки, чтобы просматривать интересные мне новости.
1. Как пользователь, я хочу иметь возможность редактировать подписки, чтобы не просматривать неактуальную информацию.

## API

### Работа со списком подписок

**Добавить подписку в список**

**URL** : `/api/subscribtion/add`

**Метод** : `POST`

**Данные запроса**

```json
{
    "url": "[валидный путь до rss сайта, с которого нужно получать новости]"
}
```

**Пример запроса**

```json
{
    "url": "https://habr.com/rss"
}
```

**Пример ответа:**

**Code** : `200 OK`

**Content example**

```json
{
    "success": "true"
}
```


**Удалить подписку из списка**

**URL** : `/api/subscribtion/remove`

**Метод** : `POST`

**Данные запроса**

```json
{
    "url": "[валидный путь до rss сайта, подписку на который нужно удалить]"
}
```

**Пример запроса**

```json
{
    "url": "https://habr.com/rss"
}
```

**Пример ответа:**

**Code** : `200 OK`

**Content example**

```json
{
    "success": "true"
}
```

**Получить список подписок**

**URL** : `/api/subscribtion/list`

**Метод** : `GET`

**Параметры запроса**: возможно передать 2 параметра. skip - нижняя граница подписок и limit - верхняя граница подписок

**Пример ответа:**

**Code** : `200 OK`

**Content example**

```json
{
    "success": "true",
    "items": [
        {
            "id": "00db4f604a1068ed4dfc",
            "url":"https://habr.com/rss",
            "title":"Хабр / Интересные публикации",
            "updated": "Mon, 09 Jul 2018 18:33:35"
        },
        {
            "id": "00db4f604a1068ed4dfc",
            "url":"https://overclockers.ru/rss/news.rss",
            "title":"Overclockers.ru / Новости",
            "updated": "Mon, 09 Jul 2018 19:44:12"
        }
    ]
}
```

**Получить ленту новостей**

**URL** : `/api/feed`

**Метод** : `GET`

**Параметры запроса**: возможно передать 2 параметра. skip - нижняя граница подписок и limit - верхняя граница подписок, id - id индивидуально необходимой ленты новостей

**Пример ответа:**

**Code** : `200 OK`

**Content example**

```json
{
    "success": "true",
    "items": [
        {
            "id": "00db4f604a1068ed4dfc",
            "url":"https://overclockers.ru/rss/news.rss",
            "title":"Overclockers.ru / Новости",
            "updated": "Mon, 09 Jul 2018 19:44:12",
            "feed": [
                {
                    "title": "Россиянин обновил рекорды GPUPI при помощи дуэта NVIDIA TITAN Xp",
                    "pubDate": "2018-07-09 05:54:09",
                    "link": "https://overclockers.ru/hardnews/show/92459/rossiyanin-obnovil-rekordy-gpupi-pri-pomoschi-dueta-nvidia-titan-xp",
                    "thumbnail": "https://st.overclockers.ru/images/news/2018/07/09/nazar_01.jpg",
                    "content": "Холодный душ в летнюю жару приятен и компьютерным компонентам."
                }
            ]
        }
    ]
}
```