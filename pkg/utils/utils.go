package utils

import (
	"github.com/persistentsys/mysql-go-operator/pkg/apis/mysql/v1alpha1"
)

func Labels(v *v1alpha1.MySQL, tier string) map[string]string {
	return map[string]string{
		"app":        "MySQL",
		"MySQL_cr": v.Name,
		"tier":       tier,
	}
}

