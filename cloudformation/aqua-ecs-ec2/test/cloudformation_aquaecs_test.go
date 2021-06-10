package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ecs"
	. "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"testing"
)

func TestBasicCloudFormationTemplate(t *testing.T) {
	t.Parallel()

	uniqId := random.UniqueId()
	ecsClusterName := fmt.Sprintf("ecs-%s", uniqId)
	region := "eu-north-1"
	fmt.Sprintf("Going to create a ECS cluster, name: %s in default VPC", ecsClusterName)
	println()


	cluster, err := CreateEcsClusterE(t, region, ecsClusterName)


	if err == nil{
		//defer DeleteEcsClusterE(t, region, cluster)
		println(cluster)
	}
	ecs.CreateCapacityProviderInput{}cluster


	ecs.LaunchTypeEc2
	////define the cloudformation template file, stackName, region and parameters
	//
	//cfOptions := &m_cf.Options{
	//	CFFile:    "../aquaEcs.yaml",
	//	StackName: fmt.Sprintf("aquaEcs-%s", uniqId),
	//	AWSRegion: region,
	//	Parameters: []*cloudformation.Parameter{
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("ECSClusterName"),
	//			ParameterValue: aws.String(ecsClusterName),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("VpcId"),
	//			ParameterValue: aws.String("vpc-4dae4c24"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("VpcCidr"),
	//			ParameterValue: aws.String("172.31.0.0/16"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("EcsInstanceSubnets"),
	//			ParameterValue: aws.String("subnet-f91fff90,subnet-f4ffd2be,subnet-9ab5b6e2"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("LbSubnets"),
	//			ParameterValue: aws.String("subnet-f91fff90,subnet-f4ffd2be,subnet-9ab5b6e2"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("LBScheme"),
	//			ParameterValue: aws.String("internet-facing"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("SSLCert"),
	//			ParameterValue: aws.String("arn:aws:acm:eu-north-1:934027998561:certificate/eaad19b1-03d2-4198-91ec-1cd01c18dc46"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("AquaServerImage"),
	//			ParameterValue: aws.String("934027998561.dkr.ecr.eu-north-1.amazonaws.com/console:6.2.RC1"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("AquaGatewayImage"),
	//			ParameterValue: aws.String("934027998561.dkr.ecr.eu-north-1.amazonaws.com/gateway:6.2.RC1"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("AquaEnforcerImage"),
	//			ParameterValue: aws.String("934027998561.dkr.ecr.eu-north-1.amazonaws.com/enforcer:6.2.RC1"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("BatchInstallToken"),
	//			ParameterValue: aws.String("6589db6a-1ee5-43d1-a06a-14a6abc38c2b"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("ActiveActive"),
	//			ParameterValue: aws.String("false"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("RdsInstanceClass"),
	//			ParameterValue: aws.String("db.t3.medium"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("RdsStorage"),
	//			ParameterValue: aws.String("50"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("MultiAzDatabase"),
	//			ParameterValue: aws.String("false"),
	//		},
	//		&cloudformation.Parameter{
	//			ParameterKey:   aws.String("EcsSecurityGroupId"),
	//			ParameterValue: aws.String("false"),
	//		},
	//	},
	//	Tag: []*cloudformation.Tag{
	//		&cloudformation.Tag{
	//			Key: aws.String("Deploy"),
	//			Value: aws.String("Deploy from aqua-aws tests"),
	//		},
	//	},
	//}
	//
	//defer m_cf.DeleteStack(t, cfOptions)
	//
	//// create the cloudformation template
	//
	//m_cf.CreateStack(t, cfOptions)
	//
	//
	//outputList := m_cf.ListOutputs(t, cfOptions)
	//aquaConsole := m_cf.FilterOutputs(outputList, "AquaConsole")
	//
	//fmt.Sprintf("AquaConSole URL: %s", aquaConsole)
	//assert.NotEqual(t, aquaConsole, nil)

}
