// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/stats"

	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type ApiSimpleEmitter interface {
	Snapshots(ctx context.Context) (<-chan *ApiSnapshot, <-chan error, error)
}

func NewApiSimpleEmitter(aggregatedWatch clients.ResourceWatch) ApiSimpleEmitter {
	return NewApiSimpleEmitterWithEmit(aggregatedWatch, make(chan struct{}))
}

func NewApiSimpleEmitterWithEmit(aggregatedWatch clients.ResourceWatch, emit <-chan struct{}) ApiSimpleEmitter {
	return &apiSimpleEmitter{
		aggregatedWatch: aggregatedWatch,
		forceEmit:       emit,
	}
}

type apiSimpleEmitter struct {
	forceEmit       <-chan struct{}
	aggregatedWatch clients.ResourceWatch
}

func (c *apiSimpleEmitter) Snapshots(ctx context.Context) (<-chan *ApiSnapshot, <-chan error, error) {
	snapshots := make(chan *ApiSnapshot)
	errs := make(chan error)

	untyped, watchErrs, err := c.aggregatedWatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	go errutils.AggregateErrs(ctx, errs, watchErrs, "api-emitter")

	go func() {
		currentSnapshot := ApiSnapshot{}
		timer := time.NewTicker(time.Second * 1)
		var previousHash uint64
		sync := func() {
			currentHash := currentSnapshot.Hash()
			if previousHash == currentHash {
				return
			}

			previousHash = currentHash

			stats.Record(ctx, mApiSnapshotOut.M(1))
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		defer func() {
			close(snapshots)
			close(errs)
		}()

		for {
			record := func() { stats.Record(ctx, mApiSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case untypedList := <-untyped:
				record()

				currentSnapshot = ApiSnapshot{}
				for _, res := range untypedList {
					switch typed := res.(type) {
					case *Artifact:
						currentSnapshot.Artifacts = append(currentSnapshot.Artifacts, typed)
					case *Endpoint:
						currentSnapshot.Endpoints = append(currentSnapshot.Endpoints, typed)
					case *Proxy:
						currentSnapshot.Proxies = append(currentSnapshot.Proxies, typed)
					case *UpstreamGroup:
						currentSnapshot.UpstreamGroups = append(currentSnapshot.UpstreamGroups, typed)
					case *Secret:
						currentSnapshot.Secrets = append(currentSnapshot.Secrets, typed)
					case *Upstream:
						currentSnapshot.Upstreams = append(currentSnapshot.Upstreams, typed)
					default:
						select {
						case errs <- fmt.Errorf("ApiSnapshotEmitter "+
							"cannot process resource %v of type %T", res.GetMetadata().Ref(), res):
						case <-ctx.Done():
							return
						}
					}
				}

			}
		}
	}()
	return snapshots, errs, nil
}
