package main_test

import (
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

	fixturePath := func(name string) string {
		wd, _ := os.Getwd()
		return filepath.Join(wd, "fixtures", name)
	}

	BeforeEach(func() {
		var err error
		pathToExecutable, err = gexec.Build("github.com/akshaymankar/int-yaml")
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	Describe("Basics", func() {
		It("should exit with 0 exit code", func() {
			testFile := fixturePath("test.yml")
			command := exec.Command(pathToExecutable, testFile)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should print the yaml contents", func() {
			testFile := fixturePath("test.yml")
			command := exec.Command(pathToExecutable, testFile)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("is_this_test: true"))
		})
	})

	Describe("Interpolation", func() {
		It("should interpolate with all options", func() {
			testFile := fixturePath("interpolate.yml")
			_ = os.Setenv("ENVVAR_InEnv", "Magical")
			command := exec.Command(pathToExecutable, testFile,
				"--vars-file", fixturePath("vars-file.yml"),
				"--var-file", "getThisFromVarFile="+fixturePath("var-file.txt"),
				"--vars-env", "ENVVAR",
				"--var=getThisFromVar='Got it from var'")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			//If you edit this, remember to convert tabs to spaces
			expectedOutput := `examples:
  cli-vars:
    foo: Got it from var
  env-vars:
    foo: Magical
  var-file:
    foo: |
      Tressure of the var-file
  vars-file:
    bar: Map simple entry
    baz:
    - first element
    - second element
    foo: Simple entry`

			Eventually(session.Out).Should(gbytes.Say(expectedOutput))
		})
	})

	Describe("Patching", func() {
		It("should interpolate with all options", func() {
			testFile := fixturePath("patch.yml")
			command := exec.Command(pathToExecutable, testFile, "--ops-file", fixturePath("ops-file.yml"))
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			//If you edit this, remember to convert tabs to spaces
			expectedOutput := `movie: Rang de basanti
songs:
- Khalbali
- Paathshala
- Tu bin bataye
- Luka chuppi`

			Eventually(session.Out).Should(gbytes.Say(expectedOutput))
		})
	})

	Describe("Querying", func() {
		It("should get value of given path", func() {
			testFile := fixturePath("patch.yml")
			command := exec.Command(pathToExecutable, testFile, "--path", "/movie")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
			Eventually(session.Out).Should(gbytes.Say("^Rang de basanti\n$"))
		})
	})
})
