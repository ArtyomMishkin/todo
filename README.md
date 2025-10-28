# Практическое задание 4

## ЭФМО-02-25 Мишкин Артём Дмитриевич 06.10.2025
---
# Информация о проекте
todo - Это веб-приложение для управления задачами (TODO-list) на Go с REST API, которое поддерживает создание, чтение, обновление и удаление задач с дополнительными функциями пагинации, фильтрации и автоматическим сохранением данных в JSON-файл.
## Цели занятия
-	Освоить базовую маршрутизацию HTTP-запросов в Go на примере роутера chi.
- Научиться строить REST-маршруты и обрабатывать методы GET/POST/PUT/DELETE.
- Реализовать небольшой CRUD-сервис «ToDo» (без БД, хранение в памяти).
- Добавить простое middleware (логирование, CORS).
- Научиться тестировать API запросами через curl/Postman/HTTPie

## Файловая структура проекта:

<img width="175" height="261" alt="image" src="https://github.com/user-attachments/assets/5968168c-9c67-49b2-8a3b-db9722a14996" />

## ВАЖНОЕ ПРИМЕЧАНИЕ

Так как практики переехали на сервер, то порты к ним поменялись(для данной практики порт 8101)!

## Примеры запросов:

### /health

<img width="545" height="172" alt="image" src="https://github.com/user-attachments/assets/b3e69062-0db8-4236-b1fd-6aff931dbb7e" />

### Создание таски

<img width="974" height="276" alt="image" src="https://github.com/user-attachments/assets/93caf1b1-d53d-4b70-9782-d81bdd728ccb" />

### Вывод списка

<img width="974" height="234" alt="image" src="https://github.com/user-attachments/assets/030a3c11-a54a-4f0d-8dc8-eb50e53f75a0" />

### Получение таски по ID

<img width="974" height="232" alt="image" src="https://github.com/user-attachments/assets/fa4ad4a6-dd72-429b-a7d6-1b8bbd7b5a3b" />

### Обновление таски

<img width="974" height="274" alt="image" src="https://github.com/user-attachments/assets/81052321-3cbf-4304-8eda-fe7dbac9ea6e" />

### Удаление таски

<img width="1123" height="321" alt="image" src="https://github.com/user-attachments/assets/d256b903-8bac-45d9-8162-dc92ee388197" />


## Домашнее задание

### Пагинация списка

Пагинация списка - разделения большого списка на страницы для удобства пользователья

<img width="974" height="209" alt="image" src="https://github.com/user-attachments/assets/1f24cf57-342d-4aed-9dfd-fa3854f82257" />

### Проверка фильтра по done

Создаём две невые задачи(невыполненные)

<img width="1016" height="239" alt="image" src="https://github.com/user-attachments/assets/decff0a9-e037-41c4-b9cf-fb563c1f7d02" />

Создаём лве выполненные задачи

<img width="1067" height="216" alt="image" src="https://github.com/user-attachments/assets/96ec7088-085e-4ce2-9260-be442a358c47" />

Получаем список невыполненных задач

<img width="1649" height="132" alt="image" src="https://github.com/user-attachments/assets/d9d3b97e-b6f6-44f4-b2d2-3a062b4ce1cb" />

Получаем список выполненных задач

<img width="969" height="92" alt="image" src="https://github.com/user-attachments/assets/135744e5-bf96-4be6-90d7-083a3f3f6391" />

### Файлы сохраняются в файл tasks.json

<img width="507" height="584" alt="image" src="https://github.com/user-attachments/assets/a6ceb012-9f95-4482-9840-85be1ea9cdf4" />

## Объяснение, как обрабатываются ошибки и коды ответа

Ошибки обрабатываются через единую функцию httpError(), которая возвращает JSON с описанием ошибки и соответствующие HTTP-статусы:

- 400 Bad Request - невалидные данные (короткий/длинный title, неверный ID)

- 404 Not Found - задача не найдена

- 201 Created - успешное создание задачи

- 204 No Content - успешное удаление или CORS preflight

- 200 OK - успешные GET/PUT запросы


