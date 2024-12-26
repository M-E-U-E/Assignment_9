package controllers

import (
    "net/http"
    "net/http/httptest"

    "github.com/beego/beego/v2/server/web/context"
	
)

// Shared setupTestContext function
func setupTestContext(req *http.Request) (*httptest.ResponseRecorder, *context.Context) {
    w := httptest.NewRecorder()
    ctx := &context.Context{
        Input:  &context.BeegoInput{Context: &context.Context{}},
        Output: &context.BeegoOutput{Context: &context.Context{}},
    }
    ctx.Reset(w, req)
    return w, ctx
}
