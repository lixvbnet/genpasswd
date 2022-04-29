# genpasswd

A simple utility for generating Unix password hash.

## install
```shell
go get https://github.com/lixvbnet/genpasswd
```

or download pre-compiled binaries from [Releases](https://github.com/lixvbnet/genpasswd/releases) page.

## usage
```shell
$ genpasswd -h
Usage: genpasswd [options] [password]
options
  -1	use MD5 based Unix password algorithm 1
  -5	use SHA256 based Unix password algorithm 5
  -6	use SHA512 based Unix password algorithm 6 (default)
  -h	show help and exit
  -s string
    	salt
  -v	show version
```

## format
```shell
$id$salt$hash
```
> If salt is not given, a random one will be generated.

## example
```shell
$ genpasswd 
Enter password: 
Confirm password: 
$6$VVa90yWsdH7LCzB7$BebmBFPXOP6.yip1P8kTNiLV8I.viWmaFiMHGOZEeIR1B.9S8wA48eRTQ3E3hgpFphsUY3taETCSEFe9H21JH1

$ genpasswd mycoolpassword
$6$2JM3A96DZP4wYU5Z$cjMpLPkfvirqGfUyyoCV0PYJHZEZAqAKChi05TQ53dUlXsg7/SvN1hXi01SB5/NfI8cX50aENvhahH2Jq5Ef3.

$ genpasswd -1 mycoolpassword
$1$j9VdnOK6$3La77YHySjAIUCtbkc463/

$ genpasswd -5 mycoolpassword
$5$Fy7t8TVAjhGyCCba$9WenuUHTFv49Y9XxgRW2FG5y1FiCB6fmgInie34/578

$ genpasswd -s mycoolsalt mycoolpassword
$6$mycoolsalt$5XN0TySv7mjdGB7WIvPx.u1uwWzBubMviWgIXFFfuHcHG5nRddzJ9wdt4Ao/hapd.ySwKJ/eRMtBPOpjJmgLH/
```