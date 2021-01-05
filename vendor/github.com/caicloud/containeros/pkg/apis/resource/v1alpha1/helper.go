package v1alpha1

// GetMachineCluster return the machine's cluster name
func GetMachineCluster(machine *Machine) string {
	if machine.Spec.ClusterName != "" {
		return machine.Spec.ClusterName
	}

	return machine.Status.ClusterName
}
