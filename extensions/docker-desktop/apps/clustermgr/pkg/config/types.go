// Copyright 2022 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

func GetDefaultPorts() []string {
	return []string{"80", "443"}
}

func RunningResponse() Response {
	return Response{
		Status:      Running,
		Description: "Cluster is running",
	}
}

func DeletingResponse() Response {
	return Response{
		Status:      Deleting,
		Description: "Cluster is deleting",
	}
}

func DeletedResponse() Response {
	return Response{
		Status:      Deleted,
		Description: "Cluster is deleted",
	}
}

func NotExistsResponse() Response {
	return Response{
		Status:      NotExists,
		Description: "Cluster does not exist",
	}
}

func EmptyStatsResponse() Response {
	return Response{
		Status:      NotExists,
		Description: "",
		Error:       false,
		Stats:       nil,
	}
}

/**
 * Structs
 */
type Status string

const (
	Unknown      Status = "Unknown"
	NotExists    Status = "NotExists"
	Initializing Status = "Initializing"
	Creating     Status = "Creating"
	Running      Status = "Running"
	Stopped      Status = "Stopped"
	Deleting     Status = "Deleting"
	Deleted      Status = "Deleted"
	Error        Status = "Error"
)

type Response struct {
	Status       Status                 `json:"status,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Error        bool                   `json:"isError,omitempty"`
	ErrorMessage string                 `json:"errorMessage,omitempty"`
	Output       string                 `json:"output,omitempty"`
	Stats        *ClusterContainerStats `json:"stats,omitempty"`
}

// type ClusterStats struct {
// 	CPU    string `json:"cpu"`
// 	Memory string `json:"memory"`
// }

type ClusterContainerStats struct {
	ID     string                 `json:"id"`
	Memory ClusterContainerMemory `json:"memory"`
	CPU    ClusterContainerCPU    `json:"cpu"`
}

type ClusterContainerMemory struct {
	Used  float64 `json:"used"`
	Total float64 `json:"total"`
	Usage float64 `json:"usage"`
}

type ClusterContainerCPU struct {
	CPUDelta       float64 `json:"cpu_delta"`
	SystemCPUDelta float64 `json:"system_cpu_delta"`
	NumberCPUs     int     `json:"number_cpus"`
	Usage          float64 `json:"usage"`
}
