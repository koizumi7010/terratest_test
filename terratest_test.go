package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraform(t *testing.T) {
	t.Parallel()

	expectedAmi := "********************"
	expectedInstanceType := "t2.micro"
	expectedVolumeType := "gp2"
	expectedVolumeSize := "20"

	terraformOption := &terraform.Options{
		TerraformDir: "../",
	}

	// ec2-a
	ec2aAmi := terraform.Output(t, terraformOption, "ec2-a_ami")
	assert.Equal(t, expectedAmi, ec2aAmi)
	ec2aInstanceType := terraform.Output(t, terraformOption, "ec2-a_instance_type")
	assert.Equal(t, expectedInstanceType, ec2aInstanceType)
	ec2aPrivateIp := terraform.Output(t, terraformOption, "ec2-a_private_ip")
	assert.Equal(t, "10.1.2.5", ec2aPrivateIp)
	ec2aVolumeType := terraform.Output(t, terraformOption, "ec2-a_volume_type")
	assert.Equal(t, expectedVolumeType, ec2aVolumeType)
	ec2aVolumeSize := terraform.Output(t, terraformOption, "ec2-a_volume_size")
	assert.Equal(t, expectedVolumeSize, ec2aVolumeSize)

	// ec2c
	ec2cAmi := terraform.Output(t, terraformOption, "ec2-c_ami")
	assert.Equal(t, expectedAmi, ec2cAmi)
	ec2cInstanceType := terraform.Output(t, terraformOption, "ec2-c_instance_type")
	assert.Equal(t, expectedInstanceType, ec2cInstanceType)
	ec2cPrivateIp := terraform.Output(t, terraformOption, "ec2-c_private_ip")
	assert.Equal(t, "10.1.12.5", ec2cPrivateIp)
	ec2cVolumeType := terraform.Output(t, terraformOption, "ec2-c_volume_type")
	assert.Equal(t, expectedVolumeType, ec2cVolumeType)
	ec2cVolumeSize := terraform.Output(t, terraformOption, "ec2-c_volume_size")
	assert.Equal(t, expectedVolumeSize, ec2cVolumeSize)
}
