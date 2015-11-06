# stellar-upgrade

Tool to upgrade your account from STR to XLM on the much improved network.

## Downloading the tool
[Prebuilt binaries](https://github.com/stellar/stellar-upgrade/releases) of the upgrade tool are available are on the [releases page](https://github.com/stellar/stellar-upgrade/releases).

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 32 bit | [stellar-upgrade-darwin-386](https://github.com/stellar/stellar-upgrade/releases)        |
| Mac OSX 64 bit | [stellar-upgrade-darwin-amd64](https://github.com/stellar/stellar-upgrade/releases)      |
| Linux 32 bit   | [stellar-upgrade-linux-386](https://github.com/stellar/stellar-upgrade/releases)         |
| Linux 64 bit   | [stellar-upgrade-linux-amd64](https://github.com/stellar/stellar-upgrade/releases)       |
| Windows 32 bit | [stellar-upgrade-windows-386.exe](https://github.com/stellar/stellar-upgrade/releases)   |
| Windows 64 bit | [stellar-upgrade-windows-amd64.exe](https://github.com/stellar/stellar-upgrade/releases) |

Alternatively, you can build the binary yourself. [gb](http://getgb.io) is used for building this upgrade tool.

## Usage
Remember to make sure that you have the correct name of the binary.

### Upgrade an account
```shell
./stellar-upgrade-[binary-suffix] upgrade
```

`upgrade` is the only argument needed when starting the tool in the command line.
The tool will ask you for the old secret key once you've started the tool.

### Check the upgrade status of an account
```shell
./stellar-upgrade-[binary-suffix] status gYourOldNetworkAddresszpwrG5QVUXqM
```

### Troubleshooting
If you get a "Permission denied" error similar to this:
```shell
-bash: ./stellar-upgrade-[binary-suffix]: Permission denied
```

Use the following command to grant permission to run the file.
```shell
chmod +x stellar-upgrade-[binary-suffix]
```

### Command line help message
```shell
stellar-upgrade upgrades your old network account

Usage:
  stellar-upgrade [command]

Available Commands:
  upgrade     Upgrade account on old network
  status      Displays your account upgrade status
  help        Help about any command

Flags:
  -h, --help=false: help for stellar-upgrade


Use "stellar-upgrade help [command]" for more information about a command.
```