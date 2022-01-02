<p align="center">

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/davidkroell/bodycomposition/Go%20build%20and%20test)](https://github.com/davidkroell/bodycomposition/actions/workflows/build.yml)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/davidkroell/bodycomposition)](https://github.com/davidkroell/bodycomposition/releases/latest)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/davidkroell/bodycomposition)](https://github.com/davidkroell/bodycomposition/blob/master/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidkroell/bodycomposition)](https://goreportcard.com/report/github.com/davidkroell/bodycomposition)
[![Go Reference](https://pkg.go.dev/badge/github.com/davidkroell/bodycomposition.svg)](https://pkg.go.dev/github.com/davidkroell/bodycomposition)
[![GitHub all releases](https://img.shields.io/github/downloads/davidkroell/bodycomposition/total?color=orange)](https://github.com/davidkroell/bodycomposition/releases)
[![GitHub license](https://img.shields.io/github/license/davidkroell/bodycomposition)](https://github.com/davidkroell/bodycomposition/blob/master/LICENSE)

</p>

# bodycompostion
Bodycomposition is a program to manage your body measurements and composition stored in
Garmin Connect Cloud (https://connect.garmin.com) from your beloved commandline.

For now, you can just add body composition values. Any other thing should be done in Garmin Connect.

## Version 2.2
Parameter `--max-tries` is obsolete and therefore no longer present.
Furthermore, two new parameters `--bmi` and `--calories` were added.

### Migrating from Version 2.x

* Remove the parameter `--max-tries`

## Version 2
After some design and command line interface changed, a new release, `v2.0.0` was created.

### Migrating from Version 1

* Parameters `--metabolicAge`, `--physiqueRating` and `--visceralFat` changed to `--metabolic-age`, `--physique-rating` and `--visceral-fat`, to create a consistent CLI.
* Parameter `--max-tries` was added because we had problems with a dependency. These bugs are now changed, so **--max-tries should not be necessary**
* Exit codes: on error, the exit code `1` is used. If no error occured, exit code `0` is used. Under the hood, error handling was improved

## Download
Releases can be found in [release](https://github.com/davidkroell/bodycomposition/releases) tab.


## Usage

Uploading weight to Garmin connect. If you do not provide a password (`--password`), it will be prompted from stdin.
```bash
$ ./bodycomposition upload --weight 80 --bone 14 --fat 13 --hydration 58 --muscle 42 --email john.doe@mail.com
```

General usage
```bash
$ ./bodycomposition -h
Bodycomposition is a program to manage your body measurements and composition stored in
Garmin Connect Cloud (https://connect.garmin.com) from your beloved commandline.

For now, you can just add body composition values. Any other thing should be done in Garmin Connect.
Version v2.0.0

Usage:
  bodycomposition [command]

Available Commands:
  help        Help about any command
  upload      Upload your body composition values to Garmin Connect

Flags:
  -h, --help   help for bodycomposition

Use "bodycomposition [command] --help" for more information about a command.
```

#### Upload command usage

```bash
$ ./bodycomposition upload -h
Upload your body composition values to Garmin Connect

Usage:
  bodycomposition upload [flags]

Aliases:
  upload, u, add

Flags:
      --bmi float               Set your BMI - body mass index
  -b, --bone float              Set your bone mass in percent
  -c, --calories float          Set your caloric intake
  -e, --email string            Email of the Garmin account
  -f, --fat float               Set your fat in percent
  -h, --help                    help for upload
      --hydration float         Set your hydration in percent
      --metabolic-age float     Set your metabolic age
  -m, --muscle float            Set your muscle mass in percent
  -p, --password string         Password of the Garmin account
      --physique-rating float   Set your physique rating (valid values: 1-9)
  -t, --unix-timestamp int      Set the timestamp of the measurement (default -1)
      --visceral-fat float      Set your visceral fat rating (valid values: 1-60)
  -w, --weight float            Set your weight in kilograms (default -1)
```
