package integration_test


var HttpMock *mocks.HttpMock
func mocks(t *testing.T)[]interface{
	HttpMock=mocks.NewHttpMock(t)
	return []interface{
		fx.Annotate(HttpMock,fx.As(new(ClientHttp)))
	}
}
func inttest(t *testing.T,r interface{}){
	app:=fxtest.New(t, app.InjectApp(),fx.Replace(mocks(t)),fx.Invoke(r))
	defer app.RequireStop()
	app.RequireStart()
}

func res(c int,b string)*http.Response{
	r:=httptest.NewRecorder()
	if r!=""{
		r.Write([]byte(r))
	}
	r.Code=c
	return r.Result()
}

func TestApp_IntegrationTest_with_no_error(t testing.T){
	inttest(t, func(app *App){
		HttpMock.On("Do",mockery.MatchedBy(func(rq *http.Request)bool{

		})).Return(res(http.StatusOK,"todo piola"))
	})
	err:=app.Run()
	assert.NoError(t,err)
}