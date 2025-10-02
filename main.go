package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/k-capehart/go-salesforce/v2"
)

func main() {
	sf, err := salesforce.Init(salesforce.Creds{
		Domain:         domain,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	},
		salesforce.WithAPIVersion("v64.0"),
	)

	if err != nil {
		panic(err)
	}

	testGetJobResults(sf)
	// testWithHeader(sf)
	// testDoRequest(sf)
	// bulkDmlAssignFile(sf)
	// bulkDmlAssign(sf)
	// bulkDmlFile(sf)
	// bulkDml(sf)
	// queryBulk(sf)
	// dmlComposite(sf)
	// dmlCollections(sf)
	// dmlSingle(sf)
	// queryStruct(sf)
	// query(sf)
	// functionalConfigExample()
	// httpConfigExample()
	// getAccessTokenAndInstanceUrl(sf)
}

func testGetJobResults(sf *salesforce.Salesforce) {
	fmt.Println("===== GetJobResults")

	type Contact struct {
		LastName string
	}

	contacts := []Contact{
		{
			LastName: "Grimm",
		},
	}
	jobIds, err := sf.InsertBulk("Contact", contacts, 1000, true)
	if err != nil {
		panic(err)
	}
	for _, id := range jobIds {
		results, err := sf.GetJobResults(id) // returns an instance of BulkJobResults
		if err != nil {
			panic(err)
		}
		fmt.Println(results)
	}
}

func testWithHeader(sf *salesforce.Salesforce) {
	fmt.Println("===== WithHeader")
	// Use If-Modified-Since for efficient caching
	resp, err := sf.DoRequest("GET", "/sobjects/Account/describe", nil,
		salesforce.WithHeader("If-Modified-Since", "Wed, 21 Oct 2015 07:28:00 GMT"))
	if err != nil {
		panic(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))

	if resp.StatusCode == 304 {
		fmt.Println("Data not modified, using cached version")
	}
}

func testDoRequest(sf *salesforce.Salesforce) {
	fmt.Println("===== DoRequest")

	resp, err := sf.DoRequest(http.MethodGet, "/limits", nil)
	if err != nil {
		panic(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))
}

func bulkDmlAssignFile(sf *salesforce.Salesforce) {
	fmt.Println("===== Bulk Insert Assign File")

	jobIds, err := sf.InsertBulkFileAssign("Lead", "data/lead_avengers.csv", 1000, false, "01QDn00000112FHMAY")
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Update Assign File")

	jobIds, err = sf.UpdateBulkFileAssign("Lead", "data/update_lead_avengers.csv", 100, true, "01QDn00000112FHMAY")
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Upsert Assign File")

	jobIds, err = sf.UpsertBulkFileAssign("Lead", "LeadExternalId__c", "data/upsert_lead_avengers.csv", 100, true, "01QDn00000112FHMAY")
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)
}

func bulkDmlAssign(sf *salesforce.Salesforce) {
	fmt.Println("===== Bulk Insert Assign")
	type Lead struct {
		LastName string
		Company  string
	}
	leads := []Lead{
		{
			LastName: "Spector",
			Company:  "The Avengers",
		},
	}
	jobIds, err := sf.InsertBulkAssign("Lead", leads, 100, true, "01QDn00000112FHMAY")
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Update Assign")

	type LeadWithId struct {
		Id       string
		LastName string
		Company  string
	}

	var queriedLeads []LeadWithId
	err = sf.Query("SELECT Id FROM Lead LIMIT 1", &queriedLeads)
	if err != nil {
		panic(err)
	}

	leadsWithId := []LeadWithId{
		{
			Id:       queriedLeads[0].Id,
			LastName: "Grant",
			Company:  "The Avengers",
		},
	}
	jobIds, err = sf.UpdateBulkAssign("Lead", leadsWithId, 100, true, "01QDn00000112FHMAY")
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Upsert Assign")

	type LeadWithExternalId struct {
		LeadExternalId__c string
		LastName          string
		Company           string
	}
	leadsWithExternalId := []LeadWithExternalId{
		{
			LeadExternalId__c: "MK3",
			LastName:          "Lockley",
			Company:           "The Avengers",
		},
	}
	jobIds, err = sf.UpsertBulkAssign("Lead", "LeadExternalId__c", leadsWithExternalId, 100, true, "00QDn0000024r6FMAQ")
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)
}

