package goframework_resp

var resultCall = func(c interface{}, httpStatus int, code int, message string, data interface{}, count int64) {

	/*
		if data == nil {
			c.JSON(httpStatus, map[string]interface{}{"code": code, "message": message})
			return
		}

		if count < 0 {
			c.JSON(httpStatus, map[string]interface{}{"code": code, "message": message, "data": data})
			return
		}

		c.JSON(httpStatus, map[string]interface{}{"code": code, "message": message, "data": data, "count": count})
	*/
}

func ReplaceResultCall(f func(c interface{}, httpStatus int, code int, message string, data interface{}, count int64)) {
	resultCall = f
}
