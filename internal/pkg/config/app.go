package config

import (
	utils "rpc-service/internal/pkg/utils"
)

// AppEnv 应用环境
// prod 生产环境
// test 测试环境
// dev 开发环境
var AppEnv = utils.GetEvnWithDefaultVal("APP_ENV", "dev")

var APIVersion = utils.GetEvnWithDefaultVal("API_VERSION", "v1")