func bulkDmlFile(sf *salesforce.Salesforce) {
	fmt.Println("===== Bulk Insert File")

	jobIds, err := sf.InsertBulkFile("Contact", "data/avengers.csv", 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Update File")

	jobIds, err = sf.UpdateBulkFile("Contact", "data/update_avengers.csv", 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Upsert File")

	jobIds, err = sf.UpsertBulkFile("Contact", "ContactExternalId__c", "data/upsert_avengers.csv", 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Delete File")
	jobIds, err = sf.DeleteBulkFile("Contact", "data/delete_avengers.csv", 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)
}

func bulkDml(sf *salesforce.Salesforce) {
	fmt.Println("===== Bulk Insert")

	type Contact struct {
		LastName string
	}

	contacts := []Contact{
		{
			LastName: "Lang",
		},
		{
			LastName: "Van Dyne",
		},
	}
	jobIds, err := sf.InsertBulk("Contact", contacts, 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Update")

	type ContactWithId struct {
		Id       string
		LastName string
	}

	var queriedContacts []ContactWithId
	err = sf.Query("SELECT Id FROM Contact LIMIT 2", &queriedContacts)
	if err != nil {
		panic(err)
	}

	contactsWithId := []ContactWithId{
		{
			Id:       queriedContacts[0].Id,
			LastName: "Strange",
		},
		{
			Id:       queriedContacts[1].Id,
			LastName: "T'Challa",
		},
	}
	jobIds, err = sf.UpdateBulk("Contact", contactsWithId, 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Upsert")

	type ContactWithExternalId struct {
		ContactExternalId__c string
		LastName             string
	}

	contactsWithExternalId := []ContactWithExternalId{
		{
			ContactExternalId__c: "Avng5",
			LastName:             "Rhodes",
		},
		{
			ContactExternalId__c: "Avng6",
			LastName:             "Quill",
		},
	}
	jobIds, err = sf.UpsertBulk("Contact", "ContactExternalId__c", contactsWithExternalId, 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)

	fmt.Println("===== Bulk Delete")

	type ContactToDelete struct {
		Id string
	}
	contactsToDelete := []ContactToDelete{
		{
			Id: queriedContacts[0].Id,
		},
		{
			Id: queriedContacts[0].Id,
		},
	}
	jobIds, err = sf.DeleteBulk("Contact", contactsToDelete, 1000, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobIds)
}

func queryBulk(sf *salesforce.Salesforce) {

	fmt.Println("===== Query Bulk Export")

	err := sf.QueryBulkExport("SELECT Id, FirstName, LastName FROM Contact", "data/export.csv")
	if err != nil {
		panic(err)
	}

	fmt.Println("===== Query Struct Bulk Export")

	type ContactSoql struct {
		Id        string `soql:"selectColumn,fieldName=Id" json:"Id"`
		FirstName string `soql:"selectColumn,fieldName=FirstName" json:"FirstName"`
		LastName  string `soql:"selectColumn,fieldName=LastName" json:"LastName"`
	}

	type ContactSoqlQuery struct {
		SelectClause ContactSoql `soql:"selectClause,tableName=Contact"`
	}

	soqlStruct := ContactSoqlQuery{
		SelectClause: ContactSoql{},
	}
	err = sf.QueryStructBulkExport(soqlStruct, "data/export2.csv")
	if err != nil {
		panic(err)
	}

	fmt.Println("===== Bulk Iterator")

	type Contact struct {
		Id        string `json:"Id" csv:"Id"`
		FirstName string `json:"FirstName" csv:"FirstName"`
		LastName  string `json:"LastName" csv:"LastName"`
	}

	it, err := sf.QueryBulkIterator("SELECT Id, FirstName, LastName FROM Contact")
	if err != nil {
		panic(err)
	}

	for it.Next() {
		var data []Contact
		if err := it.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Println(data)
	}

	if err := it.Error(); err != nil {
		panic(err)
	}

	type User struct {
		Name string
	}

	fmt.Println("===== Querying Relationships")

	type ContactWithAltOwner struct {
		Id                 string
		AlternateOwnerName User `csv:"Alternate_Owner__r.,inline"`
	}

	it, err = sf.QueryBulkIterator("SELECT Id, Alternate_Owner__r.Name FROM Contact")
	if err != nil {
		panic(err)
	}

	for it.Next() {
		var data []ContactWithAltOwner
		if err := it.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Println(data)
	}

	if err := it.Error(); err != nil {
		panic(err)
	}
}

func dmlComposite(sf *salesforce.Salesforce) {
	fmt.Println("===== Insert Composite")
	type Contact struct {
		LastName string
	}

	contacts := []Contact{
		{
			LastName: "Parker",
		},
		{
			LastName: "Murdock",
		},
	}
	results, err := sf.InsertComposite("Contact", contacts, 200, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("===== Update Composite")

	type ContactWithId struct {
		Id       string
		LastName string
	}

	contactsWithId := []ContactWithId{
		{
			Id:       results.Results[0].Id,
			LastName: "Richards",
		},
		{
			Id:       results.Results[1].Id,
			LastName: "Storm",
		},
	}
	results, err = sf.UpdateComposite("Contact", contactsWithId, 200, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("===== Upsert Composite")

	type ContactWithExternalId struct {
		ContactExternalId__c string
		LastName             string
	}

	contactsWithExternalId := []ContactWithExternalId{
		{
			ContactExternalId__c: "Avng3",
			LastName:             "Maximoff",
		},
		{
			ContactExternalId__c: "Avng4",
			LastName:             "Wilson",
		},
	}
	results, err = sf.UpsertComposite("Contact", "ContactExternalId__c", contactsWithExternalId, 200, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("===== Delete Composite")

	type ContactToDelete struct {
		Id string
	}

	contactsToDelete := []ContactToDelete{
		{
			Id: results.Results[0].Id,
		},
		{
			Id: results.Results[1].Id,
		},
	}
	results, err = sf.DeleteComposite("Contact", contactsToDelete, 200, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}

func dmlCollections(sf *salesforce.Salesforce) {
	fmt.Println("===== Insert Collection")
	type Contact struct {
		LastName string
	}

	contacts := []Contact{
		{
			LastName: "Barton",
		},
		{
			LastName: "Romanoff",
		},
	}
	results, err := sf.InsertCollection("Contact", contacts, 200)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("===== Update Collection")

	type ContactWithId struct {
		Id       string
		LastName string
	}
	contactsWithId := []ContactWithId{
		{
			Id:       results.Results[0].Id,
			LastName: "Fury",
		},
		{
			Id:       results.Results[1].Id,
			LastName: "Odinson",
		},
	}
	results, err = sf.UpdateCollection("Contact", contactsWithId, 200)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("===== Upsert Collection")

	type ContactWithExternalId struct {
		ContactExternalId__c string
		LastName             string
	}

	contactsWithExternalId := []ContactWithExternalId{
		{
			ContactExternalId__c: "Avng34",
			LastName:             "Danvers",
		},
		{
			ContactExternalId__c: "Avng35",
			LastName:             "Pym",
		},
	}
	results, err = sf.UpsertCollection("Contact", "ContactExternalId__c", contactsWithExternalId, 200)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("===== Delete Collection")

	type ContactToDelete struct {
		Id string
	}

	contactsToDelete := []ContactToDelete{
		{
			Id: results.Results[0].Id,
		},
		{
			Id: results.Results[0].Id,
		},
	}
	results, err = sf.DeleteCollection("Contact", contactsToDelete, 200)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}

func dmlSingle(sf *salesforce.Salesforce) {
	fmt.Println("===== Insert One")
	type Contact struct {
		LastName string
	}

	contact := Contact{
		LastName: "Stark",
	}
	result, err := sf.InsertOne("Contact", contact)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	fmt.Println("===== Update One")

	type ContactWithId struct {
		Id       string
		LastName string
	}

	contactWithId := ContactWithId{
		Id:       result.Id,
		LastName: "Banner",
	}
	err = sf.UpdateOne("Contact", contactWithId)
	if err != nil {
		panic(err)
	}

	fmt.Println("===== Upsert One")

	type ContactWithExternalId struct {
		ContactExternalId__c string
		LastName             string
	}
	contactWithExternalId := ContactWithExternalId{
		ContactExternalId__c: "Avng0",
		LastName:             "Rogers",
	}
	result, err = sf.UpsertOne("Contact", "ContactExternalId__c", contactWithExternalId)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	fmt.Println("===== Delete One")

	type ContactToDelete struct {
		Id string
	}

	contactToDelete := ContactToDelete{
		Id: result.Id,
	}
	err = sf.DeleteOne("Contact", contactToDelete)
	if err != nil {
		panic(err)
	}
}

func queryStruct(sf *salesforce.Salesforce) {
	fmt.Println("===== QueryStruct")
	type Contact struct {
		Id       string `soql:"selectColumn,fieldName=Id" json:"Id"`
		LastName string `soql:"selectColumn,fieldName=LastName" json:"LastName"`
	}

	type ContactQueryCriteria struct {
		LastName string `soql:"equalsOperator,fieldName=LastName"`
	}

	type ContactSoqlQuery struct {
		SelectClause Contact              `soql:"selectClause,tableName=Contact"`
		WhereClause  ContactQueryCriteria `soql:"whereClause"`
	}

	soqlStruct := ContactSoqlQuery{
		SelectClause: Contact{},
		WhereClause: ContactQueryCriteria{
			LastName: "Bond",
		},
	}
	contacts := []Contact{}
	err := sf.QueryStruct(soqlStruct, &contacts)

	if err != nil {
		panic(err)
	}

	fmt.Println(contacts)
}

func query(sf *salesforce.Salesforce) {
	fmt.Println("===== Query")
	type Contact struct {
		Id       string
		LastName string
	}

	contacts := []Contact{}
	err := sf.Query("SELECT Id, LastName FROM Contact", &contacts)
	if err != nil {
		panic(err)
	}
	fmt.Println(contacts)
}

func getAccessTokenAndInstanceUrl(sf *salesforce.Salesforce) {
	fmt.Println("===== GetAccessToken")
	token := sf.GetAccessToken()
	fmt.Println(token)

	fmt.Println("===== GetInstanceUrl")
	url := sf.GetInstanceUrl()
	fmt.Println(url)
}

func functionalConfigExample() {
	fmt.Println("===== functionalConfig")
	// Example of the new functional configuration approach

	// Basic usage with default configuration (same as before for backwards compatibility)
	creds := salesforce.Creds{
		Domain:         domain,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}

	// Initialize with default configuration
	sf, err := salesforce.Init(creds)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Default API Version: %s\n", sf.GetAPIVersion())
	fmt.Printf("Default Batch Size Max: %d\n", sf.GetBatchSizeMax())
	fmt.Printf("Default Bulk Batch Size Max: %d\n", sf.GetBulkBatchSizeMax())
	fmt.Printf("Compression Headers: %v\n", sf.GetCompressionHeaders())
	fmt.Printf("Auth Flow: %s\n", sf.GetAuthFlow())

	// Example with custom configuration using functional options
	sfCustom, err := salesforce.Init(creds,
		salesforce.WithAPIVersion("v58.0"),
		salesforce.WithBatchSizeMax(150),
		salesforce.WithBulkBatchSizeMax(8000),
		salesforce.WithCompressionHeaders(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCustom API Version: %s\n", sfCustom.GetAPIVersion())
	fmt.Printf("Custom Batch Size Max: %d\n", sfCustom.GetBatchSizeMax())
	fmt.Printf("Custom Bulk Batch Size Max: %d\n", sfCustom.GetBulkBatchSizeMax())
	fmt.Printf("Custom Compression Headers: %v\n", sfCustom.GetCompressionHeaders())
	fmt.Printf("Auth Flow: %s\n", sfCustom.GetAuthFlow())

	// Example with error handling in functional options
	_, err = salesforce.Init(creds,
		salesforce.WithAPIVersion(""), // This will cause an error
	)
	if err != nil {
		fmt.Printf("\nExpected error: %s\n", err)
	}

	// Example with invalid batch size
	_, err = salesforce.Init(creds,
		salesforce.WithBatchSizeMax(300), // This will cause an error (max is 200)
	)
	if err != nil {
		fmt.Printf("Expected error: %s\n", err)
	}
}

func httpConfigExample() {
	fmt.Println("===== httpConfig")
	// Example 1: Using a custom transport with timeout and TLS configuration

	creds := salesforce.Creds{
		Domain:         domain,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}

	// Initialize Salesforce client with custom HTTP client
	sf, err := salesforce.Init(creds,
		salesforce.WithRoundTripper(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false, // Set to true if you need to skip SSL verification
			},
			MaxIdleConns:       20,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: false,
		}),
	)
	if err != nil {
		log.Fatal("Failed to initialize Salesforce client:", err)
	}

	fmt.Printf("Salesforce client initialized with custom HTTP client\n")
	fmt.Printf("HTTP Client timeout: %v\n", sf.GetHTTPClient().Timeout)
	fmt.Printf("API Version: %s\n", sf.GetAPIVersion())

	// Example 2: Using a custom round tripper
	customRoundTripper := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		DisableCompression:  false,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 5,
	}

	sf2, err := salesforce.Init(creds,
		salesforce.WithRoundTripper(customRoundTripper),
		salesforce.WithAPIVersion("v64.0"),
	)
	if err != nil {
		log.Fatal("Failed to initialize Salesforce client with round tripper:", err)
	}

	fmt.Printf("Salesforce client initialized with custom round tripper\n")
	fmt.Printf("API Version: %s\n", sf2.GetAPIVersion())

	// Example 3: Using default configuration
	sf3, err := salesforce.Init(creds)
	if err != nil {
		log.Fatal("Failed to initialize Salesforce client with defaults:", err)
	}

	fmt.Printf("Salesforce client initialized with default configuration\n")
	fmt.Printf("HTTP Client timeout: %v\n", sf3.GetHTTPClient().Timeout)
	fmt.Printf("Compression headers enabled: %v\n", sf3.GetCompressionHeaders())

	// Example of combining multiple configuration options
	sf4, err := salesforce.Init(creds,
		salesforce.WithRoundTripper(http.DefaultTransport),
		salesforce.WithCompressionHeaders(true),
		salesforce.WithAPIVersion("v65.0"),
		salesforce.WithBatchSizeMax(150),
		salesforce.WithBulkBatchSizeMax(5000),
	)
	if err != nil {
		log.Fatal("Failed to initialize Salesforce client with multiple options:", err)
	}

	fmt.Printf("Salesforce client initialized with multiple configuration options\n")
	fmt.Printf("API Version: %s\n", sf4.GetAPIVersion())
	fmt.Printf("Batch Size Max: %d\n", sf4.GetBatchSizeMax())
	fmt.Printf("Bulk Batch Size Max: %d\n", sf4.GetBulkBatchSizeMax())
	fmt.Printf("Compression headers enabled: %v\n", sf4.GetCompressionHeaders())
}
