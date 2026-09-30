package main

import (
	"fmt"
	"slices"
)

type LifecycleAction string

const (
	DeliverAction             LifecycleAction = "Deliver"
	StoreAction               LifecycleAction = "Store"
	TransmitAction            LifecycleAction = "Transmit"
	ValidateAction            LifecycleAction = "Validate"
	UseAction                 LifecycleAction = "Use"
	PrepareNextRequestAction  LifecycleAction = "PrepareNextRequest"
	RefreshAction             LifecycleAction = "Refresh"
	StoreRefreshedTokenAction LifecycleAction = "StoreRefreshedToken"
	ExpireAction              LifecycleAction = "Expire"
	RevokeAction              LifecycleAction = "Revoke"
)

func (action LifecycleAction) Label() string {
	switch action {
	case DeliverAction:
		return "Доставить token клиенту"
	case StoreAction:
		return "Сохранить token"
	case TransmitAction:
		return "Отправить token в API"
	case ValidateAction:
		return "Проверить token в API"
	case UseAction:
		return "Завершить API-вызов успешно"
	case PrepareNextRequestAction:
		return "Подготовить следующий API-вызов с тем же token"
	case RefreshAction:
		return "Выполнить refresh / rotation"
	case StoreRefreshedTokenAction:
		return "Сохранить новый token"
	case ExpireAction:
		return "Дождаться истечения token"
	case RevokeAction:
		return "Отозвать token"
	default:
		return "Неизвестное действие"
	}
}

func AvailableActions(state TokenState) []LifecycleAction {
	switch state {
	case Issued:
		return []LifecycleAction{
			DeliverAction,
			ExpireAction,
			RevokeAction,
		}

	case Delivered:
		return []LifecycleAction{
			StoreAction,
			ExpireAction,
			RevokeAction,
		}

	case Stored:
		return []LifecycleAction{
			TransmitAction,
			RefreshAction,
			ExpireAction,
			RevokeAction,
		}

	case Transmitted:
		return []LifecycleAction{
			ValidateAction,
			ExpireAction,
			RevokeAction,
		}

	case Validated:
		return []LifecycleAction{
			UseAction,
			ExpireAction,
			RevokeAction,
		}

	case Used:
		return []LifecycleAction{
			PrepareNextRequestAction,
			ExpireAction,
			RevokeAction,
		}

	case Refreshed:
		return []LifecycleAction{
			StoreRefreshedTokenAction,
			ExpireAction,
			RevokeAction,
		}

	default:
		return nil
	}
}

func ApplyAction(token *Token, action LifecycleAction) error {
	if token == nil {
		return fmt.Errorf("token is required")
	}

	if slices.Contains(AvailableActions(token.State), action) {
		token.State = nextState(action)
		return nil
	}

	return fmt.Errorf(
		"action %q is not allowed when token state is %q",
		action,
		token.State,
	)
}

func nextState(action LifecycleAction) TokenState {
	switch action {
	case DeliverAction:
		return Delivered
	case StoreAction:
		return Stored
	case TransmitAction:
		return Transmitted
	case ValidateAction:
		return Validated
	case UseAction:
		return Used
	case PrepareNextRequestAction:
		return Stored
	case RefreshAction:
		return Refreshed
	case StoreRefreshedTokenAction:
		return Stored
	case ExpireAction:
		return Expired
	case RevokeAction:
		return Revoked
	default:
		return ""
	}
}
