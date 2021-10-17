package main

import (
	"fmt"
	"testing"
)

func TestTService_TString(t *testing.T) {
	service := &TService{
		Methods: map[string]*TMethod{
			"sendMessage": {
				RequestField: map[int]*TField{
					1: {
						Name: "reqSeq",
						Type: &TString{},
					},
					2: {
						Name: "message",
						Type: &TStruct{
							Name: "Message",
							Fields: map[int]*TField{
								1: {
									Name: "from",
									Type: &TString{},
								},
							},
						},
					},
				},
				Response: &TString{},
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
	fmt.Println(service.TString())
}
