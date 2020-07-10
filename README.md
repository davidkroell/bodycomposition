# bodycompostion
Bodycomposition is a program to manage your body measurements and composition stored in
Garmin Connect Cloud (https://connect.garmin.com) from your beloved commandline.

For now, you can just add body composition values. Any other thing should be done in Garmin Connect.

## Download
Releases can be found in [release](https://github.com/davidkroell/bodycomposition/releases) tab.


## Usage

General usage
```bash
$ ./bodycompsition -h
Bodycomposition is a program to manage your body measurements and composition stored in
Garmin Connect Cloud (https://connect.garmin.com) from your beloved commandline.

For now, you can just add body composition values. Any other thing should be done in Garmin Connect.
Version v1.0.0

Usage:
  bodycomposition [command]

Available Commands:
  help        Help about any command
  upload      Upload your body composition values to Garmin Connect

Flags:
  -h, --help   help for bodycomposition

Use "bodycomposition [command] --help" for more information about a command.
```

Upload command usage

```bash
$ ./bodycomposition upload -h
Upload your body composition values to Garmin Connect

Usage:
  bodycomposition upload [flags]

Aliases:
  upload, u, add

Flags:
  -b, --bone float             Set your bone mass in percent
  -e, --email string           Email of the Garmin account
  -f, --fat float              Set your fat in percent
  -h, --help                   help for upload
      --hydration float        Set your hydration in percent
      --metabolicAge float     Set your metabolic age
  -m, --muscle float           Set your muscle mass in percent
  -p, --password string        Password of the Garmin account
      --physiqueRating float   Set your physique rating (valid values: 1-9)
  -t, --unix-timestamp int     Set the timestamp of the measurement (default -1)
      --visceralFat float      Set your visceral fat rating (valid values: 1-60)
  -w, --weight float           Set your weight in kilograms (default -1)
```
