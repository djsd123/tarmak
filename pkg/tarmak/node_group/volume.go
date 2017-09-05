package node_group

import (
	"fmt"

	clusterv1alpha1 "github.com/jetstack/tarmak/pkg/apis/cluster/v1alpha1"
	"github.com/jetstack/tarmak/pkg/tarmak/interfaces"
)

var _ interfaces.Volume = &Volume{}

type Volume struct {
	conf *clusterv1alpha1.Volume

	volumeType string
	device     string
}

func NewVolumeFromConfig(pos int, provider interfaces.Provider, conf *clusterv1alpha1.Volume) (*Volume, error) {
	volume := &Volume{
		conf: conf,
	}

	volumeType, err := provider.VolumeType(conf.Type)
	if err != nil {
		return nil, err
	}
	volume.volumeType = volumeType

	if provider.Name() == clusterv1alpha1.CloudAmazon {
		letters := "defghijklmnop"
		volume.device = fmt.Sprintf("/dev/sd%c", letters[pos])
	}

	return volume, nil
}

func (v *Volume) Device() string {
	return v.device
}

func (v *Volume) Name() string {
	return v.conf.Name
}

func (v *Volume) Size() int {
	return int(v.conf.Size.Value() / 1024 / 1024 / 1024)
}

func (v *Volume) Type() string {
	return v.volumeType
}
