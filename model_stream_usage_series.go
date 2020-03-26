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

type StreamUsageSeries struct {
	// A key corresponding to a time series data point.
	Var0 float32 `json:"0,omitempty"`
	// A unix epoch time in milliseconds specifying the creation time of this flag.
	Time float32 `json:"time,omitempty"`
}
