# Zoo API 🦁

---

## Функциональность

| Функциональность                            | Хэндлер и сервис                               | Метод и роут                                 |
|---------------------------------------------|------------------------------------------------|----------------------------------------------|
| Получение всех животных                     | `AnimalHandler`, `AnimalService`               | `GET /animals`                               |
| Получение животного по ID                   | `AnimalHandler`, `AnimalService`               | `GET /animals/:id`                           |
| Создание и удаление животных                | `AnimalHandler`, `AnimalService`               | `POST /animals`, `DELETE /animals/:id`       |
| Перемещение животного между вольерами       | `AnimalHandler`, `AnimalTransportService`      | `POST /animals/:id/transport`                |
| Получение всех вольеров                     | `EnclosureHandler`, `EnclosureService`         | `GET /enclosures`                            |
| Получение вольера по ID                     | `EnclosureHandler`, `EnclosureService`         | `GET /enclosures/:id`                        |
| Создание и удаление вольеров                | `EnclosureHandler`, `EnclosureService`         | `POST /enclosures`, `DELETE /enclosures/:id` |
| Создание расписания кормления               | `FeedingHandler`, `FeedingOrganizationService` | `POST /schedules`                            |
| Получение всех расписаний кормления         | `FeedingHandler`, `FeedingOrganizationService` | `GET /schedules`                             |
| Получение расписания кормления по животному | `FeedingHandler`, `FeedingOrganizationService` | `GET /animals/:id/schedules`                 |
| Выполнение кормления                        | `FeedingHandler`, `FeedingOrganizationService` | `POST /schedules/:id/feed`                   |
| Получение статистики по животным            | `ZooStatisticsHandler`, `ZooStatisticsService` | `GET /statistics/animals`                    |
| Получение статистики по вольерам            | `ZooStatisticsHandler`, `ZooStatisticsService` | `GET /statistics/enclosures`                 |
| Получение статистики по кормлениям          | `ZooStatisticsHandler`, `ZooStatisticsService` | `GET /statistics/schedules`                  |

---

## DDD и Clean Architecture

### Разделение слоёв:
- `domain` — бизнес-сущности (`Animal`, `Enclosure`, `FeedingSchedule`), события (`AnimalMovedEvent`, `FeedingTimeEvent`)
- `application/port/in` — входные порты (интерфейсы use-case'ов)
- `application/port/out` — выходные порты (репозитории, event dispatcher)
- `application/service` — реализация бизнес-логики (сервисы)
- `infrastructure/repository` — in-memory репозитории
- `infrastructure/dispatcher` — реализация диспетчера событий
- `presentation/handler` — gin-хэндлеры
- `presentation/router` — маршруты

### Агрегаты с бизнес-логикой:
- `Animal` - (`MoveTo`, `Feed`, `Treat`) 
- `Enclosure` - (`IsFull`, `RemoveAnimal`, `AddAnimal`, `Clean`) 
- `FeedingSchedule` — (`MarkAsOccurred`,`ChangeScheduleTime`,`IsOccurred`)

### Value Object:
- `Size`, `Gender`, `HealthStatus`, `EnclosureType`

### Доменные события:
- `AnimalMovedEvent` — при перемещении животного
- `FeedingTimeEvent` — при кормлении
- Обрабатываются через `EventDispatcher`

### Порты и адаптеры:
- Входные порты: `AnimalUseCase`, `FeedingOrganizationUseCase`, `EnclosureUseCase`, `ZooStatisticsUseCase`
- Выходные порты: `AnimalRepository`, `EnclosureRepository`, `FeedingScheduleRepository`, `EventDispatcher`


## Запуск
```bash
go run ./zoo/cmd/main.go
```
Swagger со всеми ручками и их описанием доступен по адресу: http://localhost:8080/swagger/index.html
