package core

import (
	"github.com/fbiville/headache/fs_mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configuration loader", func() {

	var (
		t                   GinkgoTInterface
		fileReader          *fs_mocks.FileReader
		configurationLoader *ConfigurationLoader
	)

	BeforeEach(func() {
		t = GinkgoT()
		fileReader = new(fs_mocks.FileReader)
		configurationLoader = &ConfigurationLoader{Reader: fileReader}
	})

	AfterEach(func() {
		fileReader.AssertExpectations(t)
	})

	It("loads a simple configuration", func() {
		configurationPath := "conf.json"
		file := NewInMemoryFile(`{"headerFile": "some-file.txt", "style": "SlashStar", "includes": ["**/*.go"]}`)
		fileReader.On("Open", configurationPath).Return(file, nil)
		fileReader.On("Read", configurationPath).Return(file.Contents, nil)

		configuration, err := configurationLoader.readConfiguration(&configurationPath, "file://../docs/schema.json")

		Expect(err).To(BeNil())
		Expect(configuration).To(Equal(&Configuration{
			CommentStyle: "SlashStar",
			HeaderFile:   "some-file.txt",
			Includes:     []string{"**/*.go"},
			Path:         &configurationPath,
		}))
	})

	It("loads a compound configuration", func() {
		configurationPath := "conf.json"
		file := NewInMemoryFile(`[
  {"headerFile": "some-file.txt", "style": "SlashStar", "includes": ["**/*.go"]},
  {"headerFile": "some-file.txt", "style": "REM", "includes": ["**/*.sql"]}
]`)
		fileReader.On("Open", configurationPath).Return(file, nil)
		fileReader.On("Read", configurationPath).Return(file.Contents, nil)

		configuration, err := configurationLoader.readConfiguration(&configurationPath, "file://../docs/schema.json")

		Expect(err).To(BeNil())
		Expect(configuration).To(Equal([]Configuration{
			{
				CommentStyle: "SlashStar",
				HeaderFile:   "some-file.txt",
				Includes:     []string{"**/*.go"},
				Path:         &configurationPath,
			},
			{
				CommentStyle: "REM",
				HeaderFile:   "some-file.txt",
				Includes:     []string{"**/*.sql"},
				Path:         &configurationPath,
			},
		}))
	})
})
