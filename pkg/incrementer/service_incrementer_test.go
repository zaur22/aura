// +build integration

package incrementer

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/kylelemons/godebug/pretty"
	"github.com/zaur22/aura/pkg/config"
	"github.com/zaur22/aura/pkg/db"
	"os"
	"path"
	"testing"
)

var dbConn *sqlx.DB

type (
	IncrementTestCase struct {
		description string
		fixturesDir string
		golden      []increment
	}

	SetConfigTestCase struct {
		description string
		fixturesDir string
		dto         SetSettingsDTO
		golden      []increment
	}

	GetValueTestCase struct {
		description   string
		fixturesDir   string
		expectedValue uint64
	}
)

func TestMain(m *testing.M) {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var configs = config.GetConfigs()

	dbDTO := db.ConnectDTO{
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		UserName: configs.DBUsername,
		Password: configs.DBPassword,
		DBName:   configs.DBName,
		SSLMode:  configs.SSLMode,
	}

	dbConn = db.Connect(ctx, dbDTO)
	defer dbConn.Close()

	os.Exit(m.Run())
}

func TestIncrementNumber(t *testing.T) {
	var tCases = []IncrementTestCase{
		{
			description: "init checking",
			fixturesDir: "fixtures1",
			golden: []increment{
				{
					ID:            10001,
					CurrentValue:  1,
					MaxValue:      0,
					IncrementStep: 1,
				},
			},
		},
		{
			description: "checking increment with max value and increment step",
			fixturesDir: "fixtures2",
			golden: []increment{
				{
					ID:            10001,
					CurrentValue:  4,
					MaxValue:      11,
					IncrementStep: 10,
				},
			},
		},
	}
	ctx := context.Background()
	for _, tCase := range tCases {
		loadFixtures(tCase.fixturesDir)
		service := getService()
		t.Run(tCase.description, func(t *testing.T) {

			err := service.IncrementNumber(ctx)
			if err != nil {
				t.Errorf("increment error: %v", err)
				return
			}

			var increments = []increment{}
			err = dbConn.Select(&increments, "SELECT * FROM increments")
			if err != nil {
				t.Errorf("select error: %v", err)
				return
			}

			if diff := pretty.Compare(increments, &tCase.golden); diff != "" {
				t.Errorf("bad increment result (-got +want):\n%s", diff)
				return
			}
		})
	}
}

func TestSetConfig(t *testing.T) {
	var tCases = []SetConfigTestCase{
		{
			description: "update values after initing",
			fixturesDir: "fixtures1",
			dto: SetSettingsDTO{
				IncrementStep: 5,
				MaxValue:      10,
			},
			golden: []increment{
				{
					ID:            10001,
					CurrentValue:  0,
					MaxValue:      10,
					IncrementStep: 5,
				},
			},
		},
	}
	ctx := context.Background()
	for _, tCase := range tCases {
		loadFixtures(tCase.fixturesDir)
		service := getService()
		t.Run(tCase.description, func(t *testing.T) {

			err := service.SetSettings(ctx, tCase.dto)
			if err != nil {
				t.Errorf("set config error: %v", err)
				return
			}

			var increments = []increment{}
			err = dbConn.Select(&increments, "SELECT * FROM increments")
			if err != nil {
				t.Errorf("select error: %v", err)
				return
			}

			if diff := pretty.Compare(increments, &tCase.golden); diff != "" {
				t.Errorf("bad set config result (-got +want):\n%s", diff)
			}

		})
	}
}

func TestGetValue(t *testing.T) {
	var tCases = []GetValueTestCase{
		{
			description:   "init state",
			fixturesDir:   "fixtures1",
			expectedValue: 0,
		},
		{
			description:   "init state",
			fixturesDir:   "fixtures2",
			expectedValue: 5,
		},
	}
	ctx := context.Background()

	for _, tCase := range tCases {
		loadFixtures(tCase.fixturesDir)
		service := getService()
		t.Run(tCase.description, func(t *testing.T) {

			val, err := service.GetNumber(ctx)
			if err != nil {
				t.Errorf("get value: %v", err)
				return
			}

			if tCase.expectedValue != val {
				t.Errorf("bad result value: got %d, want: %d", val, tCase.expectedValue)
				return
			}

		})
	}
}

func loadFixtures(dirName string) {

	fixtures, err := testfixtures.New(
		testfixtures.Database(dbConn.DB),
		testfixtures.Dialect("postgres"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Directory(path.Join("testdata", dirName)),
	)
	if err != nil {
		panic(fmt.Errorf("fixtures initing error: %v", err))
	}

	if err = fixtures.Load(); err != nil {
		panic(fmt.Errorf("fixtures loading error: %v", err))
	}
}

func getService() Incrementer {
	ctx := context.Background()
	service, err := NewIncrementer(
		ctx,
		NewIncrementerDTO{
			DBClient: dbConn,
		})

	if err != nil {
		panic(fmt.Errorf("can't create interpreter service: %v", err))
		return nil
	}

	return service
}
