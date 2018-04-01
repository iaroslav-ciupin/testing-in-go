// build +intgr

package code

type AppSuite struct {
	suite.Suite
	app *App
}

func (s *AppSuite) SetupSuite() {
	s.app = App{}
	// set mocks to app
	// prepare database schema, insert test data, etc.
	go s.app.Start()
}

// func (s *AppSuite) TestBlaBla() {}

func (s *AppSuite) TearDownSuite() {
	s.app.Shutdown()
	// database cleanup
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, &AppSuite{})
}