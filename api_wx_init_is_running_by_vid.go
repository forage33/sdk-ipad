// api_wx_init_is_running_by_vid.go IsRunningByVid 判断指定 VID 的账号是否正在运行中。

package sdk_ipad

// IsRunningByVid 判断指定 VID 的账号是否正在运行中。
//
// 参数:
//   - host: 协议服务地址
//   - vid: 企业微信 VID
//
// 返回: bool 是否在运行中
func (s *SdkIPad) IsRunningByVid(host, vid string) bool {
	resp, err := s.GetRunClient(host)
	if err != nil {
		return false
	}
	for _, item := range resp.Data {
		for userID, userData := range item {
			if vid == userID && userData.IsLogin {
				return true
			}
		}
	}
	return false
}
