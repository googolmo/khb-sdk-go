package khb

import (
	"testing"
)

func TestScreenshot_invoke(t *testing.T) {
	type fields struct {
		Template       string
		URL            string
		HTML           string
		Data           map[string]interface{}
		Headers        map[string]interface{}
		Device         string
		CustomDevice   *DeviceDescriptor
		Type           string
		FullPage       bool
		Quality        int
		OmitBackground bool
	}
	type args struct {
		token string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantAPIError bool
	}{
		// TODO: Add test cases.
		{
			name:   "template url html empty error",
			fields: fields{},
			args: args{
				token: "123",
			},
			wantErr:      true,
			wantAPIError: true,
		},
		{
			name: "missing token",
			fields: fields{
				Template: "1234",
			},
			args: args{
				token: "",
			},
			wantErr:      true,
			wantAPIError: true,
		},
		{
			name: "url err",
			fields: fields{
				URL: "abd/sdf",
			},
			args:         args{},
			wantErr:      true,
			wantAPIError: true,
		},
		{
			name: "normal request",
			fields: fields{
				Template: "jdp3yn4s",
				Data: map[string]interface{}{
					"backgroundColor": "#c8f1fd",
					"user": map[string]interface{}{
						"avatar":   "https://khb-sample.oss-cn-shanghai.aliyuncs.com/sample/girl_2.jpg",
						"nickname": "筱墨",
					},
					"tip":    "快海报专业设计师提供的模板设计. ",
					"cover":  "https://khb-sample.oss-cn-shanghai.aliyuncs.com/sample/watermelon.jpg",
					"qrcode": "https://khb-sample.oss-cn-shanghai.aliyuncs.com/sample/sample_qr_0.png",
					"brand":  "快海报",
					"slogan": "小程序分享海报生成服务",
				},
			},
			args: args{
				token: "token",
			},
			wantErr:      false,
			wantAPIError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			screenshot := &Screenshot{
				Template:       tt.fields.Template,
				URL:            tt.fields.URL,
				HTML:           tt.fields.HTML,
				Data:           tt.fields.Data,
				Headers:        tt.fields.Headers,
				Device:         tt.fields.Device,
				CustomDevice:   tt.fields.CustomDevice,
				Type:           tt.fields.Type,
				FullPage:       tt.fields.FullPage,
				Quality:        tt.fields.Quality,
				OmitBackground: tt.fields.OmitBackground,
			}
			got, err := screenshot.invoke(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Screenshot.invoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			e, ok := err.(*APIError)
			if ok != tt.wantAPIError {
				t.Errorf("Screenshot.invoke() error = %v", err)
				return
			}
			t.Logf("Screenshot.invoke() error = %v", e)
			if got != nil {
				t.Logf("Screenshot.invoke() result = %v", got)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Screenshot.invoke() = %v, want %v", got, tt.want)
			// }
		})
	}
}
