package services

import (
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func InitStripe() {
    stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func CreatePaymentIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
    params := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(amount),
        Currency: stripe.String(currency),
    }
    return paymentintent.New(params)
}
