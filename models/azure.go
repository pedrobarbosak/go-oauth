package models

type Azure struct {
	ID          string `json:"id"`
	Name        string `json:"givenName"`
	Surname     string `json:"surname"`
	DisplayName string `json:"displayName"`
	Email       string `json:"userPrincipalName"`
}
