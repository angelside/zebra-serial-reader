# Zebra Serial Reader CLI

[![maintenance-status](https://img.shields.io/badge/maintenance-as--is-yellow.svg?style=for-the-badge)](https://gist.github.com/angelside/364976fbcf7001a5da7e79ad8ed91fec) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
[![Windows](https://img.shields.io/badge/Windows-0078D6.svg?style=for-the-badge&logo=Windows&logoColor=white)](https://www.microsoft.com)
[![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge&logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciICB2aWV3Qm94PSIwIDAgNDggNDgiIHdpZHRoPSI0OHB4IiBoZWlnaHQ9IjQ4cHgiPjxwYXRoIGZpbGw9IiM0Y2FmNTAiIGQ9Ik0yNCw1QzEzLjUsNSw1LDEzLjYsNSwyNC4xYzAsOC4yLDUuMSwxNS4xLDEyLjMsMTcuOWw0LjItMTEuNUMxOC44LDI5LjUsMTcsMjcsMTcsMjQgYzAtMy45LDMuMS03LDctN3M3LDMuMSw3LDdjMCwzLTEuOCw1LjUtNC41LDYuNUwzMC43LDQyQzM3LjksMzkuMiw0MywzMi4zLDQzLDI0LjFDNDMsMTMuNiwzNC41LDUsMjQsNXoiLz48cGF0aCBmaWxsPSIjMmU3ZDMyIiBkPSJNMTcuOSw0My4zbC0wLjktMC40QzkuMiw0MCw0LDMyLjQsNCwyNC4xQzQsMTMsMTMsNCwyNCw0YzExLDAsMjAsOSwyMCwyMC4xIGMwLDguMy01LjIsMTUuOS0xMi45LDE4LjhsLTAuOSwwLjRsLTQuOC0xMy4zbDAuOS0wLjRjMi4zLTAuOSwzLjgtMy4xLDMuOC01LjZjMC0zLjMtMi43LTYtNi02cy02LDIuNy02LDZjMCwyLjUsMS41LDQuNywzLjgsNS42IGwwLjksMC40TDE3LjksNDMuM3ogTTI0LDZDMTQuMSw2LDYsMTQuMSw2LDI0LjFjMCw3LjEsNC4zLDEzLjcsMTAuNywxNi41bDMuNS05LjZDMTcuNiwyOS43LDE2LDI3LDE2LDI0YzAtNC40LDMuNi04LDgtOCBzOCwzLjYsOCw4YzAsMy0xLjYsNS43LTQuMiw3bDMuNSw5LjZDMzcuNywzNy44LDQyLDMxLjMsNDIsMjQuMUM0MiwxNC4xLDMzLjksNiwyNCw2eiIvPjwvc3ZnPg==)](./LICENSE)

> The zebra-serial-reader finds serial numbers and product names of Zebra printers on the network via IP addresses.

### Tested on

- [x] Windows
- [ ] Linux
- [ ] MacOS

### TODO

- [ ] Test if we can awake printer from sleep.


## 📦 Installation

No installations, it's a single portable binary file. You can put in your binary path and use it from anywhere.

## 🔨 Usage

1. Download or build the binary file.
2. Create a `data.json` file next to the binary file following the template below. The first variable can be anything that identifies the printer. Second variable is printer's IP address.

```json
{
  "OFFICE_1": "127.0.0.1",
  "OFFICE_2": "127.0.0.2",
  "OFFICE_3": "127.0.0.3"
}
```

3. Run the binary without any parameters. It will read the data.json file, print the result to the console, write the results to the `__printers.csv` file, and then exit.

```bash
.\zebra-serial-reader.exe
```

## 🚀 Build from source

With task [Task](https://taskfile.dev/)

```bash
task build
```

With `go build`

```bash
go build -v -ldflags="-s -w" -o ./__dist/zebra-serial-reader.exe ./cmd/.
```

## 📋 Sample results
```bash
== Zebra serial number extractor

[OK] Location: OFFICE_1, Serial number: D3Jxxxxxx, Product: ZD620, IP: 127.0.0.1
[OK] Location: OFFICE_2, Serial number: D3Jxxxxxx, Product: ZD620, IP: 127.0.0.2
[ERROR] Location: OFFICE_3, IP: 127.0.0.4

CSV file created successfully: __printers.csv
```

## 🎯 Tested Zebra printer models

- ZD 620
- ZD 621
- GK 420d

## 🤝 Contributing

Before contributing issues or pull requests, could you review the [Contributing Guidelines](./.github/CONTRIBUTING.md) first?

## 🤩 Support

💙 If you like this project, give it a ⭐ and share it with friends!

## 🏛️ License

This project is open-sourced software licensed under the [MIT license](./LICENSE).
