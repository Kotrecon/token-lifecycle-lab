package main

type TokenState string

const (
	Issued      TokenState = "Issued"
	Delivered   TokenState = "Delivered"
	Stored      TokenState = "Stored"
	Transmitted TokenState = "Transmitted"
	Validated   TokenState = "Validated"
	Used        TokenState = "Used"
	Rotated     TokenState = "Rotated"
	Refreshed   TokenState = "Refreshed"
	Expired     TokenState = "Expired"
	Revoked     TokenState = "Revoked"
)

func (state TokenState) Description() string {
	switch state {
	case Issued:
		return "Токен создан и выпущен системой."
	case Delivered:
		return "Токен доставлен получателю."
	case Stored:
		return "Токен сохранён для последующего использования."
	case Transmitted:
		return "Токен передаётся в целевую систему."
	case Validated:
		return "API завершил проверку token и определил результат запроса."
	case Used:
		return "Токен был успешно использован."
	case Rotated:
		return "Токен заменён новым токеном."
	case Refreshed:
		return "Authorization server выдал новый token через refresh или rotation flow."
	case Expired:
		return "Срок действия токена истёк."
	case Revoked:
		return "Токен был принудительно отозван."
	default:
		return "Неизвестное состояние токена."
	}
}

func (state TokenState) Location() string {
	switch state {
	case Issued:
		return "На стороне authorization server после выпуска."
	case Delivered:
		return "Передаётся от authorization server к клиенту."
	case Stored:
		return "В хранилище клиента или в server-side session."
	case Transmitted:
		return "В запросе, который клиент отправляет в API."
	case Validated:
		return "На стороне API после проверки token."
	case Used:
		return "API уже обработал запрос с этим token."
	case Rotated:
		return "Старый token заменяется новым."
	case Refreshed:
		return "Новый token получен клиентом и ожидает сохранения."
	case Expired:
		return "Token завершён: срок действия закончился."
	case Revoked:
		return "Token завершён: issuer или сервер его отозвал."
	default:
		return "Местоположение token неизвестно."
	}
}

func (state TokenState) ClientUsage() string {
	switch state {
	case Issued:
		return "Нет. Token ещё не передан клиенту."
	case Delivered:
		return "Пока нет. Token находится в процессе доставки."
	case Stored:
		return "Да. Token сохранён и готов к отправке в API."
	case Transmitted:
		return "Token уже используется в текущем API-запросе."
	case Validated:
		return "Запрос находится на стороне API и ожидает результата проверки."
	case Used:
		return "Да. Token был успешно использован и может применяться повторно."
	case Rotated:
		return "Нет. Старый token заменён новым."
	case Refreshed:
		return "Пока нет. Новый token нужно сохранить перед следующим API-вызовом."
	case Expired:
		return "Нет. Срок действия token истёк."
	case Revoked:
		return "Нет. Token был досрочно отозван."
	default:
		return "Неизвестно."
	}
}

func (state TokenState) APIChecks() string {
	switch state {
	case Issued:
		return "API ещё не получил token."
	case Delivered:
		return "API ещё не получил token."
	case Stored:
		return "API проверит подпись, issuer, audience, exp, nbf и scopes, когда получит token."
	case Transmitted:
		return "API проверит подпись, issuer, audience, exp, nbf и scopes."
	case Validated:
		return "API проверил подпись, issuer, audience, exp, nbf и scopes."
	case Used:
		return "API уже выполнил проверки и принял token."
	case Rotated:
		return "API будет принимать новый token, а старый отклонит."
	case Refreshed:
		return "Новый token будет проверен API при первом использовании."
	case Expired:
		return "API отклонит token: срок действия token истёк."
	case Revoked:
		return "API сможет отклонить token, если проверяет revoke status у issuer или в deny-list."
	default:
		return "Неизвестно."
	}
}

func (state TokenState) NextEvents() string {
	switch state {
	case Issued:
		return "Доставка token клиенту, expiry или revoke."
	case Delivered:
		return "Сохранение token, expiry или revoke."
	case Stored:
		return "Отправка token в API, refresh / rotation, expiry или revoke."
	case Transmitted:
		return "Проверка token API."
	case Validated:
		return "Успешное выполнение API-запроса или отклонение token."
	case Used:
		return "Повторное использование, refresh / rotation, expiry или revoke."
	case Rotated:
		return "Сохранение нового token, затем его использование."
	case Refreshed:
		return "Сохранение нового token для следующего API-вызова."
	case Expired:
		return "Никаких. Token завершён."
	case Revoked:
		return "Никаких. Token завершён."
	default:
		return "Неизвестно."
	}
}
