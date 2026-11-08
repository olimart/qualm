package cli

// Version is the release this source builds. It has a file to itself because the release workflow
// watches this path: a change here naming an unpublished version publishes a release. The gem
// takes its version from here too.
const Version = "0.1.2"
