package tcp

import (
	"testing"

	insecure "github.com/dms3-p2p/go-conn-security/insecure"
	tptu "github.com/dms3-p2p/go-p2p-transport-upgrader"
	utils "github.com/dms3-p2p/go-p2p-transport/test"
	ma "github.com/dms3-mft/go-multiaddr"
	mplex "github.com/dms3-why/go-smux-multiplex"
)

func TestTcpTransport(t *testing.T) {
	for i := 0; i < 2; i++ {
		ta := NewTCPTransport(&tptu.Upgrader{
			Secure: insecure.New("peerA"),
			Muxer:  new(mplex.Transport),
		})
		tb := NewTCPTransport(&tptu.Upgrader{
			Secure: insecure.New("peerB"),
			Muxer:  new(mplex.Transport),
		})

		zero := "/ip4/127.0.0.1/tcp/0"
		utils.SubtestTransport(t, ta, tb, zero, "peerA")

		envReuseportVal = false
	}
	envReuseportVal = true
}

func TestTcpTransportCantListenUtp(t *testing.T) {
	for i := 0; i < 2; i++ {
		utpa, err := ma.NewMultiaddr("/ip4/127.0.0.1/udp/0/utp")
		if err != nil {
			t.Fatal(err)
		}

		tpt := NewTCPTransport(&tptu.Upgrader{
			Secure: insecure.New("peerB"),
			Muxer:  new(mplex.Transport),
		})

		_, err = tpt.Listen(utpa)
		if err == nil {
			t.Fatal("shouldnt be able to listen on utp addr with tcp transport")
		}

		envReuseportVal = false
	}
	envReuseportVal = true
}
