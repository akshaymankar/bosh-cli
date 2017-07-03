package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntYaml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntYaml Suite")
}
