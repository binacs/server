/*
 * user.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BinacsApiUserV1UserAuthResp struct {
	Code string `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data *BinacsApiUserV1UserTokenObj `json:"data,omitempty"`
}
