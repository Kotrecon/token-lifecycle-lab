# Token Lifecycle Lab

Интерактивная учебная CLI-утилита на Go для изучения жизненного цикла token.

Проект последовательно объясняет, как token:

```text
выпускается
→ доставляется
→ хранится
→ передаётся в API
→ проверяется
→ используется
→ повторно применяется для следующего запроса
→ обновляется или ротируется
→ истекает либо отзывается
```

Проект работает с демонстрационной моделью token в памяти.

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
- Постепенно перейти от линейной демонстрации к branching-сценариям, JWT Inspector, Web UI и browser observation
- Построить живую интерактивную шпаргалку по token lifecycle и security concepts
- Научиться отделять UI-логику от domain lifecycle rules
- Показать, что допустимость перехода должна проверяться не только интерфейсом, но и domain layer

## Текущая версия

## `v0.1.1 — Interactive Branching Lifecycle Lab`

`v0.1.1` — интерактивная учебная state machine для token lifecycle.

Пользователь больше не проходит заранее заданный маршрут с помощью `Enter`. Вместо этого программа:

```text
показывает текущую token card
→ показывает только допустимые actions
→ читает выбор пользователя
→ валидирует номер menu
→ валидирует lifecycle transition
→ переводит token в следующий state
   или объясняет, почему действие невозможно
```

Пример интерактивного flow:

```text
=== Token Card ===
Token ID: demo-token-001
Token state: Stored
Location: В хранилище клиента или в server-side session.
Can client use it now: Да. Token сохранён и готов к отправке в API.
What will API check: API проверит подпись, issuer, audience, exp, nbf и scopes, когда получит token.
Next possible event: Отправка token в API, refresh / rotation, expiry или revoke.
Что происходит: Токен сохранён для последующего использования.

Available actions:
1. Отправить token в API
2. Выполнить refresh / rotation
3. Дождаться истечения token
4. Отозвать token
Choose an action:
```

После terminal state пользователь может начать новый demo-run или выйти:

```text
What would you like to do?
1. Start a new demo run
2. Exit
Choose an option:
```

## Lifecycle graph v0.1.1

Текущая модель содержит branching, два допустимых цикла и два terminal states.

```text
Issued
├── Deliver → Delivered
├── Expire  → Expired
└── Revoke  → Revoked

Delivered
├── Store   → Stored
├── Expire  → Expired
└── Revoke  → Revoked

Stored
├── Transmit → Transmitted
├── Refresh  → Refreshed
├── Expire   → Expired
└── Revoke   → Revoked

Transmitted
├── Validate → Validated
├── Expire   → Expired
└── Revoke   → Revoked

Validated
├── Use     → Used
├── Expire  → Expired
└── Revoke  → Revoked

Used
├── Prepare next request → Stored
├── Expire               → Expired
└── Revoke               → Revoked

Refreshed
├── Store refreshed token → Stored
├── Expire                → Expired
└── Revoke                → Revoked

Expired → terminal
Revoked → terminal
```

### Повторное использование token

После успешного API-вызова token может использоваться для следующего request:

```text
Stored
→ Transmitted
→ Validated
→ Used
→ Stored
```

Состояние `Used` не является terminal state. Оно означает, что текущий API request успешно завершился. Пользователь может подготовить следующий request с тем же token.

### Refresh / rotation flow

Из состояния `Stored` пользователь может выбрать refresh / rotation:

```text
Stored
→ Refreshed
→ Store refreshed token
→ Stored
```

`Refreshed` означает, что authorization server выдал обновлённый token, но его ещё нужно сохранить перед следующим API-вызовом.

Это учебная абстракция:

```text
Программа не создаёт второй объект Token.
Программа не меняет Token.ID.
Программа не хранит access-token / refresh-token pair.
Программа не моделирует token family.
Программа не хранит parent/child relation старого и нового token.
```

Модель показывает смысл flow, а не реальное хранение credentials.

### Terminal states

В текущей версии terminal states:

```text
Expired
Revoked
```

Для terminal state lifecycle не возвращает доступных actions:

```go
AvailableActions(state) == nil
```

После этого CLI выводит final token card и предлагает:

```text
1. Start a new demo run
2. Exit
```

## Допустимые переходы

| Текущее состояние | Доступные действия                    | Результат                                                       |
| ----------------- | ------------------------------------- | --------------------------------------------------------------- |
| `Issued`          | Deliver, Expire, Revoke               | Token доставляется клиенту, истекает или отзывается             |
| `Delivered`       | Store, Expire, Revoke                 | Клиент сохраняет token, либо token истекает/отзывается          |
| `Stored`          | Transmit, Refresh, Expire, Revoke     | Token отправляется в API, обновляется, истекает или отзывается  |
| `Transmitted`     | Validate, Expire, Revoke              | API начинает проверку token, либо token истекает/отзывается     |
| `Validated`       | Use, Expire, Revoke                   | API request успешно завершается, либо token истекает/отзывается |
| `Used`            | Prepare next request, Expire, Revoke  | Token готовится к следующему request, либо завершается          |
| `Refreshed`       | Store refreshed token, Expire, Revoke | Обновлённый token сохраняется, либо завершается                 |
| `Expired`         | Нет                                   | TTL закончился, token недействителен                            |
| `Revoked`         | Нет                                   | Token досрочно запрещён                                         |

