/*
 * Customers API
 *
 * Customers focuses on solving authentic identification of humans who are legally able to hold and transfer currency within the US. Primarily this project solves [Know Your Customer](https://en.wikipedia.org/wiki/Know_your_customer) (KYC), [Customer Identification Program](https://en.wikipedia.org/wiki/Customer_Identification_Program) (CIP), [Office of Foreign Asset Control](https://www.treasury.gov/about/organizational-structure/offices/Pages/Office-of-Foreign-Assets-Control.aspx) (OFAC) checks and verification workflows to comply with US federal law and ensure authentic transfers. Also, Customers has an objective to be a service for detailed due diligence on individuals and companies for Financial Institutions and services in a modernized and extensible way.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// CreateCustomer struct for CreateCustomer
type CreateCustomer struct {
	// Given Name or First Name
	FirstName string `json:"firstName"`
	// Middle Name
	MiddleName string `json:"middleName,omitempty"`
	// Surname or Last Name
	LastName string `json:"lastName"`
	// Name Customer is preferred to be called
	NickName string `json:"nickName,omitempty"`
	// Customers name suffix. \"Jr\", \"PH.D.\"
	Suffix string `json:"suffix,omitempty"`
	// Legal date of birth
	BirthDate time.Time `json:"birthDate"`
	// Primary email address of customer name@domain.com
	Email string `json:"email"`
	// Customer Social Security Number (SSN)
	SSN       string                  `json:"SSN,omitempty"`
	Phones    []CreatePhone           `json:"phones,omitempty"`
	Addresses []CreateCustomerAddress `json:"addresses"`
	// Map of unique keys associated to values to act as foreign key relationships or arbitrary data associated to a Customer.
	Metadata map[string]string `json:"metadata,omitempty"`
}
