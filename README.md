## Commune Backend

### Описание проект

Проект представляет собой API для поиска и получения информации о банкоматах и отделениях банков в радиусе пользователя. API предоставляет следующие функции:
- Получение банкоматов/отделений банков в радиусе пользователя с возможностью фильтрации по различным параметрам.
- Получение информации о банкомате или отделении банка по ID.
- Получение талона для отделения.
- Получение талонов пользователя.

### Первый запуск
1. Склонируйте репозиторий: `git clone https://github.com/NoobyTheTurtle/commune_backend.git`
2. Перейдите в каталог проекта: `cd commune_backend`
3. Установите [Docker и Docker Compose](https://docs.docker.com/compose/install/)
4. Проверьте не заняты ли у вас :3000 (сервер) и :5432 (база) порты, при необходимости меняем их в `.env`
5. Соберите контейнеры с сервером и базой данных: `./commune_backend.sh start`
6. Ждём пока база и сервер соберутся
7. [Swagger документация](http://localhost:3000/swagger/index.html)

#### Доп. команды

- `./commune_backend.sh start` - запустить контейнеры
- `./commune_backend.sh stop` - остановить контейнеры
- `./commune_backend.sh rmi` - удалить контейнеры/images (советую и папку `postgres` удалить)
- `./commune_backend.sh start_db` - собрать только базу
- `./commune_backend.sh swagger` - обновить документацию в Swagger

#### ENV переменные

- `SERVER_PORT` - Порт для сервера
- `DB_NAME` - Название базы данных
- `DB_PORT` - Порт для базы данных
- `POSTGRES_PASSWORD` - Пароль для пользователя `postgres`
- `GIN_MODE` - режим сервера (cм. `.env.example`)
- `DB_LOG_LEVEL` - уровень логирования базы данных (cм. `.env.example`)

### Технологии
 - Golang
 - PostgreSQL