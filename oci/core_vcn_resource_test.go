// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v32/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig()
	s.ResourceName = "oci_core_virtual_network.t"
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestAccResourceCoreVirtualNetwork_basic() {
	var resId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create with cidr_block
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test create with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "11.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test add cidr with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "11.0.0.0/16", "12.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test remove cidr with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "12.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test modify cidr with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "11.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test update
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-vcn"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-vcn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
				),
			},
			// test a destructive update results in a new resource
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "MyTestDNSLabel"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "dns_label", "mytestdnslabel"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId2, err := fromInstanceState(s, "oci_core_virtual_network.t", "id")
						if resId == resId2 {
							return fmt.Errorf("Expected new vcn ocid, got the same")
						}
						return err
					},
				),
			},
			// DNS capitalization changes should be ignored.
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "mYtESTdnsLABEL"
					}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// DNS label change should cause a change
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "mynewlabel"
					}`,
				ExpectNonEmptyPlan: true,
				PlanOnly:           true,
			},
		},
	})
}

func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVirtualNetworkTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
