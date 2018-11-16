// Code generated by 'create-test-data.sh' on Mon Oct 22 22:22:44 PDT 2018 DO NOT EDIT.

package testdata

func getRawK8sPodTestData() []string {
	return []string{

		`{
			"container": [
				{
					"name": "contiv-crd"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "k8s-app",
					"value": "contiv-crd"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				},
				{
					"key": "controller-revision-hash",
					"value": "4259495223"
				}
			],
			"name": "contiv-crd-ghz49",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-etcd"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "contiv-etcd-695db96cd9"
				},
				{
					"key": "k8s-app",
					"value": "contiv-etcd"
				},
				{
					"key": "statefulset.kubernetes.io/pod-name",
					"value": "contiv-etcd-0"
				}
			],
			"name": "contiv-etcd-0",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-ksr"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "2617825213"
				},
				{
					"key": "k8s-app",
					"value": "contiv-ksr"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "contiv-ksr-8x5ch",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-vswitch",
					"port": [
						{
							"container_port": 9999,
							"host_port": 9999
						}
					]
				}
			],
			"host_ip_address": "10.20.0.11",
			"ip_address": "10.20.0.11",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "4021769348"
				},
				{
					"key": "k8s-app",
					"value": "contiv-vswitch"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "contiv-vswitch-hf4fk",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-vswitch",
					"port": [
						{
							"container_port": 9999,
							"host_port": 9999
						}
					]
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "4021769348"
				},
				{
					"key": "k8s-app",
					"value": "contiv-vswitch"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "contiv-vswitch-jrctj",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "contiv-vswitch",
					"port": [
						{
							"container_port": 9999,
							"host_port": 9999
						}
					]
				}
			],
			"host_ip_address": "10.20.0.10",
			"ip_address": "10.20.0.10",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "4021769348"
				},
				{
					"key": "k8s-app",
					"value": "contiv-vswitch"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "contiv-vswitch-lfnrz",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "coredns",
					"port": [
						{
							"container_port": 53,
							"name": "dns",
							"protocol": 1
						},
						{
							"container_port": 53,
							"name": "dns-tcp"
						},
						{
							"container_port": 9153,
							"name": "metrics"
						}
					]
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.1.1.2",
			"label": [
				{
					"key": "k8s-app",
					"value": "kube-dns"
				},
				{
					"key": "pod-template-hash",
					"value": "3497892450"
				}
			],
			"name": "coredns-78fcdf6894-5xwtl",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "coredns",
					"port": [
						{
							"container_port": 53,
							"name": "dns",
							"protocol": 1
						},
						{
							"container_port": 53,
							"name": "dns-tcp"
						},
						{
							"container_port": 9153,
							"name": "metrics"
						}
					]
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.1.1.3",
			"label": [
				{
					"key": "k8s-app",
					"value": "kube-dns"
				},
				{
					"key": "pod-template-hash",
					"value": "3497892450"
				}
			],
			"name": "coredns-78fcdf6894-qrhfw",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "etcd"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "etcd"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "etcd-k8s-master",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-apiserver"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "kube-apiserver"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "kube-apiserver-k8s-master",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-controller-manager"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "tier",
					"value": "control-plane"
				},
				{
					"key": "component",
					"value": "kube-controller-manager"
				}
			],
			"name": "kube-controller-manager-k8s-master",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-proxy"
				}
			],
			"host_ip_address": "10.20.0.11",
			"ip_address": "10.20.0.11",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "458499861"
				},
				{
					"key": "k8s-app",
					"value": "kube-proxy"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "kube-proxy-642xm",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-proxy"
				}
			],
			"host_ip_address": "10.20.0.10",
			"ip_address": "10.20.0.10",
			"label": [
				{
					"key": "controller-revision-hash",
					"value": "458499861"
				},
				{
					"key": "k8s-app",
					"value": "kube-proxy"
				},
				{
					"key": "pod-template-generation",
					"value": "1"
				}
			],
			"name": "kube-proxy-b2npj",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-proxy"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "pod-template-generation",
					"value": "1"
				},
				{
					"key": "controller-revision-hash",
					"value": "458499861"
				},
				{
					"key": "k8s-app",
					"value": "kube-proxy"
				}
			],
			"name": "kube-proxy-kqlr8",
			"namespace": "kube-system"
		}`,
		`{
			"container": [
				{
					"name": "kube-scheduler"
				}
			],
			"host_ip_address": "10.20.0.2",
			"ip_address": "10.20.0.2",
			"label": [
				{
					"key": "component",
					"value": "kube-scheduler"
				},
				{
					"key": "tier",
					"value": "control-plane"
				}
			],
			"name": "kube-scheduler-k8s-master",
			"namespace": "kube-system"
		}`,
	}
}