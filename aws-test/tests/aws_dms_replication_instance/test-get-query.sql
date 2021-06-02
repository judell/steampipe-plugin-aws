select replication_instance_identifier, arn, replication_instance_class, publicly_accessible, allocated_storage, auto_minor_version_upgrade, availability_zone, dns_name_servers, free_until, multi_az, preferred_maintenance_window, replication_instance_status, secondary_availability_zone, tags_src, title, tags, akas, partition, region, account_id
from aws.aws_dms_replication_instance
where arn = '{{ output.resource_aka.value }}';