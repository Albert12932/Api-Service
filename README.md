# Api-Service
REST API-сервис для работы с вопросами и ответами.
Проект написан на Go, использует PostgreSQL, GORM, Docker Compose и имеет встроенную Swagger-документацию.

Как запустить проект:
1. Установите Docker и Docker Compose
2. Создайте .env в корне проекта и укажите переменные: ```DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=```
3. запустите сервер командой ```docker compose up --build```

После запуска API Swagger доступен по адресу: http://localhost:8080/swagger/index.html
