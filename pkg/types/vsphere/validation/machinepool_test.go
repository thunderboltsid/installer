package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/types/vsphere"
)

func TestValidateMachinePool(t *testing.T) {
	cases := []struct {
		name           string
		pool           *vsphere.MachinePool
		expectedErrMsg string
	}{
		{
			name:           "empty",
			pool:           &vsphere.MachinePool{},
			expectedErrMsg: "",
		}, {
			name: "negative disk size",
			pool: &vsphere.MachinePool{
				OSDisk: vsphere.OSDisk{
					DiskSizeGB: -1,
				},
			},
			expectedErrMsg: `^test-path\.diskSizeGB: Invalid value: -1: storage disk size must be positive$`,
		}, {
			name: "negative CPUs",
			pool: &vsphere.MachinePool{
				NumCPUs: -1,
			},
			expectedErrMsg: `^test-path\.cpus: Invalid value: -1: number of CPUs must be positive$`,
		}, {
			name: "negative cores",
			pool: &vsphere.MachinePool{
				NumCoresPerSocket: -1,
			},
			expectedErrMsg: `^test-path\.coresPerSocket: Invalid value: -1: cores per socket must be positive$`,
		}, {
			name: "negative memory",
			pool: &vsphere.MachinePool{
				MemoryMiB: -1,
			},
			expectedErrMsg: `^test-path\.memoryMB: Invalid value: -1: memory size must be positive$`,
		}, {
			name: "less CPUs than cores per socket",
			pool: &vsphere.MachinePool{
				NumCPUs:           1,
				NumCoresPerSocket: 8,
			},
			expectedErrMsg: `^test-path\.coresPerSocket: Invalid value: 8: cores per socket must be less than number of CPUs$`,
		},
		{
			name: "numCPUs not a multiple of cores per socket",
			pool: &vsphere.MachinePool{
				NumCPUs:           7,
				NumCoresPerSocket: 4,
			},
			expectedErrMsg: `^test-path.cpus: Invalid value: 7: numCPUs specified should be a multiple of cores per socket$`,
		},
		{
			name: "numCPUs not a multiple of default cores per socket",
			pool: &vsphere.MachinePool{
				NumCPUs: 7,
			},
			expectedErrMsg: `^test-path.cpus: Invalid value: 7: numCPUs specified should be a multiple of cores per socket which is by default 4$`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateMachinePool(tc.pool, field.NewPath("test-path")).ToAggregate()
			if tc.expectedErrMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Regexp(t, tc.expectedErrMsg, err)
			}
		})
	}
}
