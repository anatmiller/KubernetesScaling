/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type MappedPipe struct {

	ContainerPipeName string `json:"ContainerPipeName,omitempty"`

	HostPath string `json:"HostPath,omitempty"`

	HostPathType string `json:"HostPathType,omitempty"`
}
