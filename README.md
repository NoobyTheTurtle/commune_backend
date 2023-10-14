## Commune Backend

### Первый запуск проект

1. Установить [Docker и Docker Compose](https://docs.docker.com/compose/install/)
2. Проверьте не заняты ли у вас :3000 (сервер) и :5432 (база) порты, при необходимости меняем их в `.env`
3. `./commune_backend.sh start` - собираем контейнеры с сервером и базой данных
4. Ждём пока база и сервер соберутся
5. [Документация к API](http://localhost:3000/swagger/index.html)

#### Другие команды

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