<div align="center"> <h1>TEST</h1></div>

![Downloads](https://img.shields.io/github/downloads/BeNDYGo/PredProf/total.svg)
![GitHub Repo stars](https://img.shields.io/github/stars/BeNDYGo/PredProf)

- ![TG-img](https://cdn-icons-png.freepik.com/16/15047/15047595.png) [Pipodripo](https://t.me/pipodripo)
- ![Gmail-img](https://cdn-icons-png.freepik.com/16/5968/5968534.png?ga=GA1.1.1230537149.1769259151) bendygo6@gmail.com

Веб-приложение для подготовки к экзаменам по русскому языку и математике с системой заданий, PvP-режимом и панелью администратора.

## Возможности

- **Каталог заданий** — просмотр и фильтрация задач по предмету, типу и сложности
- **PvP-режим** — соревновательные матчи 1 на 1 в реальном времени через WebSocket
- **Рейтинговая система** — Elo-рейтинг с подсчётом побед и поражений
- **Панель администратора** — добавление задач, управление ролями пользователей, просмотр информации о пользователях
- **Регистрация и авторизация** — система аккаунтов с хешированием паролей (bcrypt)

## Технологии

### Backend

| Технология | Назначение |
|---|---|
| Go | Серверная логика |
| `net/http` | HTTP-сервер |
| SQLite3 (`go-sqlite3`) | База данных |
| Gorilla WebSocket | PvP в реальном времени |
| `golang.org/x/crypto` | Хеширование паролей (bcrypt) |

### Frontend

| Технология | Назначение |
|---|---|
| HTML / CSS / JS | Интерфейс (без фреймворков) |
| WebSocket API | PvP-матчи |
| LocalStorage | Хранение сессии |
| SVG | Визуализация статистики |

## Структура проекта

```
PredProf/
├── Back/                          # Серверная часть (Go)
│   ├── main.go                    # Точка входа, маршрутизация
│   ├── handlers/                  # Обработчики запросов
│   │   ├── admin.go               # Админ-эндпоинты
│   │   ├── login.go               # Авторизация
│   │   ├── register.go            # Регистрация
│   │   ├── tasks.go               # Работа с заданиями
│   │   └── pvp.go                 # PvP-матчи (WebSocket)
│   ├── middleware/
│   │   └── auth.go                # Проверка прав администратора
│   └── databases/
│       ├── tasks.json             # Данные задач для импорта
│       ├── tasksDatabase/         # БД заданий (русский, математика)
│       └── usersDatabase/         # БД пользователей
└── Front/                         # Клиентская часть
    ├── main.html                  # Главная — статистика пользователя
    ├── login.html                 # Страница входа
    ├── register.html              # Страница регистрации
    ├── tasks.html                 # Каталог заданий
    ├── pvp.html                   # PvP-режим
    ├── admin.html                 # Панель администратора
    ├── js/                        # Скрипты
    │   ├── components.js          # Общие компоненты (хедер, навигация)
    │   ├── auth.js                # Проверка авторизации
    │   └── ...                    # Скрипты для каждой страницы
    └── css/                       # Стили
```

## API

### Публичные эндпоинты

| Метод | URL | Описание |
|---|---|---|
| `POST` | `/api/register` | Регистрация (username, email, password) |
| `POST` | `/api/login` | Авторизация (username, password) |
| `GET` | `/api/getAllTasks` | Получение задач (query: `subject`, `taskType`, `difficulty`) |
| `GET` | `/api/userInfo` | Информация о пользователе (query: `username`) |
| `WS` | `/api/ws` | WebSocket для PvP-матчей (query: `username`) |

### Эндпоинты администратора

| Метод | URL | Описание |
|---|---|---|
| `POST` | `/api/addTask` | Добавление задачи (query: `subject`) |
| `GET` | `/api/getUserAllInfo` | Полная информация о пользователе |
| `GET` | `/api/changeRole` | Изменение роли пользователя |

Админ-эндпоинты требуют заголовок `X-Username` с именем пользователя-администратора.

## База данных

### Пользователи (`users`)

| Поле | Тип | Описание |
|---|---|---|
| `username` | TEXT (PK) | Имя пользователя |
| `email` | TEXT | Email |
| `password` | TEXT | Хеш пароля (bcrypt) |
| `role` | TEXT | Роль (`student` / `admin`) |
| `rating` | INT | Elo-рейтинг (по умолчанию 1000) |
| `wins` | INT | Количество побед |
| `losses` | INT | Количество поражений |

### Задания (`tasks`)

| Поле | Тип | Описание |
|---|---|---|
| `id` | INTEGER (PK) | Идентификатор |
| `task` | TEXT | Текст задания |
| `answer` | TEXT | Правильный ответ |
| `taskType` | TEXT | Тип задания |
| `difficulty` | TEXT | Сложность (`easy`, `medium`, `hard`) |

Задания хранятся в двух раздельных БД: `tasks_rus.db` (русский язык) и `tasks_math.db` (математика).

## PvP-система

- Поиск соперника через очередь ожидания
- Оба игрока получают одинаковое задание
- Побеждает первый, кто даст правильный ответ
- Рейтинг обновляется по формуле Elo (K-фактор = 32)

## Запуск

### Требования

- Go 1.25+
- Поддержка CGo (для SQLite3)

### Запуск сервера

```bash
cd Back
go mod download
go run main.go
```

Сервер запустится на `http://localhost:8080`.

Базы данных и таблицы создаются автоматически при первом запуске.

### Открытие клиента

Откройте файл `Front/main.html` в браузере или используйте любой локальный HTTP-сервер для раздачи статических файлов из папки `Front/`.

# Связь со мной!
- ![TG-img](https://cdn-icons-png.freepik.com/16/15047/15047595.png) [Pipodripo](https://t.me/pipodripo)
- ![Gmail-img](https://cdn-icons-png.freepik.com/16/5968/5968534.png?ga=GA1.1.1230537149.1769259151) bendygo6@gmail.com
