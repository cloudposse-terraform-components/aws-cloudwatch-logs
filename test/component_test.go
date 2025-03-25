package test

import (
	"strings"
	"testing"

	"github.com/cloudposse/test-helpers/pkg/atmos"
	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	"github.com/stretchr/testify/assert"
)

type ComponentSuite struct {
	helper.TestSuite
}

func (s *ComponentSuite) TestBasic() {
	const component = "cloudwatch-logs/basic"
	const stack = "default-test"
	const awsRegion = "us-east-2"

	defer s.DestroyAtmosComponent(s.T(), component, stack, nil)
	options, _ := s.DeployAtmosComponent(s.T(), component, stack, nil)
	assert.NotNil(s.T(), options)

	logGroupArn := atmos.Output(s.T(), options, "log_group_arn")
	assert.True(s.T(), strings.HasPrefix(logGroupArn, "arn:aws:logs:us-east-2"))

	logGroupName := atmos.Output(s.T(), options, "log_group_name")
	assert.True(s.T(), strings.HasPrefix(logGroupName, "eg-default-ue2-test-cloudwatch-logs-"))

	roleArn := atmos.Output(s.T(), options, "role_arn")
	assert.True(s.T(), strings.HasPrefix(roleArn, "arn:aws:iam::"))

	roleName := atmos.Output(s.T(), options, "role_name")
	assert.True(s.T(), strings.HasPrefix(roleName, "eg-default-ue2-test-cloudwatch-logs-"))

	streamArns := atmos.OutputList(s.T(), options, "stream_arns")
	assert.Len(s.T(), streamArns, 2)

	s.DriftTest(component, stack, nil)
}

func (s *ComponentSuite) TestEnabledFlag() {
	const component = "cloudwatch-logs/disabled"
	const stack = "default-test"
	s.VerifyEnabledFlag(component, stack, nil)
}

func TestRunSuite(t *testing.T) {
	suite := new(ComponentSuite)
	helper.Run(t, suite)
}
