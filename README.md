# Verx GM and GN Bot

## Функционал

- Генерация сообщений с Добрым утром по шаблонам, со случайными эмодзи
- Генерация сообщений с Доброй ночи по шаблонам, со случайными эмодзи
- Отправка случайных сообщений с Добрым утром через TG аккаунт, каждый день в 10:00
- Отправка случайных сообщений с Доброй ночи через TG аккаунт, каждый день в 23:00

## Быстрый старт

1. Получить TG API_ID, API_HASH [Telegram Apps](https://my.telegram.org/auth?to=apps)
2. Создать .env по примеру .env.example, либо добавить переменные окружения без .env

```bash
export TG_ACC_API_ID=111
export TG_ACC_API_HASH=hash
export TG_ACC_TARGET_USERNAME=@username # кому будет отправляться сообщения
```

3. Подготовка сессии (вход в TG аккаунт)

```bash
make build-cli
```

**Вход без пароля**

```bash
./bin/verx-tg-cli -id <api_id> -h <api_hash> -p <phone_number>
```

**Вход с паролем (2FA)**

```bash
./bin/verx-tg-cli -id <api_id> -h <api_hash> -p <phone_number> -pw
```

4. Запуск

```bash
make run
```

## Сборка

**Для Linux AMD 64**

```bash
make build-linux-amd
```

**Для текущей OS**

```bash
make build
```

**CLI утилита для логина**

```bash
make build-cli
```