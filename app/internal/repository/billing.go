package repository

import (
	"log"

	"example.com/internal/models"
)

func (db *DB) CreateSubscription(CustomerID string, Status string, SubscriptionID string, PlanID string, CurrentPeriodStart int64, CurrentPeriodEnd int64) error {

	organisationID := 0
	subscriptionPlanID := 0

	err := db.QueryRow(`
		SELECT id
		FROM organisations
		WHERE stripe_customer_id = $1`, CustomerID).
		Scan(&organisationID)

	if err != nil {
		log.Println("No organisation found", CustomerID)
		return err
	}

	err = db.QueryRow(`
		SELECT id
		FROM subscription_plans
		WHERE stripe_plan_id = $1`, PlanID).
		Scan(&subscriptionPlanID)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println("No subscription plan found", PlanID)
			return err
		}
		return err
	}

	stmt, err := db.Prepare("INSERT INTO subscriptions (organisation_id, stripe_status, stripe_subscription_id, subscription_plan_id, current_period_start, current_period_end) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		organisationID,
		Status,
		SubscriptionID,
		subscriptionPlanID,
		CurrentPeriodStart,
		CurrentPeriodEnd,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateCancelledAtSubscription(SubscriptionID string, CancelledAt int64) error {

	stmt, err := db.Prepare(`
		UPDATE subscriptions
		SET cancelled_at = $1
		WHERE stripe_subscription_id = $2
		`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		CancelledAt,
		SubscriptionID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateInvoice(SubscriptionID string, StripeInvoiceID string, StripeStatus string, StripePDFLink string) error {
	var subscriptionID int

	err := db.QueryRow(`
		SELECT id
		FROM subscriptions
		WHERE stripe_subscription_id = $1`, SubscriptionID).
		Scan(&subscriptionID)

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO stripe_invoices (subscription_id, stripe_invoice_id, stripe_status, stripe_pdf_link) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		subscriptionID,
		StripeInvoiceID,
		StripeStatus,
		StripePDFLink,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateInvoice(ID string, status string) error {

	stmt, err := db.Prepare(`
		UPDATE stripe_invoices
		SET stripe_status = $1
		WHERE stripe_invoice_id = $2
		`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		status,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateSubscription(SubscriptionID string, Status string, CurrentPeriodStart int64, CurrentPeriodEnd int64) error {

	stmt, err := db.Prepare(`
		UPDATE subscriptions
		SET stripe_status = $1, current_period_start = $2, current_period_end = $3
		WHERE stripe_subscription_id = $4
		`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		Status,
		CurrentPeriodStart,
		CurrentPeriodEnd,
		SubscriptionID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetSubscriptionPlanByID(ID int) (models.SubscriptionPlans, error) {
	var subscriptionPlan models.SubscriptionPlans
	err := db.QueryRow(`
    SELECT id, stripe_plan_id, stripe_payment_link, description, price, duration, created_at, updated_at
    FROM subscription_plans
    WHERE id = $1`, ID).
		Scan(&subscriptionPlan.ID, &subscriptionPlan.StripePlanId, &subscriptionPlan.StripePaymentLink, &subscriptionPlan.Description, &subscriptionPlan.Price, &subscriptionPlan.Duration, &subscriptionPlan.CreatedAt, &subscriptionPlan.UpdatedAt)

	if err != nil {
		log.Println("Error querying database:", err)
		return models.SubscriptionPlans{}, err
	}

	return subscriptionPlan, nil
}

func (db *DB) UpdateStripeInvoiceAndSubscriptions(SubscriptionID string, Status string, CurrentPeriodStart int64, CurrentPeriodEnd int64) error {
	stmt, err := db.Prepare(`
	UPDATE stripe_invoices
	SET stripe_status = $1
	WHERE stripe_invoice_id = $2
	`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		Status,
		SubscriptionID,
	)

	if err != nil {
		return err
	}

	stmt, err = db.Prepare(`
	UPDATE subscriptions
	SET stripe_status = $1, current_period_start = $2, current_period_end = $3
	WHERE stripe_subscription_id = $4
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		Status,
		CurrentPeriodStart,
		CurrentPeriodEnd,
		SubscriptionID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetSubscriptionByOrgID(ID int) (models.Subscriptions, error) {
	var subscription models.Subscriptions
	err := db.QueryRow(`
	SELECT id, organisation_id, subscription_plan_id, stripe_subscription_id, stripe_status, current_period_start, current_period_end, cancelled_at
	FROM subscriptions
	WHERE organisation_id = $1`, ID).
		Scan(&subscription.ID, &subscription.OrganisationID, &subscription.SubscriptionPlanID, &subscription.StripeSubscriptionID, &subscription.StripeStatus, &subscription.CurrentPeriodStart, &subscription.CurrentPeriodEnd, &subscription.CancelledAt)

	if err != nil {
		log.Println("Error querying database:", err)
		return models.Subscriptions{}, err
	}

	return subscription, nil
}

func (db *DB) GetAllStripeInvoicesBySubscriptionID(ID int) ([]models.StripeInvoices, error) {
	var stripeInvoices []models.StripeInvoices
	rows, err := db.Query(`
	SELECT id, subscription_id, stripe_invoice_id, stripe_status, stripe_pdf_link
	FROM stripe_invoices
	WHERE subscription_id = $1`, ID)

	if err != nil {
		log.Println("Error querying database:", err)
		return []models.StripeInvoices{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var stripeInvoice models.StripeInvoices
		err := rows.Scan(&stripeInvoice.ID, &stripeInvoice.SubscriptionID, &stripeInvoice.StripeInvoiceID, &stripeInvoice.StripeStatus, &stripeInvoice.StripePDFLink)
		if err != nil {
			log.Println("Error scanning database:", err)
			return []models.StripeInvoices{}, err
		}
		stripeInvoices = append(stripeInvoices, stripeInvoice)
	}

	return stripeInvoices, nil
}
