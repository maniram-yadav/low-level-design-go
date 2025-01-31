package visitor

type Visitor interface {
	VisitDepositTransaction(depositTransaction *DepositTransaction)
	VisitWithdrawTransaction(withdrawTransaction *WithdrawTransaction)
	VisitEarnedInterestTransaction(earnedInterestTransaction *EarnedInterestTransaction)
	BuildTransactionReport() string
}
