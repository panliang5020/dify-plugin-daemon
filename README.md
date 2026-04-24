# dify-plugin-daemon

## Dify 1.14 Compatibility

This repository tracks compatibility fixes for the `dify-plugin-daemon` when used with Dify 1.14.

### Required Daemon Version

Dify 1.14 requires **dify-plugin-daemon 0.5.8 or later**. Upgrade to **0.5.9** for the best experience.

### Known Issues Fixed in 0.5.8 and 0.5.9

#### 1. `PluginDaemonInternalServerError: no proper instance` (Fixed in 0.5.9)

**Symptom:** When adding model credentials, validating a plugin, or invoking a tool/model, the following error appears:

```
PluginDaemonInternalServerError: no proper instance
```

**Root Cause:** A race condition in plugin instance startup. The daemon notified callers that a plugin instance was "ready" *before* actually adding it to the active instances list. Any request dispatched between the notification and list insertion would fail to find a running instance.

**Fix (committed in 0.5.9, PR #709):** The `OnInstanceReadyImpl` callback in `internal/core/local_runtime/subprocess.go` was corrected to:
1. Mark the instance as started
2. Add it to the instances list
3. *Then* fire the ready notification

This eliminates the window where a request could arrive and find no available instance.

#### 2. Optimistic Lock Race in Redis Cache (Fixed in 0.5.9)

**Symptom:** Intermittent plugin failures under concurrent load, especially on multi-worker deployments.

**Root Cause:** The Redis `WATCH` keys were not being passed to the transaction handler, so the optimistic lock was silently bypassed and concurrent writes could overwrite each other.

**Fix (committed in 0.5.9, PR #701):** Watch keys are now correctly threaded into the `Transaction` call.

#### 3. Binary Link Datasource Support (New in 0.5.8)

**Symptom:** Datasource plugins returning binary file links fail with a parse error on 0.5.7 or earlier.

**Fix (0.5.8, PR #669):** Added `DataSourceResponseChunkTypeBinaryLink` to handle binary link chunks from datasource plugins.

#### 4. Question Classifier Node Error (Fixed in Dify main)

**Symptom:** Running a workflow with a Question Classifier node throws:

```
LLMNode.invoke_llm() got an unexpected keyword argument 'structured_output_enabled'
```

**Root Cause:** A refactoring oversight in Dify 1.14 RC1 — not a daemon bug.

**Fix:** Already merged in the Dify repository (PR #32902). Update your Dify API image.

### Version Compatibility Matrix

| Dify version | Min daemon version | Recommended |
|---|---|---|
| 1.14.x | 0.5.8 | 0.5.9 |
| 1.13.x | 0.5.5 | 0.5.8 |
| 1.12.x | 0.5.0 | 0.5.5 |

### Upgrade Instructions

#### Docker Compose

```bash
docker compose pull
docker compose up -d
```

#### Manual binary replacement

Download the latest release from [langgenius/dify-plugin-daemon releases](https://github.com/langgenius/dify-plugin-daemon/releases/tag/0.5.9) and replace the existing binary.

### Upstream Repository

All fixes described here are available in the upstream repository:
[langgenius/dify-plugin-daemon](https://github.com/langgenius/dify-plugin-daemon)
