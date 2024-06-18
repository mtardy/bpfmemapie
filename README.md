# bpfmemapie 🥧

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

After opening http://localhost:8080 (you can change the port with `--port`), you
should see something like this:

<img width="1292" alt="output of bpfmemapie: a pie chart with the maps memory" 
  src="https://github.com/mtardy/bpfmemapie/assets/11256051/3b987551-34d2-4270-9421-4b84eff56415">

The pie chart is interactive, you can click on the legend to hide/show entries
or hover on the pie to see the tooltip. You can also change how the "others"
category using the `threshold` URL query, for example 
http://localhost:8080/?threshold=1.
