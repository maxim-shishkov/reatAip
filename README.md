# reatAip
Небольшое REST API на Go Lang.

Необходимо написать небольшое REST API на Go Lang. <br>
Ниже будут представлены эндпойнты с их описанием и форматом отображения. <br>
Все обращения в JSON, возвращаемый Content-Type application/json. <br>
Для http сервера необходимо использовать github.com/labstack/echo. <br>
Для создания запросов использовать net/http.<br>
GET /data Возвращает данные с ресурса https://randomuser.me/api/<br>
Добавить GET параметры from и to (time.Time), которые будут фильтровать список по дате registered.date<br>
Возвращаемый формат:Status: 200Body: <br>
{<br>
  "data": [<br>
    {<br>
      "gender": "male",<br>
      "first_name": "Joshua",<br>
      "last_name": "Davies",<br>
      "postcode": 24534,<br>
      "created_at": "2012-11-09T07:47:23.904Z"<br>
    },<br>
    {<br>
      "gender": "female",<br>
      "first_name": "John",<br>
      "last_name": "Doe",<br>
      "postcode": 245223,<br>
      "created_at": "2018-11-09T07:47:23.904Z"<br>
    },<br>
  ]<br>
}<br>
POST /dataПринимает параметры from и to (time.Time). <br>
Внутри метода исполняется POST запрос на любой адрес. <br>
Возвращаемый формат:Status: 200Body:<br>
{<br>
  "status": "Success",<br>
  "from": "{from}",<br>
  "to": "{to}"<br>
}<br>
