/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.33
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UserRecord struct {
	LastPing string `json:"lastPing,omitempty"`
	EnvironmentId string `json:"environmentId,omitempty"`
	// The unique resource id.
	OwnerId string `json:"ownerId,omitempty"`
	User *User `json:"user,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
