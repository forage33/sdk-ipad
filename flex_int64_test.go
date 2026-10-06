package sdk_ipad

import (
	"encoding/json"
	"testing"
)

func TestFlexInt64UnmarshalNumber(t *testing.T) {
	var f FlexInt64
	if err := json.Unmarshal([]byte(`1746252388`), &f); err != nil {
		t.Fatalf("unmarshal number failed: %v", err)
	}
	if int64(f) != 1746252388 {
		t.Fatalf("got %d, want 1746252388", int64(f))
	}
}

func TestFlexInt64UnmarshalString(t *testing.T) {
	var f FlexInt64
	if err := json.Unmarshal([]byte(`"1746252388"`), &f); err != nil {
		t.Fatalf("unmarshal string failed: %v", err)
	}
	if int64(f) != 1746252388 {
		t.Fatalf("got %d, want 1746252388", int64(f))
	}
}

func TestFlexInt64UnmarshalNullAndEmpty(t *testing.T) {
	for _, raw := range []string{`null`, `""`} {
		var f FlexInt64
		if err := json.Unmarshal([]byte(raw), &f); err != nil {
			t.Fatalf("unmarshal %s failed: %v", raw, err)
		}
		if int64(f) != 0 {
			t.Fatalf("unmarshal %s got %d, want 0", raw, int64(f))
		}
	}
}

func TestFlexInt64UnmarshalInvalid(t *testing.T) {
	var f FlexInt64
	if err := json.Unmarshal([]byte(`"abc"`), &f); err == nil {
		t.Fatal("unmarshal invalid string should fail")
	}
	if err := json.Unmarshal([]byte(`"1746252388 "`), &f); err == nil {
		t.Fatal("unmarshal string with space should fail")
	}
	if err := json.Unmarshal([]byte(`1746252388.5`), &f); err == nil {
		t.Fatal("unmarshal non-integer float should fail")
	}
	if err := json.Unmarshal([]byte(`1.7e-1`), &f); err == nil {
		t.Fatal("unmarshal non-integer float exponent should fail")
	}
}

func TestFlexInt64UnmarshalFloatInt(t *testing.T) {
	for _, raw := range []string{`0.0`, `1746252388.0`, `1.7e9`} {
		var f FlexInt64
		if err := json.Unmarshal([]byte(raw), &f); err != nil {
			t.Fatalf("unmarshal %s failed: %v", raw, err)
		}
		want := int64(0)
		switch raw {
		case `0.0`:
			want = 0
		case `1746252388.0`:
			want = 1746252388
		case `1.7e9`:
			want = 1700000000
		}
		if int64(f) != want {
			t.Fatalf("unmarshal %s got %d, want %d", raw, int64(f), want)
		}
	}
}

func TestFlexInt64UnmarshalNegative(t *testing.T) {
	var f FlexInt64
	if err := json.Unmarshal([]byte(`-1`), &f); err != nil {
		t.Fatalf("unmarshal negative failed: %v", err)
	}
	if int64(f) != -1 {
		t.Fatalf("got %d, want -1", int64(f))
	}
}

func TestFlexInt64UnmarshalOverflow(t *testing.T) {
	var f FlexInt64
	if err := json.Unmarshal([]byte(`99999999999999999999`), &f); err == nil {
		t.Fatal("unmarshal overflow should fail")
	}
	if err := json.Unmarshal([]byte(`1e20`), &f); err == nil {
		t.Fatal("unmarshal overflow float should fail")
	}
}

func TestFlexInt64Marshal(t *testing.T) {
	f := FlexInt64(1746252388)
	b, err := json.Marshal(f)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if string(b) != `1746252388` {
		t.Fatalf("marshal got %s, want 1746252388", b)
	}
}

// TestBanRespBothFormats 验证协议返回 number 与 string 两种格式时 BanResp 均能解析。
func TestBanRespBothFormats(t *testing.T) {
	raws := []string{
		// 协议当前实际返回：ban_finish_time 为 number
		`{"data":{"vid":"1688856572757825","mainData":{"ban_start_time":"1786942613","ban_title":"t","ban_finish_time":0}},"errcode":0}`,
		// 协议历史格式：ban_finish_time 为字符串数字
		`{"data":{"vid":"1688856572757825","mainData":{"ban_start_time":1786942613,"ban_title":"t","ban_finish_time":"1746252388"}},"errcode":0}`,
	}
	for i, raw := range raws {
		var resp BanResp
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("raw[%d] unmarshal failed: %v", i, err)
		}
		if i == 0 {
			if int64(resp.Data.MainData.BanFinishTime) != 0 || int64(resp.Data.MainData.BanStartTime) != 1786942613 {
				t.Fatalf("raw[%d] parse mismatch: finish=%d start=%d", i,
					int64(resp.Data.MainData.BanFinishTime), int64(resp.Data.MainData.BanStartTime))
			}
		} else {
			if int64(resp.Data.MainData.BanFinishTime) != 1746252388 || int64(resp.Data.MainData.BanStartTime) != 1786942613 {
				t.Fatalf("raw[%d] parse mismatch: finish=%d start=%d", i,
					int64(resp.Data.MainData.BanFinishTime), int64(resp.Data.MainData.BanStartTime))
			}
		}
	}
}
