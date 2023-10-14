## Commune Backend

### Первый запуск проект

1. [Установить Docker](https://hub.docker.com/editions/community/docker-ce-desktop-mac)
2. Проверьте не заняты ли у вас :3000 (сервер) и :5432 (база) порты, при необходимости меняем их в `.env`
3. `./commune_backend.sh start` - собираем контейнеры с сервером и базой данных
4. Ждём пока база и сервер соберутся
5. [Документация к API](http://localhost:3000/swagger/index.html)

#### Другие команды

- `./commune_backend.sh stop` - остановить контейнеры
- `./commune_backend.sh rmi` - удалить контейнеры/images (советую и папку `postgres` удалить)
- `./commune_backend.sh start_db` - собрать только базу
- `./commune_backend.sh swagger` - обновить документацию в Swagger