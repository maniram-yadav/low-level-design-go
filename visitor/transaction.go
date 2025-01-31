package visitor

import "time"

const (
	WITHDRAW = "Withdraw"
	DEPOSIT  = "Deposit"
	INTEREST = "Interest"
)

type CalculateSum interface {
	AddOldValue(oldvalue float64)
}

type Transaction interface {
	Accept(v Visitor)
	getTransactionType() string
}

type BaseTransaction struct {
	Amount          float64
	TransactionDate time.Time
}

type DepositTransaction struct {
	TransactionType string
	BaseTransaction
}

type WithdrawTransaction struct {
	TransactionType string
	BaseTransaction
}

type EarnedInterestTransaction struct {
	TransactionType    string
	CalculatedInterest float64
	BaseTransaction
}

func (et *EarnedInterestTransaction) calculateInterest(value float64) float64 {
	interestEarned := value * et.Amount / 100
	et.CalculatedInterest = interestEarned
	return float64(interestEarned)
}

func (dt *DepositTransaction) Accept(v Visitor) {
	v.VisitDepositTransaction(dt)
}

func (dt *DepositTransaction) getTransactionType() string {
	return DEPOSIT
}

func (wt *WithdrawTransaction) Accept(v Visitor) {
	v.VisitWithdrawTransaction(wt)
}

func (dt *WithdrawTransaction) getTransactionType() string {
	return WITHDRAW
}

func (et *EarnedInterestTransaction) Accept(v Visitor) {
	v.VisitEarnedInterestTransaction(et)
}

func (dt *EarnedInterestTransaction) getTransactionType() string {
	return INTEREST
}
