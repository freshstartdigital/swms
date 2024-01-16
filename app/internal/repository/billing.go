package repository

import (
	"log"

	"example.com/internal/models"
)

func (db *DB) CreateSubscription(CustomerID string, Status string, SubscriptionID string, ProductID string, CurrentPeriodStart int64, CurrentPeriodEnd int64) error {

	organisationID := 0
	subscriptionProductID := 0

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
		WHERE stripe_product_id = $1`, ProductID).
		Scan(&subscriptionProductID)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println("No subscription plan found", subscriptionProductID)
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
		subscriptionProductID,
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
    SELECT id, stripe_product_id, stripe_payment_link, name, description, price, created_at, updated_at
    FROM subscription_plans
    WHERE id = $1`, ID).
		Scan(&subscriptionPlan.ID, &subscriptionPlan.StripeProductId, &subscriptionPlan.StripePaymentLink, &subscriptionPlan.Name, &subscriptionPlan.Description, &subscriptionPlan.Price, &subscriptionPlan.CreatedAt, &subscriptionPlan.UpdatedAt)

	if err != nil {
		log.Println("Error querying database:", err)
		return models.SubscriptionPlans{}, err
	}

	return subscriptionPlan, nil
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

func (db *DB) CreateProduct(ProductID string, name string, description string) error {
	stmt, err := db.Prepare("INSERT INTO subscription_plans (stripe_product_id, name, description) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println("Error preparing statement:", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ProductID,
		name,
		description,
	)

	if err != nil {
		log.Println("Error executing statement:", err)
		return err
	}
	return nil
}

func (db *DB) UpdateSubscriptionPricing(ID string, price int) error {
	stmt, err := db.Prepare(`
		UPDATE subscription_plans
		SET price = $1
		WHERE stripe_product_id = $2
		`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		price,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateStripePaymentLink(ID string, link string) error {
	stmt, err := db.Prepare(`
		UPDATE subscription_plans
		SET stripe_payment_link = $1
		WHERE stripe_product_id = $2
		`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		link,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}