## Восемь учебных тем

|   № | Тема                 | Краткое описание                                                                           |
| --: | -------------------- | ------------------------------------------------------------------------------------------ |
|   1 | Выпуск               | Authorization server создаёт token, назначает ID, TTL, claims и scopes                     |
|   2 | Доставка             | Token передаётся клиенту через response body, cookie или другой разрешённый механизм       |
|   3 | Хранение             | Клиент или сервер хранит token в памяти, session, cookie или другом хранилище              |
|   4 | Передача             | Клиент отправляет token в API, например через `Authorization: Bearer ...`                  |
|   5 | Проверка             | API получает token и начинает validation: signature, issuer, audience, expiration и scopes |
|   6 | Использование        | При успешной validation API выполняет request                                              |
|   7 | Обновление и ротация | Клиент получает обновлённый token и сохраняет его перед новым request                      |
|   8 | Завершение lifecycle | Token истекает естественно или досрочно отзывается                                         |

## Что реализовано

### Доменная модель

- `Token` — Entity с идентификатором `ID` и текущим `State`
- `TokenState` — enum-like тип на основе `string`
- `LifecycleAction` — enum-like тип на основе `string`
- `AvailableActions()` — возвращает actions, допустимые в текущем state
- `ApplyAction()` — валидирует lifecycle transition и меняет token state
- `nextState()` — определяет целевое состояние для допустимого action

Текущие states:

```go
Issued
Delivered
Stored
Transmitted
Validated
Used
Refreshed
Expired
Revoked
```

Текущие actions:

```go
DeliverAction
StoreAction
TransmitAction
ValidateAction
UseAction
PrepareNextRequestAction
RefreshAction
StoreRefreshedTokenAction
ExpireAction
RevokeAction
```

### Lifecycle validation

Lifecycle layer самостоятельно проверяет переход, даже если UI показывает только разрешённые варианты.

```go
func ApplyAction(token *Token, action LifecycleAction) error
```

Логика validation:

```text
Допустимый action
→ token.State меняется на next state.

Недопустимый action
→ возвращается error.
→ token.State остаётся прежним.

nil token
→ возвращается error.
```

Например:

```text
Stored + TransmitAction
→ Transmitted
```

```text
Issued + RefreshAction
→ error
→ Issued остаётся Issued
```

```text
Expired + TransmitAction
→ error
→ Expired остаётся Expired
```

### Interactive CLI

CLI выполняет один demo-run следующим образом:

```text
Создать token в состоянии Issued
→ показать token card
→ получить AvailableActions(currentState)
→ показать actions в menu
→ прочитать номер действия
→ применить ApplyAction()
→ повторять до terminal state
```

После `Expired` или `Revoked` текущий demo-run заканчивается.

Затем пользователь выбирает:

```text
1. Start a new demo run
2. Exit
```

При `Start a new demo run` приложение создаёт новый демонстрационный token:

```go
Token{
    ID:    "demo-token-001",
    State: Issued,
}
```

При `Exit` программа завершает работу.

### Карточка token

Для каждого состояния CLI выводит:

```text
Token ID
Token state
Location
Can client use it now
What will API check
Next possible event
Что происходит
```

### Валидация ввода

CLI обрабатывает пользовательские ошибки отдельно от lifecycle errors.

| Ситуация                 | Поведение                                                 |
| ------------------------ | --------------------------------------------------------- |
| Пустой ввод              | Выводится сообщение о необходимости ввести номер          |
| Введён нечисловой текст  | Выводится сообщение о необходимости ввести номер          |
| Номер вне диапазона menu | Выводится сообщение с допустимым диапазоном               |
| Допустимый номер         | Возвращается выбранный `LifecycleAction`                  |
| Invalid lifecycle action | `ApplyAction()` возвращает error, token state не меняется |
| Terminal token           | Lifecycle menu не показывается; доступны restart или exit |

### Цвета терминала

| Цвет                   | Значение                                                                       |
| ---------------------- | ------------------------------------------------------------------------------ |
| Cyan                   | Заголовок приложения, заголовок token card, terminal result и финальная строка |
| Yellow                 | Prompts для выбора action и restart/exit option                                |
| Orange                 | Terminal state `Expired`                                                       |
| Red                    | Ошибки ввода, ошибки lifecycle validation и terminal state `Revoked`           |
| Default terminal color | Обычные состояния, menu и учебные пояснения                                    |

### Unit tests

В `lifecycle_test.go` проверяются domain rules, а не ANSI-цвета или точное форматирование консольного текста.

Текущие tests покрывают:

```text
Stored + TransmitAction → Transmitted
```

```text
Issued + RefreshAction → error
```

```text
Expired + TransmitAction → error
```

```text
Invalid action → token state не меняется
```

Также проверяется корректный результат `AvailableActions()` для:

```text
Issued
Stored
Refreshed
Expired
Revoked
```

## Структура проекта

