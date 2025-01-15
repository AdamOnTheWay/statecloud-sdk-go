// Code generated by crafter generator. DO NOT EDIT.

package eci

var baseDomain = "https://eci-global.ctapi.ctyun.cn"

type ClientSet interface {
	ContainerGroup() ContainerGroupClient
	ImageCache() ImageCacheClient
	DataCache() DataCacheClient
	VirtualNode() VirtualNodeClient
	Price() PriceClient
	CommitContainerTask() CommitContainerTaskClient
	Tag() TagClient
	Region() RegionClient
	EnterpriseProject() EnterpriseProjectClient
}

type clientSet struct {
	containerGroupCli      ContainerGroupClient
	imageCacheCli          ImageCacheClient
	dataCacheCli           DataCacheClient
	virtualNodeCli         VirtualNodeClient
	priceCli               PriceClient
	commitContainerTaskCli CommitContainerTaskClient
	tagCli                 TagClient
	regionCli              RegionClient
	enterpriseProjectCli   EnterpriseProjectClient
}

func NewClientSet(baseDomain string, options ...Option) (ClientSet, error) {
	containerGroupCli, err := NewContainerGroupClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	imageCacheCli, err := NewImageCacheClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	dataCacheCli, err := NewDataCacheClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	virtualNodeCli, err := NewVirtualNodeClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	priceCli, err := NewPriceClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	commitContainerTaskCli, err := NewCommitContainerTaskClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	tagCli, err := NewTagClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	regionCli, err := NewRegionClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}
	enterpriseProjectCli, err := NewEnterpriseProjectClient(baseDomain, options...)
	if err != nil {
		return nil, err
	}

	return &clientSet{
		containerGroupCli:      containerGroupCli,
		imageCacheCli:          imageCacheCli,
		dataCacheCli:           dataCacheCli,
		virtualNodeCli:         virtualNodeCli,
		priceCli:               priceCli,
		commitContainerTaskCli: commitContainerTaskCli,
		tagCli:                 tagCli,
		regionCli:              regionCli,
		enterpriseProjectCli:   enterpriseProjectCli,
	}, nil
}

func (cs *clientSet) ContainerGroup() ContainerGroupClient {
	return cs.containerGroupCli
}

func (cs *clientSet) ImageCache() ImageCacheClient {
	return cs.imageCacheCli
}

func (cs *clientSet) DataCache() DataCacheClient {
	return cs.dataCacheCli
}

func (cs *clientSet) VirtualNode() VirtualNodeClient {
	return cs.virtualNodeCli
}

func (cs *clientSet) Price() PriceClient {
	return cs.priceCli
}

func (cs *clientSet) CommitContainerTask() CommitContainerTaskClient {
	return cs.commitContainerTaskCli
}

func (cs *clientSet) Tag() TagClient {
	return cs.tagCli
}

func (cs *clientSet) Region() RegionClient {
	return cs.regionCli
}

func (cs *clientSet) EnterpriseProject() EnterpriseProjectClient {
	return cs.enterpriseProjectCli
}
