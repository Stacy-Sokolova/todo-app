# Todo App

REST API для управления задачами

# Для запуска приложения

- определить переменные окружения для запуска сервера и базы данных
- перед запуском приложения впервые необходимо применить миграции к базе данных
```
make migrate
```
- для запуска приложения
```
make run
```

# Примеры запросов

## Добавление новых задач
![Добавление новой задачи](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/post1.png)

![Добавление новой задачи](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/post2.png)

## Получение списка задач
![Получение списка задач](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/getall.png)

## Изменение задачи
![Изменение задачи](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/put.png)

## Проверка изменений
![Проверка изменений](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/getafterchange.png)

## Удаление задачи
![Удаление записи](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/delete.png)

## Проверка изменений
![Проверка изменений](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/getafterdelete.png)

## Итоговая база данных postgres
![Проверка изменений postgres](https://github.com/Stacy-Sokolova/todo-app/blob/main/pngs/posgres.png)