```text
token-lifecycle-lab/
├── go.mod
├── main.go
├── token.go
├── token_state.go
├── lifecycle.go
├── ui.go
├── lifecycle_test.go
└── README.md
```

### `main.go`

Точка входа приложения:

- создаёт `bufio.Reader` для stdin;
- запускает interactive demo-run;
- создаёт демонстрационный token в состоянии `Issued`;
- показывает token card;
- получает actions через `AvailableActions()`;
- читает действие пользователя;
- передаёт действие в `ApplyAction()`;
- завершает текущий run в terminal state;
- предлагает `Start a new demo run` или `Exit`.

### `token.go`

Доменная Entity:

```go
type Token struct {
    ID    string
    State TokenState
}
```

### `token_state.go`

Доменный тип состояния, constants и учебные пояснения:

```go
type TokenState string
```

Файл содержит методы и данные для token card:

```text
Description()
Location()
ClientUsage()
APIChecks()
NextEvents()
```

### `lifecycle.go`

Lifecycle rules:

```text
LifecycleAction
Label()
AvailableActions()
ApplyAction()
nextState()
```

Этот файл является источником истины для допустимых transitions.

### `ui.go`

Консольное представление и обработка пользовательского ввода:

```text
ANSI-цвета
colorize()
printTokenCard()
printActionsMenu()
readActionChoice()
readRestartChoice()
```

### `lifecycle_test.go`

Unit tests lifecycle rules:

```text
valid transition
invalid transition
terminal-state behavior
available actions for selected states
```

Все Go-файлы находятся в одной папке и относятся к одному package:

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

Перед фиксацией версии выполните:

```bash
go fmt ./...
go vet ./...
go test ./...
```

- `go fmt` форматирует Go-код
- `go vet` ищет подозрительные конструкции
- `go test` собирает package и запускает unit tests

Для подробного вывода tests:

```bash
go test -v ./...
```

## Границы v0.1.1

`v0.1.1` — учебная lifecycle model, а не production authentication system.

Утилита:

- не принимает и не сохраняет реальные token values;
- не расшифровывает и не проверяет реальные JWT;
- не проверяет cryptographic signature;
- не подключается к issuer, OAuth server или API;
- не выполняет introspection;
- не читает browser cookies, localStorage или sessionStorage;
- не является Identity Provider, API Gateway, JWT validator или production auth service;
- не выполняет настоящие API requests;
- не хранит access-token / refresh-token pair;
- не моделирует несколько одновременно существующих token objects;
- не создаёт token family;
- не реализует реальные refresh-token rotation policies;
- не использует persistent storage или database.

Важно:

```text
JWT может быть ещё не истёкшим по exp,
но уже быть отозванным server-side.
```

Реальный revoke status нельзя достоверно определить только по строке JWT. Для этого нужен issuer, introspection endpoint, deny-list или другое server-side state.

## История: v0.1

## `v0.1 — Linear CLI Lifecycle Demo`

`v0.1` была первой версией проекта. Пользователь нажимал `Enter`, а утилита последовательно показывала учебный lifecycle token.

На каждой стадии CLI отображала карточку token и отвечала на вопросы:

1. Что сейчас происходит с token?
2. Где token находится?
3. Может ли клиент использовать token сейчас?
4. Что проверит API?
5. Какое следующее событие возможно?

Пример начала исторического walkthrough:

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

Исторический linear scenario `v0.1`:

```text
Issued
→ Delivered
→ Stored
→ Transmitted
→ Used
→ Rotated
→ Expired
```

`Revoked` уже присутствовал в доменной модели как terminal state, но не входил в linear scenario `v0.1`.

Причина:

```text
Expired — token стал недействительным естественно:
          наступил exp или закончился TTL.

Revoked — token был досрочно запрещён:
          logout, security incident, смена пароля,
          ручной revoke, отключение account и так далее.
```

В `v0.1.1` линейный walkthrough заменён на interactive branching lifecycle.

## План развития

| Версия   | Результат                                                                                                                   |
| -------- | --------------------------------------------------------------------------------------------------------------------------- |
| `v0.1`   | Линейное CLI-демо lifecycle walkthrough с переходом по `Enter`                                                              |
| `v0.1.1` | Interactive Branching Lifecycle Lab: actions, validation, terminal states, restart и exit                                   |
| `v0.2`   | Scenario Simulator: явное моделирование expiry, storage, `401`, `403`, security scenarios и более реалистичных API outcomes |
| `v0.3`   | JWT Inspector: parsing JWT, claims, `exp`, `nbf`, `iat`, `jti`, TTL и объяснение decode versus verify                       |
| `v0.4`   | Локальный Web UI на Go и чистом HTML/CSS/JavaScript                                                                         |
| `v0.5`   | Browser Token Discovery для выбранного browser и домена                                                                     |
| `v1.0`   | Local real-time observer: CLI и Web UI показывают наблюдаемые изменения token                                               |

## Статус

**`v0.1` завершена.**

**`v0.1.1` реализует интерактивный branching lifecycle: menu действий, domain validation, terminal states, restart/exit и unit tests.**

Текущая версия:

```text
v0.1.1 — Interactive Branching Lifecycle Lab
```

---
