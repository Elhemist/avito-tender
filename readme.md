# Avito tender
Решение тестового задания для Avito на позицию Intern Backend разработчик
# Инструкция по запуску
Изначальное задание было с условием поделючения к уже существующей бд, связи с чем просто запускаем проект в docker.
```bash
docker build -t tender .
docker run
```
/////////////////////////////
# API
Реализованы все методы для публикации и управленем тендарами и предложениями
## Проверка доступности сервера
**Эндпоинт:** GET "/api/ping"
Возвращает OK.
### Пример
```bash
curl --request GET \
  --url http://localhost:8080/api/tenders/
```
## Получение списка тендеров
**Эндпоинт:** GET "/api/tenders"
Возвращает список всех тендеров.
### Пример
```bash
curl --request GET \
  --url http://localhost:8080/api/tenders/
```
## Создание нового тендера
**Эндпоинт:** POST "/api/tenders/new/"
Данный метод принимает параметры тендера и возвращает получившиуюся запись.
Публикация тендера не проходит, если в базе отношений сотрудников и организаций не имеется записи в которой содержится данные organizationId и creatorUsername.
### Пример
```bash
curl --request POST \
  --url http://localhost:8080/api/tenders/new \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Тендер 3",
    "description": "Описание тендера",
    "service_type": "Construction",
    "status": "Open",
    "organizationId": "00000000-0000-0000-0000-000000000000",
    "creatorUsername": "ivanov"
  }'
```
## Получение списка тендеров
**Эндпоинт:** GET "/api/tenders/my"
Возвращает список тендеров пользователя по параметру username.
### Пример
```bash
curl --request GET \
  --url 'http://localhost:8080/api/tenders/my?username=ivanov'
```
## Редактирование тендера
**Эндпоинт:** PATCH  "/api/tenders/{tenderId}/edit"
Изменение параметров тендера по id.
### Пример
```bash
curl --request PATCH \
  --url http://localhost:8080/api/tenders/1/edit \
  --header 'Content-Type: application/json' \
  --data '{
  "name": "Обновленный Тендер 5",
  "description": "Обновленное описание"
}'
```
## Откат версии тендера
**Эндпоинт:** PUT "/api/tenders/{tenderId}/rollback/{version}"
Изменение параметров тендера по id на одну из ранее существующих версий. Новой записи в истории при изменении версии не создаётся.
### Пример
```bash
curl --request PUT \
  --url http://localhost:8080/api/tenders/1/rollback/2
```

## Создание нового предложения
**Эндпоинт:** POST "/api/bids/new'
Создает новое предложение для существующего тендера..
### Пример
```bash
curl --request POST \
  --url http://localhost:8080/api/bids/new \
  --header 'Content-Type: application/json' \
  --data '{
  "name": "Предложение 1",
  "description": "Описание предложения",
  "status": "Submitted",
  "tenderId": 1,
  "organizationId": "00000000-0000-0000-0000-000000000000",
  "creatorUsername": "ivanov"
}'
```
## Получение списка предложений пользователя
**Эндпоинт:** GET "/api/bids/my"
Возвращает список предложений созданных пользователем.
### Пример
```bash
curl --request GET \
  --url 'http://localhost:8080/api/bids/my?username=ivanov'
```
## Получение списка предложений для тендера
**Эндпоинт:** GET "/api/bids/my"
Возвращает список предложений созданных для данного тендера.
### Пример
```bash
curl --request GET \
  --url http://localhost:8080/api/bids/1/list
```
## Редактирование предложения
**Эндпоинт:** PATCH "/api/bids/{bidId}/edit"
Редактирование параметров для существующего предложения.
### Пример
```bash
curl --request PATCH \
  --url http://localhost:8080/api/bids/1/edit \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/10.0.0' \
  --data '{
  "name": "Обновленное Предложение 4",
  "description": "Обновленное описание"
}'
```
## Откат версии предложения
**Эндпоинт:** PUT "/api/bids/{bidId}/rollback/{version}"
Редактирование параметров для существующего предложения.
### Пример
```bash
curl --request PUT \
  --url http://localhost:8080/api/bids/1/rollback/3
```