package main

import (
	"fmt"
	"testing"
)

func TestFormatThriftService(t *testing.T) {
	message := &TStruct{
		Name: "Message",
		Fields: map[int]*TField{
			1: {
				Name: "from",
				Type: &TString{},
			},
			2: {
				Name: "location",
				Type: &TStruct{
					Name: "Location",
					Fields: map[int]*TField{
						1: {
							Name: "address",
							Type: &TString{},
						},
					},
				},
			},
		},
	}
	service := &TService{
		Methods: map[string]*TMethod{
			"sendMessage": {
				Parameters: map[int]*TField{
					1: {
						Name: "reqSeq",
						Type: &TString{},
					},
					2: {
						Name: "message",
						Type: message,
					},
				},
				Response: message,
				Name:     "sendMessage",
				Exception: &TException{
					TStruct: &TStruct{
						Name: "TalkException",
						Fields: map[int]*TField{
							1: {
								Name: "errorCode",
								Type: &TEnum{
									Fields: map[int]*TEnumField{
										0: {
											Name: "ILLEGAL_ARGUMENT",
										},
										1: {
											Name: "AUTHENTICATION_FAILED",
										},
										2: {
											Name: "DB_FAILED",
										},
									},
									Name: "TalkErrorCode",
								},
							},
							2: {
								Name: "reason",
								Type: &TString{},
							},
							3: {
								Name: "parameterMap",
								Type: &TMap{
									KeyType:   &TString{},
									ValueType: &TString{},
								},
							},
						},
					},
				},
			},
		},
		Name: "TalkService",
	}
	fmt.Println(FormatThriftService(service))
}
