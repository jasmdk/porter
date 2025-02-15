package porter

import (
	"context"
	"fmt"

	"get.porter.sh/porter/pkg/cache"
	cnabtooci "get.porter.sh/porter/pkg/cnab/cnab-to-oci"
	"get.porter.sh/porter/pkg/tracing"
)

type BundleResolver struct {
	Cache    cache.BundleCache
	Registry cnabtooci.RegistryProvider
}

// Resolves a bundle from the cache, or pulls it and caches it
// Returns the location of the bundle or an error
func (r *BundleResolver) Resolve(ctx context.Context, opts BundlePullOptions) (cache.CachedBundle, error) {
	log := tracing.LoggerFromContext(ctx)

	if !opts.Force {
		cachedBundle, ok, err := r.Cache.FindBundle(opts.GetReference())
		if err != nil {
			return cache.CachedBundle{}, log.Error(fmt.Errorf("unable to load bundle %s from cache: %w", opts.Reference, err))
		}
		// If we found the bundle, return the path to the bundle.json
		if ok {
			return cachedBundle, nil
		}
	}

	bundleRef, err := r.Registry.PullBundle(ctx, opts.GetReference(), opts.InsecureRegistry)
	if err != nil {
		return cache.CachedBundle{}, err
	}

	return r.Cache.StoreBundle(bundleRef)
}
