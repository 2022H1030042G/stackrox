#!/usr/bin/env bash
#
# Sets some necessary variables up for local testing.
# JIRA_TOKEN still has to be set manually.
#
# Usage example:
#     bash local-env.sh create-cluster flavor name ...
#
export GITHUB_STEP_SUMMARY=/dev/stdout
GITHUB_ACTOR="$(git config --get user.email)"
export GITHUB_ACTOR
export GITHUB_REPOSITORY=stackrox/stackrox
export GITHUB_SERVER_URL=https://github.com

main_branch="$(gh repo view --json defaultBranchRef --jq .defaultBranchRef.name)"
export main_branch

if [ -z "$DRY_RUN" ]; then
    export DRY_RUN=true
fi

SCRIPTS_ROOT="$(git rev-parse --show-toplevel)/.github/workflows/scripts"

CI="false" # true if running in GitHub context
export CI

# Supress shellcheck false warning:
# shellcheck source=/dev/null
source "$SCRIPTS_ROOT/common.sh"

SCRIPT="$SCRIPTS_ROOT/$1.sh"
shift
bash "$SCRIPT" "$@"
