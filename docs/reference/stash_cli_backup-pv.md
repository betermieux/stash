---
title: Stash Cli Backup-Pv
menu:
  product_stash_0.8.3:
    identifier: stash-cli-backup-pv
    name: Stash Cli Backup-Pv
    parent: reference
product_name: stash
menu_name: product_stash_0.8.3
section_menu_id: reference
---
## stash cli backup-pv

Backup persistent volume

### Synopsis

Backup persistent volume using BackupConfiguration Template

```
stash cli backup-pv [flags]
```

### Options

```
      --directories strings   List of target directories.
  -h, --help                  help for backup-pv
      --kubeconfig string     Path of the Kube config file.
      --mountpath string      Mount path for PVC.
      --namespace string      Namespace for Persistent Volume Claim. (default "default")
      --template string       Name of the BackupConfigurationTemplate.
      --volume string         Name of the Persistent volume.
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --bypass-validating-webhook-xray   if true, bypasses validating webhook xray checks
      --enable-analytics                 Send analytical events to Google Analytics (default true)
      --enable-status-subresource        If true, uses sub resource for crds.
      --log-flush-frequency duration     Maximum number of seconds between log flushes (default 5s)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default true)
      --service-name string              Stash service name. (default "stash-operator")
      --stderrthreshold severity         logs at or above this threshold go to stderr
      --use-kubeapiserver-fqdn-for-aks   if true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [stash cli](/docs/reference/stash_cli.md)	 - Stash CLI

