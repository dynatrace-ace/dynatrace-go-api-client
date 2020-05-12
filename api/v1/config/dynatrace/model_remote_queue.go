/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// RemoteQueue Define a local definition of a remote queue owned by another queue manager. The local definition can be made visible in one or more clusters.
type RemoteQueue struct {
	// The name of the local definition of the remote queue.
	LocalQueue string `json:"localQueue"`
	// The name of the remote queue.
	RemoteQueue string `json:"remoteQueue"`
	// The name of the remote queue manager.
	RemoteQueueManager string `json:"remoteQueueManager"`
	// The local definition of the remote queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters.
	ClusterVisibility []string `json:"clusterVisibility"`
}