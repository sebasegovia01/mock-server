package models

type User struct {
	PersonalIdentification PersonalIdentification `json:"personalIdentification"`
	Accounts               Accounts               `json:"accounts"`
}

type PersonalIdentification struct {
	CustomerIdentification string                `json:"customerIdentification"`
	CustomerFirstName      string                `json:"customerFirstName"`
	CustomerMiddleName     string                `json:"customerMiddleName"`
	CustomerLastName       string                `json:"customerLastName"`
	CustomerSecondLastName string                `json:"customerSecondLastName"`
	CustomerInitialDate    string                `json:"customerInitialDate"`
	AdditionalInformation  AdditionalInformation `json:"additionalInformation"`
}

type AdditionalInformation struct {
	CustomerStreetName                  string `json:"customerStreetName"`
	CustomerBuildingNumber              string `json:"customerBuildingNumber"`
	CustomerDistrictName                string `json:"customerDistrictName"`
	CustomerCountrySubDivisionMajorName string `json:"customerCountrySubDivisionMajorName"`
	CustomerEmail                       string `json:"customerEmail"`
	CustomerPhoneNumber                 string `json:"customerPhoneNumber"`
}

type Accounts struct {
	AccountInformation AccountInformation `json:"accountInformation"`
	Balances           Balances           `json:"balances"`
	OverdraftLimits    OverdraftLimits    `json:"overdraftLimits"`
	Transactions       Transactions       `json:"transactions"`
}

type AccountInformation struct {
	ProductIdentification  string `json:"productIdentification"`
	CustomerIdentification string `json:"customerIdentification"`
	ProductType            string `json:"productType"`
	CommercialCategory     string `json:"commercialCategory"`
	ContractDate           string `json:"contractDate"`
}

type Balances struct {
	BalanceAmount float64 `json:"balanceAmount"`
	CurrencyCode  string  `json:"currencyCode"`
}

type Transactions struct {
	OperationDate              string  `json:"operationDate"`
	AccountingDate             string  `json:"accountingDate"`
	OperationType              string  `json:"operationType"`
	OperationAmount            float64 `json:"operationAmount"`
	OperationCurrencyCode      string  `json:"operationCurrencyCode"`
	PartyType                  string  `json:"partyType"`
	CounterpartyIdentification string  `json:"counterpartyIdentification"`
	CounterpartyName           string  `json:"counterpartyName"`
	OperationCategory          string  `json:"operationCategory"`
}

type OverdraftLimits struct {
	CreditLineUsedAmount                float64 `json:"creditLineUsedAmount"`
	CreditLineTotalAmount               float64 `json:"creditLineTotalAmount"`
	CreditLineAvailableAmount           float64 `json:"creditLineAvailableAmount"`
	EndOfMonthCreditLineUsedAmount      float64 `json:"endOfMonthCreditLineUsedAmount"`
	EndOfMonthCreditLineTotalAmount     float64 `json:"endOfMonthCreditLineTotalAmount"`
	EndOfMonthCreditLineAvailableAmount float64 `json:"endOfMonthCreditLineAvailableAmount"`
}
