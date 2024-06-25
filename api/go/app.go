package main

// API usage Dependencies
import (
	"fmt"
	"net/http"
	"os"
	"context"
	"github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"strconv"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
)

// This example utilizes a range of configuration values. These values are conveniently fetched from the ENV for simplicity, but you also have the option to define them directly or store them in an alternative location.
var PUB_KEY = os.Getenv("MT_PUB_KEY")
var API_KEY = os.Getenv("MT_API_KEY")
var ORG_ID = os.Getenv("MT_ORG_ID")
var FS_KEY = os.Getenv("FS_KEY")

// Instantiate a configured Modern Treasury client
var client = moderntreasury.NewClient(option.WithOrganizationID(ORG_ID),option.WithAPIKey(API_KEY))

// Session storage
var store = sessions.NewFilesystemStore("", []byte(FS_KEY))

//POST route to handle a new account collection form
func createCpAcf(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "modern-treasury")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", 405)
	}

	r.ParseForm()

	counterParty, err := client.Counterparties.New(context.TODO(), moderntreasury.CounterpartyNewParams{
		Name: moderntreasury.F(r.PostFormValue("name")),
	})
	if err != nil {
		panic(err.Error())
	}

	accountCollectionFlow, err := client.AccountCollectionFlows.New(context.TODO(), moderntreasury.AccountCollectionFlowNewParams{CounterpartyID: moderntreasury.F(counterParty.ID), PaymentTypes: moderntreasury.F(r.Form["rails[]"])})
	if err != nil {
		panic(err.Error())
	}
	// Set some session values.
	session.Values["clientToken"] = accountCollectionFlow.ClientToken
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/embed.html", 303)

}

//POST route to handle a new payment form
func createCpPf(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "modern-treasury")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", 405)
	}

	r.ParseForm()

	counterParty, err := client.Counterparties.New(context.TODO(), moderntreasury.CounterpartyNewParams{
		Name: moderntreasury.F(r.PostFormValue("name")),
	})
	if err != nil {
		panic(err.Error())
	}

	float_amount, err := strconv.ParseFloat(r.PostFormValue("amount"), 64)
	if err != nil {
        // ... handle error
        panic(err.Error())
    }

    int_amount := int64(float_amount * 100)

	paymentFlow, err := client.PaymentFlows.New(context.TODO(), moderntreasury.PaymentFlowNewParams{
		Amount: moderntreasury.F(int_amount),
		Currency: moderntreasury.F(r.PostFormValue("currency")),
		Direction: moderntreasury.F(moderntreasury.PaymentFlowNewParamsDirectionDebit),
		OriginatingAccountID: moderntreasury.F(r.PostFormValue("originating_account_id")),
		CounterpartyID: moderntreasury.F(counterParty.ID),
	})
	if err != nil {
		panic(err.Error())
	}
	// Set some session values.
	session.Values["clientToken"] = paymentFlow.ClientToken
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/embed.html", 303)

}

func config(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/javascript")

	session, err := store.Get(r, "modern-treasury")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("window.mtConfig = { publishableKey: '%s' , clientToken: '%s'}", PUB_KEY, session.Values["clientToken"])

	fmt.Fprintf(w, response)
}

func main() {
	publicDirPath := os.Getenv("PUBLIC_DIR_PATH")

	if publicDirPath == "" {
		publicDirPath = "../../public"
	}

	http.Handle("/", http.FileServer(http.Dir(publicDirPath)))

	http.HandleFunc("/config", config)

	http.HandleFunc("/api/create-cp-acf", createCpAcf)
	http.HandleFunc("/api/create-cp-pf", createCpPf)

	http.ListenAndServe(":9001", nil)
}