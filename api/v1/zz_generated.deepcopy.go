//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/easystack/cluster-api-provider-openstack/api/v1alpha6"
	"github.com/kubean-io/kubean-api/apis/clusteroperation/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOperationSet) DeepCopyInto(out *ClusterOperationSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOperationSet.
func (in *ClusterOperationSet) DeepCopy() *ClusterOperationSet {
	if in == nil {
		return nil
	}
	out := new(ClusterOperationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterOperationSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOperationSetList) DeepCopyInto(out *ClusterOperationSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterOperationSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOperationSetList.
func (in *ClusterOperationSetList) DeepCopy() *ClusterOperationSetList {
	if in == nil {
		return nil
	}
	out := new(ClusterOperationSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterOperationSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOperationSetSpec) DeepCopyInto(out *ClusterOperationSetSpec) {
	*out = *in
	if in.ClusterOperations != nil {
		in, out := &in.ClusterOperations, &out.ClusterOperations
		*out = make([]ClusterOps, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SideMap != nil {
		in, out := &in.SideMap, &out.SideMap
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOperationSetSpec.
func (in *ClusterOperationSetSpec) DeepCopy() *ClusterOperationSetSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterOperationSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOperationSetStatus) DeepCopyInto(out *ClusterOperationSetStatus) {
	*out = *in
	if in.ClusterOperationStatusList != nil {
		in, out := &in.ClusterOperationStatusList, &out.ClusterOperationStatusList
		*out = make([]ClusterOperationStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOperationSetStatus.
func (in *ClusterOperationSetStatus) DeepCopy() *ClusterOperationSetStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterOperationSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOperationStatus) DeepCopyInto(out *ClusterOperationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOperationStatus.
func (in *ClusterOperationStatus) DeepCopy() *ClusterOperationStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterOperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOps) DeepCopyInto(out *ClusterOps) {
	*out = *in
	if in.PreHook != nil {
		in, out := &in.PreHook, &out.PreHook
		*out = make([]v1alpha1.HookAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostHook != nil {
		in, out := &in.PostHook, &out.PostHook
		*out = make([]v1alpha1.HookAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOps.
func (in *ClusterOps) DeepCopy() *ClusterOps {
	if in == nil {
		return nil
	}
	out := new(ClusterOps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostConf) DeepCopyInto(out *HostConf) {
	*out = *in
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]*AnsibleNode, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AnsibleNode)
				**out = **in
			}
		}
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(AnsibleNode)
		**out = **in
	}
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeMaster != nil {
		in, out := &in.KubeMaster, &out.KubeMaster
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeNode != nil {
		in, out := &in.KubeNode, &out.KubeNode
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeIngress != nil {
		in, out := &in.KubeIngress, &out.KubeIngress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubePrometheus != nil {
		in, out := &in.KubePrometheus, &out.KubePrometheus
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeLog != nil {
		in, out := &in.KubeLog, &out.KubeLog
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostConf.
func (in *HostConf) DeepCopy() *HostConf {
	if in == nil {
		return nil
	}
	out := new(HostConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraMachine) DeepCopyInto(out *InfraMachine) {
	*out = *in
	if in.PortIDs != nil {
		in, out := &in.PortIDs, &out.PortIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraMachine.
func (in *InfraMachine) DeepCopy() *InfraMachine {
	if in == nil {
		return nil
	}
	out := new(InfraMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infras) DeepCopyInto(out *Infras) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = new(Subnet)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]*volume, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(volume)
				**out = **in
			}
		}
	}
	if in.ServerMetadata != nil {
		in, out := &in.ServerMetadata, &out.ServerMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infras.
func (in *Infras) DeepCopy() *Infras {
	if in == nil {
		return nil
	}
	out := new(Infras)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	if in.AllowedCIDRs != nil {
		in, out := &in.AllowedCIDRs, &out.AllowedCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineFailureReason) DeepCopyInto(out *MachineFailureReason) {
	*out = *in
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineFailureReason.
func (in *MachineFailureReason) DeepCopy() *MachineFailureReason {
	if in == nil {
		return nil
	}
	out := new(MachineFailureReason)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetReconcile) DeepCopyInto(out *MachineSetReconcile) {
	*out = *in
	if in.Infra != nil {
		in, out := &in.Infra, &out.Infra
		*out = make([]*Infras, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Infras)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetReconcile.
func (in *MachineSetReconcile) DeepCopy() *MachineSetReconcile {
	if in == nil {
		return nil
	}
	out := new(MachineSetReconcile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetStatus) DeepCopyInto(out *MachineSetStatus) {
	*out = *in
	if in.ReadyMachines != nil {
		in, out := &in.ReadyMachines, &out.ReadyMachines
		*out = make([]MachineStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetStatus.
func (in *MachineSetStatus) DeepCopy() *MachineSetStatus {
	if in == nil {
		return nil
	}
	out := new(MachineSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1alpha6.OpenStackMachineStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorConfig) DeepCopyInto(out *MonitorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorConfig.
func (in *MonitorConfig) DeepCopy() *MonitorConfig {
	if in == nil {
		return nil
	}
	out := new(MonitorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plan) DeepCopyInto(out *Plan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plan.
func (in *Plan) DeepCopy() *Plan {
	if in == nil {
		return nil
	}
	out := new(Plan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanList) DeepCopyInto(out *PlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanList.
func (in *PlanList) DeepCopy() *PlanList {
	if in == nil {
		return nil
	}
	out := new(PlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanSpec) DeepCopyInto(out *PlanSpec) {
	*out = *in
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NeedKeepAlive != nil {
		in, out := &in.NeedKeepAlive, &out.NeedKeepAlive
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NeedLoadBalancer != nil {
		in, out := &in.NeedLoadBalancer, &out.NeedLoadBalancer
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MachineSets != nil {
		in, out := &in.MachineSets, &out.MachineSets
		*out = make([]*MachineSetReconcile, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MachineSetReconcile)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.Monitor = in.Monitor
	if in.OtherAnsibleOpts != nil {
		in, out := &in.OtherAnsibleOpts, &out.OtherAnsibleOpts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.UserInfo = in.UserInfo
	if in.HostConf != nil {
		in, out := &in.HostConf, &out.HostConf
		*out = new(HostConf)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanSpec.
func (in *PlanSpec) DeepCopy() *PlanSpec {
	if in == nil {
		return nil
	}
	out := new(PlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanStatus) DeepCopyInto(out *PlanStatus) {
	*out = *in
	if in.ServerGroupID != nil {
		in, out := &in.ServerGroupID, &out.ServerGroupID
		*out = new(Servergroups)
		**out = **in
	}
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = make(map[PlanType]PlanPhase, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VMFailureReason != nil {
		in, out := &in.VMFailureReason, &out.VMFailureReason
		*out = make(map[string]MachineFailureReason, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.InfraMachine != nil {
		in, out := &in.InfraMachine, &out.InfraMachine
		*out = make(map[string]InfraMachine, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PlanLoadBalancer != nil {
		in, out := &in.PlanLoadBalancer, &out.PlanLoadBalancer
		*out = make([]*LoadBalancer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LoadBalancer)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(v1alpha6.Instance)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanStatus.
func (in *PlanStatus) DeepCopy() *PlanStatus {
	if in == nil {
		return nil
	}
	out := new(PlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servergroups) DeepCopyInto(out *Servergroups) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servergroups.
func (in *Servergroups) DeepCopy() *Servergroups {
	if in == nil {
		return nil
	}
	out := new(Servergroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}
