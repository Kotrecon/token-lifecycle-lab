package main

import (
	"slices"
	"testing"
)

func TestApplyActionChangesTokenState(t *testing.T) {
	token := Token{
		ID:    "test-token-001",
		State: Stored,
	}

	action := TransmitAction
	expectedState := Transmitted

	t.Logf("Token ID: %s", token.ID)
	t.Logf("Initial state: %s", token.State)
	t.Logf("Action: %s", action)
	t.Logf("Expected state: %s", expectedState)

	err := ApplyAction(&token, action)

	if err != nil {
		t.Fatalf("ApplyAction returned an unexpected error: %v", err)
	}

	if token.State != expectedState {
		t.Fatalf(
			"actual state = %q, want %q",
			token.State,
			expectedState,
		)
	}

	t.Logf("Result: action accepted, token moved to %s", token.State)
}

func TestApplyActionRejectsInvalidTransition(t *testing.T) {
	token := Token{
		ID:    "test-token-002",
		State: Issued,
	}

	action := RefreshAction
	expectedState := Issued

	t.Logf("Token ID: %s", token.ID)
	t.Logf("Initial state: %s", token.State)
	t.Logf("Action: %s", action)
	t.Logf("Expected result: action rejected, state remains %s", expectedState)

	err := ApplyAction(&token, action)

	if err == nil {
		t.Fatal("ApplyAction returned nil, want an error")
	}

	if token.State != expectedState {
		t.Fatalf(
			"actual state = %q, want %q",
			token.State,
			expectedState,
		)
	}

	t.Logf("Result: action rejected: %v", err)
	t.Logf("Actual state: %s", token.State)
}

func TestApplyActionRejectsActionForExpiredToken(t *testing.T) {
	token := Token{
		ID:    "test-token-003",
		State: Expired,
	}

	action := TransmitAction
	expectedState := Expired

	t.Logf("Token ID: %s", token.ID)
	t.Logf("Initial state: %s", token.State)
	t.Logf("Action: %s", action)
	t.Logf("Expected result: terminal token rejects the action")

	err := ApplyAction(&token, action)

	if err == nil {
		t.Fatal("ApplyAction returned nil, want an error")
	}

	if token.State != expectedState {
		t.Fatalf(
			"actual state = %q, want %q",
			token.State,
			expectedState,
		)
	}

	t.Logf("Result: action rejected: %v", err)
	t.Logf("Actual state: %s", token.State)
}

func TestAvailableActions(t *testing.T) {
	tests := []struct {
		name     string
		state    TokenState
		expected []LifecycleAction
	}{
		{
			name:  "issued token can be delivered expired or revoked",
			state: Issued,
			expected: []LifecycleAction{
				DeliverAction,
				ExpireAction,
				RevokeAction,
			},
		},
		{
			name:  "stored token can be transmitted refreshed expired or revoked",
			state: Stored,
			expected: []LifecycleAction{
				TransmitAction,
				RefreshAction,
				ExpireAction,
				RevokeAction,
			},
		},
		{
			name:  "refreshed token can be stored expired or revoked",
			state: Refreshed,
			expected: []LifecycleAction{
				StoreRefreshedTokenAction,
				ExpireAction,
				RevokeAction,
			},
		},
		{
			name:     "expired token has no available actions",
			state:    Expired,
			expected: nil,
		},
		{
			name:     "revoked token has no available actions",
			state:    Revoked,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := AvailableActions(test.state)

			if !slices.Equal(actual, test.expected) {
				t.Fatalf(
					"AvailableActions(%q) = %v, want %v",
					test.state,
					actual,
					test.expected,
				)
			}
		})
	}
}
