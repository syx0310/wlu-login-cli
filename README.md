# wlu-login-cli

> A CLI tool for web auth of westlake university, based on [hdu-cli](https://github.com/hduhelp/hdu-cli) matained by [hduhelp](https://github.com/hduhelp/)

## Installation Or Upgrade

```shell
go install github.com/syx0310/wlu-login-cli@latest
```

or direct download the release file which suffix match your platform.

## Startup

use command like
the ac_id **might** be 1 if you are using wireless network, and 6 if you are using wired network

```
wlu-login-cli net login --username {Your student number} --password {Your Westlake University SSO Password} -a {Your ac_id} --save
```

or manually use the .wlu-login-cli.yaml and fill according the comments

<details>
<summary>Trouble shoot</summary>

> The Command may need root privilege
>
> and sometimes go env is not install completely on your root account (sudo mode)
>
> so try like `sudo $GOROOT/bin/go install github.com/syx0310/wlu-login-cli@latest`
> 
> By the way, if you follow the offical installation guide of GO, The goroot will be like /usr/local/go/
</details>

## Usage

### wlu-login-cli [sub command]

### Available Sub Commands:

- completion  
  - generate the autocompletion script for the specified shell
- help        
  - Help about any command
- net         
  - westlake university network auth cli

### Flags:

- --config string   
  - config file (default is $HOME/.wlu-login-cli.yaml)
  - more detail see comments at [wlu-login-cli.yaml example](./.wlu-login-cli.yaml.example)
- -h, --help            
  - help for wlu-login-cli
- -s, --save            
  - save config
- -V, --verbose         
  - show more info
- -v, --version         
  - version for wlu-login-cli


Use `wlu-login-cli [sub command] --help` for more information about a command.


