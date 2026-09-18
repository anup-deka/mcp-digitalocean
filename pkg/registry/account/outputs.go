package account

import (
	"github.com/digitalocean/godo"

	"mcp-digitalocean/pkg/registry/common"
)

// Structured output contracts for this package's tools. Each var is declared
// once and used twice — Schema() at registration and Result() in the handler —
// so a tool's declared outputSchema and the structuredContent it emits cannot
// disagree. Collections are published under the plural resource name, single
// resources under the singular one.
//
// Tools whose only result is a confirmation message (key-delete) have no
// payload to describe and stay text-only.
var (
	accountOut        = common.NewOutput[*godo.Account]("account")
	actionOut         = common.NewOutput[*godo.Action]("action")
	actionsOut        = common.NewOutput[[]godo.Action]("actions")
	balanceOut        = common.NewOutput[*godo.Balance]("balance")
	billingHistoryOut = common.NewOutput[*godo.BillingHistory]("billing_history")
	invoiceOut        = common.NewOutput[*godo.Invoice]("invoice")
	invoiceListOut    = common.NewOutput[*godo.InvoiceList]("invoices")
	keyOut            = common.NewOutput[*godo.Key]("key")
	keysOut           = common.NewOutput[[]godo.Key]("keys")
)
