# Cloud Foundation Toolkit CLI Project

## Overview

### Requirements 

 * go 1.12 and above 

### Build and Run 

    make build
    
After build find binary at bin/cft location     

### Usage

Follow cft --help instructions


Google Cloud Foundation Toolkit CLI

```bash
Usage:
  cft [flags]
  cft [command] [flags]

Available Commands:
  help        Help about any command
  version     Print version information
  report      Generate inventory reports based on CAI outputs in a directory.

Flags:
  -h, --help   help for cft

Use "cft [command] --help" for more information about a command.
```
#### Report
Generate inventory reports for resources in Cloud Asset Inventory (CAI) output files, with reports defined in rego (in '<path_to_cloud-foundation-toolkit>/reports/sample' folder).

Example:

```bash
	
	  cft report --query-path <path_to_cloud-foundation-toolkit>/reports/sample \
		--dir-path <path-to-directory-containing-cai-export> \
		--output-path <path-to-directory-for-report-output>
```

## License

Apache 2.0 - See [LICENSE](LICENSE) for more information.