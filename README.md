# webTink
API
Данные окружения находятся в папке config.yml, так же необходимо создать файл окружения в котором храниться пароль(DB_PASSWORD, SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD). По умолчанию стартует на 8000 порту.
RESOURCE	HTTP METHOD	ROUTE	DESCRIPTION
auth	Post	/auth/sign-up	Регистрация пользователя
auth	Post	/auth/sign-in	Вход в аккаунт
auth	Post	/auth/log-out	Выход из учётной записи
auth	Post	/auth/refresh	Получение новых токенов(refresh, access)
api-guest	Get	/api-guest/posts	Получение всех карточек участников(для всех пользователей)
api-guest	Get	/api-guest/posts/:id	Получение конкретной карточки пользователя(для всех пользователей)
api-guest	Get	/api-guest/post/photo/:id	Получение фото карточки(для всех пользователей)
api	Post	/api/create-posts	Создание карточки(для авторизированного пользователя)
api	Put	/api/posts/:id	Обновление карточки пользователя(для авторизированного пользователя)
api	Delete	/api/posts/:id	Удаление карточки пользователя(для авторизированного пользователя)

1.	/auth/sign-up
Тело JSON, которое должно приходить на сервер, от сервера приходит JSON с ID пользователя
{
    "name" : "nfnfnvfdsfn",
    "email" : "tghkfjrjrjlysdfbjrjjfhu@gmail.com",
    "username" : "stazdfiss",
    "surname" : "dfddd",
    "password" : "1234567",
    "grade" : 1
}
2.	/auth/sign-in
Тело JSON, которое должно приходить на сервер, от сервера приходит джейсон с токенами
{
    "email" : "tghkfjrjrjlysdfbjrjjfhu@gmail.com",
    "password" : "1234567"
}
 
3.	/auth/log-out
Ничего отправлять на сервер не нужно
4.	/auth/refresh
Ничего отправлять на сервер не нужно, токены приходят в cookies
 
5.	/api-guest/posts
Плучаете JSON вида 
{
    "data": [
        {
            "id": "4351c301-f397-4bc9-90b9-6a05010a1848",
            "author_id": "b4ca9d58-4548-48ea-9ca8-1dd925bd9d23",
            "name": "hdhdhhd",
            "surname": "kekekek",
            "role": "fdgdgd",
            "education": "fdgdfgs",
            "additional": "fddfgds",
            "created_at": "2022-11-25T11:07:46.56Z",
            "updated_at": "2022-11-25T11:07:46.56Z"
        },
        {
            "id": "68668a85-e564-4c47-ab59-da40d0f020aa",
            "author_id": "842b94ab-ca33-456e-9908-9fb526c11fd7",
            "name": "hdhdhhd",
            "surname": "kekekek",
            "role": "fdgdgd",
            "education": "fdgdfgs",
            "additional": "fddfgds",
            "created_at": "2022-11-25T11:09:55.373Z",
            "updated_at": "2022-11-25T11:09:55.373Z"
        }
    ]
}

6.	/api-guest/posts/:id
Получаете JSON вида
{
    "id": "4351c301-f397-4bc9-90b9-6a05010a1848",
    "author_id": "b4ca9d58-4548-48ea-9ca8-1dd925bd9d23",
    "name": "hdhdhhd",
    "surname": "kekekek",
    "role": "fdgdgd",
    "education": "fdgdfgs",
    "additional": "fddfgds",
    "created_at": "2022-11-25T11:07:46.56Z",
    "updated_at": "2022-11-25T11:07:46.56Z"
}
7.	 /api/create-posts
Отправляете JSON вида, в ответ приходит id поста
{
    "name" : "hdhdhhd",
    "surname" : "kekekek",
    "photo" : "",
    "role" : "fdgdgd",
    "education" : "fdgdfgs",
    "additional" : "fddfgds"
}

8.	 PUT /api/posts/:id
Отправляете JSON вида, в ответ приходит “Status” : “ok”
2.	
3.	    "name" : "hdhdhhd",
4.	    "surname" : "kekekek",
5.	    "photo" : "",
6.	    "role" : "fdgdgd",
7.	    "education" : "fdgdfgs",
8.	    "additional" : "fddfgds"
9.	}
9.	DELETE /api/posts/:id
Отправлять ничего не нужно, в ответ приходит “Status” : “ok”

10.	GET /api-guest/post/photo/:id
Получаете json с полем photo со строкой.

"photo": "{source}",
"name": "Матвей",
"surname": "Серегин",
"role": "Android-разработчик",
"education": "HITs, 1 курс",
"additional": "amogus sus bruh"
