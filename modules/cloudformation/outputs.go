package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	aws_cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

// ListOutputs return a list of Outputs of
func ListOutputs(t *testing.T, CFOptions *Options) []*aws_cf.Output {
	svc := NewCFClient(t, CFOptions.AWSRegion)
	resOutput, err := svc.DescribeStacks(&aws_cf.DescribeStacksInput{StackName: &CFOptions.StackName})
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return resOutput.Stacks[0].Outputs
}

// ListExports returns a list of Exports of a specified stack
func ListExports(t *testing.T, CFOptions *Options) []*aws_cf.Export {
	svc := NewCFClient(t, CFOptions.AWSRegion)

	resInput := &aws_cf.ListExportsInput{
		NextToken: aws.String("")}
	resOutput, err := svc.ListExports(resInput)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return resOutput.Exports
}

func FilterOutputs(list []*aws_cf.Output, CmpID string) string {
	res := ""
	for _, s := range list {
		if *s.OutputKey == CmpID {
			res = *s.OutputValue
		}
	}
	return res
}

func FilterExports(list []*aws_cf.Output, CmpID string) string {
	res := ""
	for _, s := range list {
		if *s.OutputKey == CmpID {
			res = *s.OutputValue
		}
	}
	return res
}
