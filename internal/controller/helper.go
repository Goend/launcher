// Copyright 2022 EasyStack, Inc.
package controller

const (
	NETATTDEFFINALIZERNAME = "plan.finalizers.eks.io"

	AnsiblePlanStartEvent = "AnsiblePlanStart"
	AnsiblePlanCreatedEvent = "AnsiblePlanCreated"
	AnsiblePlanDeleteEvent = "AnsiblePlanDeleted"
	AnsiblePlanDeleteSshKeyEvent = "AnsiblePlanDeleteSshKey"

	PlanStartEvent = "PlanStart"
	PlanCreatedEvent = "PlanCreated"
	PlanDeleteEvent = "PlanDeleted"
	PlanDeleteSshKeyEvent = "AnsiblePlanDeleteSshKey"

)


func StringInArray(val string, array []string) bool {
	for i := range array {
		if array[i] == val {
			return true
		}
	}
	return false
}


func RemoveString(s string, slice []string) (result []string, found bool) {
	if len(slice) != 0 {
		for _, item := range slice {
			if item == s {
				found = true
				continue
			}
			result = append(result, item)
		}
	}
	return
}
