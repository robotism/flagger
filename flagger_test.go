package flagger

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/robotism/flagger/example/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestBinder(test *testing.T) {
	f := New()

	c := &config.AppConfig{}

	os.Setenv("SERVER_PORT", "9090")

	err := f.Parse(c, []string{"-d=true", "--timezone=Asia/Shanghai", "--database.default.host=127.0.0.1"}...)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", c)

	assert.True(test, c.Debug)
	assert.Equal(test, "Asia/Shanghai", c.Timezone)
	assert.Equal(test, 1, len(c.Database))
	assert.Equal(test, "127.0.0.1", c.Database["default"].Host)
	assert.Equal(test, 9090, c.Server.Port)
	assert.Equal(test, 3306, c.Database["default"].Port)
	assert.Equal(test, "root", c.Database["default"].User)
	assert.Equal(test, "", c.Database["default"].Pass)
}

func TestCobra(test *testing.T) {

	cmd := &cobra.Command{
		Use:   "test",
		Short: "test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	f := New()
	c := &config.AppConfig{}

	f.UseFlags(cmd.Flags())

	err := f.Parse(c, []string{"-d=false", "--timezone=US/Pacific", "--database.default.host=1.1.1.1"}...)
	if err != nil {
		panic(err)
	}

	usage := cmd.UsageString()

	log.Printf("%s\n", usage)

	assert.NotEmpty(test, usage)

	assert.True(test, strings.Contains(usage, "database.<dbkey>.host"))

	log.Printf("%+v\n", c)
}
