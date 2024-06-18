# bpfmemapie

A tool to render a pie chart of memory usage (bytes_memlock) of BPF maps on the
system. Sorry for the naming, it's for **mem**apie, me**map**ie and mema**pie**,
so BPF mem map pie.

## Installation

```shell
go install github.com/mtardy/bpfmemapie@latest
```

## Usage

It needs bpftool to be installed and to execute as root for bpftool to list the
maps and their info.

```shell
sudo bpfmemapie
```

