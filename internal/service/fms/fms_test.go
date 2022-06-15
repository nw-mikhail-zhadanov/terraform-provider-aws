package fms_test

import (
	"testing"
)

func TestAccFMS_serial(t *testing.T) {
	testCases := map[string]map[string]func(t *testing.T){
		"AdminAccount": {
			"basic": testAccAdminAccount_basic,
		},
		"Policy": {
			"basic":                  TestAccPolicy_basic,
			"cloudfrontDistribution": TestAccPolicy_cloudFrontDistribution,
			"includeMap":             TestAccPolicy_includeMap,
			"update":                 TestAccPolicy_update,
			"resourceTags":           TestAccPolicy_resourceTags,
			"tags":                   TestAccPolicy_tags,
		},
	}

	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			for name, tc := range m {
				tc := tc
				t.Run(name, func(t *testing.T) {
					tc(t)
				})
			}
		})
	}
}
