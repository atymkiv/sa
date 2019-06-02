swagger_api_gen:
	cd ${APP_PATH} && swagger generate spec -b ./cmd/api -o ${APP_PATH}/package/api/swaggerui/swagger.json
swagger_api_serve:
	cd ${APP_PATH}/package/api/swaggerui && swagger serve swagger.json
