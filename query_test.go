package aws_ip

import (
	"testing"
)

func TestQuery(t *testing.T) {
	// ref. https://docs.aws.amazon.com/ja_jp/general/latest/gr/aws-ip-ranges.html
	t.Run("Get all IPv4 addresses for the service in a particular region", func(t *testing.T) {
		ret := Query().Region("us-east-1").Service("CODEBUILD").Select()
		want := "34.228.4.208/28"
		got := ret[0].IPPrefix
		if got != want {
			t.Errorf("want: %s, got: %s", want, got)
		}
	})
	t.Run("Get information about a specific network boundary group", func(t *testing.T) {
		ret := Query().Region("us-west-2").NetworkBorderGroup("us-west-2-lax-1").Select()
		want := "us-west-2-lax-1"
		got := ret[0].NetworkBorderGroup
		if got != want {
			t.Errorf("want: %s, got: %s", want, got)
		}
	})
}
