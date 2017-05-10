
# Stpl API
https://github.com/luebken/stpl

Table of Contents

1. [Get an analysis for a given pom.xml](#analysis)
1. [Get all reference stacks known to the system](#referencestacks)

<a name="analysis"></a>

## analysis

| Specification | Value |
|-----|-----|
| Resource Path | /analysis |
| API Version | 0.0.1 |
| BasePath for the API | {{.}} |
| Consumes | application/xml |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /analysis/ | [GET](#getAnalysis) | Get an analysis for a given pom.xml |



<a name="getAnalysis"></a>

#### API: /analysis/ (GET)


Get an analysis for a given pom.xml



<a name="referencestacks"></a>

## referencestacks

| Specification | Value |
|-----|-----|
| Resource Path | /referencestacks |
| API Version | 0.0.1 |
| BasePath for the API | {{.}} |
| Consumes | application/xml |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /referencestacks/ | [GET](#getReferenceStacks) | Get all reference stacks known to the system |



<a name="getReferenceStacks"></a>

#### API: /referencestacks/ (GET)


Get all reference stacks known to the system



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [ReferenceStack](#github.com.luebken.stpl.pkg.stpl.stack.ReferenceStack) |  |




### Models

<a name="github.com.blang.semver.PRVersion"></a>

#### PRVersion

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| IsNum | bool |  |
| VersionNum | uint64 |  |
| VersionStr | string |  |

<a name="github.com.blang.semver.Version"></a>

#### Version

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Build | array |  |
| Major | uint64 |  |
| Minor | uint64 |  |
| Patch | uint64 |  |
| Pre | array |  |

<a name="github.com.luebken.stpl.pkg.stpl.stack.Dependency"></a>

#### Dependency

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| ArtefactID | string |  |
| GroupID | string |  |
| SemVer | github.com.blang.semver.Version |  |
| VersionString | string |  |

<a name="github.com.luebken.stpl.pkg.stpl.stack.ReferenceStack"></a>

#### ReferenceStack

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Dependencies | array |  |
| Description | string |  |
| Name | string |  |


