// api_wx_init_get_running_by_vid.go GetRunningByVid 根据 VID 查找运行中的实例，返回用户数据和是否找到。

package sdk_ipad

// GetRunningByVid 根据 VID 查找运行中的实例，返回用户数据和是否找到。
//
// 参数:
//   - host: 协议服务地址
//   - vid: 企业微信 VID
//
// 返回:
//   - *UserData: 用户数据，未找到时为 nil
//   - bool: 是否找到
func (s *SdkIPad) GetRunningByVid(host, vid string) (*UserData, bool) {
	resp, err := s.GetRunClient(host)
	if err != nil {
		return nil, false
	}
	for _, item := range resp.Data {
		for userID, userData := range item {
			if vid == userID && userData.IsLogin {
				ud := userData
				return &ud, true
			}
		}
	}
	return nil, false
}
