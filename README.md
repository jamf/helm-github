# helm-github

This helm plugin allows to download helm charts from private GitHub repositories.

## How to test locally

You can test your changes locally by installing the locally built binary.

1. Uninstall current version of the plugin.

    ```bash
    helm plugin remove github
    ```

2. Build your local version.

    ```bash
    make build
    ```

3. Install your local version.

    ```bash
    export HELM_GITHUB_PLUGIN_NO_INSTALL_HOOK=1
    helm plugin install .
    ```

4. Check the plugin has been installed.

    ```bash
    helm plugin list
    ```

## How to create new version

1. Create a new branch from `master` branch with your changes.
2. Make sure that the version in `plugin.yaml` file is updated.
3. Create a pull request to `master` branch.
4. Make sure that the pull request is reviewed and merged.
5. Tag the commit with the same version as in `plugin.yaml` file.
6. Release workflow will be triggered and the new version will be published to GitHub releases.
