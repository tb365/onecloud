package shell

import (
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/mcclient/options"
)

func init() {
	R(&options.LoadbalancerCertificateGetOptions{}, "lbcert-cache-show", "Show cached lbcert", func(s *mcclient.ClientSession, opts *options.LoadbalancerCertificateGetOptions) error {
		lbcert, err := modules.LoadbalancerCachedCertificates.Get(s, opts.ID, nil)
		if err != nil {
			return err
		}
		printObject(lbcert)
		return nil
	})
	R(&options.LoadbalancerCertificateListOptions{}, "lbcert-cached-list", "List cached lbcerts", func(s *mcclient.ClientSession, opts *options.LoadbalancerCertificateListOptions) error {
		params, err := options.ListStructToParams(opts)
		if err != nil {
			return err
		}
		result, err := modules.LoadbalancerCachedCertificates.List(s, params)
		if err != nil {
			return err
		}
		printList(result, modules.LoadbalancerCachedCertificates.GetColumns(s))
		return nil
	})
	R(&options.LoadbalancerCertificateDeleteOptions{}, "lbcert-cached-delete", "Delete cached lbcert", func(s *mcclient.ClientSession, opts *options.LoadbalancerCertificateDeleteOptions) error {
		lbcert, err := modules.LoadbalancerCachedCertificates.Delete(s, opts.ID, nil)
		if err != nil {
			return err
		}
		printObject(lbcert)
		return nil
	})
	R(&options.LoadbalancerCertificateDeleteOptions{}, "lbcert-cached-purge", "Purge cached lbcert", func(s *mcclient.ClientSession, opts *options.LoadbalancerCertificateDeleteOptions) error {
		lbcert, err := modules.LoadbalancerCachedCertificates.PerformAction(s, opts.ID, "purge", nil)
		if err != nil {
			return err
		}
		printObject(lbcert)
		return nil
	})
}
