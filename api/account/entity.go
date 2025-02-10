// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

// Package account implements account entities and services
package account // import "github.com/brunomvsouza/ynab.go/api/account"
import (
	"time"
)

// Account represents an account for a budget
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Type The type of account
	Type     Type `json:"type"`
	OnBudget bool `json:"on_budget"`
	// Balance The current balance of the account in milliunits format
	Balance int64 `json:"balance"`
	// ClearedBalance The current cleared balance of the account in milliunits format
	ClearedBalance int64 `json:"cleared_balance"`
	// ClearedBalance The current uncleared balance of the account in milliunits format
	UnclearedBalance int64 `json:"uncleared_balance"`
	// Closed Whether this account is closed or not
	Closed bool `json:"closed"`
	// Deleted Whether or not the account has been deleted. Deleted accounts will only be included in delta requests.
	Deleted bool `json:"deleted"`

	Note *string `json:"note"`

	// TransferPayeeID The payee id which should be used when transferring to this account
	TransferPayeeID     string `json:"transfer_payee_id"`
	DirectImportLinked  bool   `json:"direct_import_linked"`
	DirectImportInError bool   `json:"direct_import_in_error"`
	// LastReconciledAt A date/time specifying when the account was last reconciled.
	LastReconciledAt    *time.Time `json:"last_reconciled_at"`
	DebtOriginalBalance *int64     `json:"debt_original_balance"`
	DebtInterestRates   struct {
	} `json:"debt_interest_rates"`
	DebtMinimumPayments struct {
	} `json:"debt_minimum_payments"`
	DebtEscrowAmounts struct {
	} `json:"debt_escrow_amounts"`
}

// SearchResultSnapshot represents a versioned snapshot for an account search
type SearchResultSnapshot struct {
	Accounts        []*Account
	ServerKnowledge uint64
}
