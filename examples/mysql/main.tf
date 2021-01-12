// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
}

resource "oci_mysql_mysql_backup" "test_mysql_backup" {
  db_system_id = oci_mysql_mysql_db_system.test_mysql_backup_db_system.id
}

resource "oci_mysql_mysql_db_system" "test_mysql_backup_db_system" {
  #Required
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  shape_name          = "VM.Standard.E2.2"
  subnet_id           = oci_core_subnet.test_subnet.id

  #Optional
  data_storage_size_in_gb = "50"
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
  #Required
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  shape_name          = "VM.Standard.E2.2"
  subnet_id           = oci_core_subnet.test_subnet.id

  #Optional
  backup_policy {
    is_enabled        = "false"
    retention_in_days = "10"
    window_start_time = "01:00-00:00"
  }

  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.mysql_defined_tags_value}"}
  #freeform_tags = var.mysql_freeform_tags
  description = "MySQL Database Service"

  display_name   = "DBSystem001"
  fault_domain   = "FAULT-DOMAIN-1"
  hostname_label = "hostnameLabel"
  ip_address     = "10.0.0.8"

  maintenance {
    window_start_time = "sun 01:00"
  }

  port          = "3306"
  port_x        = "33306"

  # Creating DB System using a backup
  source {
    backup_id   = oci_mysql_mysql_backup.test_mysql_backup.id
    source_type = "BACKUP"
  }
}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid

  #Optional
  state        = "ACTIVE"
  display_name = "VM.Standard.E2.2.Built-in"
  shape_name   = "VM.Standard.E2.2"
}

data "oci_mysql_shapes" "test_shapes" {
  compartment_id = var.compartment_ocid
  availability_domain = lower(
    data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name,
  )
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

output "configuration_id" {
  value = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
}

