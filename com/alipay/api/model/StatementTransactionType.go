package model

type StatementTransactionType string

const (
	StatementTransactionType_PAYMENT            StatementTransactionType = "PAYMENT"
	StatementTransactionType_PAYMENT_REFUND     StatementTransactionType = "PAYMENT_REFUND"
	StatementTransactionType_CHARGE             StatementTransactionType = "CHARGE"
	StatementTransactionType_CHARGE_REFUND      StatementTransactionType = "CHARGE_REFUND"
	StatementTransactionType_TOPUP              StatementTransactionType = "TOPUP"
	StatementTransactionType_SETTLEMENT         StatementTransactionType = "SETTLEMENT"
	StatementTransactionType_WITHDRAW           StatementTransactionType = "WITHDRAW"
	StatementTransactionType_WITHDRAW_RETURN    StatementTransactionType = "WITHDRAW_RETURN"
	StatementTransactionType_TRANSFER           StatementTransactionType = "TRANSFER"
	StatementTransactionType_TRANSFER_RETURN    StatementTransactionType = "TRANSFER_RETURN"
	StatementTransactionType_TRANSFER_TO_CHINA  StatementTransactionType = "TRANSFER_TO_CHINA"
	StatementTransactionType_TRANSFER_RECIPIENT StatementTransactionType = "TRANSFER_RECIPIENT"
	StatementTransactionType_EXCHANGE           StatementTransactionType = "EXCHANGE"
	StatementTransactionType_CREDIT_LOAN        StatementTransactionType = "CREDIT_LOAN"
	StatementTransactionType_CREDIT_REPAY       StatementTransactionType = "CREDIT_REPAY"
	StatementTransactionType_CREDIT_REPAYMENT   StatementTransactionType = "CREDIT_REPAYMENT"
	StatementTransactionType_DIRECT_PAYMENT     StatementTransactionType = "DIRECT_PAYMENT"
	StatementTransactionType_DIRECT_REFUND      StatementTransactionType = "DIRECT_REFUND"
	StatementTransactionType_CARD_PAYMENT       StatementTransactionType = "CARD_PAYMENT"
	StatementTransactionType_CARD_REFUND        StatementTransactionType = "CARD_REFUND"
	StatementTransactionType_OVERFLOW_DEBIT     StatementTransactionType = "OVERFLOW_DEBIT"
	StatementTransactionType_OVERFLOW_CREDIT    StatementTransactionType = "OVERFLOW_CREDIT"
	StatementTransactionType_CASH_BACK          StatementTransactionType = "CASH_BACK"
)
