/*
 * Management service
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models_api

import "time"

type History struct {
	Id           uint64    `json:"id"`
	Username     string    `json:"username,omitempty"`
	UserIp       string    `json:"user-ip,omitempty"`
	Command      string    `json:"command,omitempty"`
	NeName       string    `json:"ne-name,omitempty"`
	Result       bool      `json:"result,omitempty"`
	ExecutedTime time.Time `json:"executed-time,omitempty"`
	Mode         string    `json:"mode"`
}
