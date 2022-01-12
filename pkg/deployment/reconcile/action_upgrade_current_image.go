//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package reconcile

import (
	"context"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"github.com/rs/zerolog"
)

func init() {
	registerAction(api.ActionTypeSetCurrentImage, newSetCurrentImageAction)
}

// newSetCurrentImageAction creates a new Action that implements the given
// planned SetCurrentImage action.
func newSetCurrentImageAction(log zerolog.Logger, action api.Action, actionCtx ActionContext) Action {
	a := &setCurrentImageAction{}

	a.actionImpl = newActionImplDefRef(log, action, actionCtx, upgradeMemberTimeout)

	return a
}

// setCurrentImageAction implements an SetCurrentImage.
type setCurrentImageAction struct {
	// actionImpl implement timeout and member id functions
	actionImpl
}

// Start performs the start of the action.
// Returns true if the action is completely finished, false in case
// the start time needs to be recorded and a ready condition needs to be checked.
func (a *setCurrentImageAction) Start(ctx context.Context) (bool, error) {
	ready, _, err := a.CheckProgress(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ready, nil
}

// CheckProgress checks the progress of the action.
// Returns true if the action is completely finished, false otherwise.
func (a *setCurrentImageAction) CheckProgress(ctx context.Context) (bool, bool, error) {
	log := a.log

	imageInfo, found := a.actionCtx.GetImageInfo(a.action.Image)
	if !found {
		return false, false, nil
	}
	if err := a.actionCtx.SetCurrentImage(ctx, imageInfo); err != nil {
		log.Error().Err(err).Msg("Unable to set current image")
		return false, false, nil
	}
	log.Info().Str("image", a.action.Image).Str("to", imageInfo.Image).Msg("Changed current main image")
	return true, false, nil
}
