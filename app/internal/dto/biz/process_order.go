package biz

type OrderInfo struct {
	BaseFormData      map[string]interface{}              `json:"baseFormData"`
	FormData          map[string]interface{}              `json:"formData"`
	GroupFormDataInfo map[string][]map[string]interface{} `json:"groupFormDataInfo"`
}

type Task struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type TeamInfo struct {
	Name     string `json:"name"`
	Leader   string `json:"leader"`
	Director string `json:"director"`
}
