/*
Copyright 2020 The Crossplane Authors.

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

package main

// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

// +kubebuilder:rbac:groups=acm.aws.crossplane.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acm.aws.crossplane.io,resources=certificates/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=acm.aws.crossplane.io,resources=certificates/finalizers,verbs=update

// +kubebuilder:rbac:groups=acmpca.aws.crossplane.io,resources=certificateauthorities,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acmpca.aws.crossplane.io,resources=certificateauthorities/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=acmpca.aws.crossplane.io,resources=certificateauthorities/finalizers,verbs=update

// +kubebuilder:rbac:groups=acmpca.aws.crossplane.io,resources=certificateauthoritypermissions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acmpca.aws.crossplane.io,resources=certificateauthoritypermissions/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=acmpca.aws.crossplane.io,resources=certificateauthoritypermissions/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=apimappings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=apimappings/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=apimappings/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=apis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=apis/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=apis/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=authorizers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=authorizers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=authorizers/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=deployments/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=domainnames,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=domainnames/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=domainnames/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=integrationresponses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=integrationresponses/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=integrationresponses/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=integrations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=integrations/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=integrations/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=models,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=models/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=models/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=routeresponses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=routeresponses/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=routeresponses/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=routes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=routes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=routes/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=stages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=stages/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=stages/finalizers,verbs=update

// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=vpclinks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=vpclinks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apigatewayv2.aws.crossplane.io,resources=vpclinks/finalizers,verbs=update

// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providerconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providerconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providerconfigs/finalizers,verbs=update

// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providerconfigusages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providerconfigusages/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providerconfigusages/finalizers,verbs=update

// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=aws.crossplane.io,resources=providers/finalizers,verbs=update

// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=cacheclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=cacheclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=cacheclusters/finalizers,verbs=update

// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=cachesubnetgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=cachesubnetgroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=cachesubnetgroups/finalizers,verbs=update

// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=replicationgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=replicationgroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cache.aws.crossplane.io,resources=replicationgroups/finalizers,verbs=update

// +kubebuilder:rbac:groups=database.aws.crossplane.io,resources=dbsubnetgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=database.aws.crossplane.io,resources=dbsubnetgroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=database.aws.crossplane.io,resources=dbsubnetgroups/finalizers,verbs=update

// +kubebuilder:rbac:groups=database.aws.crossplane.io,resources=rdsinstances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=database.aws.crossplane.io,resources=rdsinstances/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=database.aws.crossplane.io,resources=rdsinstances/finalizers,verbs=update

// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=backups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=backups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=backups/finalizers,verbs=update

// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=globaltables,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=globaltables/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=globaltables/finalizers,verbs=update

// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=tables,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=tables/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=dynamodb.aws.crossplane.io,resources=tables/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=elasticips,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=elasticips/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=elasticips/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=internetgateways,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=internetgateways/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=internetgateways/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=natgateways,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=natgateways/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=natgateways/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=routetables,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=routetables/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=routetables/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=securitygroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=securitygroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=securitygroups/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=subnets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=subnets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=subnets/finalizers,verbs=update

// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=vpcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=vpcs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.aws.crossplane.io,resources=vpcs/finalizers,verbs=update

// +kubebuilder:rbac:groups=ecr.aws.crossplane.io,resources=repositories,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ecr.aws.crossplane.io,resources=repositories/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ecr.aws.crossplane.io,resources=repositories/finalizers,verbs=update

// +kubebuilder:rbac:groups=efs.aws.crossplane.io,resources=filesystems,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=efs.aws.crossplane.io,resources=filesystems/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=efs.aws.crossplane.io,resources=filesystems/finalizers,verbs=update

// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=clusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=clusters/finalizers,verbs=update

// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=fargateprofiles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=fargateprofiles/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=fargateprofiles/finalizers,verbs=update

// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=nodegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=nodegroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=eks.aws.crossplane.io,resources=nodegroups/finalizers,verbs=update

// +kubebuilder:rbac:groups=elasticloadbalancing.aws.crossplane.io,resources=elbattachments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=elasticloadbalancing.aws.crossplane.io,resources=elbattachments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=elasticloadbalancing.aws.crossplane.io,resources=elbattachments/finalizers,verbs=update

// +kubebuilder:rbac:groups=elasticloadbalancing.aws.crossplane.io,resources=elbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=elasticloadbalancing.aws.crossplane.io,resources=elbs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=elasticloadbalancing.aws.crossplane.io,resources=elbs/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamaccesskeys,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamaccesskeys/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamaccesskeys/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgrouppolicyattachments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgrouppolicyattachments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgrouppolicyattachments/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgroups/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgroupusermemberships,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgroupusermemberships/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamgroupusermemberships/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iampolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iampolicies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iampolicies/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamrolepolicyattachments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamrolepolicyattachments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamrolepolicyattachments/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamroles/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamroles/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamuserpolicyattachments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamuserpolicyattachments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamuserpolicyattachments/finalizers,verbs=update

// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamusers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=identity.aws.crossplane.io,resources=iamusers/finalizers,verbs=update

// +kubebuilder:rbac:groups=kms.aws.crossplane.io,resources=keys,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kms.aws.crossplane.io,resources=keys/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kms.aws.crossplane.io,resources=keys/finalizers,verbs=update

// +kubebuilder:rbac:groups=notification.aws.crossplane.io,resources=snssubscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=notification.aws.crossplane.io,resources=snssubscriptions/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=notification.aws.crossplane.io,resources=snssubscriptions/finalizers,verbs=update

// +kubebuilder:rbac:groups=notification.aws.crossplane.io,resources=snstopics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=notification.aws.crossplane.io,resources=snstopics/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=notification.aws.crossplane.io,resources=snstopics/finalizers,verbs=update

// +kubebuilder:rbac:groups=rds.aws.crossplane.io,resources=dbclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rds.aws.crossplane.io,resources=dbclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=rds.aws.crossplane.io,resources=dbclusters/finalizers,verbs=update

// +kubebuilder:rbac:groups=redshift.aws.crossplane.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=redshift.aws.crossplane.io,resources=clusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=redshift.aws.crossplane.io,resources=clusters/finalizers,verbs=update

// +kubebuilder:rbac:groups=route53.aws.crossplane.io,resources=hostedzones,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=route53.aws.crossplane.io,resources=hostedzones/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=route53.aws.crossplane.io,resources=hostedzones/finalizers,verbs=update

// +kubebuilder:rbac:groups=route53.aws.crossplane.io,resources=resourcerecordsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=route53.aws.crossplane.io,resources=resourcerecordsets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=route53.aws.crossplane.io,resources=resourcerecordsets/finalizers,verbs=update

// +kubebuilder:rbac:groups=s3.aws.crossplane.io,resources=bucketpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=s3.aws.crossplane.io,resources=bucketpolicies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=s3.aws.crossplane.io,resources=bucketpolicies/finalizers,verbs=update

// +kubebuilder:rbac:groups=s3.aws.crossplane.io,resources=buckets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=s3.aws.crossplane.io,resources=buckets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=s3.aws.crossplane.io,resources=buckets/finalizers,verbs=update

// +kubebuilder:rbac:groups=secretsmanager.aws.crossplane.io,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=secretsmanager.aws.crossplane.io,resources=secrets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=secretsmanager.aws.crossplane.io,resources=secrets/finalizers,verbs=update

// +kubebuilder:rbac:groups=sfn.aws.crossplane.io,resources=activities,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sfn.aws.crossplane.io,resources=activities/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=sfn.aws.crossplane.io,resources=activities/finalizers,verbs=update

// +kubebuilder:rbac:groups=sfn.aws.crossplane.io,resources=statemachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sfn.aws.crossplane.io,resources=statemachines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=sfn.aws.crossplane.io,resources=statemachines/finalizers,verbs=update

// +kubebuilder:rbac:groups=sqs.aws.crossplane.io,resources=queues,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sqs.aws.crossplane.io,resources=queues/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=sqs.aws.crossplane.io,resources=queues/finalizers,verbs=update
