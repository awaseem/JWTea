package main

import (
	"bytes"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	bucketJWTea = "jwtea"
	region      = "us-west-2"
)

var svc *s3.S3

// Initialize start store
func Initialize() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	svc = s3.New(sess)
}

// Set set datastore value
func Set(key string, value []byte) error {
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketJWTea),
		Key:    aws.String(key),
		Body:   bytes.NewReader(value),
	}
	_, err := svc.PutObject(params)
	if err != nil {
		return err
	}
	return nil
}

// Get datastore value
func Get(key string) ([]byte, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketJWTea),
		Key:    aws.String(key),
	}
	resp, err := svc.GetObject(params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
