/*
Package aws implements a steampipe plugin for aws.

This plugin provides data that Steampipe uses to present foreign
tables that represent Amazon AWS resources.
*/
package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-aws"

// Plugin creates this (aws) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{
					"NoSuchEntity",
					"NotFoundException",
					"ResourceNotFoundException",
					"InvalidParameter",
					"InvalidParameterValue",
					"InvalidParameterException",
					"InvalidParameterValueException",
					"ValidationError",
					"ValidationException",
				}),
			},
		},
		// Default ignore config for the plugin
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrorPluginDefault(),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"aws_accessanalyzer_analyzer":                                  tableAwsAccessAnalyzer(ctx),
			"aws_account":                                                  tableAwsAccount(ctx),
			"aws_account_alternate_contact":                                tableAwsAccountAlternateContact(ctx),
			"aws_account_contact":                                          tableAwsAccountContact(ctx),
			"aws_acm_certificate":                                          tableAwsAcmCertificate(ctx),
			"aws_amplify_app":                                              tableAwsAmplifyApp(ctx),
			"aws_api_gateway_api_key":                                      tableAwsAPIGatewayAPIKey(ctx),
			"aws_api_gateway_authorizer":                                   tableAwsAPIGatewayAuthorizer(ctx),
			"aws_api_gateway_rest_api":                                     tableAwsAPIGatewayRestAPI(ctx),
			"aws_api_gateway_stage":                                        tableAwsAPIGatewayStage(ctx),
			"aws_api_gateway_usage_plan":                                   tableAwsAPIGatewayUsagePlan(ctx),
			"aws_api_gatewayv2_api":                                        tableAwsAPIGatewayV2Api(ctx),
			"aws_api_gatewayv2_domain_name":                                tableAwsAPIGatewayV2DomainName(ctx),
			"aws_api_gatewayv2_integration":                                tableAwsAPIGatewayV2Integration(ctx),
			"aws_api_gatewayv2_stage":                                      tableAwsAPIGatewayV2Stage(ctx),
			"aws_appautoscaling_target":                                    tableAwsAppAutoScalingTarget(ctx),
			"aws_appconfig_application":                                    tableAwsAppConfigApplication(ctx),
			"aws_auditmanager_assessment":                                  tableAwsAuditManagerAssessment(ctx),
			"aws_auditmanager_control":                                     tableAwsAuditManagerControl(ctx),
			"aws_auditmanager_evidence":                                    tableAwsAuditManagerEvidence(ctx),
			"aws_auditmanager_evidence_folder":                             tableAwsAuditManagerEvidenceFolder(ctx),
			"aws_auditmanager_framework":                                   tableAwsAuditManagerFramework(ctx),
			"aws_availability_zone":                                        tableAwsAvailabilityZone(ctx),
			"aws_backup_framework":                                         tableAwsBackupFramework(ctx),
			"aws_backup_plan":                                              tableAwsBackupPlan(ctx),
			"aws_backup_protected_resource":                                tableAwsBackupProtectedResource(ctx),
			"aws_backup_recovery_point":                                    tableAwsBackupRecoveryPoint(ctx),
			"aws_backup_selection":                                         tableAwsBackupSelection(ctx),
			"aws_backup_vault":                                             tableAwsBackupVault(ctx),
			"aws_cloudcontrol_resource":                                    tableAwsCloudControlResource(ctx),
			"aws_cloudformation_stack":                                     tableAwsCloudFormationStack(ctx),
			"aws_cloudfront_cache_policy":                                  tableAwsCloudFrontCachePolicy(ctx),
			"aws_cloudfront_distribution":                                  tableAwsCloudFrontDistribution(ctx),
			"aws_cloudfront_function":                                      tableAwsCloudFrontFunction(ctx),
			"aws_cloudfront_origin_access_identity":                        tableAwsCloudFrontOriginAccessIdentity(ctx),
			"aws_cloudfront_origin_request_policy":                         tableAwsCloudFrontOriginRequestPolicy(ctx),
			"aws_cloudfront_response_headers_policy":                       tableAwsCloudFrontResponseHeadersPolicy(ctx),
			"aws_cloudtrail_trail":                                         tableAwsCloudtrailTrail(ctx),
			"aws_cloudtrail_trail_event":                                   tableAwsCloudtrailTrailEvent(ctx),
			"aws_cloudwatch_alarm":                                         tableAwsCloudWatchAlarm(ctx),
			"aws_cloudwatch_log_event":                                     tableAwsCloudwatchLogEvent(ctx),
			"aws_cloudwatch_log_group":                                     tableAwsCloudwatchLogGroup(ctx),
			"aws_cloudwatch_log_metric_filter":                             tableAwsCloudwatchLogMetricFilter(ctx),
			"aws_cloudwatch_log_resource_policy":                           tableAwsCloudwatchLogResourcePolicy(ctx),
			"aws_cloudwatch_log_stream":                                    tableAwsCloudwatchLogStream(ctx),
			"aws_cloudwatch_log_subscription_filter":                       tableAwsCloudwatchLogSubscriptionFilter(ctx),
			"aws_cloudwatch_metric":                                        tableAwsCloudWatchMetric(ctx),
			"aws_codeartifact_domain":                                      tableAwsCodeArtifactDomain(ctx),
			"aws_codeartifact_repository":                                  tableAwsCodeArtifactRepository(ctx),
			"aws_codebuild_project":                                        tableAwsCodeBuildProject(ctx),
			"aws_codebuild_source_credential":                              tableAwsCodeBuildSourceCredential(ctx),
			"aws_codecommit_repository":                                    tableAwsCodeCommitRepository(ctx),
			"aws_codedeploy_app":                                           tableAwsCodeDeployApplication(ctx),
			"aws_codepipeline_pipeline":                                    tableAwsCodepipelinePipeline(ctx),
			"aws_config_aggregate_authorization":                           tableAwsConfigAggregateAuthorization(ctx),
			"aws_config_configuration_recorder":                            tableAwsConfigConfigurationRecorder(ctx),
			"aws_config_conformance_pack":                                  tableAwsConfigConformancePack(ctx),
			"aws_config_rule":                                              tableAwsConfigRule(ctx),
			"aws_cost_by_account_daily":                                    tableAwsCostByLinkedAccountDaily(ctx),
			"aws_cost_by_account_monthly":                                  tableAwsCostByLinkedAccountMonthly(ctx),
			"aws_cost_by_record_type_daily":                                tableAwsCostByRecordTypeDaily(ctx),
			"aws_cost_by_record_type_monthly":                              tableAwsCostByRecordTypeMonthly(ctx),
			"aws_cost_by_service_daily":                                    tableAwsCostByServiceDaily(ctx),
			"aws_cost_by_service_monthly":                                  tableAwsCostByServiceMonthly(ctx),
			"aws_cost_by_service_usage_type_daily":                         tableAwsCostByServiceUsageTypeDaily(ctx),
			"aws_cost_by_service_usage_type_monthly":                       tableAwsCostByServiceUsageTypeMonthly(ctx),
			"aws_cost_forecast_daily":                                      tableAwsCostForecastDaily(ctx),
			"aws_cost_forecast_monthly":                                    tableAwsCostForecastMonthly(ctx),
			"aws_cost_usage":                                               tableAwsCostAndUsage(ctx),
			"aws_dax_cluster":                                              tableAwsDaxCluster(ctx),
			"aws_dax_subnet_group":                                         tableAwsDaxSubnetGroup(ctx),
			"aws_directory_service_directory":                              tableAwsDirectoryServiceDirectory(ctx),
			"aws_dlm_lifecycle_policy":                                     tableAwsDLMLifecyclePolicy(ctx),
			"aws_dms_replication_instance":                                 tableAwsDmsReplicationInstance(ctx),
			"aws_docdb_cluster":                                            tableAwsDocDBCluster(ctx),
			"aws_dynamodb_backup":                                          tableAwsDynamoDBBackup(ctx),
			"aws_dynamodb_global_table":                                    tableAwsDynamoDBGlobalTable(ctx),
			"aws_dynamodb_metric_account_provisioned_read_capacity_util":   tableAwsDynamoDBMetricAccountProvisionedReadCapacityUtilization(ctx),
			"aws_dynamodb_metric_account_provisioned_write_capacity_util":  tableAwsDynamoDBMetricAccountProvisionedWriteCapacityUtilization(ctx),
			"aws_dynamodb_table":                                           tableAwsDynamoDBTable(ctx),
			"aws_dynamodb_table_export":                                    tableAwsDynamoDBTableExport(ctx),
			"aws_ebs_snapshot":                                             tableAwsEBSSnapshot(ctx),
			"aws_ebs_volume":                                               tableAwsEBSVolume(ctx),
			"aws_ebs_volume_metric_read_ops":                               tableAwsEbsVolumeMetricReadOps(ctx),
			"aws_ebs_volume_metric_read_ops_daily":                         tableAwsEbsVolumeMetricReadOpsDaily(ctx),
			"aws_ebs_volume_metric_read_ops_hourly":                        tableAwsEbsVolumeMetricReadOpsHourly(ctx),
			"aws_ebs_volume_metric_write_ops":                              tableAwsEbsVolumeMetricWriteOps(ctx),
			"aws_ebs_volume_metric_write_ops_daily":                        tableAwsEbsVolumeMetricWriteOpsDaily(ctx),
			"aws_ebs_volume_metric_write_ops_hourly":                       tableAwsEbsVolumeMetricWriteOpsHourly(ctx),
			"aws_ec2_ami":                                                  tableAwsEc2Ami(ctx),
			"aws_ec2_ami_shared":                                           tableAwsEc2AmiShared(ctx),
			"aws_ec2_application_load_balancer":                            tableAwsEc2ApplicationLoadBalancer(ctx),
			"aws_ec2_application_load_balancer_metric_request_count":       tableAwsEc2ApplicationLoadBalancerMetricRequestCount(ctx),
			"aws_ec2_application_load_balancer_metric_request_count_daily": tableAwsEc2ApplicationLoadBalancerMetricRequestCountDaily(ctx),
			"aws_ec2_autoscaling_group":                                    tableAwsEc2ASG(ctx),
			"aws_ec2_capacity_reservation":                                 tableAwsEc2CapacityReservation(ctx),
			"aws_ec2_classic_load_balancer":                                tableAwsEc2ClassicLoadBalancer(ctx),
			"aws_ec2_gateway_load_balancer":                                tableAwsEc2GatewayLoadBalancer(ctx),
			"aws_ec2_instance":                                             tableAwsEc2Instance(ctx),
			"aws_ec2_instance_availability":                                tableAwsInstanceAvailability(ctx),
			"aws_ec2_instance_metric_cpu_utilization":                      tableAwsEc2InstanceMetricCpuUtilization(ctx),
			"aws_ec2_instance_metric_cpu_utilization_daily":                tableAwsEc2InstanceMetricCpuUtilizationDaily(ctx),
			"aws_ec2_instance_metric_cpu_utilization_hourly":               tableAwsEc2InstanceMetricCpuUtilizationHourly(ctx),
			"aws_ec2_instance_type":                                        tableAwsInstanceType(ctx),
			"aws_ec2_key_pair":                                             tableAwsEc2KeyPair(ctx),
			"aws_ec2_launch_configuration":                                 tableAwsEc2LaunchConfiguration(ctx),
			"aws_ec2_load_balancer_listener":                               tableAwsEc2ApplicationLoadBalancerListener(ctx),
			"aws_ec2_managed_prefix_list":                                  tableAwsEc2ManagedPrefixList(ctx),
			"aws_ec2_network_interface":                                    tableAwsEc2NetworkInterface(ctx),
			"aws_ec2_network_load_balancer":                                tableAwsEc2NetworkLoadBalancer(ctx),
			"aws_ec2_network_load_balancer_metric_net_flow_count":          tableAwsEc2NetworkLoadBalancerMetricNetFlowCount(ctx),
			"aws_ec2_network_load_balancer_metric_net_flow_count_daily":    tableAwsEc2NetworkLoadBalancerMetricNetFlowCountDaily(ctx),
			"aws_ec2_regional_settings":                                    tableAwsEc2RegionalSettings(ctx),
			"aws_ec2_reserved_instance":                                    tableAwsEc2ReservedInstance(ctx),
			"aws_ec2_spot_price":                                           tableAwsEc2SpotPrice(ctx),
			"aws_ec2_ssl_policy":                                           tableAwsEc2SslPolicy(ctx),
			"aws_ec2_target_group":                                         tableAwsEc2TargetGroup(ctx),
			"aws_ec2_transit_gateway":                                      tableAwsEc2TransitGateway(ctx),
			"aws_ec2_transit_gateway_route":                                tableAwsEc2TransitGatewayRoute(ctx),
			"aws_ec2_transit_gateway_route_table":                          tableAwsEc2TransitGatewayRouteTable(ctx),
			"aws_ec2_transit_gateway_vpc_attachment":                       tableAwsEc2TransitGatewayVpcAttachment(ctx),
			"aws_ecr_image":                                                tableAwsEcrImage(ctx),
			"aws_ecr_image_scan_finding":                                   tableAwsEcrImageScanFinding(ctx),
			"aws_ecr_repository":                                           tableAwsEcrRepository(ctx),
			"aws_ecrpublic_repository":                                     tableAwsEcrpublicRepository(ctx),
			"aws_ecs_cluster":                                              tableAwsEcsCluster(ctx),
			"aws_ecs_cluster_metric_cpu_utilization":                       tableAwsEcsClusterMetricCpuUtilization(ctx),
			"aws_ecs_cluster_metric_cpu_utilization_daily":                 tableAwsEcsClusterMetricCpuUtilizationDaily(ctx),
			"aws_ecs_cluster_metric_cpu_utilization_hourly":                tableAwsEcsClusterMetricCpuUtilizationHourly(ctx),
			"aws_ecs_container_instance":                                   tableAwsEcsContainerInstance(ctx),
			"aws_ecs_service":                                              tableAwsEcsService(ctx),
			"aws_ecs_task":                                                 tableAwsEcsTask(ctx),
			"aws_ecs_task_definition":                                      tableAwsEcsTaskDefinition(ctx),
			"aws_efs_access_point":                                         tableAwsEfsAccessPoint(ctx),
			"aws_efs_file_system":                                          tableAwsElasticFileSystem(ctx),
			"aws_efs_mount_target":                                         tableAwsEfsMountTarget(ctx),
			"aws_eks_addon":                                                tableAwsEksAddon(ctx),
			"aws_eks_addon_version":                                        tableAwsEksAddonVersion(ctx),
			"aws_eks_cluster":                                              tableAwsEksCluster(ctx),
			"aws_eks_identity_provider_config":                             tableAwsEksIdentityProviderConfig(ctx),
			"aws_eks_node_group":                                           tableAwsEksNodeGroup(ctx),
			"aws_elastic_beanstalk_application":                            tableAwsElasticBeanstalkApplication(ctx),
			"aws_elastic_beanstalk_environment":                            tableAwsElasticBeanstalkEnvironment(ctx),
			"aws_elasticache_cluster":                                      tableAwsElastiCacheCluster(ctx),
			"aws_elasticache_parameter_group":                              tableAwsElastiCacheParameterGroup(ctx),
			"aws_elasticache_redis_metric_cache_hits_hourly":               tableAwsElasticacheRedisMetricCacheHitsHourly(ctx),
			"aws_elasticache_redis_metric_curr_connections_hourly":         tableAwsElasticacheRedisMetricCurrConnectionsHourly(ctx),
			"aws_elasticache_redis_metric_engine_cpu_utilization_daily":    tableAwsElasticacheRedisEngineCPUUtilizationDaily(ctx),
			"aws_elasticache_redis_metric_engine_cpu_utilization_hourly":   tableAwsElasticacheRedisEngineCPUUtilizationHourly(ctx),
			"aws_elasticache_redis_metric_get_type_cmds_hourly":            tableAwsElasticacheRedisMetricGetTypeCmdsHourly(ctx),
			"aws_elasticache_redis_metric_list_based_cmds_hourly":          tableAwsElasticacheRedisMetricListBasedCmdsHourly(ctx),
			"aws_elasticache_redis_metric_new_connections_hourly":          tableAwsElasticacheRedisMetricNewConnectionsHourly(ctx),
			"aws_elasticache_replication_group":                            tableAwsElastiCacheReplicationGroup(ctx),
			"aws_elasticache_reserved_cache_node":                          tableAwsElastiCacheReservedCacheNode(ctx),
			"aws_elasticache_subnet_group":                                 tableAwsElastiCacheSubnetGroup(ctx),
			"aws_elasticsearch_domain":                                     tableAwsElasticsearchDomain(ctx),
			"aws_emr_cluster":                                              tableAwsEmrCluster(ctx),
			"aws_emr_cluster_metric_is_idle":                               tableAwsEmrClusterMetricIsIdle(ctx),
			"aws_emr_instance":                                             tableAwsEmrInstance(ctx),
			"aws_emr_instance_fleet":                                       tableAwsEmrInstanceFleet(ctx),
			"aws_emr_instance_group":                                       tableAwsEmrInstanceGroup(ctx),
			"aws_eventbridge_bus":                                          tableAwsEventBridgeBus(ctx),
			"aws_eventbridge_rule":                                         tableAwsEventBridgeRule(ctx),
			"aws_fsx_file_system":                                          tableAwsFsxFileSystem(ctx),
			"aws_glacier_vault":                                            tableAwsGlacierVault(ctx),
			"aws_globalaccelerator_accelerator":                            tableAwsGlobalAcceleratorAccelerator(ctx),
			"aws_globalaccelerator_endpoint_group":                         tableAwsGlobalAcceleratorEndpointGroup(ctx),
			"aws_globalaccelerator_listener":                               tableAwsGlobalAcceleratorListener(ctx),
			"aws_glue_catalog_database":                                    tableAwsGlueCatalogDatabase(ctx),
			"aws_glue_catalog_table":                                       tableAwsGlueCatalogTable(ctx),
			"aws_glue_connection":                                          tableAwsGlueConnection(ctx),
			"aws_glue_crawler":                                             tableAwsGlueCrawler(ctx),
			"aws_glue_data_catalog_encryption_settings":                    tableAwsGlueDataCatalogEncryptionSettings(ctx),
			"aws_glue_dev_endpoint":                                        tableAwsGlueDevEndpoint(ctx),
			"aws_glue_job":                                                 tableAwsGlueJob(ctx),
			"aws_glue_security_configuration":                              tableAwsGlueSecurityConfiguration(ctx),
			"aws_guardduty_detector":                                       tableAwsGuardDutyDetector(ctx),
			"aws_guardduty_filter":                                         tableAwsGuardDutyFilter(ctx),
			"aws_guardduty_finding":                                        tableAwsGuardDutyFinding(ctx),
			"aws_guardduty_ipset":                                          tableAwsGuardDutyIPSet(ctx),
			"aws_guardduty_member":                                         tableAwsGuardDutyMember(ctx),
			"aws_guardduty_publishing_destination":                         tableAwsGuardDutyPublishingDestination(ctx),
			"aws_guardduty_threat_intel_set":                               tableAwsGuardDutyThreatIntelSet(ctx),
			"aws_iam_access_advisor":                                       tableAwsIamAccessAdvisor(ctx),
			"aws_iam_access_key":                                           tableAwsIamAccessKey(ctx),
			"aws_iam_account_password_policy":                              tableAwsIamAccountPasswordPolicy(ctx),
			"aws_iam_account_summary":                                      tableAwsIamAccountSummary(ctx),
			"aws_iam_action":                                               tableAwsIamAction(ctx),
			"aws_iam_credential_report":                                    tableAwsIamCredentialReport(ctx),
			"aws_iam_group":                                                tableAwsIamGroup(ctx),
			"aws_iam_policy":                                               tableAwsIamPolicy(ctx),
			"aws_iam_policy_attachment":                                    tableAwsIamPolicyAttachment(ctx),
			"aws_iam_policy_simulator":                                     tableAwsIamPolicySimulator(ctx),
			"aws_iam_role":                                                 tableAwsIamRole(ctx),
			"aws_iam_saml_provider":                                        tableAwsIamSamlProvider(ctx),
			"aws_iam_server_certificate":                                   tableAwsIamServerCertificate(ctx),
			"aws_iam_service_specific_credential":                          tableAwsIamUserServiceSpecificCredential(ctx),
			"aws_iam_user":                                                 tableAwsIamUser(ctx),
			"aws_iam_virtual_mfa_device":                                   tableAwsIamVirtualMfaDevice(ctx),
			"aws_identitystore_group":                                      tableAwsIdentityStoreGroup(ctx),
			"aws_identitystore_user":                                       tableAwsIdentityStoreUser(ctx),
			"aws_inspector_assessment_run":                                 tableAwsInspectorAssessmentRun(ctx),
			"aws_inspector_assessment_target":                              tableAwsInspectorAssessmentTarget(ctx),
			"aws_inspector_assessment_template":                            tableAwsInspectorAssessmentTemplate(ctx),
			"aws_inspector_exclusion":                                      tableAwsInspectorExclusion(ctx),
			"aws_inspector_finding":                                        tableAwsInspectorFinding(ctx),
			"aws_kinesis_consumer":                                         tableAwsKinesisConsumer(ctx),
			"aws_kinesis_firehose_delivery_stream":                         tableAwsKinesisFirehoseDeliveryStream(ctx),
			"aws_kinesis_stream":                                           tableAwsKinesisStream(ctx),
			"aws_kinesis_video_stream":                                     tableAwsKinesisVideoStream(ctx),
			"aws_kinesisanalyticsv2_application":                           tableAwsKinesisAnalyticsV2Application(ctx),
			"aws_kms_key":                                                  tableAwsKmsKey(ctx),
			"aws_kms_alias":                                                tableAwsKmsAlias(ctx),
			"aws_lambda_alias":                                             tableAwsLambdaAlias(ctx),
			"aws_lambda_function":                                          tableAwsLambdaFunction(ctx),
			"aws_lambda_function_metric_duration_daily":                    tableAwsLambdaFunctionMetricDurationDaily(ctx),
			"aws_lambda_function_metric_errors_daily":                      tableAwsLambdaFunctionMetricErrorsDaily(ctx),
			"aws_lambda_function_metric_invocations_daily":                 tableAwsLambdaFunctionMetricInvocationsDaily(ctx),
			"aws_lambda_layer":                                             tableAwsLambdaLayer(ctx),
			"aws_lambda_layer_version":                                     tableAwsLambdaLayerVersion(ctx),
			"aws_lambda_version":                                           tableAwsLambdaVersion(ctx),
			"aws_lightsail_instance":                                       tableAwsLightsailInstance(ctx),
			"aws_macie2_classification_job":                                tableAwsMacie2ClassificationJob(ctx),
			"aws_media_store_container":                                    tableAwsMediaStoreContainer(ctx),
			"aws_msk_cluster":                                              tableAwsMSKCluster(ctx),
			"aws_msk_serverless_cluster":                                   tableAwsMSKServerlessCluster(ctx),
			"aws_neptune_db_cluster":                                       tableAwsNeptuneDBCluster(ctx),
			"aws_networkfirewall_firewall_policy":                          tableAwsNetworkFirewallPolicy(ctx),
			"aws_networkfirewall_rule_group":                               tableAwsNetworkFirewallRuleGroup(ctx),
			"aws_opensearch_domain":                                        tableAwsOpenSearchDomain(ctx),
			"aws_organizations_account":                                    tableAwsOrganizationsAccount(ctx),
			"aws_pinpoint_app":                                             tableAwsPinpointApp(ctx),
			"aws_pricing_product":                                          tableAwsPricingProduct(ctx),
			"aws_pricing_service_attribute":                                tableAwsPricingServiceAttribute(ctx),
			"aws_ram_principal_association":                                tableAwsRAMPrincipalAssociation(ctx),
			"aws_ram_resource_association":                                 tableAwsRAMResourceAssociation(ctx),
			"aws_rds_db_cluster":                                           tableAwsRDSDBCluster(ctx),
			"aws_rds_db_cluster_parameter_group":                           tableAwsRDSDBClusterParameterGroup(ctx),
			"aws_rds_db_cluster_snapshot":                                  tableAwsRDSDBClusterSnapshot(ctx),
			"aws_rds_db_event_subscription":                                tableAwsRDSDBEventSubscription(ctx),
			"aws_rds_db_instance":                                          tableAwsRDSDBInstance(ctx),
			"aws_rds_db_instance_metric_connections":                       tableAwsRdsInstanceMetricConnections(ctx),
			"aws_rds_db_instance_metric_connections_daily":                 tableAwsRdsInstanceMetricConnectionsDaily(ctx),
			"aws_rds_db_instance_metric_connections_hourly":                tableAwsRdsInstanceMetricConnectionsHourly(ctx),
			"aws_rds_db_instance_metric_cpu_utilization":                   tableAwsRdsInstanceMetricCpuUtilization(ctx),
			"aws_rds_db_instance_metric_cpu_utilization_daily":             tableAwsRdsInstanceMetricCpuUtilizationDaily(ctx),
			"aws_rds_db_instance_metric_cpu_utilization_hourly":            tableAwsRdsInstanceMetricCpuUtilizationHourly(ctx),
			"aws_rds_db_instance_metric_read_iops":                         tableAwsRdsInstanceMetricReadIops(ctx),
			"aws_rds_db_instance_metric_read_iops_daily":                   tableAwsRdsInstanceMetricReadIopsDaily(ctx),
			"aws_rds_db_instance_metric_read_iops_hourly":                  tableAwsRdsInstanceMetricReadIopsHourly(ctx),
			"aws_rds_db_instance_metric_write_iops":                        tableAwsRdsInstanceMetricWriteIops(ctx),
			"aws_rds_db_instance_metric_write_iops_daily":                  tableAwsRdsInstanceMetricWriteIopsDaily(ctx),
			"aws_rds_db_instance_metric_write_iops_hourly":                 tableAwsRdsInstanceMetricWriteIopsHourly(ctx),
			"aws_rds_db_option_group":                                      tableAwsRDSDBOptionGroup(ctx),
			"aws_rds_db_parameter_group":                                   tableAwsRDSDBParameterGroup(ctx),
			"aws_rds_db_proxy":                                             tableAwsRDSDBProxy(ctx),
			"aws_rds_db_snapshot":                                          tableAwsRDSDBSnapshot(ctx),
			"aws_rds_db_subnet_group":                                      tableAwsRDSDBSubnetGroup(ctx),
			"aws_rds_reserved_db_instance":                                 tableAwsRDSReservedDBInstance(ctx),
			"aws_redshift_cluster":                                         tableAwsRedshiftCluster(ctx),
			"aws_redshift_cluster_metric_cpu_utilization_daily":            tableAwsRedshiftClusterMetricCpuUtilizationDaily(ctx),
			"aws_redshift_event_subscription":                              tableAwsRedshiftEventSubscription(ctx),
			"aws_redshift_parameter_group":                                 tableAwsRedshiftParameterGroup(ctx),
			"aws_redshift_snapshot":                                        tableAwsRedshiftSnapshot(ctx),
			"aws_redshift_subnet_group":                                    tableAwsRedshiftSubnetGroup(ctx),
			"aws_redshiftserverless_namespace":                             tableAwsRedshiftServerlessNamespace(ctx),
			"aws_redshiftserverless_workgroup":                             tableAwsRedshiftServerlessWorkgroup(ctx),
			"aws_region":                                                   tableAwsRegion(ctx),
			"aws_resource_explorer_index":                                  tableAWSResourceExplorerIndex(ctx),
			"aws_resource_explorer_search":                                 tableAWSResourceExplorerSearch(ctx),
			"aws_resource_explorer_supported_resource_type":                tableAWSResourceExplorerSupportedResourceType(ctx),
			"aws_route53_domain":                                           tableAwsRoute53Domain(ctx),
			"aws_route53_health_check":                                     tableAwsRoute53HealthCheck(ctx),
			"aws_route53_record":                                           tableAwsRoute53Record(ctx),
			"aws_route53_resolver_endpoint":                                tableAwsRoute53ResolverEndpoint(ctx),
			"aws_route53_resolver_rule":                                    tableAwsRoute53ResolverRule(ctx),
			"aws_route53_traffic_policy":                                   tableAwsRoute53TrafficPolicy(ctx),
			"aws_route53_traffic_policy_instance":                          tableAwsRoute53TrafficPolicyInstance(ctx),
			"aws_route53_zone":                                             tableAwsRoute53Zone(ctx),
			"aws_s3_access_point":                                          tableAwsS3AccessPoint(ctx),
			"aws_s3_account_settings":                                      tableAwsS3AccountSettings(ctx),
			"aws_s3_bucket":                                                tableAwsS3Bucket(ctx),
			"aws_sagemaker_app":                                            tableAwsSageMakerApp(ctx),
			"aws_sagemaker_domain":                                         tableAwsSageMakerDomain(ctx),
			"aws_sagemaker_endpoint_configuration":                         tableAwsSageMakerEndpointConfiguration(ctx),
			"aws_sagemaker_model":                                          tableAwsSageMakerModel(ctx),
			"aws_sagemaker_notebook_instance":                              tableAwsSageMakerNotebookInstance(ctx),
			"aws_sagemaker_training_job":                                   tableAwsSageMakerTrainingJob(ctx),
			"aws_secretsmanager_secret":                                    tableAwsSecretsManagerSecret(ctx),
			"aws_securityhub_action_target":                                tableAwsSecurityHubActionTarget(ctx),
			"aws_securityhub_finding":                                      tableAwsSecurityHubFinding(ctx),
			"aws_securityhub_finding_aggregator":                           tableAwsSecurityHubFindingAggregator(ctx),
			"aws_securityhub_hub":                                          tableAwsSecurityHub(ctx),
			"aws_securityhub_insight":                                      tableAwsSecurityHubInsight(ctx),
			"aws_securityhub_member":                                       tableAwsSecurityHubMember(ctx),
			"aws_securityhub_product":                                      tableAwsSecurityhubProduct(ctx),
			"aws_securityhub_standards_control":                            tableAwsSecurityHubStandardsControl(ctx),
			"aws_securityhub_standards_subscription":                       tableAwsSecurityHubStandardsSubscription(ctx),
			"aws_serverlessapplicationrepository_application":              tableAwsServerlessApplicationRepositoryApplication(ctx),
			"aws_servicequotas_default_service_quota":                      tableAwsServiceQuotasDefaultServiceQuota(ctx),
			"aws_servicequotas_service_quota":                              tableAwsServiceQuotasServiceQuota(ctx),
			"aws_servicequotas_service_quota_change_request":               tableAwsServiceQuotasServiceQuotaChangeRequest(ctx),
			"aws_ses_domain_identity":                                      tableAwsSESDomainIdentity(ctx),
			"aws_ses_email_identity":                                       tableAwsSESEmailIdentity(ctx),
			"aws_sfn_state_machine":                                        tableAwsStepFunctionsStateMachine(ctx),
			"aws_sfn_state_machine_execution":                              tableAwsStepFunctionsStateMachineExecution(ctx),
			"aws_sfn_state_machine_execution_history":                      tableAwsStepFunctionsStateMachineExecutionHistory(ctx),
			"aws_sns_topic":                                                tableAwsSnsTopic(ctx),
			"aws_sns_topic_subscription":                                   tableAwsSnsTopicSubscription(ctx),
			"aws_sqs_queue":                                                tableAwsSqsQueue(ctx),
			"aws_ssm_association":                                          tableAwsSSMAssociation(ctx),
			"aws_ssm_document":                                             tableAwsSSMDocument(ctx),
			"aws_ssm_inventory":                                            tableAwsSSMInventory(ctx),
			"aws_ssm_maintenance_window":                                   tableAwsSSMMaintenanceWindow(ctx),
			"aws_ssm_managed_instance":                                     tableAwsSSMManagedInstance(ctx),
			"aws_ssm_managed_instance_compliance":                          tableAwsSSMManagedInstanceCompliance(ctx),
			"aws_ssm_parameter":                                            tableAwsSSMParameter(ctx),
			"aws_ssm_patch_baseline":                                       tableAwsSSMPatchBaseline(ctx),
			"aws_ssoadmin_instance":                                        tableAwsSsoAdminInstance(ctx),
			"aws_ssoadmin_managed_policy_attachment":                       tableAwsSsoAdminManagedPolicyAttachment(ctx),
			"aws_ssoadmin_permission_set":                                  tableAwsSsoAdminPermissionSet(ctx),
			"aws_tagging_resource":                                         tableAwsTaggingResource(ctx),
			"aws_vpc":                                                      tableAwsVpc(ctx),
			"aws_vpc_customer_gateway":                                     tableAwsVpcCustomerGateway(ctx),
			"aws_vpc_dhcp_options":                                         tableAwsVpcDhcpOptions(ctx),
			"aws_vpc_egress_only_internet_gateway":                         tableAwsVpcEgressOnlyIGW(ctx),
			"aws_vpc_eip":                                                  tableAwsVpcEip(ctx),
			"aws_vpc_endpoint":                                             tableAwsVpcEndpoint(ctx),
			"aws_vpc_endpoint_service":                                     tableAwsVpcEndpointService(ctx),
			"aws_vpc_flow_log":                                             tableAwsVpcFlowlog(ctx),
			"aws_vpc_flow_log_event":                                       tableAwsVpcFlowLogEvent(ctx),
			"aws_vpc_internet_gateway":                                     tableAwsVpcInternetGateway(ctx),
			"aws_vpc_nat_gateway":                                          tableAwsVpcNatGateway(ctx),
			"aws_vpc_network_acl":                                          tableAwsVpcNetworkACL(ctx),
			"aws_vpc_peering_connection":                                   tableAwsVpcPeeringConnection(ctx),
			"aws_vpc_route":                                                tableAwsVpcRoute(ctx),
			"aws_vpc_route_table":                                          tableAwsVpcRouteTable(ctx),
			"aws_vpc_security_group":                                       tableAwsVpcSecurityGroup(ctx),
			"aws_vpc_security_group_rule":                                  tableAwsVpcSecurityGroupRule(ctx),
			"aws_vpc_subnet":                                               tableAwsVpcSubnet(ctx),
			"aws_vpc_vpn_connection":                                       tableAwsVpcVpnConnection(ctx),
			"aws_vpc_vpn_gateway":                                          tableAwsVpcVpnGateway(ctx),
			"aws_waf_rate_based_rule":                                      tableAwsWafRateBasedRule(ctx),
			"aws_waf_rule":                                                 tableAwsWAFRule(ctx),
			"aws_waf_rule_group":                                           tableAwsWafRuleGroup(ctx),
			"aws_waf_web_acl":                                              tableAwsWafWebAcl(ctx),
			"aws_wafregional_rule":                                         tableAwsWAFRegionalRule(ctx),
			"aws_wafv2_ip_set":                                             tableAwsWafv2IpSet(ctx),
			"aws_wafv2_regex_pattern_set":                                  tableAwsWafv2RegexPatternSet(ctx),
			"aws_wafv2_rule_group":                                         tableAwsWafv2RuleGroup(ctx),
			"aws_wafv2_web_acl":                                            tableAwsWafv2WebAcl(ctx),
			"aws_wellarchitected_workload":                                 tableAwsWellArchitectedWorkload(ctx),
			"aws_workspaces_workspace":                                     tableAwsWorkspace(ctx),
		},
	}

	return p
}
