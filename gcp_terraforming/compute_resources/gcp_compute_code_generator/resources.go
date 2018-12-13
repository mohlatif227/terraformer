// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// metadata for generate code for GCP compute service
var terraformResources = map[string]gcpResourceRenderble{
	"addresses": basicGCPResource{
		terraformName: "google_compute_address",
		ignoreKeys: []string{
			"type",
			"users",
			"address",
		},
	},
	"autoscalers": basicGCPResource{
		terraformName: "google_compute_autoscaler",
	},
	"backendBuckets": basicGCPResource{
		terraformName: "google_compute_backend_bucket",
	},
	"backendServices": backendServices{
		basicGCPResource{terraformName: "google_compute_backend_service",
			ignoreKeys: []string{"region"},
		},
	},
	"disks": basicGCPResource{
		terraformName: "google_compute_disk",
		ignoreKeys: []string{
			"last_attach_timestamp",
			"last_detach_timestamp",
			"users",
			"source_image_id",
			"source_snapshot_id",
		},
	},
	"firewalls": basicGCPResource{
		terraformName: "google_compute_firewall",
	},
	"forwardingRules": basicGCPResource{
		terraformName: "google_compute_forwarding_rule",
		ignoreKeys:    []string{"service_name"},
	},
	"globalAddresses": basicGCPResource{
		terraformName: "google_compute_global_address",
		ignoreKeys:    []string{"address"},
	},
	"globalForwardingRules": globalForwardingRules{
		basicGCPResource{
			terraformName: "google_compute_global_forwarding_rule",
			ignoreKeys:    []string{"region"},
		},
	},
	"healthChecks": basicGCPResource{
		terraformName: "google_compute_health_check",
		ignoreKeys:    []string{"type"},
	},
	"httpHealthChecks": basicGCPResource{
		terraformName: "google_compute_http_health_check",
	},
	"httpsHealthChecks": basicGCPResource{
		terraformName: "google_compute_https_health_check",
	},
	"images": basicGCPResource{
		terraformName: "google_compute_image",
	},
	"instanceGroupManagers": instanceGroupManagers{
		basicGCPResource{
			terraformName:    "google_compute_instance_group_manager",
			ignoreKeys:       []string{"^instance_group$"},
			allowEmptyValues: []string{"^version.[0-9].name", "^auto_healing_policies.[0-9].health_check"},
		},
	},
	"instanceGroups": basicGCPResource{
		terraformName: "google_compute_instance_group",
		ignoreKeys:    []string{"size"},
	},
	"instanceTemplates": basicGCPResource{
		terraformName: "google_compute_instance_template",
		ignoreKeys:    []string{"tags_fingerprint"},
	},
	"instances": instances{
		basicGCPResource{
			terraformName:    "google_compute_instance",
			allowEmptyValues: []string{"labels."},
			ignoreKeys: []string{
				"^instance_id",
				"^metadata_fingerprint",
				"^tags_fingerprint",
				"^cpu_platform",
				"^network_interface.0.access_config.0.assigned_nat_ip",
				"^network_interface.0.name",
			},
			additionalFieldsForRefresh: map[string]string{
				"disk.#": "0",
			},
		},
	},
	"networks": basicGCPResource{
		terraformName: "google_compute_network",
		ignoreKeys:    []string{"gateway_ipv4"},
	},
	"regionAutoscalers": basicGCPResource{
		terraformName: "google_compute_region_autoscaler",
	},
	"regionBackendServices": basicGCPResource{
		terraformName: "google_compute_region_backend_service",
	},
	"regionDisks": basicGCPResource{
		terraformName: "google_compute_region_disk",
		ignoreKeys: []string{
			"last_attach_timestamp",
			"last_detach_timestamp",
			"users",
			"source_snapshot_id",
		},
	},
	"regionInstanceGroupManagers": basicGCPResource{
		terraformName:    "google_compute_region_instance_group_manager",
		ignoreKeys:       []string{"instance_group"},
		allowEmptyValues: []string{"name", "health_check"},
	},
	"routers": basicGCPResource{
		terraformName: "google_compute_router",
	},
	"routes": basicGCPResource{
		terraformName: "google_compute_route",
		ignoreKeys:    []string{"google_compute_route", "next_hop_network"},
	},
	"securityPolicies": basicGCPResource{
		terraformName: "google_compute_security_policy",
	},
	/*"snapshots": {
		terraformName: "google_compute_snapshot",
		ignoreKeys: []string{
			"snapshot_encryption_key_sha256",
			"source_disk_encryption_key_sha256",
			"source_disk_link",
		},
	},*/
	/*"sslCertificates": {
		terraformName:       "google_compute_ssl_certificate",
		ignoreKeys: []string{"certificate_id"},
	},*/
	"sslPolicies": basicGCPResource{
		terraformName: "google_compute_ssl_policy",
		ignoreKeys: []string{
			"enabled_features",
		},
	},
	"subnetworks": basicGCPResource{
		terraformName: "google_compute_subnetwork",
		ignoreKeys: []string{
			"gateway_address",
		},
	},
	"targetHttpProxies": basicGCPResource{
		terraformName: "google_compute_target_http_proxy",
		ignoreKeys:    []string{"proxy_id"},
	},
	"targetHttpsProxies": basicGCPResource{
		terraformName: "google_compute_target_https_proxy",
		ignoreKeys:    []string{"proxy_id"},
	},
	"targetSslProxies": basicGCPResource{
		terraformName: "google_compute_target_ssl_proxy",
		ignoreKeys:    []string{"proxy_id"},
	},
	"targetTcpProxies": basicGCPResource{
		terraformName: "google_compute_target_tcp_proxy",
		ignoreKeys:    []string{"proxy_id"},
	},
	"urlMaps": basicGCPResource{
		terraformName: "google_compute_url_map",
		ignoreKeys:    []string{"map_id"},
	},
	"vpnTunnels": basicGCPResource{
		terraformName: "google_compute_vpn_tunnel",
		ignoreKeys: []string{
			"shared_secret_hash",
			"detailed_status",
		},
	},
}