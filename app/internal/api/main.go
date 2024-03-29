package api

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/internal/data"
	"example.com/internal/models"
	"example.com/internal/repository"
	"example.com/internal/services"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SwmsSchemaHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.SwmsSchema)
}

type CreateSwmsRequest struct {
	ProjectAddress   string          `json:"projectAddress"`
	ScopeOfWork      string          `json:"scopeOfWork"`
	DateDeveloped    string          `json:"dateDeveloped"`
	ApprovalDate     string          `json:"approvalDate"`
	DateLastReviewed string          `json:"dateLastReviewed"`
	NextReviewDate   string          `json:"nextReviewDate"`
	TableData        json.RawMessage `json:"tableData"`
}

func CreateSwms(w http.ResponseWriter, r *http.Request) {
	user, err := services.Authenticate(r)

	if err != nil {
		log.Printf("Error authenticating user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createSwmsRequest := CreateSwmsRequest{}
	err = json.NewDecoder(r.Body).Decode(&createSwmsRequest)

	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := repository.ConnectDB()
	defer db.Close()
	var swmsID int
	swmsID, err = db.CreateSwms(user, createSwmsRequest.ProjectAddress, createSwmsRequest.TableData, "construction")

	if err != nil {
		log.Printf("Error creating swms: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	conn, err := amqp.Dial("amqp://admin:adminpassword@swms_rabbitmq_container")
	if err != nil {
		log.Printf("Error connecting to RabbitMQ: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error opening channel: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"pdf", // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Printf("Error declaring queue: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type CreatePdfData struct {
		ProjectAddress   string          `json:"projectAddress"`
		ApprovalDate     string          `json:"approvalDate"`
		DateDeveloped    string          `json:"dateDeveloped"`
		DateLastReviewed string          `json:"dateLastReviewed"`
		NextReviewDate   string          `json:"nextReviewDate"`
		ScopeOfWork      string          `json:"scopeOfWork"`
		TableData        json.RawMessage `json:"tableData"`
		ABN              string          `json:"abn"`
		BusinessName     string          `json:"businessName"`
		BusinessPhone    string          `json:"businessPhone"`
		BusinessEmail    string          `json:"businessEmail"`
	}

	org := models.Organisations{}

	org, err = db.GetOrganisation(user.OrganisationID)

	type SwmsRequest struct {
		Id   int           `json:"id"`
		Data CreatePdfData `json:"data"`
		Pdf  string        `json:"pdf"`
	}

	PdfData := CreatePdfData{
		ProjectAddress:   createSwmsRequest.ProjectAddress,
		ScopeOfWork:      createSwmsRequest.ScopeOfWork,
		ApprovalDate:     createSwmsRequest.ApprovalDate,
		DateDeveloped:    createSwmsRequest.DateDeveloped,
		DateLastReviewed: createSwmsRequest.DateLastReviewed,
		NextReviewDate:   createSwmsRequest.NextReviewDate,
		TableData:        createSwmsRequest.TableData,
		ABN:              org.ABN,
		BusinessName:     org.Name,
		BusinessPhone:    org.BusinessPhone,
		BusinessEmail:    org.BusinessEmail,
	}

	body := SwmsRequest{
		Id:   swmsID,
		Data: PdfData,
		Pdf:  "swms",
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error marshaling body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ch.PublishWithContext(
		r.Context(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.SwmsSchema)
}

type UpdateSwmsRequest struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	ID       int    `json:"id"`
}

func UpdateFileHandler(w http.ResponseWriter, r *http.Request) {

	updateSwmsRequest := UpdateSwmsRequest{}
	err := json.NewDecoder(r.Body).Decode(&updateSwmsRequest)

	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := repository.ConnectDB()
	defer db.Close()
	err = db.UpdateFile(updateSwmsRequest.ID, updateSwmsRequest.FileName, updateSwmsRequest.FilePath)

	if err != nil {
		log.Printf("Error updating file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.SwmsSchema)

}
