# haversine CLI

A small program to calculate the distance between two coordinates using the haversine formula.
A coordinate exists of 2 values, Latitude and Longitude.

## Usage

```bash
cli for harversine formula

Usage:
  haversine-cli [command]

Available Commands:
  calc        calculate distance between two coordinates
  help        Help about any command

Flags:
  -h, --help     help for haversine-cli
  -t, --toggle   Help message for toggle

Use "haversine-cli [command] --help" for more information about a command.
```

### Usage of calc command

```bash
calculate distance between two coordinates

Usage:
  haversine-cli calc [lat1] [lng1] [lat2] [lng2] [flags]

Flags:
  -h, --help         help for calc
  -o, --out string   output format (km, m, cm) (default "m")
  -s, --short        prints only the value
```
