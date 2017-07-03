package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	var pathToExecutable string
	wd, _ := os.Getwd()

	BeforeEach(func() {
		var err error
		pathToExecutable, err = gexec.Build("github.com/akshaymankar/int-yaml")
		Expect(err).ShouldNot(HaveOccurred())
		fmt.Println(pathToExecutable)
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("should exit with 0 exit code", func() {
		testFile := filepath.Join(wd, "fixtures", "test.yml")
		command := exec.Command(pathToExecutable, testFile)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
	})

	It("should print the yaml contents", func() {
		testFile := filepath.Join(wd, "fixtures", "test.yml")
		command := exec.Command(pathToExecutable, testFile)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("is_this_test: true"))
	})
})
