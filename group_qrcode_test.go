package sdk_ipad

import "testing"

func TestNormalizeQrCodePath(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "windows 反斜杠路径转正斜杠",
			in:   `http://8.148.226.80:8060/download\RoomQrCode\10956327057102598.jpg`,
			want: "http://8.148.226.80:8060/download/RoomQrCode/10956327057102598.jpg",
		},
		{
			name: "正斜杠 URL 保持不变",
			in:   "http://47.94.7.218:8083/RoomQrCode/10696052955016166.jpg",
			want: "http://47.94.7.218:8083/RoomQrCode/10696052955016166.jpg",
		},
		{
			name: "混合分隔符全部转正斜杠",
			in:   `http://host/a\b/c\d.jpg`,
			want: "http://host/a/b/c/d.jpg",
		},
		{
			name: "URL 编码反斜杠转正斜杠",
			in:   "http://host/download%5CRoomQrCode%5C10956327057102598.jpg",
			want: "http://host/download/RoomQrCode/10956327057102598.jpg",
		},
		{
			name: "小写 URL 编码反斜杠转正斜杠",
			in:   "http://host/download%5cRoomQrCode%5c10956327057102598.jpg",
			want: "http://host/download/RoomQrCode/10956327057102598.jpg",
		},
		{
			name: "双反斜杠全部转正斜杠",
			in:   `http://host/download\\RoomQrCode\\10956327057102598.jpg`,
			want: "http://host/download//RoomQrCode//10956327057102598.jpg",
		},
		{
			name: "空字符串保持为空",
			in:   "",
			want: "",
		},
		{
			name: "无分隔符不变化",
			in:   "http://host/qrcode.jpg",
			want: "http://host/qrcode.jpg",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NormalizeQrCodePath(tc.in); got != tc.want {
				t.Errorf("NormalizeQrCodePath(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
