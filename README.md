<h1 align="center">elastic-books</h1>
<p align="center">
    <a href="https://github.com/blyndusk/elastic-books/releases">
      <img src="https://img.shields.io/github/v/release/blyndusk/elastic-books"/>
    </a>
    <a href="https://github.com/blyndusk/elastic-books/commits/main">
      <img src="https://img.shields.io/github/release-date/blyndusk/elastic-books"/>
    </a>
    <a href="https://github.com/blyndusk/elastic-books/issues">
      <img src="https://img.shields.io/github/issues/blyndusk/elastic-books"/>
    </a>
    <a href="https://github.com/blyndusk/elastic-books/pulls">
      <img src="https://img.shields.io/github/issues-pr/blyndusk/elastic-books"/>
    </a>
    <a href="https://github.com/blyndusk/elastic-books/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/blyndusk/elastic-books"/>
    </a>
</p>

<p align="center">
  <a href="https://github.com/blyndusk/elastic-books/actions/workflows/go.yml">
      <img src="https://github.com/blyndusk/elastic-books/actions/workflows/go.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/elastic-books/actions/workflows/docker.yml">
      <img src="https://github.com/blyndusk/elastic-books/actions/workflows/docker.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/elastic-books/actions/workflows/release.yml">
      <img src="https://github.com/blyndusk/elastic-books/actions/workflows/release.yml/badge.svg"/>
    </a>
</p>

## I - Goal

