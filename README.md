# Token Lifecycle Lab

Интерактивная учебная CLI-утилита на Go для изучения жизненного цикла token.

Проект последовательно объясняет, как token:

```text
выпускается
→ доставляется
→ хранится
→ передаётся в API
→ используется
→ обновляется или ротируется
→ истекает либо отзывается
```

`v0.1` работает только с демонстрационной моделью в памяти.

```text
Нет реальных token values.
Нет browser integration.
Нет сети.
Нет базы данных.
Нет внешних API или Identity Provider.
```

## Цели проекта

- Понять жизненный цикл access token, refresh token, JWT, API key и browser session
- Изучить Go на маленьком, локальном и законченном CLI-проекте
- Постепенно перейти от линейной демонстрации к сценариям, JWT Inspector, Web UI и browser observation
- Построить живую интерактивную шпаргалку по token lifecycle и security concepts

## Текущая версия

## `v0.1 — CLI Lifecycle Demo`

Пользователь нажимает `Enter`, а утилита последовательно показывает учебный lifecycle token.

На каждой стадии CLI отображает карточку token и отвечает на вопросы:

1. Что сейчас происходит с token?
2. Где token находится?
3. Может ли клиент использовать token сейчас?
4. Что проверит API?
5. Какое следующее событие возможно?

Пример начала walkthrough:

```text
=== Welcome to Token Lifecycle Lab ===

=== Stage 1/7 ===
Token ID: demo-token-001
Token state: Issued
Location: На стороне authorization server после выпуска.
Can client use it now: Нет. Token ещё не передан клиенту.
What will API check: API ещё не получил token.
Next possible event: Доставка token клиенту.
Что происходит: Токен создан и выпущен системой.

Press Enter to continue...
```

## Линейный сценарий v0.1

Текущий walkthrough проходит по заранее определённой последовательности:

```text
Issued
→ Delivered
→ Stored
→ Transmitted
→ Used
→ Rotated
→ Expired
```

`Revoked` уже присутствует в доменной модели как terminal state, но не входит в линейный сценарий `v0.1`.

Причина: `Expired` и `Revoked` — разные финалы lifecycle.

```text
Expired — token стал недействительным естественно:
          наступил exp или закончился TTL.

Revoked — token был досрочно запрещён:
          logout, security incident, смена пароля,
          ручной revoke, отключение account и так далее.
```

Интерактивный выбор веток, включая revoke, появится в `v0.1.1`.

## Восемь учебных стадий

|   № | Стадия                   | Краткое описание                                                                     |
| --: | ------------------------ | ------------------------------------------------------------------------------------ |
|   1 | Выпуск                   | Authorization server создаёт token, назначает ID, TTL, claims и scopes               |
|   2 | Доставка                 | Token передаётся клиенту через response body, cookie или другой разрешённый механизм |
|   3 | Хранение                 | Клиент или сервер хранит token в памяти, session, cookie или другом хранилище        |
|   4 | Передача                 | Клиент отправляет token в API, например через `Authorization: Bearer ...`            |
|   5 | Проверка и использование | API проверяет token и при успешном результате выполняет запрос                       |
|   6 | Обновление и ротация     | Старый token заменяется новым credential или token pair                              |
|   7 | Истечение                | Token становится недействительным после наступления `exp` или окончания TTL          |
|   8 | Отзыв                    | Token досрочно запрещается issuer, authorization server или session registry         |

## Что реализовано

### Доменная модель

- `Token` — Entity с идентификатором `ID` и текущим `State`
- `TokenState` — enum-like тип на основе `string`
- Константы состояний:

```go
Issued
Delivered
Stored
Transmitted
Used
Rotated
Expired
Revoked
```

### CLI walkthrough

- Сценарий задаётся слайсом `lifecycleStages`
- Пользователь нажимает `Enter` для перехода к следующему этапу
- CLI показывает прогресс вида `Stage 3/7`
- Ошибки чтения stdin обрабатываются корректно
- После `Expired` приложение выводит финальное сообщение и завершается

### Карточка token

Для каждого состояния выводятся:

```text
Token ID
Token state
Location
Can client use it now
What will API check
Next possible event
Что происходит
```

### Цвета терминала

| Цвет                   | Значение                                              |
| ---------------------- | ----------------------------------------------------- |
| Cyan                   | Заголовки приложения, номер стадии и финальная строка |
| Yellow                 | Подсказка `Press Enter to continue...`                |
| Orange                 | Terminal state `Expired`                              |
| Red                    | Ошибка ввода и terminal state `Revoked`               |
| Default terminal color | Обычные состояния и пояснения                         |

## Структура проекта

```text
token-lifecycle-lab/
├── go.mod
├── main.go
├── token.go
├── token_state.go
├── ui.go
└── README.md
```

### `main.go`

Точка входа приложения:

- создаёт демонстрационный token;
- содержит `lifecycleStages`;
- читает Enter из stdin;
- запускает линейный walkthrough;
- завершает приложение после `Expired`.

### `token.go`

Доменная Entity:

```go
type Token struct {
    ID    string
    State TokenState
}
```

### `token_state.go`

Доменный тип состояния, константы и учебные пояснения:

```go
type TokenState string
```

Файл содержит:

```text
Description()
Location()
ClientUsage()
APIChecks()
NextEvents()
```

### `ui.go`

Консольное представление:

```text
ANSI-цвета
colorize()
printTokenCard()
```

Все Go-файлы `v0.1` находятся в одной папке и относятся к одному package:

```go
package main
```

## Запуск

Требуется установленный Go.

Из корня проекта:

```bash
go run .
```

## Проверка кода

Перед фиксацией версии полезно выполнить:

```bash
go fmt ./...
go vet ./...
go test ./...
```

- `go fmt` форматирует Go-код
- `go vet` ищет подозрительные конструкции
- `go test` собирает package и запускает тесты, когда они появятся

## Границы v0.1

`v0.1` — учебная модель, а не production authentication system.

Утилита:

- не принимает и не сохраняет реальные token values;
- не расшифровывает и не проверяет реальные JWT;
- не проверяет cryptographic signature;
- не подключается к issuer, OAuth server или API;
- не выполняет introspection;
- не читает browser cookies, localStorage или sessionStorage;
- не является Identity Provider, API Gateway, JWT validator или production auth service.

Важно:

```text
JWT может быть ещё не истёкшим по exp,
но уже быть отозванным server-side.
```

Реальный revoke status нельзя достоверно определить только по строке JWT. Для этого нужен issuer, introspection endpoint, deny-list или другое server-side state.

## План развития

| Версия   | Результат                                                                                             |
| -------- | ----------------------------------------------------------------------------------------------------- |
| `v0.1`   | Интерактивное CLI-демо линейного lifecycle walkthrough                                                |
| `v0.1.1` | Branching Lifecycle Lab: пользователь выбирает допустимое следующее событие и результаты validation   |
| `v0.2`   | Scenario Simulator: expiry, refresh, rotation, revoke, storage, `401`, `403` и security scenarios     |
| `v0.3`   | JWT Inspector: parsing JWT, claims, `exp`, `nbf`, `iat`, `jti`, TTL и объяснение decode versus verify |
| `v0.4`   | Локальный Web UI на Go и чистом HTML/CSS/JavaScript                                                   |
| `v0.5`   | Browser Token Discovery для выбранного browser и домена                                               |
| `v1.0`   | Local real-time observer: CLI и Web UI показывают наблюдаемые изменения token                         |

## Статус

**`v0.1` завершена.**

Следующая версия:

```text
v0.1.1 — Validation Branching Lab
```
