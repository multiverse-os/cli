package is

import (
	"regexp"
	"unicode"

	is "github.com/multiverse-os/cli/data/is"
	validate "github.com/multiverse-os/cli/data/validate"
)

var (
	// Email validates if a string is an email or not.
	Email = validate.NewStringRule(is.IsEmail, "must be a valid email address")
	// URL validates if a string is a valid URL
	URL = validate.NewStringRule(is.IsURL, "must be a valid URL")
	// RequestURL validates if a string is a valid request URL
	RequestURL = validate.NewStringRule(is.IsRequestURL, "must be a valid request URL")
	// RequestURI validates if a string is a valid request URI
	RequestURI = validate.NewStringRule(is.IsRequestURI, "must be a valid request URI")
	// Alpha validates if a string contains English letters only (a-zA-Z)
	Alpha = validate.NewStringRule(is.IsAlpha, "must contain English letters only")
	// Digit validates if a string contains digits only (0-9)
	Digit = validate.NewStringRule(isDigit, "must contain digits only")
	// Alphanumeric validates if a string contains English letters and digits only (a-zA-Z0-9)
	Alphanumeric = validate.NewStringRule(is.IsAlphanumeric, "must contain English letters and digits only")
	// UTFLetter validates if a string contains unicode letters only
	UTFLetter = validate.NewStringRule(is.IsUTFLetter, "must contain unicode letter characters only")
	// UTFDigit validates if a string contains unicode decimal digits only
	UTFDigit = validate.NewStringRule(is.IsUTFDigit, "must contain unicode decimal digits only")
	// UTFLetterNumeric validates if a string contains unicode letters and numbers only
	UTFLetterNumeric = validate.NewStringRule(is.IsUTFLetterNumeric, "must contain unicode letters and numbers only")
	// UTFNumeric validates if a string contains unicode number characters (category N) only
	UTFNumeric = validate.NewStringRule(isUTFNumeric, "must contain unicode number characters only")
	// LowerCase validates if a string contains lower case unicode letters only
	LowerCase = validate.NewStringRule(is.IsLowerCase, "must be in lower case")
	// UpperCase validates if a string contains upper case unicode letters only
	UpperCase = validate.NewStringRule(is.IsUpperCase, "must be in upper case")
	// Hexadecimal validates if a string is a valid hexadecimal number
	Hexadecimal = validate.NewStringRule(is.IsHexadecimal, "must be a valid hexadecimal number")
	// HexColor validates if a string is a valid hexadecimal color code
	HexColor = validate.NewStringRule(is.IsHexcolor, "must be a valid hexadecimal color code")
	// RGBColor validates if a string is a valid RGB color in the form of rgb(R, G, B)
	RGBColor = validate.NewStringRule(is.IsRGBcolor, "must be a valid RGB color code")
	// Int validates if a string is a valid integer number
	Int = validate.NewStringRule(is.IsInt, "must be an integer number")
	// Float validates if a string is a floating point number
	Float = validate.NewStringRule(is.IsFloat, "must be a floating point number")
	// UUIDv3 validates if a string is a valid version 3 UUID
	UUIDv3 = validate.NewStringRule(is.IsUUIDv3, "must be a valid UUID v3")
	// UUIDv4 validates if a string is a valid version 4 UUID
	UUIDv4 = validate.NewStringRule(is.IsUUIDv4, "must be a valid UUID v4")
	// UUIDv5 validates if a string is a valid version 5 UUID
	UUIDv5 = validate.NewStringRule(is.IsUUIDv5, "must be a valid UUID v5")
	// UUID validates if a string is a valid UUID
	UUID = validate.NewStringRule(is.IsUUID, "must be a valid UUID")
	// CreditCard validates if a string is a valid credit card number
	CreditCard = validate.NewStringRule(is.IsCreditCard, "must be a valid credit card number")
	// ISBN10 validates if a string is an ISBN version 10
	ISBN10 = validate.NewStringRule(is.IsISBN10, "must be a valid ISBN-10")
	// ISBN13 validates if a string is an ISBN version 13
	ISBN13 = validate.NewStringRule(is.IsISBN13, "must be a valid ISBN-13")
	// ISBN validates if a string is an ISBN (either version 10 or 13)
	ISBN = validate.NewStringRule(isISBN, "must be a valid ISBN")
	// JSON validates if a string is in valid JSON format
	JSON = validate.NewStringRule(is.IsJSON, "must be in valid JSON format")
	// ASCII validates if a string contains ASCII characters only
	ASCII = validate.NewStringRule(is.IsASCII, "must contain ASCII characters only")
	// PrintableASCII validates if a string contains printable ASCII characters only
	PrintableASCII = validate.NewStringRule(is.IsPrintableASCII, "must contain printable ASCII characters only")
	// Multibyte validates if a string contains multibyte characters
	Multibyte = validate.NewStringRule(is.IsMultibyte, "must contain multibyte characters")
	// FullWidth validates if a string contains full-width characters
	FullWidth = validate.NewStringRule(is.IsFullWidth, "must contain full-width characters")
	// HalfWidth validates if a string contains half-width characters
	HalfWidth = validate.NewStringRule(is.IsHalfWidth, "must contain half-width characters")
	// VariableWidth validates if a string contains both full-width and half-width characters
	VariableWidth = validate.NewStringRule(is.IsVariableWidth, "must contain both full-width and half-width characters")
	// Base64 validates if a string is encoded in Base64
	Base64 = validate.NewStringRule(is.IsBase64, "must be encoded in Base64")
	// DataURI validates if a string is a valid base64-encoded data URI
	DataURI = validate.NewStringRule(is.IsDataURI, "must be a Base64-encoded data URI")
	// E164 validates if a string is a valid ISO3166 Alpha 2 country code
	E164 = validate.NewStringRule(isE164Number, "must be a valid E164 number")
	// CountryCode2 validates if a string is a valid ISO3166 Alpha 2 country code
	CountryCode2 = validate.NewStringRule(is.IsISO3166Alpha2, "must be a valid two-letter country code")
	// CountryCode3 validates if a string is a valid ISO3166 Alpha 3 country code
	CountryCode3 = validate.NewStringRule(is.IsISO3166Alpha3, "must be a valid three-letter country code")
	// CurrencyCode validates if a string is a valid IsISO4217 currency code.
	CurrencyCode = validate.NewStringRule(is.IsISO4217, "must be valid ISO 4217 currency code")
	// DialString validates if a string is a valid dial string that can be passed to Dial()
	DialString = validate.NewStringRule(is.IsDialString, "must be a valid dial string")
	// MAC validates if a string is a MAC address
	MAC = validate.NewStringRule(is.IsMAC, "must be a valid MAC address")
	// IP validates if a string is a valid IP address (either version 4 or 6)
	IP = validate.NewStringRule(is.IsIP, "must be a valid IP address")
	// IPv4 validates if a string is a valid version 4 IP address
	IPv4 = validate.NewStringRule(is.IsIPv4, "must be a valid IPv4 address")
	// IPv6 validates if a string is a valid version 6 IP address
	IPv6 = validate.NewStringRule(is.IsIPv6, "must be a valid IPv6 address")
	// Subdomain validates if a string is valid subdomain
	Subdomain = validate.NewStringRule(isSubdomain, "must be a valid subdomain")
	// Domain validates if a string is valid domain
	Domain = validate.NewStringRule(isDomain, "must be a valid domain")
	// DNSName validates if a string is valid DNS name
	DNSName = validate.NewStringRule(is.IsDNSName, "must be a valid DNS name")
	// Host validates if a string is a valid IP (both v4 and v6) or a valid DNS name
	Host = validate.NewStringRule(is.IsHost, "must be a valid IP address or DNS name")
	// Port validates if a string is a valid port number
	Port = validate.NewStringRule(is.IsPort, "must be a valid port number")
	// MongoID validates if a string is a valid Mongo ID
	MongoID = validate.NewStringRule(is.IsMongoID, "must be a valid hex-encoded MongoDB ObjectId")
	// Latitude validates if a string is a valid latitude
	Latitude = validate.NewStringRule(is.IsLatitude, "must be a valid latitude")
	// Longitude validates if a string is a valid longitude
	Longitude = validate.NewStringRule(is.IsLongitude, "must be a valid longitude")
	// SSN validates if a string is a social security number (SSN)
	SSN = validate.NewStringRule(is.IsSSN, "must be a valid social security number")
	// Semver validates if a string is a valid semantic version
	Semver = validate.NewStringRule(is.IsSemver, "must be a valid semantic version")
)

var (
	reDigit = regexp.MustCompile("^[0-9]+$")
	// Subdomain regex source: https://stackoverflow.com/a/7933253
	reSubdomain = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9\-]{0,61}[A-Za-z0-9])?$`)
	// E164 regex source: https://stackoverflow.com/a/23299989
	reE164 = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	// Domain regex source: https://stackoverflow.com/a/7933253
	// Slightly modified: Removed 255 max length validate since Go regex does not
	// support lookarounds. More info: https://stackoverflow.com/a/38935027
	reDomain = regexp.MustCompile(`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-z]{1,63}| xn--[a-z0-9]{1,59})$`)
)

func isISBN(value string) bool {
	return is.IsISBN(value, 10) || is.IsISBN(value, 13)
}

func isDigit(value string) bool {
	return reDigit.MatchString(value)
}

func isE164Number(value string) bool {
	return reE164.MatchString(value)
}

func isSubdomain(value string) bool {
	return reSubdomain.MatchString(value)
}

func isDomain(value string) bool {
	if len(value) > 255 {
		return false
	}

	return reDomain.MatchString(value)
}

func isUTFNumeric(value string) bool {
	for _, c := range value {
		if unicode.IsNumber(c) == false {
			return false
		}
	}
	return true
}
