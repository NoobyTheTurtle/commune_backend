## Commune Backend

### API Документация

### Первый запуск проект

1. [Установить Docker](https://hub.docker.com/editions/community/docker-ce-desktop-mac)
2. Создать ``.env`` на основе ``.env.example``
3. Проверьте не заняты ли у вас :3000 (сервер) и :5432 (база) порты
4. ``./commune_backend.sh start`` - собираем контейнеры с сервером и базой данных
5. Ждём пока база и сервер соберутся
6. http://localhost:3000

#### [Документация к API](http://localhost:3000/swagger/index.html)

#### Другие команды

- ``./commune_backend.sh stop`` - остановить контейнеры
- ``./commune_backend.sh rmi`` - удалить контейнеры/images
- ``./commune_backend.sh start_db`` - собрать только базу
- ``./commune_backend.sh swagger`` - обновить документацию в Swagger