# haversine CLI

A small program to calculate the distance between two coordinates using the haversine formula.
A coordinate exists of 2 values, Latitude and Longitude.

**Important:**

If the coordinates contains negative values add `--` before the values. This tells the shell that ALL what follows is not an option but a param. That means if you want to use flags you have to write these before the `--`.

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

## Example

Distance from NewYork to Long Beach LA

command:
```bash
haversine-cli calc -o km -- 40.704691303042246 -74.05109689631233 33.76909535040377 -118.19343237522592
```

output:
```bash
Distance between 40.704691, -74.051097 and 33.769095, -118.193432
3940.707km
```
