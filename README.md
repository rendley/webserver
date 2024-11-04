### Установка утилиты golang-migrate

```
https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
```

### Создание миграционного репозитория
В данном репозитории будут находится up/down пары sql миграционных запросов к бд.
```
migrate create -ext sql -dir migrations UsersCreationMigration
```

### Создание up/down sql файлов
См. ```migrations/....up.sql``` и ```migrations/...down.sql```

### Применить миграцию

```
migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=qwerty" up
```

### Откат миграции
```
Для выполнения отката ```migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down```

```