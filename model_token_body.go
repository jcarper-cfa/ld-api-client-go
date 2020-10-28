/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.8.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type TokenBody struct {
	// A human-friendly name for the access token
	Name string `json:"name,omitempty"`
	// The name of a built-in role for the token
	Role string `json:"role,omitempty"`
	// A list of custom role IDs to use as access limits for the access token
	CustomRoleIds []string `json:"customRoleIds,omitempty"`
	InlineRole []Statement `json:"inlineRole,omitempty"`
	// Whether the token will be a service token https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens
	ServiceToken bool `json:"serviceToken,omitempty"`
	// The default API version for this token
	DefaultApiVersion int32 `json:"defaultApiVersion,omitempty"`
}
