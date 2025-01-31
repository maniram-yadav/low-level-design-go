package visitor

import (
	"strconv"
	"strings"
)

type SumaryTransactionVisitor struct {
	report         strings.Builder
	totalDeposits  float64
	totalWithdrwas float64
	totalInterests float64
	totalSum       float64
}

func (sv *SumaryTransactionVisitor) VisitDepositTransaction(depositTransaction *DepositTransaction) {
	sv.totalDeposits += depositTransaction.Amount
	sv.totalSum += depositTransaction.Amount
}
func (sv *SumaryTransactionVisitor) VisitWithdrawTransaction(withdrawTransaction *WithdrawTransaction) {
	sv.totalWithdrwas += withdrawTransaction.Amount
	sv.totalSum -= withdrawTransaction.Amount
}
func (sv *SumaryTransactionVisitor) VisitEarnedInterestTransaction(earnedInterestTransaction *EarnedInterestTransaction) {
	earnedInterestTransaction.calculateInterest(sv.totalSum)
	sv.totalInterests += earnedInterestTransaction.CalculatedInterest
	sv.totalSum += earnedInterestTransaction.CalculatedInterest
}

func (sv *SumaryTransactionVisitor) BuildTransactionReport() string {
	sv.report.WriteString("\n\tTransaction Summary Report \n")
	sv.report.WriteString("Total Deposits : " + strconv.FormatFloat(sv.totalDeposits, 'f', 3, 64) + "\n")
	sv.report.WriteString("Total Withdraw : " + strconv.FormatFloat(sv.totalWithdrwas, 'f', 3, 64) + "\n")
	sv.report.WriteString("Total Interest : " + strconv.FormatFloat(sv.totalInterests, 'f', 3, 64) + "\n")
	sv.report.WriteString("Total Sum : " + strconv.FormatFloat(sv.totalSum, 'f', 3, 64) + "\n")
	return sv.report.String()
}
