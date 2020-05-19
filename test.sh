#!/bin/bash

source <( yamltoenv -f test.yaml )

echo $cluster_cloud_provisioner_node_group_1_additional_security_group_ids
