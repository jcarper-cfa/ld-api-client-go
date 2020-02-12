/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.29
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagStatusAcrossEnvironments struct {
	Links *Links `json:"_links,omitempty"`
	Key string `json:"key,omitempty"`
	Environments map[string]FeatureFlagStatusForQueriedEnvironment `json:"environments,omitempty"`
}