The goal of this project is to use the [Elastic search](https://www.elastic.co/) tool with a [Golang](https://golang.org/) API in order to be able to search for books by title, author or summary, using this API. The development environment is build with [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/), using a [Makefile](<https://en.wikipedia.org/wiki/Make_(software)>).

This repository provides **commit writting** and **branch naming conventions**, **issues** and **pull request templates**, and a **custom issues labels preset**.

But also **CI/CD** and **release** using [GitHub Actions](https://github.com/features/actions), [GitHub Container Registry](https://github.com/features/packages) and [Docker](https://www.docker.com/).

## II - Table of content

- [I - Goal](#i---goal)
- [II - Table of content](#ii---table-of-content)
- [III - Conventions, templates and labels](#iii---conventions-templates-and-labels)
  - [A - Commit conventions](#a---commit-conventions)
  - [B - Branch naming convention](#b---branch-naming-convention)
  - [C - Issue template](#c---issue-template)
  - [D - Pull request template](#d---pull-request-template)
  - [E - Custom issues labels preset](#e---custom-issues-labels-preset)
- [IV - CI/CD, release and container registry](#iv---cicd-release-and-container-registry)
  - [A - CI](#a---ci)
  - [B - CD](#b---cd)
  - [C - Release](#c---release)
- [V - Elastic-books product](#v---elastic-books-product)
  - [A - Stack](#a---stack)
  - [B - Makefile](#b---makefile)
- [VI - License](#vi---license)

## III - Conventions, templates and labels

### A - Commit conventions

```
tag(scope): #issue_id - message
```

See [commit_conventions.md](.github/commit_conventions.md) for more informations.

### B - Branch naming convention

```
type_scope-of-the-work
```

See [branch_naming_convention.md](.github/branch_naming_convention.md) for more informations.

### C - Issue template

See [user-story.md](.github/ISSUE_TEMPLATE/user-story.md) for more informations.

### D - Pull request template

See [pull_request_template.md](.github/pull_request_template.md) for more informations.

### E - Custom issues labels preset

The labels preset is located at [.github/settings.yml](.github/settings.yml).

You can **add, edit or remove** them. To automatically update these labels, you need to **install** the ["Settings" GitHub app](https://github.com/apps/settings), which will **syncs repository settings defined in the file above to your repository**.

## IV - CI/CD, release and container registry

### A - CI

[![GO](https://github.com/blyndusk/elastic-books/actions/workflows/go.yml/badge.svg)](https://github.com/blyndusk/elastic-books/actions/workflows/go.yml)

The **CI** workflow is located at [.github/workflows/go.yml](.github/workflows/go.yml). It's triggered a **each push** on **all branches**.

It consist of:

- **install Golang** on the VM
- get the dependancies **using the cache of other Actions run**
- **lint the files** to check or syntax errors
- **build** the application

### B - CD

[![DOCKER](https://github.com/blyndusk/elastic-books/actions/workflows/docker.yml/badge.svg)](https://github.com/blyndusk/elastic-books/actions/workflows/docker.yml)

The **CD** workflow is located at [.github/workflows/docker.yml](.github/workflows/docker.yml). It's triggered a **each push** on **`main` branch**.

It consist of:

- **login** into the GitHub container registry (ghcr.io)
- **build and push** the Golang api using the **production Dockerfile** located at [.docker/api/prod.Dockerfile](.docker/api/prod.Dockerfile)

After that, you can check the **pushed container** at: `https://github.com/<username>?tab=packages&repo_name=<repository-name>`

> IMPORTANT: you need to **update the production Dockerfile** with your **username** AND **_repository name_**. Otherwise, there will be errors at push:

```bash
LABEL org.opencontainers.image.source = "https://github.com/<username>/<repository-name>"
```

### C - Release

[![RELEASE](https://github.com/blyndusk/elastic-books/actions/workflows/release.yml/badge.svg)](https://github.com/blyndusk/elastic-books/actions/workflows/release.yml)

The **release** workflow is located at [.github/workflows/release.yml](.github/workflows/release.yml). It's triggered **manually by user input** at: [Actions > RELEASE](https://github.com/blyndusk/elastic-books/actions/workflows/release.yml).

> IMPORTANT: you need to set the **image tag** in the action input, for the action to be able to push the docker image and create a release **with a specific version**. The image tag is a [SemVer](https://en.wikipedia.org/wiki/Software_versioning) tag, e.g. `1.0.2`.

It consist of:

- check if the **environment match the branch**
- do the CD (docker) action again, but **with a specific image tag**
- create a release **with the same tag**, filled with the **generated changelog as closed issues since the last release**

After that, you can check the release at `https://github.com/<username>/<repository-name>/releases`.

## V - Elastic-books product

The project use **Docker** and **Docker Compose** to build and run local and distant images in our workspace.

> For more informations, check the [wiki](https://github.com/blyndusk/elastic-books/wiki).

### A - Stack

All the images use the **same network**, more informations at [docker-compose.yml](docker-compose.yml)

| CONTAINER | PORT        | IMAGE                                                                                                        |
| :-------- | :---------- | :----------------------------------------------------------------------------------------------------------- |
| GOLANG    | `3333:3333` | [.docker/api/dev.Dockerfile](.docker/api/dev.Dockerfile)                                                     |
| ELASTIC   | `9200:9200` | [docker.elastic.co/elasticsearch/elasticsearch:7.13.2](docker.elastic.co/elasticsearch/elasticsearch:7.13.2) |
| KIBANA    | `5601:5601` | [docker.elastic.co/kibana/kibana:7.13.2](docker.elastic.co/kibana/kibana:7.13.2)                             |

### B - Makefile

#### TL;DR <!-- omit in toc -->

```bash
make lint start logs
```

> You will have errors at 1st build, because the Golang API is ready before the ES service. Just `Ctrl + S` on a Golang file will fix the problem (hot-reload)

#### `make help` <!-- omit in toc -->

**Display** informations about other commands.

#### `make start` <!-- omit in toc -->

Up the containers with **full cache reset** to avoid cache errors.

#### `make stop` <!-- omit in toc -->

**Down** the containers.

#### `make logs` <!-- omit in toc -->

**Display and follow** the logs.

#### `make lint` <!-- omit in toc -->

**Lint** the Go files using `gofmt`.

#### `make seed` <!-- omit in toc -->

**Seed** the Elasticsearch database with sample data

## VI - License

Under [MIT](./LICENSE) license.
