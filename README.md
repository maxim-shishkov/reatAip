# reatAip
Небольшое REST API на Go Lang.

Необходимо написать небольшое REST API на Go Lang. 
Ниже будут представлены эндпойнты с их описанием и форматом отображения. 
Все обращения в JSON, возвращаемый Content-Type application/json. 
Для http сервера необходимо использовать github.com/labstack/echo. 
Для создания запросов использовать net/http.
GET /data Возвращает данные с ресурса https://randomuser.me/api/
Добавить GET параметры from и to (time.Time), которые будут фильтровать 
список по дате registered.date
Возвращаемый формат:Status: 200Body: 
{
  "data": [
    {
      "gender": "male",
      "first_name": "Joshua",
      "last_name": "Davies",
      "postcode": 24534,
      "created_at": "2012-11-09T07:47:23.904Z"
    },
    {
      "gender": "female",
      "first_name": "John",
      "last_name": "Doe",
      "postcode": 245223,
      "created_at": "2018-11-09T07:47:23.904Z"
    },
  ]
}
POST /dataПринимает параметры from и to (time.Time). 
Внутри метода исполняется POST запрос на любой адрес.
Возвращаемый формат:Status: 200Body:
{
  "status": "Success",
  "from": "{from}",
  "to": "{to}"
}
