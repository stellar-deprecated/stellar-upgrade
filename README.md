# stellar-upgrade

Tool to upgrade your account from STR to XLM on the much improved network.

## Downloading the tool
[Prebuilt binaries](https://github.com/stellar/stellar-upgrade/releases) of the upgrade tool are available are on the [releases page](https://github.com/stellar/stellar-upgrade/releases).

| Platform       | Binary file name                  |
|----------------|-----------------------------------|
| Mac OSX 32 bit | stellar-upgrade-darwin-386        |
| Mac OSX 64 bit | stellar-upgrade-darwin-amd64      |
| Linux 32 bit   | stellar-upgrade-linux-386         |
| Linux 64 bit   | stellar-upgrade-linux-amd64       |
| Windows 32 bit | stellar-upgrade-windows-386.exe   |
| Windows 64 bit | stellar-upgrade-windows-amd64.exe |

Alternatively, you can build the binary yourself. [gb](http://getgb.io) is used for building this upgrade tool.

## Usage
Remember to make sure that you have the correct name of the binary.

### Check the upgrade status of an account
```shell
./stellar-upgrade-[binary-suffix] status ganVp9o5emfzpwrG5QVUXqMv8AgLcdvySb
```

### Upgrade an account
```shell
./stellar-upgrade-[binary-suffix] upgrade
```

### Command line help
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
