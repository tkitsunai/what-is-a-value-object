package vo

// You should convert this code to ValueObject
type _Company struct {
	name         string
	address1     string
	address2     string
	listedFlag   bool
	securityCode int
	phoneNumber  string
}

type Address struct {
	value string
}
type ListedStatus struct {
	Listed bool
}
type PhoneNumber struct {
	value string
}
type PersonName struct {
	value string
}
type CompanyName struct {
	value string
}

// Phase1: Do not using Primitive value
type Company struct {
	name         Name
	address1     Address
	address2     Address
	listedFlag   ListedStatus
	securityCode SecurityCode
	phoneNumber  PhoneNumber
}

// Phase2: Is the name appropriate?
type ListedCompany struct {
	name         CompanyName
	address1     Address
	address2     Address
	securityCode SecurityCode
	phoneNumber  PhoneNumber
}
type UnlistedCompany struct {
	name        CompanyName
	address1    Address
	address2    Address
	phoneNumber PhoneNumber
}
