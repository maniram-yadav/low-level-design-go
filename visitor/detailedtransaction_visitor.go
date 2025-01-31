package visitor

import (
	"strconv"
	"strings"
)

type DetailedTransactionVisitor struct {
	report   strings.Builder
	totalSum float64
}

func (dv *DetailedTransactionVisitor) VisitDepositTransaction(depositTransaction *DepositTransaction) {
	dv.report.WriteString("Deposit : +" + strconv.FormatFloat(depositTransaction.Amount, 'f', 3, 64) + "\n")
	dv.totalSum += depositTransaction.Amount

}
func (dv *DetailedTransactionVisitor) VisitWithdrawTransaction(withdrawTransaction *WithdrawTransaction) {
	dv.report.WriteString("Withdraw : -" + strconv.FormatFloat(withdrawTransaction.Amount, 'f', 3, 64) + "\n")
	dv.totalSum -= withdrawTransaction.Amount
}
func (dv *DetailedTransactionVisitor) VisitEarnedInterestTransaction(earnedInterestTransaction *EarnedInterestTransaction) {
	earnedInterestTransaction.calculateInterest(dv.totalSum)
	dv.report.WriteString("Interest paid : +" + strconv.FormatFloat(earnedInterestTransaction.CalculatedInterest, 'f', 3, 64) + "\n")
	dv.totalSum += earnedInterestTransaction.CalculatedInterest
}

func (dv *DetailedTransactionVisitor) BuildTransactionReport() string {

	detailedReport := "\n\tDetailed Transaction Summary Report \n"
	dv.report.WriteString("Total Sum : " + strconv.FormatFloat(dv.totalSum, 'f', 3, 64) + "\n")

	return detailedReport + dv.report.String()
}
