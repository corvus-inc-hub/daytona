/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

import { parseDockerImage } from './docker-image.util'

describe('parseDockerImage', () => {
  it('preserves digest separators when rebuilding an image reference', () => {
    const digest = 'a'.repeat(64)
    const image = parseDockerImage(`ghcr.io/corvus-inc-hub/manifest-sandbox@sha256:${digest}`)

    expect(image.getFullName()).toBe(`ghcr.io/corvus-inc-hub/manifest-sandbox@sha256:${digest}`)
  })
})
