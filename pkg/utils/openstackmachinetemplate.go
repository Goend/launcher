package utils

import (
	"context"
	ecnsv1 "easystack.com/plan/api/v1"
	"fmt"
	clusteropenstack "github.com/easystack/cluster-api-provider-openstack/api/v1alpha6"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// TemplateClonedFromNameAnnotation is the infrastructure machine annotation that stores the name of the infrastructure template resource
	// that was cloned for the machine. This annotation is set only during cloning a template. Older/adopted machines will not have this annotation.
	TemplateClonedFromNameAnnotation = "cluster.x-k8s.io/cloned-from-name"
	// DeleteMachineAnnotation marks control plane and worker nodes that will be given priority for deletion
	// when KCP or a machineset scales down. This annotation is given top priority on all delete policies.
	DeleteMachineAnnotation = "cluster.x-k8s.io/delete-machine"
)

type AdoptInfra struct {
	Diff              int64
	MachineHasDeleted int64
}

// GetAdoptInfra TODO return need to up or down replicas and infra uid name.key is uid of infra,value is replicas,include negative number
func GetAdoptInfra(ctx context.Context, cli client.Client, target *ecnsv1.MachineSetReconcile, plan *ecnsv1.Plan) (map[string]AdoptInfra, error) {
	var result = make(map[string]AdoptInfra)
	for _, in := range target.Infra {
		var om clusteropenstack.OpenStackMachineList
		// filter openstackMachine by
		labels := map[string]string{ecnsv1.MachineSetClusterLabelName: plan.Spec.ClusterName}
		err := cli.List(ctx, &om, client.InNamespace(plan.Namespace), client.MatchingLabels(labels))
		if err != nil {
			return nil, err
		}
		var hasCheck int64 = 0
		var machineHasDeleted int64 = 0
		for _, o := range om.Items {
			if o.Annotations[TemplateClonedFromNameAnnotation] == fmt.Sprintf("%s%s%s", plan.Spec.ClusterName, target.Role, in.UID) {
				hasCheck++
				if o.Annotations[DeleteMachineAnnotation] == "true" {
					machineHasDeleted++
				}
			}
		}
		data := int64(in.Replica) - hasCheck
		//skip add map if data is 0
		if data == 0 {
			continue
		}
		result[in.UID] = AdoptInfra{
			Diff:              data,
			MachineHasDeleted: machineHasDeleted,
		}
	}
	return result, nil
}
