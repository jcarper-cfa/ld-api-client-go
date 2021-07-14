/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.3.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UserSegmentBody struct {
	// A human-friendly name for the user segment.
	Name string `json:"name"`
	// A unique key that will be used to reference the user segment in feature flags.
	Key string `json:"key"`
	// A description for the user segment.
	Description string `json:"description,omitempty"`
	// Controls whether this is considered a \"big segment\" which can support an unlimited numbers of users. Include/exclude lists sent with this payload are not used in big segments. Contact your account manager for early access to this feature.
	Unbounded bool `json:"unbounded,omitempty"`
	// Tags for the user segment.
	Tags []string `json:"tags,omitempty"`
}
