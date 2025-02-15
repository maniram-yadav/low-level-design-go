package test

import (
	"lld/visitor"
	"testing"
)

func TestVisitor(t *testing.T) {

	transactions := []visitor.Transaction{
		&visitor.DepositTransaction{BaseTransaction: visitor.BaseTransaction{Amount: 1000.20}},
		&visitor.DepositTransaction{BaseTransaction: visitor.BaseTransaction{Amount: 134.897}},
		&visitor.WithdrawTransaction{BaseTransaction: visitor.BaseTransaction{Amount: 200}},
		&visitor.EarnedInterestTransaction{BaseTransaction: visitor.BaseTransaction{Amount: 10}},
		&visitor.WithdrawTransaction{BaseTransaction: visitor.BaseTransaction{Amount: 110}},
	}

	summaryVisitor := &visitor.SumaryTransactionVisitor{}

	for _, transaction := range transactions {
		transaction.Accept(summaryVisitor)
	}
	t.Log(summaryVisitor.BuildTransactionReport())

	detailedSummaryVisitor := &visitor.DetailedTransactionVisitor{}

	for _, transaction := range transactions {
		transaction.Accept(detailedSummaryVisitor)
	}
	t.Log(detailedSummaryVisitor.BuildTransactionReport())
}
