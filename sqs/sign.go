package sqs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/goamz/goamz/aws"
	"sort"
	"strings"
)

var b64 = base64.StdEncoding

func sign(auth *aws.Auth, method, path string, params map[string]string, host string) {
	accessKey, secretKey, token := auth.Credentials()
	params["AWSAccessKeyId"] = accessKey
	params["SignatureVersion"] = "2"
	params["SignatureMethod"] = "HmacSHA256"
	if token != "" {
		params["SecurityToken"] = token
	}

	var sarray []string
	for k, v := range params {
		sarray = append(sarray, aws.Encode(k)+"="+aws.Encode(v))
	}
	sort.StringSlice(sarray).Sort()
	joined := strings.Join(sarray, "&")
	payload := method + "\n" + host + "\n" + path + "\n" + joined
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum(nil))

	params["Signature"] = string(signature)
}
