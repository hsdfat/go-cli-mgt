/*
 * Management service
 *
 * No description provided (generated by models_api Codegen https://github.com/models_api-api/models_api-codegen)
 *
 * API version: 1.0
 * Generated by: models_api Codegen (https://github.com/models_api-api/models_api-codegen.git)
 */
package models_api

type ChangePassWord struct {
	Username string `json:"username,omitempty"`

	OldPassword string `json:"old-password,omitempty"`

	NewPassword string `json:"new-password,omitempty"`
}
