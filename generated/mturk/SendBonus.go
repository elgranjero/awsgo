package mturk

// SendBonus is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The SendBonus operation issues a payment of money from your account to a
//
// Worker. This payment happens separately from the reward you pay to the Worker
// when you approve the Worker's assignment. The SendBonus operation requires the
// Worker's ID and the assignment ID as parameters to initiate payment of the
// bonus. You must include a message that explains the reason for the bonus
// payment, as the Worker may not be expecting the payment. Amazon Mechanical Turk
// collects a fee for bonus payments, similar to the HIT listing fee. This
// operation fails if your account does not have enough funds to pay for both the
// bonus and the fees.
