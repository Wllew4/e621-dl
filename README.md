# e621-dl
Download posts, pools, and queries in bulk from [e621.net](https://e621.net). Inspired by [jacklul/e621-Pool-Downloader](https://github.com/jacklul/e621-Pool-Downloader).

# ⚠️🔞 CONTENT WARNING 🔞⚠️
For those unfamiliar, [e621.net](https://e621.net) hosts drawn/animated adult media.
If you found this repository haphazardly while perusing my catalog of personal
projects, you may want to skip over this one.

# Usage
1. Download the latest release for your platform [from here](https://github.com/Wllew4/e621-dl/releases).
2. Open a terminal in the directory with the executable, or add the executable to your PATH.

## Fetch a post
```bash
# By URL
e621-dl post --url https://e621.net/posts/3076214
# By ID
e621-dl post --id 3076214
# Specify the output file
e621-dl post --id 3076214 --out federal_agents.jpg
```

## Download a pool
```bash
# By URL
e621-dl pool --url https://e621.net/pools/<pool-id>
# By ID
e621-dl pool --id <pool-id>
# Specify the output directory
e621-dl post --id <pool-id> --out wowie
```

## Download from search
```bash
# Search for tags
e621-dl query --tags male,female
# Get the first n results
e621-dl query --tags male,female --first 20
# Specify the output directory
e621-dl query --tags male,female --out wowie
```

# Friendly reminder to give me money
if you want https://ko-fi.com/soupsu
