package aws

import (
	"net/http"
	"time"
)

// Auth:
// Exporting methods for testing

func (a *Auth) AccessKey() string {
	return a.accessKey
}

func (a *Auth) SecretKey() string {
	return a.secretKey
}

func (a *Auth) Token() string {
	return a.token
}

func (a *Auth) Expiration() time.Time {
	return a.expiration
}

// V4Signer:
// Exporting methods for testing

func (s *V4Signer) RequestTime(req *http.Request) time.Time {
	return s.requestTime(req)
}

func (s *V4Signer) CanonicalRequest(req *http.Request) string {
	return s.canonicalRequest(req)
}

func (s *V4Signer) StringToSign(t time.Time, creq string) string {
	return s.stringToSign(t, creq)
}

func (s *V4Signer) Signature(t time.Time, sts, secretKey string) string {
	return s.signature(t, sts, secretKey)
}

func (s *V4Signer) Authorization(header http.Header, t time.Time, signature, accessKey string) string {
	return s.authorization(header, t, signature, accessKey)
}
