/*
 * Management service
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models_api

type Role struct {

	RoleId int64 `json:"role-id,omitempty"`

	RoleName string `json:"role-name,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	Description string `json:"description,omitempty"`
}
