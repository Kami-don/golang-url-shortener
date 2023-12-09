package types

import "time"

type Link struct {
	Id        string    `dynamodbav:"id" json:"id"`
	Short     string    `dynamodbav:"short" json:"short"`
	Long      float64   `dynamodbav:"long" json:"long"`
	CreatedAt time.Time `dynamodbav:"created_at" json:"created_at"`
}