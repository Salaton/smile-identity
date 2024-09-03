package smileidentity_test

import (
	"context"
	"log"
	"testing"

	smileidentity "github.com/Salaton/smile-identity"
)

func initializeClient(t *testing.T) *smileidentity.Client {
	client, err := smileidentity.NewClientFromEnvVars()
	if err != nil {
		t.Errorf("init test client: %v", err)
	}

	return client
}

func TestClient_VerifyPhoneNumber(t *testing.T) {
	c := initializeClient(t)

	type args struct {
		ctx   context.Context
		input *smileidentity.PhoneNumberVerification
	}
	tests := []struct {
		name    string
		args    args
		want    *smileidentity.PhoneNumberVerificationResponse
		wantErr bool
	}{
		{
			name: "Happy Case: Verify kenyan phonenumber",
			args: args{
				ctx: context.Background(),
				input: &smileidentity.PhoneNumberVerification{
					CallbackURL: "/",
					Country:     "KE",
					PhoneNumber: "0000000000",
					MatchFields: smileidentity.MatchFields{
						IDNumber: "00000000",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.VerifyPhoneNumber(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.VerifyPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			log.Printf("the response is %+v", got)
		})
	}
}
