package bilierrorcode

// ErrorCode ErrorCode
type ErrorCode int32

// ErrorCodeDetail Detailed information
type ErrorCodeDetail struct {
	Code        ErrorCode
	Message     string
	Description string
}

// GetRegion Get which part your error code is
func (code ErrorCode) GetRegion() string {
	if code >= 0 && code <= 990000 {
		return "main_or_ep"
	} else if code >= 1000000 && code <= 1999999 {
		return "live"
	} else if code >= 5000000 && code < 6000000 {
		return "bbq"
	} else if code >= 2000000 && code <= 2099999 {
		return "ticket"
	} else if code >= 2000000 && code <= 2999999 {
		return "open_platform"
	}
	return "unknown"
}

// GetDetail Get detail information about the ErrorCode given
func (code ErrorCode) GetDetail() ErrorCodeDetail {
	var result ErrorCodeDetail

	switch code.GetRegion() {
	case "main_or_ep":
		// 尝试匹配主站
		result = getMainSiteDetail(code)
		if result.Message != "" {
			return result
		}
		// TODO: ep
		break
	case "live":
		result = getLiveSiteDetail(code)
		if result.Message != "" {
			return result
		}
		break
	case "bbq":
		// result = getBBQSiteDetail(code)
		if result.Message != "" {
			return result
		}
		break
	case "ticket":
		// result = getTicketSiteDetail(code)
		if result.Message != "" {
			return result
		}
		break
	case "open_platform":
		// result = getOpenPlatformDetail(code)
		if result.Message != "" {
			return result
		}
		break

	}

	// 默认情况 返回空
	result.Code = code
	result.Message = ""
	result.Description = ""
	return result
}
